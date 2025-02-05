package main

import (
	"flag"

	"github.com/RyoJerryYu/schemax/pkg/schemagen"
	"github.com/RyoJerryYu/schemax/proto/schemax/plugin"
	"github.com/golang/glog"
	"github.com/octago/sflags/gen/gflag"
)

type Option struct {
}

// GetName implements schemagen.TableGenerator.
func (o *Option) GetName() string {
	panic("unimplemented")
}

// Run implements schemagen.TableGenerator.
func (o *Option) Run(req *plugin.TableGeneratorRequest) (*plugin.TableGeneratorResponse, error) {
	glog.V(3).Infof("req: %+v", req)
	return &plugin.TableGeneratorResponse{
		Files: []*plugin.TableGeneratorResponse_File{
			{
				Name:    "test.go",
				Content: `package main`,
			},
		},
	}, nil
}

func DefaultOption() *Option {
	return &Option{}
}

var _ schemagen.TableGenerator = (*Option)(nil)

func main() {
	opt := DefaultOption()
	gflag.ParseToDef(opt)
	flag.Parse()
	defer glog.Flush()

	schemagen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(opt)
}
