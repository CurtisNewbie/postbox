package postbox

import (
	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/gocommon/goauth"
	"github.com/curtisnewbie/miso/miso"
	"github.com/gin-gonic/gin"
)

const (
	ResourceQueryNotification   = "postbox:notification:query"
	ResourceReceiveNotification = "postbox:notification:receive"
	ResourceCreateNotification  = "postbox:notification:create"
)

func RegisterRoutes(rail miso.Rail) error {
	goauth.ReportOnBoostrapped(miso.EmptyRail(), []goauth.AddResourceReq{
		{Code: ResourceQueryNotification, Name: "Query Notifications"},
		{Code: ResourceReceiveNotification, Name: "Receive Notifications"},
		{Code: ResourceCreateNotification, Name: "Create Notifications"},
	})

	miso.BaseRoute("/open/api/v1").With(
		miso.SubPath("/notification").Group(
			miso.IPost[CreateNotificationReq]("/create", CreateNotificationEp).
				Desc("Create platform notification").
				Resource(ResourceCreateNotification),
		),
	)
	return nil
}

type CreateNotificationReq struct {
	Title           string `valid:"maxLen:255"`
	Message         string `valid:"maxLen:1000"`
	ReceiverUserNos []string
}

func CreateNotificationEp(c *gin.Context, rail miso.Rail, req CreateNotificationReq) (any, error) {
	return nil, CreateNotification(rail, miso.GetMySQL(), req, common.GetUser(rail))
}
