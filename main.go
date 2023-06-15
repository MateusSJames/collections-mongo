package main

import (
	"files/service"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	service.ConnectDb()
	dates, err := service.FindClassByStudent("648a1422c875a5d4d496134d")

	if(err != nil) {
		fmt.Println(err)
	}

	fmt.Println(dates)
}