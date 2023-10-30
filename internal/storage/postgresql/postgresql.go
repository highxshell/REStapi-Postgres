package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib" // init postgresql driver
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgresql.NewStorage"

	db, err := sql.Open("pgx", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	//stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS url( id INTEGER PRIMARY KEY, alias TEXT NOT NULL UNIQUE, url TEXT NOT NULL); CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);")
	//if err != nil {
	//	return nil, fmt.Errorf("%s: %w", op, err)
	//}

	//_, err = stmt.Exec()
	//if err != nil {
	//	return nil, fmt.Errorf("%s: %w", op, err)
	//}
	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgresql.SaveUrl"
	//stmt, err := s.db.Prepare(`INSERT INTO url(url, alias) VALUES ($1, $2) RETURNING id`)
	//if err != nil {
	//	return 0, fmt.Errorf("%s: prepare statement: %w", op, err)
	//}
	// Выполняем запрос
	//res, err := stmt.Exec(urlToSave, alias)
	//if err != nil {
	//	return 0, fmt.Errorf("%s: execute statement: %w", op, err)
	//}
	//_ = res
	// Получаем ID созданной записи
	id := 0
	err := s.db.QueryRow(`INSERT INTO url(url, alias) VALUES ($1, $2) RETURNING id`, urlToSave, alias).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: query row: %w", op, err)
	}
	// Возвращаем ID
	return int64(id), nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.postgresql.GetUrl"
	var resURL string
	err := s.db.QueryRow("SELECT url FROM url WHERE alias = $1", alias).Scan(&resURL)
	if err != nil {
		return "", fmt.Errorf("%s: query row: %w", op, err)
	}
	return resURL, nil
}

func (s *Storage) DeleteURL(alias string) error {
	const op = "storage.postgresql.DeleteUrl"
	err := s.db.QueryRow("DELETE FROM url WHERE alias = $1", alias)
	if err.Err() != nil {
		return fmt.Errorf("%s: query row: %w", op, err.Err())
	}
	return nil
}
