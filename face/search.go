package face

/**
 * API文档 https://cloud.baidu.com/doc/FACE/s/zk25rde8b
 */

import (
	"net/url"
)

// FaceSearch .
type FaceSearch struct {
	Image           string
	ImageType       string
	GroupIdList     string
	UserId          string
	QualityControl  string
	LivenessControl string
	MaxUserNum      string
}

// SearchRes .
type SearchRes struct {
	FaceToken string `json:"face_token"`
	UserList  []struct {
		GroupId  string  `json:"group_id"`
		UserId   string  `json:"user_id"`
		UserInfo string  `json:"user_info"`
		Score    float64 `json:"score"`
	} `json:"user_list"`
}

// 人脸搜索
func (f *Face) SearchFace(search *FaceSearch) (res *SearchRes, err error) {
	res = &SearchRes{}

	v := url.Values{
		"image":         {search.Image},
		"image_type":    {search.ImageType},
		"group_id_list": {search.GroupIdList},
	}
	if search.UserId != "" {
		v.Add("user_id", search.UserId)
	}
	if search.QualityControl != "" {
		v.Add("quality_control", search.QualityControl)
	}
	if search.LivenessControl != "" {
		v.Add("liveness_control", search.LivenessControl)
	}
	if search.MaxUserNum != "" {
		v.Add("max_user_num", search.MaxUserNum)
	}

	err = f.PostForm(FACE_SEARCH, v, &Reply{
		Result: res,
	})
	return
}
