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
	query := `INSERT INTO prays(language, text) VALUES($1, $2)`

	_, err := p.db.Exec(query, pray.Language, pray.Definition)
	if err != nil {
		return fmt.Errorf("failed to create pray: %v", err)
	}

	return nil
}

func (p *prayRepo) Read(prayId int) (*model.Pray, error) {
	query := `SELECT p.language, p.definition from prays p WHERE id = $1`

	rows, err := p.db.Query(query, prayId)
	var pray model.Pray
	if err = rows.Scan(&prayId, &pray.Language, &pray.Definition); err != nil {
		return nil, fmt.Errorf("can't read pray %v", err)
	}
	return &pray, nil
}

func (p *prayRepo) Put(prayId int, language, definition string) error {
	query := `UPDATE prays SET language = $2 AND definition = $3 WHERE prayId = $1`

	_, err := p.db.Exec(query, prayId, language, definition)
	if err != nil {
		return fmt.Errorf("failed to update pray %v", err)
	}
	return nil
}

func (p *prayRepo) Delete(prayId int) error {
	query := `DELETE FROM prays WHERE prayId = $1`

	_, err := p.db.Exec(query, prayId)
	if err != nil {
		return err
	}
	return nil
}
