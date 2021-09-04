package user

import (
	"context"

	"github.com/lemon-mint/open-backend/ent"
)

func CreateUser(c *ent.Client, ctx context.Context, Username, Password []byte) error {
	return nil
}
