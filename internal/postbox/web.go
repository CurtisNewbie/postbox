package postbox

import (
	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/gocommon/goauth"
	"github.com/curtisnewbie/miso/miso"
	"github.com/curtisnewbie/postbox/api"
	"github.com/gin-gonic/gin"
)

const (
	ResourceQueryNotification  = "postbox:notification:query"
	ResourceCreateNotification = "postbox:notification:create"
)

func RegisterRoutes(rail miso.Rail) error {
	goauth.ReportOnBoostrapped(miso.EmptyRail(), []goauth.AddResourceReq{
		{Code: ResourceQueryNotification, Name: "Query Notifications"},
		{Code: ResourceCreateNotification, Name: "Create Notifications"},
	})

	miso.BaseRoute("/open/api/v1").With(
		miso.SubPath("/notification").Group(
			miso.IPost[api.CreateNotificationReq]("/create", CreateNotificationEp).
				Desc("Create platform notification").
				Resource(ResourceCreateNotification),

			miso.IPost[QueryNotificationReq]("/query", QueryNotificationEp).
				Desc("Query platform notification").
				Resource(ResourceQueryNotification),

			miso.Get("/count", CountNotificationEp).
				Desc("Count received platform notification").
				Resource(ResourceQueryNotification),
		),
	)
	return nil
}

func CreateNotificationEp(c *gin.Context, rail miso.Rail, req api.CreateNotificationReq) (any, error) {
	return nil, CreateNotification(rail, miso.GetMySQL(), req, common.GetUser(rail))
}

type QueryNotificationReq struct {
	Page   miso.Paging
	Status string
}

func QueryNotificationEp(c *gin.Context, rail miso.Rail, req QueryNotificationReq) (any, error) {
	return QueryNotification(rail, miso.GetMySQL(), req, common.GetUser(rail))
}

func CountNotificationEp(c *gin.Context, rail miso.Rail) (any, error) {
	return CountNotification(rail, miso.GetMySQL(), common.GetUser(rail))
}
