package face

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// DetectReq .
type DetectReq struct {
	Image           string `json:"image"`
	ImageType       string `json:"image_type"`
	FaceField       string `json:"face_field"`
	MaxFaceNum      uint32 `json:"max_face_num"`
	FaceType        string `json:"face_type"`
	LivenessControl string `json:"liveness_control"`
}

// DetectResp .
type DetectResp struct {
	FaceNum  int `json:"face_num"`
	FaceList []struct {
		FaceToken       string          `json:"face_token"`
		FaceProbability float64         `json:"face_probability"`
		Age             int64           `json:"age"`
		Beauty          float64         `json:"beauty"`
		Spoofing        float64         `json:"spoofing"`
		Expression      TypeProbability `json:"expression"`
		Glasses         TypeProbability `json:"glasses"`
		FaceType        TypeProbability `json:"face_type"`
		Emotion         TypeProbability `json:"emotion"`
		Gender          TypeProbability `json:"gender"`
		FaceShape       TypeProbability `json:"face_shape"`
		Mask            struct {
			Type        int     `json:"type"`
			Probability float64 `json:"probability"`
		} `json:"mask"`
		EyeStatus struct {
			LeftEye  float64 `json:"left_eye"`
			RightEye float64 `json:"right_eye"`
		} `json:"eye_status"`
		Location struct {
			Left     float64 `json:"left"`
			Top      float64 `json:"top"`
			Width    float64 `json:"width"`
			Height   float64 `json:"height"`
			Rotation float64 `json:"rotation"`
		} `json:"location"`
		Angle struct {
			Yaw   float64 `json:"yaw"`
			Pitch float64 `json:"pitch"`
			Roll  float64 `json:"roll"`
		} `json:"angle"`
		Quality struct {
			Occlusion struct {
				LeftEye    float64 `json:"left_eye"`
				RightEye   float64 `json:"right_eye"`
				Nose       float64 `json:"nose"`
				Mouth      float64 `json:"mouth"`
				LeftCheek  float64 `json:"left_cheek"`
				RightCheek float64 `json:"right_cheek"`
				Chin       float64 `json:"chin"`
			} `json:"occlusion"`
			Blur         float64 `json:"blur"`
			Illumination float64 `json:"illumination"`
			Completeness int64   `json:"completeness"`
		} `json:"quality"`
	} `json:"face_list"`
}

// TypeProbability .
type TypeProbability struct {
	Type        string  `json:"type"`
	Probability float64 `json:"probability"`
}

// 人脸检测
func (f *Face) Detect(d *DetectReq) (res *DetectResp, err error) {
	res = &DetectResp{}

	marshal, err := json.Marshal(d)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = f.PostJson(DETECT_URL, marshal, &Reply{
		Result: res,
	})
	return
}
