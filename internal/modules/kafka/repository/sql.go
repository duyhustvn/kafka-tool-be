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
