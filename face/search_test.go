package face

import (
	"fmt"
	"testing"
)

var search = &FaceSearch{
	//Image:           "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1573837226842&di=311e9cb79691395e0c6243b9b4751b96&imgtype=0&src=http%3A%2F%2Fn.sinaimg.cn%2Fsinacn10104%2F719%2Fw358h361%2F20181223%2Fea63-hqqzpku4309788.bmp",
	Image:           "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=2385696140,1893414624&fm=15&gp=0.jpg",
	ImageType:       "URL",
	GroupIdList:     "demo1",
	UserId:          "1",
	QualityControl:  "NONE",
	LivenessControl: "NONE",
	MaxUserNum:      "2",
}

func TestFace_SearchFace(t *testing.T) {
	if res, err := f.SearchFace(search); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", res)
	}
}
