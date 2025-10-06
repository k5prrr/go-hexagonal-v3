// Тут константы, таймауты
package app

import "net/http"

const (
	httpAddr = ":8080"
)

type App struct {
	diContainer *diContainer
	httpServer  *http.Server
}

func New(*App) (*App, error) {}
