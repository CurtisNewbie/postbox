package api

import (
	"testing"

	"github.com/curtisnewbie/miso/miso"
)

func _clientPreTest(t *testing.T) miso.Rail {
	miso.SetProp("client.addr.postbox.host", "localhost")
	miso.SetProp("client.addr.postbox.port", "8092")
	rail := miso.EmptyRail()
	return rail
}

func TestCreateNotification(t *testing.T) {
	rail := _clientPreTest(t)
	err := CreateNotification(rail, CreateNotificationReq{
		ReceiverUserNos: []string{"UE1049787455160320075953"},
		Title:           "Some message",
		Message:         "Notification should be saved",
	})
	if err != nil {
		t.Fatal(err)
	}
}
