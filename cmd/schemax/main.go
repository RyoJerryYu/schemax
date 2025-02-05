package main

import (
	"context"
	"flag"
	"os"

	// drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"buf.build/go/protoyaml"
	"github.com/RyoJerryYu/schemax/pkg/schemax"
	protorecipe "github.com/RyoJerryYu/schemax/proto/schemax/recipe"
	"github.com/golang/glog"
	"github.com/octago/sflags/gen/gflag"
)

type Option struct {
	Config string `flag:"c" desc:"config file"`
	Dsn    string `flag:"dsn" desc:"database dsn"`
	Schema string `flag:"schema" desc:"database schema"`
}

func DefaultOption() *Option {
	return &Option{
		// Config: "schemax.gen.yaml",
	}
}

func (o *Option) Run() error {
	recipe := protorecipe.Recipe{
		Dsn:    o.Dsn,
		Schema: o.Schema,
	}

	if o.Config != "" {
		glog.V(3).Infof("config: %s", o.Config)

		cfgRaw, err := os.ReadFile(o.Config)
		if err != nil {
			return err
		}

		err = protoyaml.Unmarshal(cfgRaw, &recipe)
		if err != nil {
			return err
		}
	}

	if len(recipe.TableSets) == 0 {
		recipe.TableSets = append(recipe.TableSets, &protorecipe.TableSet{
			SetName: "default",
			Tables:  []string{"*"},
		})
	}

	return schemax.Gen(context.Background(), &recipe)
}

func main() {
	opt := DefaultOption()
	gflag.ParseToDef(opt)
	flag.Parse()
	defer glog.Flush()

	if err := opt.Run(); err != nil {
		glog.Errorf("error: %v", err)
		os.Exit(1)
	}
}
