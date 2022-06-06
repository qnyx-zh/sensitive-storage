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
	"sensitive-storage/module/entity"
	"time"
)

var client *gorm.DB

func InitDataBase(conf *ini.File) *sql.DB {
	var err error
	dbName := conf.Section("sqlite").Key("db_name").String()
	client, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{
		SkipDefaultTransaction: false, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 设置为true时，表名为复数形式 User的表名应该是user
			TablePrefix:   "t_",  //表名前缀 User的表名应该是t_user
		},
		DisableForeignKeyConstraintWhenMigrating: true, //设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)
	})
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
		return 0
	}
	return result.RowsAffected
}

type GeneralQ struct {
	eq    map[string]interface{}
	ne    map[string]interface{}
	in    map[string]interface{}
	notIn map[string]interface{}
	or    map[string]interface{}
	q     []string     //查询字段
	asc   []string     //升序字段
	desc  []string     //降序
	page  *entity.Page //分页
	sql   string       //自定义sql
	group []string     //分组
}
type generalDB struct {
}
