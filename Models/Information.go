package Models

import (
	"fmt"
	"github.com/wzhyyds123/golibrary/log"
	"go-gin/Global"
	"time"
)

type Information_C_F struct {
	User_id      int64     `json:"user_id"`
	Active_name  string    `json:"active_name"`
	Active_place string    `json:"active_place"`
	Time_begin   time.Time `json:"time_begin"`
}
func (t *Information_C_F) TableName() string {
	return "information"
}
func WzhFind(active_name string) (res Information_C_F, err error) {
	res = Information_C_F{}
	result := Global.DB.Where("active_name = ?", active_name).Find(&res)
	fmt.Println(res)
	if result.Error != nil {
		log.Error.Println(result.Error)
		return res, result.Error
	} else {
		return res, result.Error
	}
}
