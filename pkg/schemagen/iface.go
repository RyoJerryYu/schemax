package schemagen

import protoplugin "github.com/RyoJerryYu/schemax/proto/schemax/plugin"

type TableGenerator interface {
	GetName() string
	Run(req *protoplugin.TableGeneratorRequest) (*protoplugin.TableGeneratorResponse, error)
}
