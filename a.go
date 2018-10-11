package a

import (
	"fmt"

	"github.com/pkg/errors"
)

func Feed(args ...string) error {
	if len(args) == 0 {
		return errors.New("so hungry~")
	}
	for _, food := range args {
		fmt.Println("eat ", food, " ...")
	}
	fmt.Println("Ge~~~")
	return nil
}

func mockAddFeature1() {

}
