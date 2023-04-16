package infrastructures

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mochafiqri/simple-crud/commons/entities"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = ""
	dbName   = "medioker_crud"
	addRedis = "localhost:6379"
)

func InitMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(entities.Content{})
	return db, err
}

func InitPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(entities.Content{})
	return db, err
}

func InitRedis() (*redis.Client, error) {
	var opt = redis.Options{
		Addr:     addRedis,
		Password: "",
		DB:       0,
	}
	var ctx = context.Background()
	rdc := redis.NewClient(&opt)
	err := rdc.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return rdc, nil
}
