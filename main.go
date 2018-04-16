package main

import (
	"os"

	"groupup/src/system/serv"

	"github.com/joho/godotenv"
)

var port string

func init() {
	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	} else {
		panic("Missing PORT in config.ini")
	}
}

func main() {

	s := serv.NewServer()
	s.Init(port)
	s.Start()
}
