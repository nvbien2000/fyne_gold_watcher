package repository

import "time"

// TestRepository is a test repo
type TestRepository struct{}

// NewTestRepository returns a new repo
func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

// Migrate creates `holdings` table
func (repo *TestRepository) Migrate() error {
	return nil
}

// InsertHolding inserts new record to `holdings` tbl
func (repo *TestRepository) InsertHolding(rec Holdings) (*Holdings, error) {
	return &rec, nil
}

// AllHoldings lists all holdings from sqlite
func (repo *TestRepository) AllHoldings() ([]Holdings, error) {
	var all []Holdings
	h1 := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	h2 := Holdings{
		Amount:        2,
		PurchaseDate:  time.Now(),
		PurchasePrice: 2000,
	}
	all = append(all, h1, h2)
	return all, nil
}

// GetHoldingByID finds a record from holdings tbl by ID
func (repo *TestRepository) GetHoldingByID(id int) (*Holdings, error) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	return &h, nil
}

// UpdateHolding updates a record from holdings tbl by ID
func (repo *TestRepository) UpdateHolding(id int64, rec Holdings) error {
	return nil
}

// DeleteHolding deletes a rec from holdings tbl by ID
func (repo *TestRepository) DeleteHolding(id int64) error {
	return nil
}
