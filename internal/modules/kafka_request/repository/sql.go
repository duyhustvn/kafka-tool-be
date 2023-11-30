package kafkareqrepo

import (
	"context"
	"fmt"
	"kafkatool/internal/logger"
	kafkareqmodel "kafkatool/internal/modules/kafka_request/models"

	"github.com/jmoiron/sqlx"
)

type SqlRepo struct {
	sqlClient *sqlx.DB
	log       logger.Logger
}

func NewSqlRepo(client *sqlx.DB, log logger.Logger) *SqlRepo {
	return &SqlRepo{sqlClient: client, log: log}
}

func (s *SqlRepo) ListRequest(ctx context.Context) ([]kafkareqmodel.Request, error) {
	query := `SELECT id, title, topic, quantity, type, message FROM requests`

	rows, err := s.sqlClient.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("[ListRequest] query requests failed %+v", err)
	}
	var request kafkareqmodel.Request
	requests := []kafkareqmodel.Request{}
	for rows.Next() {
		if err := rows.StructScan(&request); err != nil {
			return nil, fmt.Errorf("[ListRequest] parsed request failed %+v", err)
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (s *SqlRepo) CreateRequest(ctx context.Context, request kafkareqmodel.Request) error {
	query := `INSERT INTO requests
					(title, topic, quantity, type, message)
				VALUES
					(:title, :topic, :quantity, :type, :message)
			`
	_, err := s.sqlClient.NamedExec(query, &request)
	if err != nil {
		return fmt.Errorf("[UpdateRequest] failed to save request %+v", err)
	}
	return nil
}

func (s *SqlRepo) UpdateRequest(ctx context.Context, requestID int, request kafkareqmodel.Request) error {
	query := `UPDATE requests
				SET
					title=:title, topic=:topic, quantity=:quantity, type=:type, message=:message
				WHERE
					id = :id
			`
	arg := map[string]interface{}{
		"id":       requestID,
		"title":    request.Title,
		"topic":    request.Topic,
		"quantity": request.Quantity,
		"type":     request.Type,
		"message":  request.Message,
	}
	_, err := s.sqlClient.NamedExec(query, arg)
	if err != nil {
		return fmt.Errorf("[UpdateRequest] failed to save request %+v", err)
	}
	return nil
}
