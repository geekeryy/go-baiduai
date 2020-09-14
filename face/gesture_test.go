package face

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFace_Gesture(t *testing.T) {
	image, _ := ioutil.ReadFile("gesture.png")
	imageBase64 := base64.StdEncoding.EncodeToString(image)

	res, err := f.Gesture(imageBase64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", res)
	}
}
