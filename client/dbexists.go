package client

import (
	"context"

	kivik "github.com/IG-Soft/kivik/v3"
	"github.com/IG-Soft/kiviktest/v3/kt"
)

func init() {
	kt.Register("DBExists", dbExists)
}

func dbExists(ctx *kt.Context) {
	ctx.RunAdmin(func(ctx *kt.Context) {
		for _, dbName := range ctx.MustStringSlice("databases") {
			checkDBExists(ctx, ctx.Admin, dbName)
		}
	})
	ctx.RunNoAuth(func(ctx *kt.Context) {
		for _, dbName := range ctx.MustStringSlice("databases") {
			checkDBExists(ctx, ctx.NoAuth, dbName)
		}
	})
	ctx.RunRW(func(ctx *kt.Context) {
		dbName := ctx.TestDB()
		defer ctx.DestroyDB(dbName)
		ctx.Run("group", func(ctx *kt.Context) {
			ctx.RunAdmin(func(ctx *kt.Context) {
				checkDBExists(ctx, ctx.Admin, dbName)
			})
			ctx.RunNoAuth(func(ctx *kt.Context) {
				checkDBExists(ctx, ctx.NoAuth, dbName)
			})
		})
	})
}

func checkDBExists(ctx *kt.Context, client *kivik.Client, dbName string) {
	ctx.Run(dbName, func(ctx *kt.Context) {
		ctx.Parallel()
		exists, err := client.DBExists(context.Background(), dbName)
		if !ctx.IsExpectedSuccess(err) {
			return
		}
		if ctx.MustBool("exists") != exists {
			ctx.Errorf("Expected: %t, actual: %t", ctx.Bool("exists"), exists)
		}
	})
}
