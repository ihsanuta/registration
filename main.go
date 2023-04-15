package main

import (
	"registration/app/handler"
	"registration/app/repository"
	"registration/app/usecase"
	"registration/module/mysql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// Business Layer
	repo *repository.Repository
	uc   *usecase.Usecase

	h handler.Handler
)

func main() {
	// konek to mysql
	db := mysql.GetMysqlConnection()

	// Business layer Initialization
	repo = repository.Init(
		db,
	)
	uc = usecase.Init(repo)
	handler.Init(uc)
}
