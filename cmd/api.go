package cmd

import (
	"fmt"
	"github.com/hasanbakirci/mongodb-kafka-connector/internal/mongokafka"
	"github.com/hasanbakirci/mongodb-kafka-connector/pkg/mongoclient"
	"github.com/labstack/echo/v4"
)

func StartApi() {
	instance := echo.New()

	db, err := mongoclient.ConnectDb()
	if err != nil {
		fmt.Println("Db connection error")
	}

	repository := mongokafka.NewKafkalogRepository(db)
	service := mongokafka.NewService(repository)
	controller := mongokafka.NewController(service)
	mongokafka.RegisterHandler(instance, controller)

	fmt.Println("Api starting")
	if err := instance.Start(":9494"); err != nil {
		fmt.Println("Api fatal error")
		fmt.Println(err)
	}
}
