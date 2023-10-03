package binanceP2PAPI

type PriceBody struct {
	Asset         string   `json:"asset"`
	Fiat          string   `json:"fiat"`
	MerchantCheck bool     `json:"merchantCheck"`
	Page          int64    `json:"page"`
	PayTypes      []string `json:"payTypes"`
	PublisherType *string  `json:"publisherType"`
	Rows          int64    `json:"rows"`
	TradeType     string   `json:"tradeType"`
}

type PriceResponse struct {
	Code          string      `json:"code"`
	Message       interface{} `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          []struct {
		Adv struct {
			AdvNo                 string      `json:"advNo"`
			Classify              string      `json:"classify"`
			TradeType             string      `json:"tradeType"`
			Asset                 string      `json:"asset"`
			FiatUnit              string      `json:"fiatUnit"`
			AdvStatus             interface{} `json:"advStatus"`
			PriceType             interface{} `json:"priceType"`
			PriceFloatingRatio    interface{} `json:"priceFloatingRatio"`
			RateFloatingRatio     interface{} `json:"rateFloatingRatio"`
			CurrencyRate          interface{} `json:"currencyRate"`
			Price                 string      `json:"price"`
			InitAmount            string      `json:"initAmount"`
			SurplusAmount         string      `json:"surplusAmount"`
			AmountAfterEditing    string      `json:"amountAfterEditing"`
			MaxSingleTransAmount  string      `json:"maxSingleTransAmount"`
			MinSingleTransAmount  string      `json:"minSingleTransAmount"`
			BuyerKycLimit         interface{} `json:"buyerKycLimit"`
			BuyerRegDaysLimit     interface{} `json:"buyerRegDaysLimit"`
			BuyerBtcPositionLimit interface{} `json:"buyerBtcPositionLimit"`
			Remarks               interface{} `json:"remarks"`
			AutoReplyMsg          string      `json:"autoReplyMsg"`
			PayTimeLimit          int         `json:"payTimeLimit"`
			TradeMethods          []struct {
				PayId                interface{} `json:"payId"`
				PayMethodId          string      `json:"payMethodId"`
				PayType              string      `json:"payType"`
				PayAccount           interface{} `json:"payAccount"`
				PayBank              interface{} `json:"payBank"`
				PaySubBank           interface{} `json:"paySubBank"`
				Identifier           string      `json:"identifier"`
				IconUrlColor         string      `json:"iconUrlColor"`
				TradeMethodName      string      `json:"tradeMethodName"`
				TradeMethodShortName string      `json:"tradeMethodShortName"`
				TradeMethodBgColor   string      `json:"tradeMethodBgColor"`
			} `json:"tradeMethods"`
			UserTradeCountFilterTime        interface{}   `json:"userTradeCountFilterTime"`
			UserBuyTradeCountMin            interface{}   `json:"userBuyTradeCountMin"`
			UserBuyTradeCountMax            interface{}   `json:"userBuyTradeCountMax"`
			UserSellTradeCountMin           interface{}   `json:"userSellTradeCountMin"`
			UserSellTradeCountMax           interface{}   `json:"userSellTradeCountMax"`
			UserAllTradeCountMin            interface{}   `json:"userAllTradeCountMin"`
			UserAllTradeCountMax            interface{}   `json:"userAllTradeCountMax"`
			UserTradeCompleteRateFilterTime interface{}   `json:"userTradeCompleteRateFilterTime"`
			UserTradeCompleteCountMin       interface{}   `json:"userTradeCompleteCountMin"`
			UserTradeCompleteRateMin        interface{}   `json:"userTradeCompleteRateMin"`
			UserTradeVolumeFilterTime       interface{}   `json:"userTradeVolumeFilterTime"`
			UserTradeType                   interface{}   `json:"userTradeType"`
			UserTradeVolumeMin              interface{}   `json:"userTradeVolumeMin"`
			UserTradeVolumeMax              interface{}   `json:"userTradeVolumeMax"`
			UserTradeVolumeAsset            interface{}   `json:"userTradeVolumeAsset"`
			CreateTime                      interface{}   `json:"createTime"`
			AdvUpdateTime                   interface{}   `json:"advUpdateTime"`
			FiatVo                          interface{}   `json:"fiatVo"`
			AssetVo                         interface{}   `json:"assetVo"`
			AdvVisibleRet                   interface{}   `json:"advVisibleRet"`
			AssetLogo                       interface{}   `json:"assetLogo"`
			AssetScale                      int           `json:"assetScale"`
			FiatScale                       int           `json:"fiatScale"`
			PriceScale                      int           `json:"priceScale"`
			FiatSymbol                      string        `json:"fiatSymbol"`
			IsTradable                      bool          `json:"isTradable"`
			DynamicMaxSingleTransAmount     string        `json:"dynamicMaxSingleTransAmount"`
			MinSingleTransQuantity          string        `json:"minSingleTransQuantity"`
			MaxSingleTransQuantity          string        `json:"maxSingleTransQuantity"`
			DynamicMaxSingleTransQuantity   string        `json:"dynamicMaxSingleTransQuantity"`
			TradableQuantity                string        `json:"tradableQuantity"`
			CommissionRate                  string        `json:"commissionRate"`
			TradeMethodCommissionRates      []interface{} `json:"tradeMethodCommissionRates"`
			LaunchCountry                   interface{}   `json:"launchCountry"`
		} `json:"adv"`
		Advertiser struct {
			UserNo           string        `json:"userNo"`
			RealName         interface{}   `json:"realName"`
			NickName         string        `json:"nickName"`
			Margin           interface{}   `json:"margin"`
			MarginUnit       interface{}   `json:"marginUnit"`
			OrderCount       interface{}   `json:"orderCount"`
			MonthOrderCount  int           `json:"monthOrderCount"`
			MonthFinishRate  float64       `json:"monthFinishRate"`
			AdvConfirmTime   int           `json:"advConfirmTime"`
			Email            interface{}   `json:"email"`
			RegistrationTime interface{}   `json:"registrationTime"`
			Mobile           interface{}   `json:"mobile"`
			UserType         string        `json:"userType"`
			TagIconUrls      []interface{} `json:"tagIconUrls"`
			UserGrade        int           `json:"userGrade"`
			UserIdentity     string        `json:"userIdentity"`
			ProMerchant      interface{}   `json:"proMerchant"`
			IsBlocked        interface{}   `json:"isBlocked"`
		} `json:"advertiser"`
	} `json:"data"`
	Total   int  `json:"total"`
	Success bool `json:"success"`
}
