package body

import (
	"net/url"
)

// AttrReq .
type AttrReq struct {
	Image string `json:"image"`
}

// AttrResp .
type AttrResp struct {
	BaseReply
	PersonNum  uint32 `json:"person_num"`
	PersonInfo []struct {
		Location struct {
			Left   float64 `json:"left"`
			Top    float64 `json:"top"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
			Score  float64 `json:"score"`
		} `json:"location"`
		Attributes struct {
			Gender           NameScore `json:"gender"`
			Age              NameScore `json:"age"`
			UpperColor       NameScore `json:"upper_color"`
			LowerColor       NameScore `json:"lower_color"`
			Cellphone        NameScore `json:"cellphone"`
			LowerWear        NameScore `json:"lower_wear"`
			UpperWear        NameScore `json:"upper_wear"`
			Headwear         NameScore `json:"headwear"`
			FaceMask         NameScore `json:"face_mask"`
			Glasses          NameScore `json:"glasses"`
			UpperWearFg      NameScore `json:"upper_wear_fg"`
			UpperWearTexture NameScore `json:"upper_wear_texture"`
			Orientation      NameScore `json:"orientation"`
			Umbrella         NameScore `json:"umbrella"`
			Bag              NameScore `json:"bag"`
			Smoke            NameScore `json:"smoke"`
			Vehicle          NameScore `json:"vehicle"`
			UpperCut         NameScore `json:"upper_cut"`
			LowerCut         NameScore `json:"lower_cut"`
			Occlusion        NameScore `json:"occlusion"`
			IsHuman          NameScore `json:"is_human"`
		} `json:"attributes"`
	} `json:"person_info"`
}

// NameScore .
type NameScore struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

// Attr 人体检测和属性识别.
func (b *Body) Attr(attrReq *AttrReq) (*AttrResp, error) {
	res := &AttrResp{}

	v := url.Values{
		"image": {attrReq.Image},
	}

	err := b.PostForm(ATTR_URL, v, res)

	return res, b.checkErr(res.BaseReply, err)
}
