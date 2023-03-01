package models

import (
	"github.com/basegin/db"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Id    uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Key   string `gorm:"column:key"`
	Value string `gorm:"column:value"`
	Desc  string `gorm:"column:desc"`
}

func (Config) TableName() string {
	return "config"
}

func GetConfigValue(c *gin.Context, key string) (value Config, err error) {
	var configs []Config
	err = db.MysqlClient.Ctx(c).Where("`key` = ?", key).Find(&configs).Error
	if err != nil {
		return value, err
	}

	if len(configs) == 0 {
		return value, err
	}
	return configs[0], nil
}
