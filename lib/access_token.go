package lib

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// ACCESS_TOKEN_URL .
	ACCESS_TOKEN_URL = "https://aip.baidubce.com/oauth/2.0/token"
)

// AccessToken .
type AccessToken struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int64  `json:"expires_in"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	SessionSecret    string `json:"session_secret"`
	UpdateTimeStamp  int64  `json:"update_time_stamp"`
	appKey           string
	appSecret        string
}

// NewToken .
func NewToken(appKey string, appSecret string) *AccessToken {
	return &AccessToken{
		appKey:    appKey,
		appSecret: appSecret,
	}
}

// SetAccessToken .
func (t *AccessToken) SetAccessToken() (err error) {

	if time.Now().Unix()-t.UpdateTimeStamp < t.ExpiresIn {
		return nil
	}
	resp, err := http.PostForm(ACCESS_TOKEN_URL, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {t.appKey},
		"client_secret": {t.appSecret},
	})
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.Unmarshal(bytes, t); err != nil {
		return
	}
	if t.Error != "" {
		return errors.New(fmt.Sprintf("%s:%s", t.Error, t.ErrorDescription))
	}
	t.UpdateTimeStamp = time.Now().Unix()

	return nil
}
