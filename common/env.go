package common

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

//第一个字母必须大写才能导出
func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("loading env from .env success")
	}

}
