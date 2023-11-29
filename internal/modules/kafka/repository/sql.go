package kafkarepo

import (
	"context"
	"fmt"
	"kafkatool/internal/logger"
	kafkamodel "kafkatool/internal/modules/kafka/models"

	"github.com/jmoiron/sqlx"
)

type SqlRepo struct {
	sqlClient *sqlx.DB
	log       logger.Logger
}

func NewSqlRepo(client *sqlx.DB, log logger.Logger) *SqlRepo {
	return &SqlRepo{sqlClient: client, log: log}
}

func (s *SqlRepo) ListRequest(ctx context.Context) ([]kafkamodel.Request, error) {
	query := `SELECT title, topic, quantity, type, message FROM requests`

	rows, err := s.sqlClient.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("[ListRequest] query requests failed %+v", err)
	}
	var request kafkamodel.Request
	requests := []kafkamodel.Request{}
	for rows.Next() {
		if err := rows.StructScan(&request); err != nil {
			return nil, fmt.Errorf("[ListRequest] parsed request failed %+v", err)
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (s *SqlRepo) SaveRequest(ctx context.Context, request kafkamodel.Request) error {
	query := `INSERT INTO requests
					(title, topic, quantity, type, message)
				VALUES
					(:title, :topic, :quantity, :type, :message)
			`
	_, err := s.sqlClient.NamedExec(query, &request)
	if err != nil {
		return fmt.Errorf("[SaveRequest] failed to save request %+v", err)
	}
	return nil
}
