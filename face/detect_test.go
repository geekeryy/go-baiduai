package face

import (
	"fmt"
	"testing"
)

func TestFace_Detect(t *testing.T) {
	detect, err := f.Detect(&DetectReq{
		Image:           "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1573837226842&di=311e9cb79691395e0c6243b9b4751b96&imgtype=0&src=http%3A%2F%2Fn.sinaimg.cn%2Fsinacn10104%2F719%2Fw358h361%2F20181223%2Fea63-hqqzpku4309788.bmp",
		ImageType:       "URL",
		FaceField:       "age,beauty,expression,face_shape,gender,glasses,landmark,landmark150,race,quality,eye_status,emotion,face_type,mask,spoofing",
		MaxFaceNum:      10,
		FaceType:        "LIVE",
		LivenessControl: "LOW",
	})
	fmt.Printf("%+v \n %+v", detect, err)
}
