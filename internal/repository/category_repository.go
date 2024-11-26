package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/tiborm/barefoot-bear/internal/model"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) CreateTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS categories (
        id VARCHAR(255) PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        url VARCHAR(255) NOT NULL,
        im VARCHAR(255),
        subs JSONB
    )`
	_, err := r.DB.Exec(query)
	return err
}

func (r *CategoryRepository) Insert(category model.Category) error {
	subs, err := json.Marshal(category.Subs)
	if err != nil {
		return err
	}

	query := `
    INSERT INTO categories (id, name, url, im, subs)
    VALUES ($1, $2, $3, $4, $5)
    ON CONFLICT (id) DO UPDATE
    SET name = EXCLUDED.name,
        url = EXCLUDED.url,
        im = EXCLUDED.im,
        subs = EXCLUDED.subs`
	_, err = r.DB.Exec(query, category.ID, category.Name, category.URL, category.IM, subs)
	return err
}

func (r *CategoryRepository) GetByID(id string) (*model.Category, error) {
	query := `SELECT id, name, url, im, subs FROM categories WHERE id = $1`
	row := r.DB.QueryRow(query, id)

	var category model.Category
	var subs []byte
	err := row.Scan(&category.ID, &category.Name, &category.URL, &category.IM, &subs)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(subs, &category.Subs)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	query := `SELECT id, name, url, im, subs FROM categories`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		var subs []byte
		err := rows.Scan(&category.ID, &category.Name, &category.URL, &category.IM, &subs)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(subs, &category.Subs)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) DeleteByID(id string) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
