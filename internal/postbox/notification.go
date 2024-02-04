package postbox

import (
	"fmt"

	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/miso/miso"
	uservault "github.com/curtisnewbie/user-vault/api"
	"gorm.io/gorm"
)

func CreateNotification(rail miso.Rail, db *gorm.DB, req CreateNotificationReq, user common.User) error {
	if len(req.ReceiverUserNos) < 1 {
		return nil
	}

	// check whether the userNos are leegal
	req.ReceiverUserNos = miso.Distinct(req.ReceiverUserNos)
	aw := miso.NewAwaitFutures[string](PostboxPool)

	for i := range req.ReceiverUserNos {
		userNo := req.ReceiverUserNos[i]
		aw.SubmitAsync(func() (string, error) {
			_, err := uservault.FindUser(rail.NextSpan(), uservault.FindUserReq{
				UserNo: &userNo,
			})
			if err != nil {
				rail.Errorf("failed to FindUser, %v", err)
				return userNo, err
			}
			return userNo, err
		})
	}

	futures := aw.Await()
	for _, f := range futures {
		if userNo, err := f.Get(); err != nil {
			return miso.NewErr("User not found", "failed to FindUser, userNo: %v, %v", userNo, err)
		}
	}

	for _, u := range req.ReceiverUserNos {
		sr := SaveNotifiReq{
			UserNo:  u,
			Title:   req.Title,
			Message: req.Message,
		}
		if err := SaveNotification(rail, db, sr, user); err != nil {
			return fmt.Errorf("failed to save notification, %+v, %v", sr, err)
		}
	}

	return nil
}

type SaveNotifiReq struct {
	UserNo  string
	Title   string
	Message string
}

func SaveNotification(rail miso.Rail, db *gorm.DB, req SaveNotifiReq, user common.User) error {
	notifiNo := NotifiNo()
	err := db.Exec(`insert into notification (user_no, notifi_no, title, message, create_by) values (?, ?, ?, ?, ?)`,
		req.UserNo, notifiNo, req.Title, req.Message, user.Username).Error
	if err != nil {
		return fmt.Errorf("failed to save notifiication record, %+v", req)
	}
	return nil
}

func NotifiNo() string {
	return miso.GenIdP("notif_")
}
