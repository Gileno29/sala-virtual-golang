package main

import (
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		print("erro ao carregar variaveis de ambiente")
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf")

	if err := cmd.Run(); err != nil {
		//print("Erro ao executar")
		panic("erro ao exectar o tern")
	}
}
