package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/meguriri/OnlineAssessmentSystem/constant"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Mysql struct {
		Dsn string `yaml:"dsn"`
	}
}

func GetEnvironment(name string) string {
	return os.Getenv(name)
}

func GetConfig() Config {
	environment := GetEnvironment(constant.EnvironmentVariableIdentification)
	data, err := ioutil.ReadFile(fmt.Sprintf(constant.ProfilePath, environment))
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}

	return config
}
