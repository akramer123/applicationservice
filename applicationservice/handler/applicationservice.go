package handler

import (
	"applicationMicroservice/applicationservice/db"
	"context"
	"fmt"

	"github.com/micro/go-micro/util/log"

	applicationservice "applicationMicroservice/applicationservice/proto/applicationservice"
)

type Applicationservice struct {
	db db.PostgresRepository
}

func NewApplicationService(repository *db.PostgresRepository) *Applicationservice {
	return &Applicationservice{db: *repository}
}

func (e *Applicationservice) GetAllPersons(ctx context.Context, req *applicationservice.Message, rsp *applicationservice.AllPersonResponse) error {
	log.Log("Received Applicationservice.GetAllPersons request")

	applications, err := e.db.ListAllApplications(ctx)

	if err != nil {
		fmt.Println(err)
	}
	var persons []*applicationservice.Person
	for _, application := range applications {
		pers := applicationservice.Person{
			Name: application.Name,
			Age:  application.Age,
			Job:  application.Job,
			Id:   application.ID,
		}
		persons = append(persons, &pers)
	}
	rsp.Persons = persons
	return nil
}
