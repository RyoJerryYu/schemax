package schemagen

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/RyoJerryYu/schemax/proto/schemax/plugin"
	"google.golang.org/protobuf/proto"
)

type Options struct {
	// If ParamFunc is non-nil, it will be called with each unknown
	// generator parameter.
	//
	// Plugins for protoc can accept parameters from the command line,
	// passed in the --<lang>_out protoc, separated from the output
	// directory with a colon; e.g.,
	//
	//   --go_out=<param1>=<value1>,<param2>=<value2>:<output_directory>
	//
	// Parameters passed in this fashion as a comma-separated list of
	// key=value pairs will be passed to the ParamFunc.
	//
	// The (flag.FlagSet).Set method matches this function signature,
	// so parameters can be converted into flags as in the following:
	//
	//   var flags flag.FlagSet
	//   value := flags.Bool("param", false, "")
	//   opts := &protogen.Options{
	//     ParamFunc: flags.Set,
	//   }
	//   opts.Run(func(p *protogen.Plugin) error {
	//     if *value { ... }
	//   })
	ParamFunc func(name, value string) error
}

func (opts Options) Run(generator TableGenerator) {
	err := opts.run(generator)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}

func (opts Options) run(generator TableGenerator) error {
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	req := &plugin.TableGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		return err
	}
	opts.InitParams(req)
	resp, err := generator.Run(req)
	if err != nil {
		// Errors from the plugin function are reported by setting the
		// error field in the TableGeneratorResponse.
		//
		// In contrast, errors that indicate a problem in protoc
		// itself (unparsable input, I/O errors, etc.) are reported
		// to stderr.
		if resp.Error == nil {
			resp.Error = proto.String(err.Error())
		}
	}
	out, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(out); err != nil {
		return err
	}
	return nil
}

func (opts Options) InitParams(req *plugin.TableGeneratorRequest) {
	if opts.ParamFunc == nil {
		return
	}

	for _, param := range strings.Split(req.GetParameter(), ",") {
		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = param[i+1:]
			param = param[0:i]
		}

		switch param {
		case "":
			// Ignore.
		default:
			opts.ParamFunc(param, value)
		}
	}
}
