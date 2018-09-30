package ocr

import (
	"encoding/json"
	"io/ioutil"
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

func (c *Client) doRequest(api string, img *vision.Image, r interface{}, params ...RequestParam) (buf json.RawMessage, err error) {
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

	buf, err = ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(buf, r)
	return
}

// GeneralBasic ...
func (c *Client) GeneralBasic(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	_, err = c.doRequest(OcrGeneralBasicURL, img, resp, params...)
	return
}

// GeneralBasicWithRaw 只解析错误信息
func (c *Client) GeneralBasicWithRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp, raw json.RawMessage, err error) {
	resp = &ErrorResp{}
	raw, err = c.doRequest(OcrGeneralBasicURL, img, resp, params...)
	return
}

// General ...
func (c *Client) General(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	_, err = c.doRequest(OcrGeneralURL, img, resp, params...)
	return
}

// GeneralRaw 只解析错误信息
func (c *Client) GeneralRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp, raw json.RawMessage, err error) {
	resp = &ErrorResp{}
	raw, err = c.doRequest(OcrGeneralURL, img, resp, params...)

	return
}

// Accurate ...
func (c *Client) Accurate(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	_, err = c.doRequest(OcrAccurateURL, img, resp, params...)

	return
}

// AccurateRaw 只解析错误信息
func (c *Client) AccurateRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp, raw json.RawMessage, err error) {
	resp = &ErrorResp{}
	raw, err = c.doRequest(OcrAccurateURL, img, resp, params...)

	return
}

// GeneralEnhanced ...
func (c *Client) GeneralEnhanced(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	_, err = c.doRequest(OcrGeneralEnhancedURL, img, resp, params...)

	return
}

// GeneralEnhancedRaw 只解析错误信息
func (c *Client) GeneralEnhancedRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp, raw json.RawMessage, err error) {
	resp = &ErrorResp{}
	raw, err = c.doRequest(OcrGeneralEnhancedURL, img, resp, params...)

	return
}

// Webimage ...
func (c *Client) Webimage(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	_, err = c.doRequest(OcrWebimageURL, img, resp, params...)

	return
}

// WebimageRaw 只解析错误信息
func (c *Client) WebimageRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp, raw json.RawMessage, err error) {
	resp = &ErrorResp{}
	raw, err = c.doRequest(OcrWebimageURL, img, resp, params...)

	return
}

// Handwriting ...
func (c *Client) Handwriting(img *vision.Image, params ...RequestParam) (resp *GeneralResp, err error) {
	resp = &GeneralResp{}
	_, err = c.doRequest(OcrHandwritingURL, img, resp, params...)

	return
}

// HandwritingRaw 只解析错误信息
func (c *Client) HandwritingRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp, raw json.RawMessage, err error) {
	resp = &ErrorResp{}
	raw, err = c.doRequest(OcrHandwritingURL, img, resp, params...)

	return
}

// Idcard ...
func (c *Client) Idcard(img *vision.Image, params ...RequestParam) (resp *IdcardResp, err error) {
	resp = &IdcardResp{}
	_, err = c.doRequest(OcrIdcardURL, img, resp, params...)

	return
}

// IdcardRaw 只解析错误信息
func (c *Client) IdcardRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp, raw json.RawMessage, err error) {
	resp = &ErrorResp{}
	raw, err = c.doRequest(OcrIdcardURL, img, resp, params...)

	return
}

// Bankcard ...
func (c *Client) Bankcard(img *vision.Image, params ...RequestParam) (resp *BankcardResp, err error) {
	resp = &BankcardResp{}
	_, err = c.doRequest(OcrBankcardURL, img, resp, params...)

	return
}

// BankcardRaw 只解析错误信息
func (c *Client) BankcardRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp, raw json.RawMessage, err error) {
	resp = &ErrorResp{}
	raw, err = c.doRequest(OcrBankcardURL, img, resp, params...)

	return
}

// DrivingLicense ...
func (c *Client) DrivingLicense(img *vision.Image, params ...RequestParam) (resp *DrivingLicenseResp, err error) {
	resp = &DrivingLicenseResp{}
	_, err = c.doRequest(OcrDrivingLicenseURL, img, resp, params...)

	return
}

// DrivingLicenseRaw 只解析错误信息，ErrorResp2
func (c *Client) DrivingLicenseRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp2, raw json.RawMessage, err error) {
	resp = &ErrorResp2{}
	raw, err = c.doRequest(OcrDrivingLicenseURL, img, resp, params...)

	return
}

// VehiclLicense ...
func (c *Client) VehiclLicense(img *vision.Image, params ...RequestParam) (resp *VehiclLicenseResp, err error) {
	resp = &VehiclLicenseResp{}
	_, err = c.doRequest(OcrVehiclLicenseURL, img, resp, params...)

	return
}

// VehiclLicenseRaw 只解析错误信息，ErrorResp2
func (c *Client) VehiclLicenseRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp2, raw json.RawMessage, err error) {
	resp = &ErrorResp2{}
	raw, err = c.doRequest(OcrVehiclLicenseURL, img, resp, params...)

	return
}

// LicensePlate ...
func (c *Client) LicensePlate(img *vision.Image, params ...RequestParam) (resp *LicensePlateResp, err error) {
	resp = &LicensePlateResp{}
	_, err = c.doRequest(OcrLicensePlateURL, img, resp, params...)

	return
}

// LicensePlateRaw 只解析错误信息，ErrorResp2
func (c *Client) LicensePlateRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp2, raw json.RawMessage, err error) {
	resp = &ErrorResp2{}
	raw, err = c.doRequest(OcrLicensePlateURL, img, resp, params...)

	return
}

// BusinessLicense ...
func (c *Client) BusinessLicense(img *vision.Image, params ...RequestParam) (resp *BusinessLicenseResp, err error) {
	resp = &BusinessLicenseResp{}
	_, err = c.doRequest(OcrBusinessLicenseURL, img, resp, params...)

	return
}

// BusinessLicenseRaw 只解析错误信息，ErrorResp2
func (c *Client) BusinessLicenseRaw(img *vision.Image, params ...RequestParam) (resp *ErrorResp2, raw json.RawMessage, err error) {
	resp = &ErrorResp2{}
	raw, err = c.doRequest(OcrBusinessLicenseURL, img, resp, params...)

	return
}
