package body

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/comeonjy/go-baiduai/lib"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Interface .
type Interface interface{}

// Reply .
type Reply struct {
	BaseReply
	Interface
}

// BaseReply .
type BaseReply struct {
	ErrorCode int64  `json:"error_code"`
	LogId     int64  `json:"log_id"`
	ErrorMsg  string `json:"error_msg"`
}

// Image .
type Image struct {
	Token *lib.AccessToken
}

// New .
func New(appKey string, appSecret string) *Image {
	return &Image{
		Token: lib.NewToken(appKey, appSecret),
	}
}

// PostJson .
func (i *Image) PostJson(url string, v []byte, res *Reply) error {
	if err := i.Token.SetAccessToken(); err != nil {
		return errors.WithStack(err)
	}
	resp, err := http.Post(url+"?access_token="+i.Token.AccessToken, "application/json", bytes.NewReader(v))
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := json.Unmarshal(body, res); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// PostForm .
func (i *Image) PostForm(url string, v url.Values, res interface{}) error {
	if err := i.Token.SetAccessToken(); err != nil {
		return errors.WithStack(err)
	}

	v.Add("access_token", i.Token.AccessToken)
	resp, err := http.PostForm(url, v)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := json.Unmarshal(body, res); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (i *Image) checkErr(baseReply BaseReply, err error) error {
	if err != nil {
		return err
	}
	if baseReply.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("%d:%s", baseReply.ErrorCode, baseReply.ErrorMsg))
	}
	return nil
}
