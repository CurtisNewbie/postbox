package postbox

import (
	"testing"

	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/miso/miso"
)

func _notificationPreTest(t *testing.T) miso.Rail {
	miso.SetProp(miso.PropMySQLEnabled, true)
	miso.SetProp(miso.PropMySQLUser, "root")
	miso.SetProp(miso.PropMySQLSchema, "postbox")
	miso.SetProp(miso.PropMySQLPassword, "")
	miso.SetProp("client.addr.user-vault.host", "localhost")
	miso.SetProp("client.addr.user-vault.port", "8089")
	rail := miso.EmptyRail()
	miso.InitMySQLFromProp(rail)
	return rail
}

func TestSaveNotification(t *testing.T) {
	rail := _notificationPreTest(t)
	user := common.User{
		Username: "postbox",
	}
	err := SaveNotification(rail, miso.GetMySQL(), SaveNotifiReq{
		UserNo:  "UE1049787455160320075953",
		Title:   "Some message",
		Message: "Notification should be saved",
	}, user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateNotification(t *testing.T) {
	miso.SetLogLevel("debug")
	rail := _notificationPreTest(t)
	user := common.User{
		Username: "postbox",
	}
	err := CreateNotification(rail, miso.GetMySQL(), CreateNotificationReq{
		ReceiverUserNos: []string{"UE1049787455160320075953"},
		Title:           "Some message",
		Message:         "Notification should be saved",
	}, user)
	if err != nil {
		t.Fatal(err)
	}
}
