package face

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

// Face .
type Face struct {
	Token *lib.AccessToken
}

// New .
func New(appKey string, appSecret string, store lib.Storage) *Face {
	return &Face{
		Token: lib.NewToken(appKey, appSecret, store),
	}
}

// PostJson .
func (f *Face) PostJson(url string, v []byte, res *Reply) error {
	if err := f.Token.SetAccessToken(); err != nil {
		return errors.WithStack(err)
	}
	resp, err := http.Post(url+"?access_token="+f.Token.AccessToken, "application/json", bytes.NewReader(v))
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

	if res.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("%d:%s", res.ErrorCode, res.ErrorMsg))
	}
	return nil
}

// PostForm .
func (f *Face) PostForm(url string, v url.Values, res *Reply) error {
	if err := f.Token.SetAccessToken(); err != nil {
		return errors.WithStack(err)
	}

	v.Add("access_token", f.Token.AccessToken)
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
	if res.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("%d:%s", res.ErrorCode, res.ErrorMsg))
	}
	return nil
}
