package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/wzhyyds123/golibrary/log"
	"go-gin/Global"
	"go-gin/Router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	var err error
	dsn := Global.UserName + ":" + Global.Password + "@tcp(" + Global.Host + ")/" + Global.DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error.Println(err)
	}

	conn, err1 := redis.Dial("tcp",Global.RedisAddr,redis.DialDatabase(Global.RedisDB),redis.DialPassword(Global.RedisPassword))
	Global.RDB = conn
	if err1 != nil {
		log.Error.Println(err1)
	}

}

func main() {
	Router.InitRouter()
}
