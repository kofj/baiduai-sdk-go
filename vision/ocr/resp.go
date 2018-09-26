package ocr

// GeneralResp 通用响应
type GeneralResp struct {
	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`

	Direction   int   `json:"direction,omitempty"`
	Language    int   `json:"language,omitempty"`
	LogID       int64 `json:"log_id,omitempty"`
	WordsResult []struct {
		Location struct {
			Height int `json:"height,omitempty"`
			Left   int `json:"left,omitempty"`
			Top    int `json:"top,omitempty"`
			Width  int `json:"width,omitempty"`
		} `json:"location,omitempty"`
		Words string `json:"words,omitempty"`
		Chars []struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Char string `json:"char,omitempty"`
		} `json:"chars,omitempty"`
	} `json:"words_result,omitempty"`
	WordsResultNum int `json:"words_result_num,omitempty"`
}

// IdcardResp 身份证响应
type IdcardResp struct {
	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`

	LogID       int64  `json:"log_id,omitempty"`
	Direction   int    `json:"direction,omitempty"`
	ImageStatus string `json:"image_status,omitempty"`
	IdcardType  string `json:"idcard_type,omitempty"`
	EditTool    string `json:"edit_tool,omitempty"`
	WordsResult struct {
		Address struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"住址,omitempty"`
		ID struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"公民身份号码,omitempty"`
		Birthday struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"出生,omitempty"`
		Name struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"姓名,omitempty"`
		Gender struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"性别,omitempty"`
		Nationality struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"民族,omitempty"`
	} `json:"words_result,omitempty"`
	WordsResultNum int `json:"words_result_num,omitempty"`
}

// BankcardResp 银行卡响应
type BankcardResp struct {
	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`

	LogID  int `json:"log_id,omitempty"`
	Result struct {
		BankCardNumber string `json:"bank_card_number,omitempty"`
		BankName       string `json:"bank_name,omitempty"`
		BankCardType   int    `json:"bank_card_type,omitempty"`
	} `json:"result,omitempty"`
}

// DrivingLicenseResp 驾照响应
type DrivingLicenseResp struct {
	Errno int    `json:"errno,omitempty"`
	Msg   string `json:"msg,omitempty"`
	Data  struct {
		WordsResultNum int `json:"words_result_num,omitempty"`
		WordsResult    struct {
			No struct {
				Words string `json:"words,omitempty"`
			} `json:"证号,omitempty"`
			ExpireAt struct {
				Words string `json:"words,omitempty"`
			} `json:"有效期限,omitempty"`
			CarType struct {
				Words string `json:"words,omitempty"`
			} `json:"准驾车型,omitempty"`
			SignAt struct {
				Words string `json:"words,omitempty"`
			} `json:"有效起始日期,omitempty"`
			Address struct {
				Words string `json:"words,omitempty"`
			} `json:"住址,omitempty"`
			Name struct {
				Words string `json:"words,omitempty"`
			} `json:"姓名,omitempty"`
			Nationality struct {
				Words string `json:"words,omitempty"`
			} `json:"国籍,omitempty"`
			Birthday struct {
				Words string `json:"words,omitempty"`
			} `json:"出生日期,omitempty"`
			Gender struct {
				Words string `json:"words,omitempty"`
			} `json:"性别,omitempty"`
			GetAt struct {
				Words string `json:"words,omitempty"`
			} `json:"初次领证日期,omitempty"`
		} `json:"words_result,omitempty"`
	} `json:"data,omitempty"`
}

// VehiclLicenseResp 行驶证响应
type VehiclLicenseResp struct {
	Errno int    `json:"errno,omitempty"`
	Msg   string `json:"msg,omitempty"`
	Data  struct {
		WordsResultNum int `json:"words_result_num,omitempty"`
		WordsResult    struct {
			CarVendor struct {
				Words string `json:"words,omitempty"`
			} `json:"品牌型号,omitempty"`
			SignAt struct {
				Words string `json:"words,omitempty"`
			} `json:"发证日期,omitempty"`
			Purpose struct {
				Words string `json:"words,omitempty"`
			} `json:"使用性质,omitempty"`
			EngineNo struct {
				Words string `json:"words,omitempty"`
			} `json:"发动机号码,omitempty"`
			PlateNo struct {
				Words string `json:"words,omitempty"`
			} `json:"号牌号码,omitempty"`
			Owner struct {
				Words string `json:"words,omitempty"`
			} `json:"所有人,omitempty"`
			Address struct {
				Words string `json:"words,omitempty"`
			} `json:"住址,omitempty"`
			RegAt struct {
				Words string `json:"words,omitempty"`
			} `json:"注册日期,omitempty"`
			CarID struct {
				Words string `json:"words,omitempty"`
			} `json:"车辆识别代号,omitempty"`
			CarType struct {
				Words string `json:"words,omitempty"`
			} `json:"车辆类型,omitempty"`
		} `json:"words_result,omitempty"`
	} `json:"data,omitempty"`
}

// LicensePlateResp 车牌响应
type LicensePlateResp struct {
	Errno int    `json:"errno,omitempty"`
	Msg   string `json:"msg,omitempty"`
	Data  struct {
		LogID       string `json:"log_id,omitempty"`
		WordsResult struct {
			Color            string `json:"color,omitempty"`
			Number           string `json:"number,omitempty"`
			Probability      []int  `json:"probability,omitempty"`
			VertexesLocation []struct {
				Y int `json:"y,omitempty"`
				X int `json:"x,omitempty"`
			} `json:"vertexes_location,omitempty"`
		} `json:"words_result,omitempty"`
	} `json:"data,omitempty"`
}

// BusinessLicenseResp 营业执照响应
type BusinessLicenseResp struct {
	LogID       int `json:"log_id,omitempty"`
	WordsResult struct {
		Name struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"单位名称,omitempty"`
		Legal struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"法人,omitempty"`
		Address struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"地址,omitempty"`
		Period struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"有效期,omitempty"`
		ID struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"证件编号,omitempty"`
		SocialCreditCode struct {
			Location struct {
				Left   int `json:"left,omitempty"`
				Top    int `json:"top,omitempty"`
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"location,omitempty"`
			Words string `json:"words,omitempty"`
		} `json:"社会信用代码,omitempty"`
	} `json:"words_result,omitempty"`
	WordsResultNum int `json:"words_result_num,omitempty"`
}
