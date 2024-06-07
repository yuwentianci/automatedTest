package config

const (
	//Test = "https://api-future.biconomy.com"
	Test               = "https://future.biconomy.vip"
	DetailUrl          = Test + "/future/api/v1/detailV2?client=web"
	OpenPositionUrl    = Test + "/future/api/v1/private/position/openPositions?symbol="
	ChangeLvgUrl       = Test + "future/api/v1/private/position/changeLvg"
	RiskLimitUrl       = Test + "/future/api/v1/private/account/riskLimit"
	LvgUrl             = Test + "/future/api/v1/private/position/lvg?symbol="
	OpenOrderUrl       = Test + "/future/api/v1/private/order/list/openOrders?page_num=1&page_size=100"
	AsserUrl           = Test + "/future/api/v1/private/account/assets"
	AssetRecordUrl     = Test + "/future/api/v1/private/account/assetRecord?&pageSize=10"
	HistoryPositionUrl = Test + "/future/api/v1/private/position/historyPositions?page_size=10"
	HistoryOrderUrl    = Test + "/future/api/v1/private/order/historyOrders?page_size=10&symbol=&category=1,6&states=3,4,5"
	HistoryTradeUrl    = Test + "/future/api/v1/private/order/orderDeals?page_size=10"
	LiqHistoryOrderUrl = Test + "/future/api/v1/private/order/closeOrders?page_size=10&category=2"
	OpenOrdersUrl      = Test + "/future/api/v1/private/order/list/openOrders?page_num=1&page_size=100"
	Symbol             = "BTC_USDT"
)
