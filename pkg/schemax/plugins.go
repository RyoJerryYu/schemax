package schemax

import (
	"bytes"
	"os"
	"os/exec"
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
	var rawResp bytes.Buffer

	// Split c.plugin into command and arguments
	parts := strings.Fields(c.plugin)
	cmdName := parts[0]
	cmdArgs := parts[1:]

	cmd := exec.Command(cmdName, cmdArgs...)
	// write rawReq to stdin and read rawResp from stdout
	cmd.Stdin = strings.NewReader(string(rawReq))
	cmd.Stdout = &rawResp

	// run command
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// parse response
	resp := &protoplugin.TableGeneratorResponse{}
	err = proto.Unmarshal(rawResp.Bytes(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

var _ schemagen.TableGenerator = (*cmdTableGenerator)(nil)
