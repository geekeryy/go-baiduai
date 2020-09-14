package face

import (
	"fmt"
	"testing"
)

func TestFace_FaceVerify(t *testing.T) {
	var temp []*VerifyFace
	temp = append(temp, &VerifyFace{
		Image:     "600d3bf0edb0742e90a8e20e2d7d28f5",
		ImageType: "FACE_TOKEN",
		//FaceField: "age,beauty,expression,face_shape,gender,glasses,landmark,race,quality,face_type",
		//Option:    "COMMON",
	})
	temp = append(temp, &VerifyFace{
		Image:     "d9927337738909b5ac001ba6b253f75c",
		ImageType: "FACE_TOKEN",
		//FaceField: "age,beauty,expression,face_shape,gender,glasses,landmark,race,quality,face_type",
		//Option:    "COMMON",
	})

	if res, err := f.FaceVerify(temp); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v \n", res)
	}
}

func TestFace_PersonalVerify(t *testing.T) {
	req := &VerifyPersonal{
		Image:        "cc576b6bec4fc38f62c04e67c327da57",
		ImageType:    "FACE_TOKEN",
		Name:         "江杨",
		IdCardNumber: "513922199909281213",
	}
	if res, err := f.PersonalVerify(req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", res)
	}
}

func TestFace_IdMatch(t *testing.T) {
	req := &VerifyPersonal{
		Name:         "xx",
		IdCardNumber: "123456",
	}
	if err := f.IdMatch(req.Name, req.IdCardNumber); err != nil {
		fmt.Println(err)
	}
}
