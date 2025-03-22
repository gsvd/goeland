package main

import (
	"context"
	"sync"

	"github.com/gsvd/goeland/internal/store"
	"github.com/xmppo/go-xmpp"
)

type App struct {
	ctx         context.Context
	store       *store.Store
	connections map[string]*xmpp.Client
	messages    map[string][]string
	mu          sync.Mutex
}

func NewApp(store *store.Store) *App {
	return &App{
		store:       store,
		connections: make(map[string]*xmpp.Client),
		messages:    make(map[string][]string),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
