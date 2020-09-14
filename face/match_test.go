package face

import (
	"fmt"
	"testing"
)

func TestFace_MatchFace(t *testing.T) {
	temp := []*FaceMatch{}
	temp = append(temp, &FaceMatch{
		Image:     "600d3bf0edb0742e90a8e20e2d7d28f5",
		ImageType: "FACE_TOKEN",
	})

	temp = append(temp, &FaceMatch{
		Image:     "d9927337738909b5ac001ba6b253f75c",
		ImageType: "FACE_TOKEN",
	})
	if res, err := f.MatchFace(temp); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", res)
	}
}
