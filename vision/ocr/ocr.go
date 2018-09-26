package ocr

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/kofj/baiduai-sdk-go"
	"github.com/kofj/baiduai-sdk-go/vision"
)

const (
	// OcrGeneralBasicURL 通用文字识别
	OcrGeneralBasicURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	// AccurateBasic 通用文字识别（高精度版）
	AccurateBasic = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
	// OcrGeneralURL 通用文字识别（含位置信息版）
	OcrGeneralURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
	// OcrAccurateURL 通用文字识别（含位置高精度版）
	OcrAccurateURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate"
	// OcrGeneralEnhancedURL 通用文字识别（含生僻字版
	OcrGeneralEnhancedURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_enhanced"

	// OcrWebimageURL 网络图片文字识别,用户向服务请求识别一些网络上背景复杂，特殊字体的文字。
	OcrWebimageURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"
	// OcrHandwritingURL 手写体识别
	OcrHandwritingURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/handwriting"
	// OcrIdcardURL 身份证识别
	OcrIdcardURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"
	// OcrBankcardURL 银行卡识别
	OcrBankcardURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"
	// OcrDrivingLicenseURL 驾驶证识别
	OcrDrivingLicenseURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"
	// OcrVehiclLicenseURL 行驶证识别
	OcrVehiclLicenseURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_license"
	// OcrLicensePlateURL 车牌识别
	OcrLicensePlateURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"
	// OcrBusinessLicenseURL 营业执照识别
	OcrBusinessLicenseURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/business_license"
)

func parseRequestParam(image *vision.Image, params ...RequestParam) (r *url.Values) {
	r = &url.Values{}
	for _, fn := range params {
		fn(r)
	}

	if image.URL != "" {
		r.Set("url", image.URL)
	} else {
		r.Set("image", image.Data)
	}
	return
}

// Client 客户端
type Client struct {
	sdk.Authorizer
}

// New 客户端实例化
func New(authorizer sdk.Authorizer) *Client {
	return &Client{
		authorizer,
	}
}

func (c *Client) doRequest(api string, img *vision.Image, r interface{}, params ...RequestParam) (err error) {
	token, err := c.Authorizer.Token()
	if err != nil {
		return
	}
	api += "?access_token=" + token

	hc := &http.Client{}

	resp, err := hc.PostForm(api, *parseRequestParam(img, params...))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// var i interface{}
	err = json.NewDecoder(resp.Body).Decode(r)
	return
}

// GeneralBasic ...
func (c *Client) GeneralBasic(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	err = c.doRequest(OcrGeneralBasicURL, img, resp, params...)
	return
}

// General ...
func (c *Client) General(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	err = c.doRequest(OcrGeneralURL, img, resp, params...)

	return
}

// Accurate ...
func (c *Client) Accurate(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	err = c.doRequest(OcrAccurateURL, img, resp, params...)

	return
}

// GeneralEnhanced ...
func (c *Client) GeneralEnhanced(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	err = c.doRequest(OcrGeneralEnhancedURL, img, resp, params...)

	return
}

// Webimage ...
func (c *Client) Webimage(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	err = c.doRequest(OcrWebimageURL, img, resp, params...)

	return
}

// Handwriting ...
func (c *Client) Handwriting(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	err = c.doRequest(OcrHandwritingURL, img, resp, params...)

	return
}

// Idcard ...
func (c *Client) Idcard(img *vision.Image, params ...RequestParam) (resp *IdcardResp, err error) {
	resp = &IdcardResp{}
	err = c.doRequest(OcrIdcardURL, img, resp, params...)

	return
}

// Bankcard ...
func (c *Client) Bankcard(img *vision.Image, params ...RequestParam) (resp *BankcardResp, err error) {
	resp = &BankcardResp{}
	err = c.doRequest(OcrBankcardURL, img, resp, params...)

	return
}

// DrivingLicense ...
func (c *Client) DrivingLicense(img *vision.Image, params ...RequestParam) (resp *DrivingLicenseResp, err error) {
	resp = &DrivingLicenseResp{}
	err = c.doRequest(OcrDrivingLicenseURL, img, resp, params...)

	return
}

// VehiclLicense ...
func (c *Client) VehiclLicense(img *vision.Image, params ...RequestParam) (resp *VehiclLicenseResp, err error) {
	resp = &VehiclLicenseResp{}
	err = c.doRequest(OcrVehiclLicenseURL, img, resp, params...)

	return
}

// LicensePlate ...
func (c *Client) LicensePlate(img *vision.Image, params ...RequestParam) (resp *LicensePlateResp, err error) {
	resp = &LicensePlateResp{}
	err = c.doRequest(OcrLicensePlateURL, img, resp, params...)

	return
}

// BusinessLicense ...
func (c *Client) BusinessLicense(img *vision.Image, params ...RequestParam) (resp *BusinessLicenseResp, err error) {
	resp = &BusinessLicenseResp{}
	err = c.doRequest(OcrBusinessLicenseURL, img, resp, params...)

	return
}
