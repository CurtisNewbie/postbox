package postbox

import (
	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/miso/miso"
)

func BootstrapServer(args []string) {
	common.LoadBuiltinPropagationKeys()
	miso.PreServerBootstrap(RegisterRoutes)
	miso.BootstrapServer(args)
}
