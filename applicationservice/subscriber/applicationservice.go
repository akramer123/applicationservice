package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	applicationservice "applicationMicroservice/applicationservice/proto/applicationservice"
)

type Applicationservice struct{}

func (e *Applicationservice) Handle(ctx context.Context, msg *applicationservice.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *applicationservice.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
