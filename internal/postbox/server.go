package postbox

import (
	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/gocommon/goauth"
	"github.com/curtisnewbie/miso/miso"
)

func BootstrapServer(args []string) {
	common.LoadBuiltinPropagationKeys()
	goauth.ReportOnBoostrapped(miso.EmptyRail(), []goauth.AddResourceReq{})

	// TODO:

	miso.BootstrapServer(args)
}
