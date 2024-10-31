package handler

import (
	"be-golang-chapter-19/repository-api-pattern/model"
	"be-golang-chapter-19/repository-api-pattern/repository"
	"be-golang-chapter-19/repository-api-pattern/service"
	"database/sql"
	"encoding/json"
	"fmt"

	"io"
	"os"
)

func Login(db *sql.DB) {

	// input
	user := model.Customer{}
	file, err := os.Open("body.json")

	if err != nil {
		fmt.Println("Error : ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&user)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
	}

	// proses
	repo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(repo)

	customer, err := customerService.LoginService(user)

	// output
	if err != nil {
		response := model.Response{
			StatusCode: 404,
			Message:    "Account not found",
			Data:       nil,
		}
		jsonData, err := json.MarshalIndent(response, " ", " ")

		if err != nil {
			fmt.Println("err :", err)
		}

		fmt.Println(string(jsonData))
	} else {

		response := model.Response{
			StatusCode: 200,
			Message:    "login success",
			Data:       customer,
		}
		jsonData, err := json.MarshalIndent(response, " ", "")

		if err != nil {
			fmt.Println("err :", err)
		}

		fmt.Println(string(jsonData))
	}

}
