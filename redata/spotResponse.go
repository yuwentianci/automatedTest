package redata

// SpotOrderResult 现货创建订单响应result
type SpotOrderResult struct {
	Amount    string  `json:"amount"`
	Ctime     float64 `json:"ctime"`
	DealFee   string  `json:"deal_fee"`
	DealMoney string  `json:"deal_money"`
	DealStock string  `json:"deal_stock"`
	ID        int64   `json:"id"`
	Left      string  `json:"left"`
	MakerFee  string  `json:"maker_fee"`
	Market    string  `json:"market"`
	Mtime     float64 `json:"mtime"`
	Price     string  `json:"price"`
	Side      int     `json:"side"`
	Source    string  `json:"source"`
	TakerFee  string  `json:"taker_fee"`
	Type      int     `json:"type"`
	User      int64   `json:"user"`
}

type ResponseSpotOrderData struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Result  *SpotOrderResult `json:"result"`
}

type SpotTradeResultData struct {
	ID                  int64   `json:"id"`
	Market              string  `json:"market"`
	Type                int     `json:"type"`
	Side                int     `json:"side"`
	Filled              string  `json:"filled"`
	Total               string  `json:"total"`
	Price               string  `json:"price"`
	Time                string  `json:"time"`
	Ftime               float64 `json:"ftime"`
	Ctime               float64 `json:"ctime"`
	DealFee             string  `json:"deal_fee"`
	DealStock           string  `json:"deal_stock"`
	DealMoney           string  `json:"deal_money"`
	Status              int     `json:"status"`
	Amount              string  `json:"amount"`
	MakerFee            string  `json:"maker_fee"`
	TakerFee            string  `json:"taker_fee"`
	Role                string  `json:"role"`
	BaseAssetPrecision  int     `json:"base_asset_precision"`
	QuoteAssetPrecision int     `json:"quote_asset_precision"`
	QuoteAsset          string  `json:"quote_asset"`
	BaseAsset           string  `json:"base_asset"`
}

type SpotTradeResult struct {
	Data   []SpotTradeResultData `json:"data"`
	Total  int                   `json:"total"`
	Offset int                   `json:"offset"`
	Limit  int                   `json:"limit"`
}

type ResponseSpotTradeData struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Result  SpotTradeResult `json:"result"`
}
