package service

import (
	"database/sql"
	"errors"
	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"reflect"
	"sensitive-storage/constant"
	"sensitive-storage/module/entity"
	myReflect "sensitive-storage/util/my-reflect"
	"strings"
	"time"
)

var client *gorm.DB

var Sqlx *gorm.DB

func InitDataBase(conf *ini.File) *sql.DB {
	var err error
	dbName := conf.Section("sqlite").Key("db_name").String()
	client, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{
		SkipDefaultTransaction: false, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 设置为true时，表名为复数形式 User的表名应该是user
			TablePrefix:   "t_",  // 表名前缀 User的表名应该是t_user
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)
	})
	Sqlx = client
	if err != nil {
		log.Fatalln("数据库连接错误")
	}
	pool, err := client.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(10)
	pool.SetConnMaxLifetime(time.Minute)
	return pool
}

func (r *generalDB) GetById(entity interface{}, id interface{}) interface{} {
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
		result = client.First(entity, "id = ?", strId)
	}
	if typeOf.Kind() == reflect.Uint {
		intId = id.(uint)
		result = client.First(entity, intId)
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return entity
}
func (r *generalDB) GetByIds(entity interface{}, ids interface{}) interface{} {
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
		result = client.Where("id in ?", strIds).Find(entity)

	}
	if idsType.Kind() == reflect.Uint {
		uintIds = ids.([]uint)
		result = client.First(entity, uintIds)
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return entity
}
func (r *generalDB) Save(entity interface{}) int64 {
	if exist := client.Migrator().HasTable(entity); !exist {
		client.AutoMigrate(entity)
	}
	result := client.Create(entity)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return constant.DBSaveFail
	}
	return result.RowsAffected
}

func (r *generalDB) GetOne(e interface{}) interface{} {
	sb := &StringBuilder{sb: &strings.Builder{}}
	params := make([]interface{}, 0)
	nameAndValue := myReflect.GetNameAndValue(e)
	for k, v := range nameAndValue {
		if !myReflect.IsBlank(v) {
			sb.append("and ").append(k).append(" = ? ")
			params = append(params, v)
		}
	}
	var result entity.User
	var sql string
	if strings.HasPrefix(sb.toStr(), "and") {
		sql = sb.toStr()[3:]
	}
	first := client.Where(sql, "zhanghao3", "1234567").First(&result)
	if first.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return result
}

func (r *generalDB) GetList(e interface{}) interface{} {
	sb := &StringBuilder{sb: &strings.Builder{}}
	var params []interface{}
	nameAndValue := myReflect.GetNameAndValue(e)
	i := 0
	for k, v := range nameAndValue {
		if !myReflect.IsBlank(v) {
			sb.append("and ").append(k).append(" = ? ")
			params[i] = v
			i++
		}
	}
	var sql string
	if strings.HasPrefix(sb.toStr(), "and") {
		sql = sb.toStr()[3:]
	}
	var result []entity.User
	client.Where(sql, params).First(result)
	return result
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
	eq    map[string]interface{}
	ne    map[string]interface{}
	in    map[string]interface{}
	notIn map[string]interface{}
	or    map[string]interface{}
	q     []string     // 查询字段
	asc   []string     // 升序字段
	desc  []string     // 降序
	page  *entity.Page // 分页
	sql   string       // 自定义sql
	group []string     // 分组
}
type generalDB struct {
}
