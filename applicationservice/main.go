package main

import (
	"applicationMicroservice/applicationservice/handler"
	applicationservice "applicationMicroservice/applicationservice/proto/applicationservice"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.applicationservice"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	applicationservice.RegisterApplicationserviceHandler(service.Server(), new(handler.Applicationservice))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.applicationservice", service.Server(), new(subscriber.Applicationservice))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.applicationservice", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
