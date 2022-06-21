package service

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"reflect"
	"sensitive-storage/env"
	"sensitive-storage/module/entity"
	"sensitive-storage/util/collection"
	"strings"
	"time"
)

var Client *gorm.DB

func InitDataBase() *sql.DB {
	var err error
	db := os.Getenv(env.DB)
	var dial gorm.Dialector
	if db == "" || db == "sqlite" {
		dial = sqlite.Open("./data/sensitive.db")
	} else if db == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local&timeout=5s&charset=utf8mb4&collation=utf8mb4_unicode_ci&interpolateParams=true&parseTime=true&loc=Local",
			os.Getenv("db_username"), os.Getenv("db_password"), os.Getenv("db_host"), os.Getenv("db_port"), os.Getenv("db_database"))
		dial = mysql.Open(dsn)
	} else {
		panic("db not validate")
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	Client, err = gorm.Open(dial, &gorm.Config{
		SkipDefaultTransaction: false, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 设置为true时，表名为复数形式 User的表名应该是user
			TablePrefix:   "t_",  // 表名前缀 User的表名应该是t_user
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)
		Logger:                                   newLogger,
	})
	if err != nil {
		log.Fatalln("数据库连接错误:", err)
	}
	pool, err := Client.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(10)
	pool.SetConnMaxLifetime(time.Minute)
	return pool
}

func (r *generalDB) GetById(entity any, id any) any {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("异常=%v", e)
			errors.New("异常了原因")
		}
	}()
	if entityType := reflect.TypeOf(entity); entityType.Kind() != reflect.Ptr {
		errors.New("实体必须为结构体指针")
	}
	typeOf := reflect.TypeOf(id)
	var strId string
	var intId uint
	var result *gorm.DB
	if typeOf.Kind() == reflect.String {
		strId = id.(string)
		result = Client.First(entity, "id = ?", strId)
	}
	if typeOf.Kind() == reflect.Uint {
		intId = id.(uint)
		result = Client.First(entity, intId)
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return entity
}
func (r *generalDB) GetByIds(entity any, ids any) any {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("异常=%v", e)
		}
	}()
	if entityType := reflect.TypeOf(entity); entityType.Kind() != reflect.Ptr || entityType.Elem().Kind() != reflect.Array {
		errors.New("实体必须为结构体/结构体数组指针")
	}
	idsType := reflect.TypeOf(ids)
	var strIds []string
	var uintIds []uint
	var result *gorm.DB
	if idsType.Kind() == reflect.Array {
		strIds = ids.([]string)
		result = Client.Where("id in ?", strIds).Find(entity)
	}
	if idsType.Kind() == reflect.Uint {
		uintIds = ids.([]uint)
		result = Client.First(entity, uintIds)
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return entity
}
func (r *generalDB) Save(entity any) int64 {
	if exist := Client.Migrator().HasTable(entity); !exist {
		Client.AutoMigrate(entity)
	}
	result := Client.Create(entity)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return 0
	}
	return result.RowsAffected
}

func (r *generalDB) GetOne(e any, result any) error {
	tx := Client.Where(e).First(result)
	if tx.Error != nil {
		if tx.Error != gorm.ErrRecordNotFound {
			log.Printf("sql异常,原因=%v", tx.Error.Error())
			return tx.Error
		}
	}
	return nil
}

func (r *generalDB) GetList(e any, result any) error {
	if first := Client.Where(e).Find(result); first.Error != nil && first.Error != gorm.ErrRecordNotFound {
		log.Printf("sql异常,原因=%v", first.Error.Error())
		return first.Error
	}
	return nil
}

func (r *generalDB) Page(e any, page *entity.Page) *entity.Page {
	offset := (page.Cur - 1) * page.Size
	result := make([]reflect.Value, 0)
	find := Client.Where(e).Offset(offset).Limit(page.Size).Find(result)
	if find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		log.Printf("sql异常,原因=%v", find.Error.Error())
		return page
	}
	var count int64
	find.Count(&count)
	page.Data = result
	page.Total = count
	return page
}

func (r *generalDB) RemoveById(e any, id any) bool {
	Client.Delete(e, id)
	return true
}
func (r *generalDB) RemoveByIds(e any, ids ...any) bool {
	Client.Where("id = ?", ids).Delete(e)
	return true
}

type StringBuilder struct {
	sb *strings.Builder
}

func (s *StringBuilder) append(str string) *StringBuilder {
	s.sb.WriteString(str)
	return s
}
func (s *StringBuilder) toStr() string {
	return s.sb.String()
}

type GeneralQ struct {
	eq    map[string]any
	ne    map[string]any
	in    map[string]any
	notIn map[string]any
	or    map[string]any
	q     []string     // 查询字段
	asc   []string     // 升序字段
	desc  []string     // 降序
	page  *entity.Page // 分页
	sql   string       // 自定义sql
	group []string     // 分组
}
type generalDB struct {
}

func (r *generalDB) LambdaQuery() *LambdaQuery {
	return &LambdaQuery{}
}

type LambdaQuery struct {
	eq    map[string]any
	ne    map[string]any
	in    map[string]any
	notIn map[string]any
	q     []string // 查询字段
	asc   []string // 升序字段
	desc  []string // 降序
	group []string // 分组
}

func (l *LambdaQuery) Eq(field string, value any) *LambdaQuery {
	if l.eq == nil {
		l.eq = map[string]any{}
	}
	l.eq[field] = value
	return l
}
func (l *LambdaQuery) Ne(field string, value any) *LambdaQuery {
	if l.ne == nil {
		l.ne = map[string]any{}
	}
	l.ne[field] = value
	return l
}
func (l *LambdaQuery) In(field string, value any) *LambdaQuery {
	if l.in == nil {
		l.in = map[string]any{}
	}
	l.in[field] = value
	return l
}
func (l *LambdaQuery) NotIn(field string, value any) *LambdaQuery {
	if l.notIn == nil {
		l.notIn = map[string]any{}
	}
	l.notIn[field] = value
	return l
}
func (l *LambdaQuery) Select(field string) *LambdaQuery {
	l.q = append(l.q, field)
	return l
}
func (l *LambdaQuery) Asc(field string) *LambdaQuery {
	l.asc = append(l.asc, field)
	return l
}
func (l *LambdaQuery) Desc(field string) *LambdaQuery {
	l.desc = append(l.desc, field)
	return l
}
func (l *LambdaQuery) Group(field string) *LambdaQuery {
	l.group = append(l.group, field)
	return l
}
func (l *LambdaQuery) One(e any) {
	db := Client
	if _, arr := creatSql(l); len(arr) > 0 {
		db = db.Where(creatSql(l))
	}
	if !collection.MapIsEmpty(l.notIn) {
		db = db.Where(l.notIn)
	}
	if len(l.q) > 0 {
		db = db.Select(l.q)
	}
	db = db.Order(arrayConvSql(l.asc, Asc))
	db = db.Order(arrayConvSql(l.desc, Desc))
	db.Find(e)
}
func (l *LambdaQuery) List(e any) {
	db := Client
	if _, arr := creatSql(l); len(arr) > 0 {
		db = db.Where(creatSql(l))
	}
	if !collection.MapIsEmpty(l.notIn) {
		db = db.Not(l.notIn)
	}
	if len(l.q) > 0 {
		db = db.Select(l.q)
	}
	db = db.Order(arrayConvSql(l.asc, Asc))
	db = db.Order(arrayConvSql(l.desc, Desc))
	db.Find(e)
}
func (l *LambdaQuery) Page(e any, page *entity.Page) {
	db := Client
	if _, arr := creatSql(l); len(arr) > 0 {
		db = db.Where(creatSql(l))
	}
	if !collection.MapIsEmpty(l.notIn) {
		db = db.Not(l.notIn)
	}
	if len(l.q) > 0 {
		db = db.Select(l.q)
	}
	offset := (page.Cur - 1) * page.Size
	db = db.Order(arrayConvSql(l.asc, Asc))
	db = db.Order(arrayConvSql(l.desc, Desc))
	find := db.Offset(offset).Limit(page.Size).Find(e)
	var count int64
	find.Count(&count)
	page.Data = e
	page.Total = count
}
func creatSql(l *LambdaQuery) (string, []any) {
	vList := make([]any, 0)
	sb := StringBuilder{sb: &strings.Builder{}}
	vList = funcName(l.eq, &sb, vList, Equal)
	vList = funcName(l.ne, &sb, vList, NoEqual)
	vList = funcName(l.in, &sb, vList, IN)
	if len(vList) > 0 {
		return sb.toStr()[3:], vList
	}
	return "", vList
}

func funcName(t map[string]any, sb *StringBuilder, vList []any, enum string) []any {
	if !collection.MapIsEmpty(t) {
		for k, v := range t {
			sb.append("and ").append(k).append(enum)
			vList = append(vList, v)
		}
	}
	return vList
}
func arrayConvSql(array []string, order string) string {
	sb := StringBuilder{sb: &strings.Builder{}}
	if len(array) > 0 {
		for i := range array {
			sb.append(" ").append(array[i])
		}
		sb.append(" ").append(order)
	}
	return sb.toStr()
}

const (
	Equal   = " = ?"
	NoEqual = " <> ?"
	IN      = " in ?"
	NotIn   = " not in ?"
	Asc     = " asc "
	Desc    = " desc "
)
