package postbox

import (
	"github.com/curtisnewbie/miso/middleware/logbot"
	"github.com/curtisnewbie/miso/middleware/user-vault/common"
	"github.com/curtisnewbie/miso/miso"
)

func BootstrapServer(args []string) {
	common.LoadBuiltinPropagationKeys()
	logbot.EnableLogbotErrLogReport()
	miso.PreServerBootstrap(RegisterRoutes)
	miso.BootstrapServer(args)
}
