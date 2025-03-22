package store

import (
	"database/sql"
	"fmt"
)

type Account struct {
	ID       int
	Address  string
	Password string
}

func (s *Store) AddAccount(address string, password string) (*Account, error) {
	res, err := s.db.Exec(`
		INSERT INTO accounts (address, password)
		VALUES (?, ?)
	`, address, password)
	if err != nil {
		return nil, fmt.Errorf("AddAccount: db.Exec: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("AddAccount: LastInsertId: %w", err)
	}

	return &Account{
		ID:       int(id),
		Address:  address,
		Password: password,
	}, nil
}

func (s *Store) GetAllAccounts() ([]Account, error) {
	rows, err := s.db.Query(`SELECT id, address, password FROM accounts`)
	if err != nil {
		return nil, fmt.Errorf("GetAllAccounts: db.Query: %w", err)
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var a Account
		if err := rows.Scan(&a.ID, &a.Address, &a.Password); err != nil {
			return nil, fmt.Errorf("GetAllAccounts: rows.Scan: %w", err)
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func (s *Store) GetAccountByAddress(address string) (*Account, error) {
	row := s.db.QueryRow(`SELECT id, address, password FROM accounts WHERE address = ?`, address)

	var a Account
	if err := row.Scan(&a.ID, &a.Address, &a.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetAccountByAddress: no account found for address %s", address)
		}
		return nil, fmt.Errorf("GetAccountByAddress: row.Scan: %w", err)
	}
	return &a, nil
}

func (s *Store) DeleteAccount(id int) error {
	_, err := s.db.Exec(`DELETE FROM accounts WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("DeleteAccount: db.Exec: %w", err)
	}
	return nil
}
