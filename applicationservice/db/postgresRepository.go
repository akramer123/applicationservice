package db

import (
	"context"
	"database/sql"

	"applicationMicroservice/applicationservice/shema"
	_ "github.com/lib/pq"
)



type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) InsertApplication(ctx context.Context, app schema.Application) error {
	_, err := r.db.Exec("INSERT INTO applications(id, name, age, job) VALUES($1, $2, $3, $4)", app.ID, app.Name, app.Age, app.Job)
	return err
}

func (r *PostgresRepository) ListApplications(ctx context.Context, skip uint64, take uint64) ([]schema.Application, error) {
	rows, err := r.db.Query("SELECT * FROM applications ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of Meows
	applications := []schema.Application{}
	for rows.Next() {
		app := schema.Application{}
		if err = rows.Scan(&app.ID, &app.Name, &app.Age, &app.Job); err == nil {
			applications = append(applications, app)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}

func (r *PostgresRepository) ListAllApplications(ctx context.Context) ([]schema.Application, error) {
	rows, err := r.db.Query("SELECT * FROM applications ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse all rows into an array of Meows
	applications := []schema.Application{}
	for rows.Next() {
		app := schema.Application{}
		if err = rows.Scan(&app.ID, &app.Name, &app.Age, &app.Job); err == nil {
			applications = append(applications, app)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}