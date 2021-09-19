package main

import (
	"context"

	"github.com/lemon-mint/open-backend/ent"
)

func initDB(c *ent.Client) {
	if c.Group.Query().CountX(context.Background()) != 0 {
		return
	}
	println("Creating group...")

	admin := c.Group.Create().
		SetName("admin").
		SetDescription("Administrator group").
		SaveX(context.Background())

	user := c.Group.Create().
		SetName("user").
		SetDescription("User group").
		SaveX(context.Background())
	_ = user

	if c.Resource.Query().CountX(context.Background()) != 0 {
		return
	}

	println("Creating resources...")

	c.Resource.Create().
		SetName("auth").
		SetOwner(admin).
		SetAcls([]string{
			"user:signin",
			"user:signout",
			"admin:kick",
			"admin:ban",
			"admin:unban",
			"admin:create",
			"admin:update",
			"admin:delete",
		})
}
