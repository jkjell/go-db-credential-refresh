package static

import (
	"context"

	"github.com/jkjell/go-db-credential-refresh/driver"
	"github.com/jkjell/go-db-credential-refresh/store"
)

type Static struct {
	credentials *store.Credential
}

func NewStaticStore(username, password string) *Static {
	return &Static{
		credentials: &store.Credential{
			Username: username,
			Password: password,
		},
	}
}

func (s Static) Get(ctx context.Context) (driver.Credentials, error) {
	return s.credentials, nil
}

func (s Static) Refresh(ctx context.Context) (driver.Credentials, error) {
	return s.credentials, nil
}
