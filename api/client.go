package api

import (
	"fmt"

	"github.com/curtisnewbie/miso/miso"
)

func CreateNotification(rail miso.Rail, req CreateNotificationReq) error {
	var resp miso.GnResp[miso.Void]
	err := miso.NewDynTClient(rail, "/open/api/v1/notification/create", "postbox").
		Require2xx().
		PostJson(req).
		Json(&resp)
	if err != nil {
		return fmt.Errorf("failed to create notiication, req: %+v, %v", req, err)
	}
	return resp.Err()
}
