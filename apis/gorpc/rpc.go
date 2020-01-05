package gorpc

import (
	"github.com/Kinggo/infra"
	"github.com/Kinggo/infra/base"
)

type GoRpcApiStarter struct {
	infra.BaseStarter
}

func (g *GoRpcApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRpc))
}
