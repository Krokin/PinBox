package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Krokin/PinBox/pkg/models"
)

type PinModel struct {
	DB *sql.DB
}

func (p *PinModel) Insert(title, content, expires string) (int, error) {
	fmt.Print(title, content, expires)
	stmt := `INSERT INTO pins (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := p.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (p *PinModel) Get(id int) (*models.Pin, error) {
	stmt := `SELECT id, title, content, created, expires FROM pins
    WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := p.DB.QueryRow(stmt, id)
	pin := &models.Pin{}
	err := row.Scan(&pin.ID, &pin.Title, &pin.Content, &pin.Created, &pin.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return pin, nil
}

func (p *PinModel) Latest() ([]*models.Pin, error) {
	stmt := `SELECT id, title, content, created, expires FROM pins
    WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := p.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var pins []*models.Pin
	for rows.Next() {
		pin := &models.Pin{}
		err := rows.Scan(&pin.ID, &pin.Title, &pin.Content, &pin.Created, &pin.Expires)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, models.ErrNoRecord
			} else {
				return nil, err
			}
		}
		pins = append(pins, pin)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pins, nil
}