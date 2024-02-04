package api

type CreateNotificationReq struct {
	Title           string `valid:"maxLen:255"`
	Message         string `valid:"maxLen:1000"`
	ReceiverUserNos []string
}
