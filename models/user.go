package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	// ID           uint `gorm:"primaryKey"`
	gorm.Model
	Name  string
	Email string `gorm:"column:email"` // 将列名设为 `email`
	Age   uint8
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `users`
func (User) TableName() string {
	return "users"
}

func (user *User) connect() (db *gorm.DB, err error) {
	dsn := "root:NvGHHsQvo3!90YS@@tcp(10.38.2.74:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if nil == err {
		// GORM 使用 database/sql 来维护连接池
		sqlDB, _ := db.DB()
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime 设置了连接可复用的最大时间
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
	return
}

func (user *User) Add(name string, email string, age uint8, memeber string) int64 {
	db, err := user.connect()
	if nil != err {
		return 0
	}

	// now := time.Now()
	u := User{
		Name:  name,
		Email: email,
		Age:   age,
	}
	result := db.Create(&u) // 通过数据的指针来创建
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return result.RowsAffected
}

func (user *User) info(id uint) *User {
	db, err := user.connect()
	if nil != err {
		return nil
	}

	// 获取第一条记录（主键升序）
	// db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	// 获取一条记录，没有指定排序字段
	// db.Take(&user)
	// SELECT * FROM users LIMIT 1;
	// 获取最后一条记录（主键降序）
	// db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	db.Where("id = ?", id).First(&user)
	return user
}
