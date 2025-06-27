package repository

import (
	"database/sql"
	"fmt"
	"time"
)

// SQLiteRepository represents a type for repository that connects to sqlite database
type SQLiteRepository struct {
	Conn *sql.DB
}

// NewSQLiteRepository returns a new SQLiteRepository with database connection
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

// Migrate creates `holdings` table
func (repo *SQLiteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS holdings(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL NOT NULL,
		purchase_date INTEGER NOT NULL,
		purchase_price INTEGER NOT NULL)
	`

	_, err := repo.Conn.Exec(query)
	return err
}

// InsertHolding inserts new record to `holdings` tbl
func (repo *SQLiteRepository) InsertHolding(rec Holdings) (*Holdings, error) {
	stmt := `
	INSERT INTO holdings (amount, purchase_date, purchase_price)
	VALUES (?, ?, ?)
	`

	res, err := repo.Conn.Exec(stmt, rec.Amount, rec.PurchaseDate.Unix(), rec.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	rec.ID = id
	return &rec, nil
}

// AllHoldings lists all holdings from order by purchase_date
func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := `
	SELECT id, amount, purchase_date, purchase_price
	FROM holdings
	ORDER BY purchase_date
	`

	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Holdings
	for rows.Next() {

		var h Holdings
		var unixTime int64
		err := rows.Scan(
			&h.ID,
			&h.Amount,
			&unixTime,
			&h.PurchasePrice,
		)
		if err != nil {
			return nil, err
		}
		h.PurchaseDate = time.Unix(unixTime, 0)
		fmt.Println("h: ", h)
		all = append(all, h)
	}
	fmt.Println("all: ", all)

	return all, nil
}

// GetHoldingByID finds a record from holdings tbl by ID
func (repo *SQLiteRepository) GetHoldingByID(id int) (*Holdings, error) {
	stmt := `
	SELECT id, amount, purchase_date, purchase_price
	FROM holdings
	WHERE id = ?
	`

	row := repo.Conn.QueryRow(stmt, id)

	var h Holdings
	var unixTime int64
	err := row.Scan(
		&h.ID,
		&h.Amount,
		&unixTime,
		&h.PurchasePrice,
	)
	if err != nil {
		return nil, err
	}
	h.PurchaseDate = time.Unix(unixTime, 0)

	return &h, nil
}

// UpdateHolding updates a record from holdings tbl by ID
func (repo *SQLiteRepository) UpdateHolding(id int64, rec Holdings) error {
	stmt := `
	UPDATE holdings
	SET amount = ?, purchase_date = ?, purchase_price = ?
	WHERE id = ?
	`

	res, err := repo.Conn.Exec(stmt, rec.Amount, rec.PurchaseDate.Unix(), rec.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

// DeleteHolding deletes a rec from holdings tbl by ID
func (repo *SQLiteRepository) DeleteHolding(id int64) error {
	stmt := `DELETE FROM holdings WHERE id = ?`

	res, err := repo.Conn.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errDeleteFailed
	}

	return nil
}
