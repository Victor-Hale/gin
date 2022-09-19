package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/wzhyyds123/golibrary/log"
	"go-gin/Global"
	"go-gin/Router"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func init() {
	var err error
	cfg, err := ini.Load("Config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	dsn := cfg.Section("mysql").Key("UserName").String() + ":" + cfg.Section("mysql").Key("Password").String() + "@tcp(" +  cfg.Section("mysql").Key("Host").String() + ")/" + cfg.Section("mysql").Key("DatabaseName").String() + "?charset=utf8mb4&parseTime=True&loc=Local"
	Global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error.Println(err)
	}

	conn, err1 := redis.Dial("tcp",cfg.Section("redis").Key("RedisAddr").String(),redis.DialDatabase(cfg.Section("redis").Key("RedisDB").MustInt()),redis.DialPassword(cfg.Section("redis").Key("RedisPassword").String()))
	Global.RDB = conn
	if err1 != nil {
		log.Error.Println(err1)
	}

}

func main() {
	Router.InitRouter()
}
