package util

import (
	"fmt"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey"
)

func TestGetConfig(t *testing.T) {
	dsn := "root:256275@tcp(localhost:3306)/management_system?charset=utf8mb4&parseTime=True&loc=Local"
	gomonkey.ApplyFunc(os.Getenv, func(name string) string {
		return "yz"
	})

	gomonkey.ApplyFunc(fmt.Sprintf, func(format string, a ...interface{}) string {
		return "../config/Config_yz.yml"
	})

	config := GetConfig()

	if config.Mysql.Dsn != dsn {
		t.Errorf("[TestGetConfig] err,config.Dsn:%+v,dsn:%+v", config.Mysql.Dsn, dsn)
	}

}
