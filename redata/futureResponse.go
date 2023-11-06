package redata

// MarginAndLeverage 保证金模式及杠杆
type MarginAndLeverage struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Result  []MarginAndLeverageResult `json:"result"`
}

// MarginAndLeverageResult 保证金模式及杠杆result数据
type MarginAndLeverageResult struct {
	Symbol            string `json:"symbol"`
	ID                int    `json:"id"`
	MemberID          int    `json:"member_id"`
	FutureSymbolsID   int    `json:"future_symbols_id"`
	USDTLongLeverage  int    `json:"usdt_long_leverage"`
	USDTShortLeverage int    `json:"usdt_short_leverage"`
	CoinLongLeverage  int    `json:"coin_long_leverage"`
	CoinShortLeverage int    `json:"coin_short_leverage"`
	LimitOrMarket     string `json:"limit_or_market"`
	OpenOrClose       string `json:"open_or_close"`
	Pattern           int    `json:"pattern"`
	ForbiddenDeal     string `json:"forbidden_deal"`
	ForbiddenTransfer string `json:"forbidden_transfer"`
}

// SymbolThumb 币对数据
type SymbolThumb struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  []SymbolResult `json:"result"`
}

// SymbolResult 币对的result数据
type SymbolResult struct {
	Symbol                string `json:"symbol"`
	ID                    int    `json:"id"`
	MakerFee              string `json:"maker_fee"`
	TakerFee              string `json:"taker_fee"`
	MaintenanceMarginRate string `json:"maintenance_margin_rate"`
}

// OrdersPositions 下单
type OrdersPositions struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// OpenOrders 当前委托单
type OpenOrders struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Result  OpenOrdersResult `json:"result"`
}

// OpenOrdersResult 当前委托单result数据
type OpenOrdersResult struct {
	Data []OpenOrdersData `json:"data"`
}

// OpenOrdersData 当前委托单result的data数据
type OpenOrdersData struct {
	Amount          string  `json:"amount"`
	AmountPrecision int     `json:"amount_precision"`
	CTime           float64 `json:"ctime"`
	DealFee         string  `json:"deal_fee"`
	DealMoney       string  `json:"deal_money"`
	DealStock       string  `json:"deal_stock"`
	ID              int     `json:"id"`
	Left            string  `json:"left"`
	MakerFee        string  `json:"maker_fee"`
	Market          string  `json:"market"`
	MTime           float64 `json:"mtime"`
	Price           string  `json:"price"`
	PricePrecision  int     `json:"price_precision"`
	Side            int     `json:"side"`
	OperType        int     `json:"oper_type"`
	Source          string  `json:"source"`
	TakerFee        string  `json:"taker_fee"`
	Type            int     `json:"type"`
	User            int     `json:"user"`
	IsBlast         int     `json:"isblast"`
	Trigger         string  `json:"trigger"`
}

// Positions 仓位数据
type Positions struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Result  []PositionsResult `json:"result"`
}

// PositionsResult 仓位的result数据
type PositionsResult struct {
	PricePrecision int    `json:"price_precision"`
	Frozen         string `json:"frozen"`
	Leverage       string `json:"leverage"`
	Market         string `json:"market"`
	Pattern        int    `json:"pattern"`
	Position       string `json:"position"`
	Price          string `json:"price"`
	Principal      string `json:"principal"`
	Side           int    `json:"side"`
	ShareNumber    string `json:"share_number"`
}

// PositionsInfo 仓位有用的数据
type PositionsInfo struct {
	Positions []PositionsDetailsInfo
}

// PositionsDetailsInfo 仓位有用的详细信息
type PositionsDetailsInfo struct {
	Market    string `json:"market"`
	Pattern   int    `json:"pattern"`
	Side      int    `json:"side"`
	Position  string `json:"position"`
	Frozen    string `json:"frozen"`
	Price     string `json:"price"`
	Principal string `json:"principal"`
}

// FuturesBalance 合约资产数据
type FuturesBalance struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Result  BalanceResult `json:"result"`
}

// BalanceResult 合约资产的Result数据
type BalanceResult struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

// LeveragePattern 币对杠杆和模式
type LeveragePattern struct {
	ID       int    `json:"id"`
	Symbol   string `json:"symbol"`
	Leverage int    `json:"leverage"`
	Pattern  int    `json:"pattern"`
}
