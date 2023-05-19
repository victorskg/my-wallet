package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Entity interface {
	FromRow(rows *sql.Rows) (any, error)
}

type Repository[T Entity] struct {
	conn      Connection
	db        *sql.DB
	tableName string
	schema    string
	entity    T
}

func NewRepository[T Entity](
	databaseURL string,
	databasePort string,
	databaseDriver string,
	databaseName string,
	databaseUser string,
	databasePassword string,
	tableName string,
	schema string,
	entity T) *Repository[T] {
	conn := NewConnection(databaseURL, databasePort, databaseDriver,
		databaseName, databaseUser, databasePassword)

	return &Repository[T]{tableName: tableName, schema: schema, entity: entity, conn: conn}
}

func (r *Repository[T]) Connect() error {
	db, err := r.conn.Open()
	if err != nil {
		fmt.Printf("Connection error: %s", err.Error())
		return err
	}

	r.db = db
	return err
}

func (r *Repository[T]) SelectOne(where string) (*T, error) {
	results, err := r.Select(where)
	if err != nil {
		fmt.Printf("Select one error: %s", err.Error())
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	if len(results) != 1 {
		return nil, fmt.Errorf("Expected 1 result, got %d.", len(results))
	}

	return results[0], nil
}

func (r *Repository[T]) Select(where string) ([]*T, error) {
	queryString := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s", r.schema, r.tableName, where)
	rows, err := r.db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(rows)

	results := make([]*T, 0)
	for rows.Next() {
		result, err := r.entity.FromRow(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, result.(*T))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
