package config

const (
	Test            = "https://api-future.biconomy.com"
	DetailUrl       = Test + "/future/api/v1/detailV2?client=web"
	OpenPositionUrl = Test + "/future/api/v1/private/position/openPositions?symbol="
	ChangeLvgUrl    = Test + "future/api/v1/private/position/changeLvg"
	RiskLimitUrl    = Test + "/future/api/v1/private/account/riskLimit"
	LvgUrl          = Test + "/future/api/v1/private/position/lvg?symbol="
	OpenOrderUrl    = Test + "/future/api/v1/private/order/list/openOrders?page_num=1&page_size=100"
	AsserUrl        = Test + "/future/api/v1/private/account/assets"
	AssetRecordUrl  = Test + "/future/api/v1/private/account/assetRecord?symbol=&pageSize=10"

	Symbol = "BTC_USDT"
)
