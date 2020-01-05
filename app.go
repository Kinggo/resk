package resk

import (
	_ "github.com/Kinggo/account/core/accounts"
	"github.com/Kinggo/infra"
	"github.com/Kinggo/infra/base"
	"github.com/Kinggo/resk/apis/gorpc"
	_ "github.com/Kinggo/resk/apis/gorpc"
	_ "github.com/Kinggo/resk/apis/web"
	_ "github.com/Kinggo/resk/core/envelopes"
	"github.com/Kinggo/resk/jobs"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
