package face

/**
 * API文档 https://cloud.baidu.com/doc/FACE/s/dk25rde0o
 * API文档 https://cloud.baidu.com/doc/FACE/s/rk25rddle
 */

import (
	"encoding/json"
	"net/url"
)

// VerifyRes .
type VerifyRes struct {
	FaceLiveness float64 `json:"face_liveness"`
	FaceList     []struct {
		FaceToken       string  `json:"face_token"`
		FaceProbability float64 `json:"face_probability"`
		Age             int64   `json:"age"`
		Beauty          float64 `json:"beauty"`
		Location        struct {
			Left     float64 `json:"left"`
			Top      float64 `json:"top"`
			Width    float64 `json:"width"`
			Height   float64 `json:"height"`
			Rotation float64 `json:"rotation"`
		} `json:"location"`
		Quality struct {
			Occlusion struct {
				LeftEye     float64 `json:"left_eye"`
				RightEye    float64 `json:"right_eye"`
				Nose        float64 `json:"nose"`
				Mouth       float64 `json:"mouth"`
				LeftCheek   float64 `json:"left_cheek"`
				RightCheek  float64 `json:"right_cheek"`
				ChinContour float64 `json:"chin_contour"`
			} `json:"occlusion"`
			Blur         float64 `json:"blur"`
			Illumination float64 `json:"illumination"`
			Completeness float64 `json:"completeness"`
		} `json:"quality"`
		FaceType struct {
			Type        string  `json:"type"`
			Probability float64 `json:"probability"`
		} `json:"face_type"`
	} `json:"face_list"`
}

// VerifyFace .
type VerifyFace struct {
	Image     string `json:"image"`
	ImageType string `json:"image_type"`
	FaceField string `json:"face_field"`
	Option    string `json:"option"`
}

// 活体检测
func (f *Face) FaceVerify(verify []*VerifyFace) (res *VerifyRes, err error) {
	res = &VerifyRes{}

	var data []map[string]string
	for _, v := range verify {
		temp := map[string]string{
			"image":      v.Image,
			"image_type": v.ImageType,
		}
		if v.Option != "" {
			temp["option"] = v.Option
		}
		if v.FaceField != "" {
			temp["face_field"] = v.FaceField
		}
		data = append(data, temp)
	}
	marshal, err := json.Marshal(data)

	err = f.PostJson(FACE_VERIFY, marshal, &Reply{Result: res})

	return
}

// VerifyPersonal .
type VerifyPersonal struct {
	Image           string `json:"image"`
	ImageType       string `json:"image_type"`
	Name            string `json:"name"`
	IdCardNumber    string `json:"id_card_number"`
	QualityControl  string `json:"quality_control"`
	LivenessControl string `json:"liveness_control"`
}

// VerifyScore .
type VerifyScore struct {
	Score float64 `json:"score"`
}

// 身份证照片验证，包含了身份证号与名字验证
func (f *Face) PersonalVerify(verify *VerifyPersonal) (res *VerifyScore, err error) {
	res = &VerifyScore{}
	v := url.Values{
		"image":          {verify.Image},
		"image_type":     {verify.ImageType},
		"id_card_number": {verify.IdCardNumber},
		"name":           {verify.Name},
	}
	if verify.QualityControl != "" {
		v.Add("quality_control", verify.QualityControl)
	}
	if verify.LivenessControl != "" {
		v.Add("liveness_control", verify.LivenessControl)
	}
	err = f.PostForm(PERSON_VERIFY, v, &Reply{
		Result: res,
	})
	return
}

// 身份证号与名字验证
func (f *Face) IdMatch(name, id_card_number string) (err error) {
	err = f.PostForm(ID_MATCH, url.Values{
		"name":           {name},
		"id_card_number": {id_card_number},
	}, &Reply{})
	return
}
