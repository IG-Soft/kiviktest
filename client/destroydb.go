package client

import (
	"context"

	kivik "github.com/IG-Soft/kivik/v3"
	"github.com/IG-Soft/kiviktest/v3/kt"
)

func init() {
	kt.Register("DestroyDB", destroyDB)
}

func destroyDB(ctx *kt.Context) {
	// All DestroyDB tests are RW by nature.
	ctx.RunRW(func(ctx *kt.Context) {
		ctx.RunAdmin(func(ctx *kt.Context) {
			ctx.Parallel()
			testDestroy(ctx, ctx.Admin)
		})
		ctx.RunNoAuth(func(ctx *kt.Context) {
			ctx.Parallel()
			testDestroy(ctx, ctx.NoAuth)
		})
	})
}

func testDestroy(ctx *kt.Context, client *kivik.Client) {
	ctx.Run("ExistingDB", func(ctx *kt.Context) {
		ctx.Parallel()
		dbName := ctx.TestDB()
		defer ctx.DestroyDB(dbName)
		ctx.CheckError(client.DestroyDB(context.Background(), dbName, ctx.Options("db")))
	})
	ctx.Run("NonExistantDB", func(ctx *kt.Context) {
		ctx.Parallel()
		ctx.CheckError(client.DestroyDB(context.Background(), ctx.TestDBName(), ctx.Options("db")))
	})
}
