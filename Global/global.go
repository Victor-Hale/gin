package Global

import (
	"github.com/garyburd/redigo/redis"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	RDB           redis.Conn

	Host         = "139.196.106.241:3306"
	DatabaseName = "third"
	UserName     = "third"
	Password     = "123456"

	RedisAddr = "127.0.0.1:6379"
	RedisPassword = ""
	RedisDB = 0
)
