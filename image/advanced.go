package body

import "net/url"

// AdvancedReq .
type AdvancedReq struct {
	Image string `json:"image"`
}

// AdvancedResp .
type AdvancedResp struct {
	BaseReply
	ResultNum uint32 `json:"result_num"`
	Result    []struct {
		Keyword   string  `json:"keyword"`
		Score     float64 `json:"score"`
		Root      string  `json:"root"`
		BaikeInfo struct {
			BaikeUrl    string `json:"baike_url"`
			ImageUrl    string `json:"image_url"`
			Description string `json:"description"`
		} `json:"baike_info"`
	} `json:"result"`
}

// NameScore .
type NameScore struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

// Advanced 通用物体和场景识别高级版 .
func (i *Image) Advanced(advancedReq *AdvancedReq) (*AdvancedResp, error) {
	res := &AdvancedResp{}

	v := url.Values{
		"image": {advancedReq.Image},
	}

	return res, i.checkErr(res.BaseReply, i.PostForm(ADVANCED_URL, v, res))
}
