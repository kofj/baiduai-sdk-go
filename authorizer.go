package sdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// AuthURL Access Token 地址
const AuthURL = "https://aip.baidubce.com/oauth/2.0/token"

// AuthResp 认证 API 响应
type AuthResp struct {
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	SessionKey       string `json:"session_key"`
	AccessToken      string `json:"access_token"`
	SessionSecret    string `json:"session_secret"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// Authorizer 用于设置 AccessToken
// 可以通过 RESTful api的方式从百度方获取
// 有效期为一个月，可以存至数据库中然后从数据库中获取
type Authorizer interface {
	Auth() error
	Token() (string, error)
}

// DefaultAuthorizer 默认授权器
type DefaultAuthorizer struct {
	clientID     string
	clientSecret string
	accessToken  string
	expireAt     time.Time
	locker       sync.RWMutex
}

var _ Authorizer = &DefaultAuthorizer{}

// isExpired 检查 token 是否过期了
func (d *DefaultAuthorizer) isExpired() bool {
	return time.Now().After(d.expireAt)
}

// Auth 认证
func (d *DefaultAuthorizer) Auth() (err error) {
	c := &http.Client{}

	var params = url.Values{}
	params.Add("grant_type", "client_credentials")
	params.Add("client_id", d.clientID)
	params.Add("client_secret", d.clientSecret)

	resp, err := c.PostForm(AuthURL, params)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var data = &AuthResp{}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return
	}

	if data.Error != "" {
		return errors.New("API Error:" + data.Error)
	}

	var expire = data.ExpiresIn
	if expire == 0 {
		expire = 2592000
	}

	d.locker.Lock()
	d.accessToken = data.AccessToken
	d.expireAt = time.Now().Add(time.Duration(expire) * time.Second)
	d.locker.Unlock()

	return
}

// Token 获取 Token
func (d *DefaultAuthorizer) Token() (token string, err error) {
	if d.isExpired() || d.accessToken == "" {
		err = d.Auth()
		if err != nil {
			return
		}
	}

	return d.accessToken, nil
}

// NewAuthorizer 实例化认证器
func NewAuthorizer(APIKEY, APISECRET string) *DefaultAuthorizer {
	return &DefaultAuthorizer{
		clientID:     APIKEY,
		clientSecret: APISECRET,
	}
}
