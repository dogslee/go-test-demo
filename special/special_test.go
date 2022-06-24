package special

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
)

func TestSpecial(t *testing.T) {

	info := "123"
	patches := gomonkey.ApplyFunc(GetSomething, func() string {
		return info
	})
	defer patches.Reset()

	for i := 0; i < 10; i++ {
		name := GetSomething()
		t.Logf("name:%v", name)
	}
	// do biz test ...
}

/*
  compare with apply_func_seq_test.go
*/
// func TestSpecial(t *testing.T) {
// 	Convey("TestApplyFuncReturn", t, func() {

// 		Convey("declares the values to be returned", func() {
// 			info1 := "hello cpp"

// 			patches := v2.ApplyFunc(GetSomething, func() string {
// 				return info1
// 			})
// 			defer patches.Reset()

// 			for i := 0; i < 10; i++ {
// 				name := GetSomething()
// 				// t.Logf("name:%v", name)
// 				So(name, ShouldEqual, name)
// 			}

// 			// patches.Reset() // if not reset will occur:patch has been existed
// 			// info2 := "hello golang"
// 			// patches.ApplyFuncReturn(fake.ReadLeaf, info2, nil)
// 			// for i := 0; i < 10; i++ {
// 			// 	output, err := fake.ReadLeaf("")
// 			// 	So(err, ShouldEqual, nil)
// 			// 	So(output, ShouldEqual, info2)
// 			// }
// 		})
// 	})
// }
