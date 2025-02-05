package schemax

import (
	"os"
	"strings"

	"github.com/RyoJerryYu/schemax/pkg/schemagen"
	protoplugin "github.com/RyoJerryYu/schemax/proto/schemax/plugin"
	protorecipe "github.com/RyoJerryYu/schemax/proto/schemax/recipe"
	"google.golang.org/protobuf/proto"
)

func PreparePlugins(recipePlugin []*protorecipe.PluginSetting) []schemagen.TableGenerator {
	plugins := make([]schemagen.TableGenerator, 0, len(recipePlugin))
	for _, plugin := range recipePlugin {
		plugins = append(plugins, &cmdTableGenerator{
			plugin: plugin.Plugin,
			out:    plugin.Out,
			opts:   plugin.Opts,
		})
	}
	if len(plugins) == 0 {
		plugins = append(plugins, &stdoutTableGenerator{})
	}
	return plugins
}

type stdoutTableGenerator struct{}

// Name implements schemagen.TableGenerator.
func (m *stdoutTableGenerator) Name() string {
	return "stdout"
}

// Run implements schemagen.TableGenerator.
func (m *stdoutTableGenerator) Run(req *protoplugin.TableGeneratorRequest) (*protoplugin.TableGeneratorResponse, error) {
	rawReq, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	// write to stdout
	_, err = os.Stdout.Write(rawReq)
	if err != nil {
		return nil, err
	}
	return &protoplugin.TableGeneratorResponse{
		Files: []*protoplugin.TableGeneratorResponse_File{},
	}, nil
}

var _ schemagen.TableGenerator = (*stdoutTableGenerator)(nil)

type cmdTableGenerator struct {
	plugin string
	out    string
	opts   []string
}

// Name implements schemagen.TableGenerator.
func (c *cmdTableGenerator) Name() string {
	return c.plugin
}

// Run implements schemagen.TableGenerator.
func (c *cmdTableGenerator) Run(req *protoplugin.TableGeneratorRequest) (*protoplugin.TableGeneratorResponse, error) {
	req = proto.Clone(req).(*protoplugin.TableGeneratorRequest)
	req.Parameter = proto.String(strings.Join(c.opts, ","))

	rawReq, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	_ = rawReq // TODO
	_ = c.out
	// find in path
	// write to stdin
	// read from stdout
	// parse response

	return &protoplugin.TableGeneratorResponse{
		Files: []*protoplugin.TableGeneratorResponse_File{},
	}, nil
}

var _ schemagen.TableGenerator = (*cmdTableGenerator)(nil)
