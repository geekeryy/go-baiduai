package face

import (
	"encoding/json"
	"errors"

	"net/url"
)

var (
	gestureMap = map[string]string{
		"One":            "数字1（原食指）",
		"Five":           "数字5（原掌心向前）",
		"Fist":           "拳头",
		"OK":             "OK",
		"Prayer":         "祈祷",
		"Congratulation": "作揖",
		"Honour":         "作别",
		"Heart_single":   "单手比心",
		"Thumb_up":       "点赞",
		"Thumb_down":     "Diss",
		"ILY":            "我爱你",
		"Palm_up":        "掌心向上",
		"Heart_1":        "双手比心1",
		"Heart_2":        "双手比心2",
		"Heart_3":        "双手比心3",
		"Two":            "数字2",
		"Three":          "数字3",
		"Four":           "数字4",
		"Six":            "数字6",
		"Seven":          "数字7",
		"Eight":          "数字8",
		"Nine":           "数字9",
		"Rock":           "Rock",
		"Insult":         "竖中指",
		"Face":           "人脸",
	}
)

// GestureRes .
type GestureRes struct {
	Probability float64 `json:"probability"`
	Classname   string  `json:"classname"`
	Desc        string  `json:"desc"`
}

// 手势识别
func (f *Face) Gesture(imageBase64 string) (res []GestureRes, err error) {
	result := Reply{}
	err = f.PostForm(GESTURE_URL, url.Values{
		"image": {imageBase64},
	}, &result)
	bytes, err := json.Marshal(result.Result)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return
	}
	if len(res) == 0 {
		err = errors.New("识别结果为空")
		return
	}
	for k, v := range res {
		res[k].Desc = gestureMap[v.Classname]
	}
	return
}
