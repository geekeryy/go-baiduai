package face

import (
	"encoding/base64"
	"fmt"
	go_baiduai "github.com/comeonjy/go-baiduai"
	"io/ioutil"
	"testing"
)

var f = New(go_baiduai.FACE_APP_KEY, go_baiduai.FACE_APP_SECRET)

var reg = &Register{
	//Image:           "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1573837226842&di=311e9cb79691395e0c6243b9b4751b96&imgtype=0&src=http%3A%2F%2Fn.sinaimg.cn%2Fsinacn10104%2F719%2Fw358h361%2F20181223%2Fea63-hqqzpku4309788.bmp",
	//Image:           "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=2385696140,1893414624&fm=15&gp=0.jpg",
	Image:           "http://ww1.sinaimg.cn/large/0060QYYfly1g8zqstbp57j30dq0dqwex.jpg",
	ImageType:       "URL",
	GroupId:         "demo1",
	UserId:          "jy",
	UserInfo:        "jy",
	QualityControl:  "NONE",
	LivenessControl: "NONE",
	ActionType:      "APPEND",
}

func TestFace_AddUser(t *testing.T) {
	image, _ := ioutil.ReadFile("jy.jpg")
	imageBase64 := base64.StdEncoding.EncodeToString(image)
	reg.Image = imageBase64
	reg.ImageType = "BASE64"
	if res, err := f.AddUser(reg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestFace_GetUser(t *testing.T) {
	if res, err := f.GetUser("1", "demo1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestFace_DelUserFace(t *testing.T) {
	if err := f.DelUserFace("1", "demo1", "d9927337738909b5ac001ba6b253f75c"); err != nil {
		fmt.Println(err)
	}
}

func TestFace_UpdateUser(t *testing.T) {
	reg.ActionType = "UPDATE"
	if res, err := f.UpdateUser(reg); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestFace_ListUser(t *testing.T) {
	if res, err := f.ListUser("demo1", "0", "10"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestFace_ListUserFace(t *testing.T) {
	if res, err := f.ListUserFace("1", "demo1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestFace_CopyUser(t *testing.T) {

}
