package schemax

import (
	"context"
	"os/user"

	protoschema "github.com/RyoJerryYu/schemax/proto/schemax/schema"
	"github.com/golang/glog"
	"github.com/xo/dburl"
	"github.com/xo/dburl/passfile"
	xocmd "github.com/xo/xo/cmd"
	"github.com/xo/xo/loader"
	xotype "github.com/xo/xo/types"
)

func LoadSchema(ctx context.Context, dsn string, schema string) (*protoschema.Schema, error) {
	ctx, err := open(ctx, dsn, schema)
	if err != nil {
		return nil, err
	}

	xoSet := xotype.Set{}
	err = xocmd.LoadSchema(ctx, &xoSet, xocmd.NewArgs(""))
	if err != nil {
		return nil, err
	}

	set := protoschema.SetFromXO(&xoSet)
	if len(set.Schemas) == 0 {
		glog.V(1).Infof("no schema found")
		return nil, nil
	}
	return set.Schemas[0], nil
}

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
