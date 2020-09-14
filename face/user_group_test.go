package face

import (
	"fmt"
	"testing"
)

func TestFace_AddUserGroup(t *testing.T) {
	if err := f.AddUserGroup("demo2"); err != nil {
		fmt.Println(err)
	}
}

func TestFace_ListUserGroup(t *testing.T) {
	if list, err := f.ListUserGroup("0", "10"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v \n", list)
	}
}

func TestFace_DelUserGroup(t *testing.T) {
	if err := f.DelUserGroup("demo2"); err != nil {
		fmt.Println(err)
	}
}
