package schemax

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/RyoJerryYu/go-utilx/pkg/container/containerx"
	"github.com/RyoJerryYu/schemax/pkg/schemagen"
	protoplugin "github.com/RyoJerryYu/schemax/proto/schemax/plugin"
	protorecipe "github.com/RyoJerryYu/schemax/proto/schemax/recipe"
	protoschema "github.com/RyoJerryYu/schemax/proto/schemax/schema"
	"github.com/golang/glog"
)

func Gen(ctx context.Context, recipe *protorecipe.Recipe) error {
	plugins := PreparePlugins(recipe.Plugins)
	glog.V(2).Infof("gen for %d table sets, with %d plugins", len(recipe.TableSets), len(recipe.Plugins))

	schema, err := LoadSchema(ctx, recipe.Dsn, recipe.Schema)
	if err != nil || schema == nil {
		return err
	}

	for _, table := range schema.Tables {
		glog.V(3).Infof("db schema: %s, table: %s", schema.Name, table.Name)
	}

	return GenFor(recipe, schema, plugins...)
}

func GenFor(recipe *protorecipe.Recipe, schema *protoschema.Schema, plugins ...schemagen.TableGenerator) error {
	tableByName := containerx.MapByNames(schema.Tables)
	for _, tableSet := range recipe.TableSets {
		glog.V(1).Infof("table set: %s", tableSet.SetName)

		tables := make([]*protoschema.Table, 0)
		if len(tableSet.Tables) == 1 && tableSet.Tables[0] == "*" {
			tables = schema.Tables
		} else {
			for _, tableName := range tableSet.Tables {
				table, ok := tableByName[tableName]
				if !ok {
					glog.V(1).Infof("table not found: %s", tableName)
					continue
				}
				tables = append(tables, table)
			}
		}

		req := protoplugin.TableGeneratorRequest{
			TablesToGenerate: containerx.ToNames(tables),
			Tables:           tables,
			SourceSchema:     schema,
		}

		name2Plugin := containerx.MapByNames(plugins)

		// write to plugins
		for _, pluginCfg := range recipe.Plugins {
			plugin, ok := name2Plugin[pluginCfg.Plugin]
			if !ok {
				glog.Errorf("plugin not found: %s", pluginCfg.Plugin)
				continue
			}
			resp, err := plugin.Run(&req)
			if err != nil {
				return err
			}
			if resp.Error != nil {
				glog.Errorf("plugin %s error: %s", plugin.GetName(), resp.GetError())
				continue
			}

			glog.V(3).Infof("gen %d files for plugin %s", len(resp.Files), plugin.GetName())

			// Replace {set_name} with tableSet.SetName
			outDir := strings.Replace(pluginCfg.Out, "{set_name}", tableSet.SetName, -1)

			for _, file := range resp.Files {
				glog.V(3).Infof("gen file: %s", file.GetName())
				fullPath := filepath.Join(outDir, file.GetName())

				// Ensure the directory exists
				if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
					return err
				}

				err = os.WriteFile(fullPath, []byte(file.GetContent()), 0644)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
