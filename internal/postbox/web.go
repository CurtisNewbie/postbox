package postbox

import (
	"github.com/curtisnewbie/gocommon/auth"
	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/miso/miso"
	"github.com/curtisnewbie/postbox/api"
	"github.com/gin-gonic/gin"
)

const (
	ResourceQueryNotification  = "postbox:notification:query"
	ResourceCreateNotification = "postbox:notification:create"
)

func RegisterRoutes(rail miso.Rail) error {
	auth.ExposeResourceInfo([]auth.Resource{
		{Code: ResourceQueryNotification, Name: "Query Notifications"},
		{Code: ResourceCreateNotification, Name: "Create Notifications"},
	})

	miso.BaseRoute("/open/api/v1/notification").Group(

		miso.IPost("/create", CreateNotificationEp).
			Desc("Create platform notification").
			Resource(ResourceCreateNotification),

		miso.IPost("/query", QueryNotificationEp).
			Desc("Query platform notification").
			Resource(ResourceQueryNotification),

		miso.Get("/count", CountNotificationEp).
			Desc("Count received platform notification").
			Resource(ResourceQueryNotification),

		miso.IPost("/open", OpenNotificationEp).
			Desc("Record user opened platform notification").
			Resource(ResourceQueryNotification),
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

type OpenNotificationReq struct {
	NotifiNo string `valid:"notEmpty"`
}

func OpenNotificationEp(c *gin.Context, rail miso.Rail, req OpenNotificationReq) (any, error) {
	return nil, OpenNotification(rail, miso.GetMySQL(), req, common.GetUser(rail))
}
