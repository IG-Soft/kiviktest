package client

import (
	"context"

	kivik "github.com/IG-Soft/kivik/v3"
	"github.com/IG-Soft/kiviktest/v3/kt"
)

func init() {
	kt.Register("CreateDB", createDB)
}

func createDB(ctx *kt.Context) {
	ctx.RunRW(func(ctx *kt.Context) {
		ctx.RunAdmin(func(ctx *kt.Context) {
			testCreateDB(ctx, ctx.Admin)
		})
		ctx.RunNoAuth(func(ctx *kt.Context) {
			testCreateDB(ctx, ctx.NoAuth)
		})
	})
}

func testCreateDB(ctx *kt.Context, client *kivik.Client) {
	ctx.Parallel()
	dbName := ctx.TestDBName()
	defer ctx.DestroyDB(dbName)
	err := client.CreateDB(context.Background(), dbName, ctx.Options("db"))
	if !ctx.IsExpectedSuccess(err) {
		return
	}
	ctx.Run("Recreate", func(ctx *kt.Context) {
		err := client.CreateDB(context.Background(), dbName, ctx.Options("db"))
		ctx.CheckError(err)
	})
}
