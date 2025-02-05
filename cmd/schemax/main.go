package main

import (
	"context"
	"flag"
	"os"
	"os/user"

	// drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"

	"buf.build/go/protoyaml"
	"github.com/RyoJerryYu/schemax/proto/schemax/recipe"
	"github.com/RyoJerryYu/schemax/proto/schemax/schema"
	"github.com/golang/glog"
	"github.com/octago/sflags/gen/gflag"
	"github.com/xo/dburl"
	"github.com/xo/dburl/passfile"
	xocmd "github.com/xo/xo/cmd"
	"github.com/xo/xo/loader"
	xotype "github.com/xo/xo/types"
)

type Option struct {
	Config string `flag:"c" desc:"config file"`
	Dsn    string `flag:"dsn" desc:"database dsn"`
	Schema string `flag:"schema" desc:"database schema"`
}

func DefaultOption() *Option {
	return &Option{
		Config: "schemax.gen.yaml",
	}
}

func (o *Option) Run() error {
	glog.V(3).Infof("config: %s", o.Config)

	cfgRaw, err := os.ReadFile(o.Config)
	if err != nil {
		return err
	}

	cfgSet := recipe.Recipe{
		Dsn:    o.Dsn,
		Schema: o.Schema,
	}
	err = protoyaml.Unmarshal(cfgRaw, &cfgSet)
	if err != nil {
		return err
	}

	glog.V(2).Infof("gen for %d table sets, with %d plugins", len(cfgSet.TableSets), len(cfgSet.Plugins))
	glog.Infof("schema: %s", cfgSet.Schema)

	ctx := context.Background()
	ctx, err = open(ctx, cfgSet.Dsn, cfgSet.Schema)
	if err != nil {
		return err
	}

	xoSet := xotype.Set{}
	err = xocmd.LoadSchema(ctx, &xoSet, xocmd.NewArgs(""))
	if err != nil {
		return err
	}

	set := schema.SetFromXO(&xoSet)

	for _, schema := range set.Schemas {
		for _, table := range schema.Tables {
			glog.V(3).Infof("schema: %s, table: %s", schema.Name, table.Name)
		}
	}

	rawSet, err := proto.Marshal(set)
	if err != nil {
		return err
	}

	if len(cfgSet.Plugins) == 0 {
		// write to stdout
		_, err = os.Stdout.Write(rawSet)
		return err
	}

	// write to plugins

	return nil
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

func main() {
	opt := DefaultOption()
	gflag.ParseToDef(opt)
	flag.Parse()
	defer glog.Flush()

	glog.V(1).Infof("config: %s", opt.Config)

	if err := opt.Run(); err != nil {
		glog.Errorf("error: %v", err)
		os.Exit(1)
	}
}
