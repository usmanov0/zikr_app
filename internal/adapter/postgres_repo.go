package adapter

import (
	"fmt"
	"github.com/jackc/pgx"
	"zikar-app/internal/model"
)

type prayRepo struct {
	db *pgx.Conn
}

func NewPrayRepo(db *pgx.Conn) model.PrayRepository {
	return &prayRepo{db: db}
}

func (p *prayRepo) Save(pray *model.Pray) error {
	query := `INSERT INTO prays(language, text, count_pray) VALUES($1, $2, $3)`

	_, err := p.db.Exec(query, pray.Language, pray.Definition, pray.CountPray)
	if err != nil {
		return fmt.Errorf("failed to create pray: %w", err)
	}

	return nil
}

func (p *prayRepo) UpdaterCount() error {
	query := `UPDATE prays SET count_pray = count_pray + 1`

	_, err := p.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to update count: %w", err)
	}

	return nil
}
