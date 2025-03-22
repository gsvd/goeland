package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gsvd/goeland/internal/store"
	"github.com/gsvd/goeland/internal/validation"
	"github.com/gsvd/goeland/pkg/errorsx"
	"github.com/xmppo/go-xmpp"
	"modernc.org/sqlite"
	lib "modernc.org/sqlite/lib"
)

func (a *App) AddAccount(address string, password string) (*store.Account, error) {
	if err := validation.ValidateXMPPAddress(address); err != nil {
		if mErr, ok := err.(interface{ Marshal() error }); ok {
			return nil, mErr.Marshal()
		}
		return nil, err
	}

	if password == "" {
		return nil, errorsx.NewAPIError(errorsx.ErrCodePasswordRequired)
	}

	err := a.OpenAccount(store.Account{
		Address:  address,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	account, err := a.store.AddAccount(address, password)
	if err != nil {
		var sqliteErr *sqlite.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code() == lib.SQLITE_CONSTRAINT_UNIQUE {
			return nil, errorsx.NewAPIError(errorsx.ErrCodeAccountExists)
		}
		return nil, errorsx.NewAPIError(errorsx.ErrCodeUnknown)
	}

	return account, nil
}

func (a *App) GetAllAccounts() ([]store.Account, error) {
	accounts, err := a.store.GetAllAccounts()
	if err != nil {
		return nil, errorsx.NewAPIError(errorsx.ErrCodeUnknown)
	}
	return accounts, nil
}

func (a *App) OpenAllAccounts() error {
	accounts, err := a.store.GetAllAccounts()
	if err != nil {
		return errorsx.NewAPIError(errorsx.ErrCodeUnknown)
	}

	for _, account := range accounts {
		if err := a.OpenAccount(account); err != nil {
			return errorsx.NewAPIError(errorsx.ErrCodeUnknown)
		}
	}

	return nil
}

func (a *App) OpenAccount(account store.Account) error {
	a.mu.Lock()
	if _, exists := a.connections[account.Address]; exists {
		a.mu.Unlock()
		return nil
	}
	a.mu.Unlock()

	host := strings.Split(account.Address, "@")[1]

	options := xmpp.Options{
		Host:          host,
		User:          account.Address,
		Password:      account.Password,
		Debug:         true,
		Session:       false,
		Status:        "xa",
		StatusMessage: "Available",
		NoTLS:         true,
		StartTLS:      true,
	}

	talk, err := options.NewClient()
	if err != nil {
		return errorsx.NewAPIError(errorsx.ErrCodeAuthenticationFailed)
	}

	a.mu.Lock()
	a.connections[account.Address] = talk
	a.mu.Unlock()

	go func() {
		for {
			chat, err := talk.Recv()
			if err != nil {
				fmt.Printf("OpenAccount: Recv: %v\n", err)
			}

			var msg string
			switch v := chat.(type) {
			case xmpp.Chat:
				msg = fmt.Sprintf("%s: %s", v.Remote, v.Text)
			case xmpp.Presence:
				msg = fmt.Sprintf("%s: %s", v.From, v.Show)
			}

			a.mu.Lock()
			a.messages[account.Address] = append(a.messages[account.Address], msg)
			a.mu.Unlock()
		}
	}()

	return nil
}
