package main

import (
	"context"
	"os"
	"os/user"
	"testing"

	// drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"buf.build/go/protoyaml"
	"github.com/RyoJerryYu/schemax/proto/schemax/schema"
	"github.com/stretchr/testify/require"
	"github.com/xo/dburl"
	"github.com/xo/dburl/passfile"
	xocmd "github.com/xo/xo/cmd"
	"github.com/xo/xo/loader"
	xotype "github.com/xo/xo/types"
)

// open opens a connection to the database, returning a context for use in
// template generation.
func open(ctx context.Context, urlstr, schema string) (context.Context, error) {
	v, err := user.Current()
	if err != nil {
		return nil, err
	}
	// parse dsn
	u, err := dburl.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	// open database
	db, err := passfile.OpenURL(u, v.HomeDir, "xopass")
	if err != nil {
		return nil, err
	}
	// add driver to context
	ctx = context.WithValue(ctx, xotype.DriverKey, u.Driver)
	// add db to context
	ctx = context.WithValue(ctx, xotype.DbKey, db)
	// determine schema
	if schema == "" {
		if schema, err = loader.Schema(ctx); err != nil {
			return nil, err
		}
	}
	// add schema to context
	ctx = context.WithValue(ctx, xotype.SchemaKey, schema)
	return ctx, nil
}
func TestLoad(t *testing.T) {
	ctx := context.Background()
	var err error

	urlstr := "mysql://schemaxuser:123456@127.0.0.1:13306/schemax?loc=Local&parseTime=true&charset=utf8mb4"
	ctx, err = open(ctx, urlstr, "")
	require.NoError(t, err)
	set := xotype.Set{}
	err = xocmd.LoadSchema(ctx, &set, xocmd.NewArgs(""))
	require.NoError(t, err)

	protoSet := schema.SetFromXO(&set)
	require.NotNil(t, protoSet)

	protoSetRaw, err := protoyaml.Marshal(protoSet)
	require.NoError(t, err)

	err = os.WriteFile("schema.yaml", protoSetRaw, os.ModePerm)
	require.NoError(t, err)
}
