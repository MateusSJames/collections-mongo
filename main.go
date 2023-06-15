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

	dates2, err2 := service.ConnectWithOnlyDB("648a1422c875a5d4d496134d")

	if(err != nil) {
		fmt.Println(err)
	}

	if(err2 != nil) {
		fmt.Println(err2)
	}
	fmt.Println("----- Conexao com 1 connection string -----")
	fmt.Println(dates)
	fmt.Println("-------------------------------------------")
	fmt.Println("----- Conexao com 2 connection strings -----")
	fmt.Println(dates2)
	fmt.Println("-------------------------------------------")
}