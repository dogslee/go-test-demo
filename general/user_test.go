package general

import (
	"testing"

	"github.com/dogslee/go-test-demo/general/gomocks"
	"github.com/golang/mock/gomock"
)

func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockUser := gomocks.NewMockUser(ctrl)
	gomock.InOrder(
		mockUser.EXPECT().Name(1).Return("Bob"),
		mockUser.EXPECT().Age(1).Return(18),

		mockUser.EXPECT().Name(2).Return("Tom"),
		mockUser.EXPECT().Age(2).Return(18),
	)
	userBiz := UserBiz{Data: mockUser}
	for i := 1; i <= 2; i++ {
		name, age := userBiz.GetUserInfo(i)
		t.Logf("name:%s, age%d", name, age)
	}

	// 无法执行
	// for i := 2; i >= 1; i-- {
	// 	name, age := userBiz.GetUserInfo(i)
	// 	t.Logf("name:%s, age%d", name, age)
	// }
}
