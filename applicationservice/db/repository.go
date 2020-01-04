package db

import (
	"context"

	schema "applicationMicroservice/applicationservice/shema"
)

type Repository interface {
	Close()
	InsertApplication(ctx context.Context, meow schema.Application) error
	ListApplications(ctx context.Context, skip uint64, take uint64) ([]schema.Application, error)
	ListAllApplications(ctx context.Context) ([]schema.Application, error)
}


