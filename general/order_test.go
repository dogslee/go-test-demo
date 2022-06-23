package general

import (
	"testing"

	"github.com/dogslee/go-test-demo/general/mocks"
)

func TestOrder(t *testing.T) {
	mockOrder := mocks.NewOrder(t)

	mockOrder.On("Count", 1).Return(10)
	mockOrder.On("Count", 2).Return(20)

	biz := &OrderBiz{Data: mockOrder}

	for i := 1; i <= 2; i++ {
		count := biz.GetOrderCount(i)
		t.Logf("order count:%d", count)
	}

	for i := 2; i >= 1; i-- {
		count := biz.GetOrderCount(i)
		t.Logf("desc order count:%d", count)
	}
}
