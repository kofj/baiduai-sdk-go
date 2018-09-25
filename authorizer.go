package sdk

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/mozillazg/request"
)

// AuthURL Access Token 地址
const AuthURL = "https://aip.baidubce.com/oauth/2.0/token"

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

	req := request.NewRequest(c)
	req.Data = map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     d.clientID,
		"client_secret": d.clientSecret,
	}
	resp, err := req.Post(AuthURL)
	if err != nil {
		return
	}

	data, err := resp.Json()
	if err != nil {
		return
	}

	if data.Get("error").MustString() != "" {
		return errors.New("API Error:" + data.Get("error").MustString())
	}

	var expire = data.Get("expires_in").MustInt()
	if expire == 0 {
		expire = 2592000
	}

	d.locker.Lock()
	d.accessToken = data.Get("access_token").MustString()
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
