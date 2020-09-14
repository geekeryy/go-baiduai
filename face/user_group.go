package face

import (
	"net/url"
)

// 添加用户组
func (f *Face) AddUserGroup(group_id string) (err error) {
	err = f.PostForm(USER_GROUP_ADD, url.Values{
		"group_id": {group_id},
	}, &Reply{})
	return
}

// 删除用户组
func (f *Face) DelUserGroup(group_id string) (err error) {
	err = f.PostForm(USER_GROUP_DEL, url.Values{
		"group_id": {group_id},
	}, &Reply{})
	return
}

// GroupList .
type GroupList struct {
	GroupIdList []string `json:"group_id_list"`
}

// 用户组列表
func (f *Face) ListUserGroup(start, length string) (res *GroupList, err error) {
	res = &GroupList{}
	if start == "" {
		start = "0"
	}
	if length == "" {
		length = "100"
	}
	err = f.PostForm(USER_GROUP_LIST, url.Values{
		"start":  {start},
		"length": {length},
	}, &Reply{
		Result: res,
	})
	return
}
