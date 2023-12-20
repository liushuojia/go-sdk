package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

var Config = struct {
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

func main() {
	fmt.Println(os.Getwd())
	err := configor.Load(&Config, "./test.yaml")
	fmt.Println(err)
	fmt.Println(Config)

	err1 := configor.Load(&Config, "./test.json")
	fmt.Println(err1)
	fmt.Println(Config)
}
