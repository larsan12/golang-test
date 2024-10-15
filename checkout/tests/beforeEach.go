package tt

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

var Run = RunDecorator(beforeEach, afterEach)

func beforeEach(t *testing.T) {
	fmt.Println("Before each")
	fmt.Println(LomsClient)
	*lomsController = *gomock.NewController(t)
	fmt.Println(LomsClient)
}

func afterEach(*testing.T) {
	fmt.Println("After each")
	CleanCartItemTable()
	lomsController.Finish()
}

func RunDecorator(beforeFunc, afterFunc func(*testing.T)) func(func(*testing.T)) func(*testing.T) {
	return func(test func(*testing.T)) func(*testing.T) {
		return func(t *testing.T) {
			if beforeFunc != nil {
				beforeFunc(t)
			}
			if afterFunc != nil {
				defer afterFunc(t)
			}
			test(t)
		}
	}
}
