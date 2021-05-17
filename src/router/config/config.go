package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConexaoBanco = ""
	Porta        = 0
)

func Carregar() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 0
	}
	ConexaoBanco = fmt.Sprintf("mongodb+srv://admin:%s@cluster0.0pizl.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

}
