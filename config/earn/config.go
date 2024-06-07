package earn

const (
	//earnUrl                = "http://47.245.59.223:84"
	//earnUrl                = "https://www.biconomy.vip"
	earnUrl                = "https://www.biconomy.com"
	AllEarnProductUrl      = earnUrl + "/api/activity/earn_saving/allearnterms?"
	BuyEarnUrl             = earnUrl + "/api/activity/earn_saving/order/create"
	MyEarnUrl              = earnUrl + "/api/activity/earn_saving/myearn?"
	MyEarnAssetsDetailsUrl = earnUrl + "/api/activity/earn_saving/instancedetail?instance_id="
	MyEarnAssetsUrl        = earnUrl + "/api/activity/earn_saving/overview"
	MyEarnHistoryUrl       = earnUrl + "/api/activity/earn_saving/earnSaving/list?"
	RedeemUrl              = earnUrl + "/api/activity/earn_saving/redeem/list?"
	ProfitUrl              = earnUrl + "/api/activity/earn_saving/profit/list?"
)
