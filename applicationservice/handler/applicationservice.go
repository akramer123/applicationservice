package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	applicationservice "applicationMicroservice/applicationservice/proto/applicationservice"
)

type Applicationservice struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Applicationservice) Call(ctx context.Context, req *applicationservice.Request, rsp *applicationservice.Response) error {
	log.Log("Received Applicationservice.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Applicationservice) Stream(ctx context.Context, req *applicationservice.StreamingRequest, stream applicationservice.Applicationservice_StreamStream) error {
	log.Logf("Received Applicationservice.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&applicationservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Applicationservice) PingPong(ctx context.Context, stream applicationservice.Applicationservice_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&applicationservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func (e *Applicationservice) GetAllPersons(ctx context.Context, req *applicationservice.Message, rsp *applicationservice.AllPersonResponse) error {
	log.Log("Received Applicationservice.GetAllPersons request")
	a := applicationservice.Person{
		Name: "Tobi",
		Age:  "12",
	}
	b := applicationservice.Person{
		Name: "Salman",
		Age:  "13",
	}
	c := applicationservice.Person{
		Name: "Tobi2",
		Age:  "14",
	}
	persons := []*applicationservice.Person{&a, &b, &c}
	rsp.Persons = persons
	return nil
}
