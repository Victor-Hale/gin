package Global

import (
	"github.com/garyburd/redigo/redis"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	RDB           redis.Conn
)
