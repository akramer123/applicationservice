package main

import (
	"applicationMicroservice/applicationservice/db"
	"applicationMicroservice/applicationservice/handler"
	applicationservice "applicationMicroservice/applicationservice/proto/applicationservice"
	schema "applicationMicroservice/applicationservice/shema"
	"context"

	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

const (
	host     = "0.0.0.0"
	port     = 32768
	user     = "applications"
	password = "123456"
	dbname   = "applications"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.applicationservice"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()



	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	postgres, err := db.NewPostgres(psqlInfo)

	if err != nil {
		panic(err)
	}

	initializeDatabase(*postgres)
	// Register Handler
	applicationservice.RegisterApplicationserviceHandler(service.Server(), handler.NewApplicationService(postgres))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func initializeDatabase(postgres db.PostgresRepository) {

	a := applicationservice.Person{
		Name: "Tobi",
		Age:  "12",
		Job:  "Software Developer",
		Id:   "1",
	}
	b := applicationservice.Person{
		Name: "Salman",
		Age:  "13",
		Job:  "Software Developer",
		Id:   "2",
	}
	c := applicationservice.Person{
		Name: "Tobi2",
		Age:  "14",
		Job:  "Projektleiter",
		Id:   "3",
	}

	persons := []*applicationservice.Person{&a, &b, &c}

	for _, person := range persons {
		err := postgres.InsertApplication(context.TODO(), schema.Application{
			ID:   person.Id,
			Name: person.Name,
			Age:  person.Age,
			Job:  person.Job,
		})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Sprintf("Inserted %s", person.Name)
		}
	}
}
