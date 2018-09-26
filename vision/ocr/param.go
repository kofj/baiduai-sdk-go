package ocr

import (
	"net/url"
)

// RequestParam OCR 请求参数
type RequestParam func(*url.Values)

// LangType 识别语言类型，默认为CHN_ENG。可选值包括：
// - CHN_ENG：中英文混合；
// - ENG：英文；
// - POR：葡萄牙语；
// - FRE：法语；
// - GER：德语；
// - ITA：意大利语；
// - SPA：西班牙语；
// - RUS：俄语；
// - JAP：日语；
// - KOR：韩语
func LangType(lang string) RequestParam {
	options := []string{
		"CHN_ENG",
		"ENG",
		"POR",
		"FRE",
		"GER",
		"ITA",
		"SPA",
		"RUS",
		"JAP",
		"KOR",
	}

	illegal := true
	for _, v := range options {
		if v == lang {
			illegal = false
			break
		}
	}

	if illegal {
		lang = "CHN_ENG"
	}
	return func(m *url.Values) {
		m.Set("language_type", lang)
	}
}

// DetectDirection 是否检测图像朝向，默认不检测，即：false。朝向是指输入图像是正常方向、逆时针旋转90/180/270度。可选值包括:
// - true：检测朝向；
// - false：不检测朝向。
func DetectDirection() RequestParam {
	return func(m *url.Values) {
		m.Set("detect_direction", "true")
	}
}

// DetectLanguage 是否检测语言，默认不检测。
// 当前支持（中文、英语、日语、韩语）
func DetectLanguage() RequestParam {
	return func(m *url.Values) {
		m.Set("detect_language", "true")
	}
}

// WithProbability 是否返回识别结果中每一行的置信度
func WithProbability() RequestParam {
	return func(m *url.Values) {
		m.Set("probability", "true")
	}
}

// RecognizeGranularity 是否定位单字符位置
// - big：不定位单字符位置，默认值；
// - small：定位单字符位置
func RecognizeGranularity() RequestParam {
	return func(m *url.Values) {
		m.Set("recognize_granularity", "small")
	}
}

// WithVertexesLocation 是否返回文字外接多边形顶点位置，不支持单字位置。默认为false
func WithVertexesLocation() RequestParam {
	return func(m *url.Values) {
		m.Set("vertexes_location", "true")
	}
}

// IDCardSideFront 身份证含照片的一面
func IDCardSideFront() RequestParam {
	return func(m *url.Values) {
		m.Set("id_card_side", "front")
	}
}

// IDCardSideBack 身份证带国徽的一面
func IDCardSideBack() RequestParam {
	return func(m *url.Values) {
		m.Set("id_card_side", "back")
	}
}

// DetectRisk 是否开启身份证风险类型(身份证复印件、临时身份证、身份证翻拍、修改过的身份证)功能，默认不开启，即：false。
// true-开启；false-不开启
func DetectRisk() RequestParam {
	return func(m *url.Values) {
		m.Set("back", "true")
	}
}

// Accuracy normal 使用快速服务，1200ms左右时延；缺省或其它值使用高精度服务，1600ms左右时延
func Accuracy(opt string) RequestParam {
	if opt != "normal" && opt != "high" {
		opt = "normal"
	}
	return func(m *url.Values) {
		m.Set("back", "normal")
	}
}
