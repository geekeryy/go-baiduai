package face

/**
 * 人脸库管理
 * API文档 https://cloud.baidu.com/doc/FACE/s/Uk25rdd5p
 */

import (
	"net/url"
)

// Reply .
type Reply struct {
	BaseReply
	Result interface{} `json:"result"`
}

// BaseReply .
type BaseReply struct {
	ErrorCode int64  `json:"error_code"`
	LogId     int64  `json:"log_id"`
	ErrorMsg  string `json:"error_msg"`
}

// Register .
type Register struct {
	Image           string
	ImageType       string
	GroupId         string
	UserId          string
	UserInfo        string
	QualityControl  string
	LivenessControl string
	ActionType      string
}

// RegRes .
type RegRes struct {
	FaceToken string `json:"face_token"`
	Location  struct {
		Left     float64 `json:"left"`
		Top      float64 `json:"top"`
		Width    float64 `json:"width"`
		Height   float64 `json:"height"`
		Rotation float64 `json:"rotation"`
	} `json:"location"`
}

// 人脸注册
func (f *Face) AddUser(register *Register) (res *RegRes, err error) {
	res = &RegRes{}
	if register.QualityControl == "" {
		register.QualityControl = "NONE"
	}
	if register.LivenessControl == "" {
		register.LivenessControl = "NONE"
	}
	if register.ActionType == "" {
		register.ActionType = "APPEND"
	}
	err = f.PostForm(USER_ADD, url.Values{
		"image":            {register.Image},
		"image_type":       {register.ImageType},
		"group_id":         {register.GroupId},
		"user_id":          {register.UserId},
		"user_info":        {register.UserId},
		"quality_control":  {register.QualityControl},
		"liveness_control": {register.LivenessControl},
		"action_type":      {register.ActionType},
	}, &Reply{
		Result: res,
	})
	return
}

// 人脸更新
func (f *Face) UpdateUser(register *Register) (res *RegRes, err error) {
	res = &RegRes{}
	if register.QualityControl == "" {
		register.QualityControl = "NONE"
	}
	if register.LivenessControl == "" {
		register.LivenessControl = "NONE"
	}
	if register.ActionType == "" {
		register.ActionType = "UPDATE"
	}
	err = f.PostForm(USER_UPDATE, url.Values{
		"image":            {register.Image},
		"image_type":       {register.ImageType},
		"group_id":         {register.GroupId},
		"user_id":          {register.UserId},
		"user_info":        {register.UserId},
		"quality_control":  {register.QualityControl},
		"liveness_control": {register.LivenessControl},
		"action_type":      {register.ActionType},
	}, &Reply{
		Result: res,
	})
	return
}

// 删除人脸
func (f *Face) DelUserFace(userId, groupId, faceToken string) (err error) {
	err = f.PostForm(USER_FACE_DEL, url.Values{
		"user_id":    {userId},
		"group_id":   {groupId},
		"face_token": {faceToken},
	}, &Reply{})
	return
}

// User .
type User struct {
	UserList []struct {
		UserInfo string `json:"user_info"`
		GroupId  string `json:"group_id"`
	} `json:"user_list"`
}

// 用户信息查询
func (f *Face) GetUser(userId, groupId string) (res *User, err error) {
	res = &User{}
	err = f.PostForm(USER_GET, url.Values{
		"user_id":  {userId},
		"group_id": {groupId},
	}, &Reply{
		Result: res,
	})
	return
}

// UserFaceList .
type UserFaceList struct {
	FaceList []struct {
		FaceToken string `json:"face_token"`
		Ctime     string `json:"ctime"`
	} `json:"face_list"`
}

// 人脸列表
func (f *Face) ListUserFace(userId, groupId string) (res *UserFaceList, err error) {
	res = &UserFaceList{}
	err = f.PostForm(USER_FACE_LIST, url.Values{
		"user_id":  {userId},
		"group_id": {groupId},
	}, &Reply{
		Result: res,
	})
	return
}

// UserList .
type UserList struct {
	UserIdList []string `json:"user_id_list"`
}

// 用户列表
func (f *Face) ListUser(groupId, start, length string) (res *UserList, err error) {
	res = &UserList{}
	if start == "" {
		start = "0"
	}
	if length == "" {
		length = "100"
	}
	err = f.PostForm(USER_LIST, url.Values{
		"group_id": {groupId},
		"start":    {start},
		"length":   {length},
	}, &Reply{
		Result: res,
	})
	return
}

// 复制用户
func (f *Face) CopyUser(user_id, src_group_id, dst_group_id string) (err error) {
	err = f.PostForm(USER_COPY, url.Values{
		"user_id":      {user_id},
		"src_group_id": {src_group_id},
		"dst_group_id": {dst_group_id},
	}, &Reply{})
	return
}
