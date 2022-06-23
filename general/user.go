package general

type User interface {
	Name(userId int) string
	Age(userId int) int
}

type UserBiz struct {
	Data User
}

func (biz *UserBiz) GetUserInfo(userId int) (string, int) {
	return biz.Data.Name(userId), biz.Data.Age(userId)
}
