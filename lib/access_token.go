package lib

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	// ACCESS_TOKEN_URL .
	ACCESS_TOKEN_URL = "https://aip.baidubce.com/oauth/2.0/token"
)

// AccessToken .
type AccessToken struct {
	tokenErr

	ExpiresIn int64 `json:"expires_in"`

	AccessToken      string `json:"access_token"`
	ExpiresTimeStamp int64  `json:"expires_time_stamp"`

	appKey    string
	appSecret string

	// token持久化
	store Storage

	// 避免多次刷新token
	mu sync.Mutex

	//RefreshToken     string `json:"refresh_token"`
	//SessionKey       string `json:"session_key"`
	//Scope            string `json:"scope"`
	//SessionSecret    string `json:"session_secret"`
}

type tokenErr struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// NewToken .
func NewToken(appKey string, appSecret string, store Storage) *AccessToken {
	if store == nil {
		store = &FileStore{FilePath: "./access_token.json"}
	}

	token := &AccessToken{
		store: store,
	}
	token.store.Load(token)
	token.appKey = appKey
	token.appSecret = appSecret

	return token
}

// SetAccessToken .
func (t *AccessToken) SetAccessToken() (err error) {

	if time.Now().Unix() < t.ExpiresTimeStamp {
		return nil
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	if time.Now().Unix() < t.ExpiresTimeStamp {
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
	t.ExpiresTimeStamp = time.Now().Unix() + t.ExpiresIn

	return t.store.Store(t)

}
