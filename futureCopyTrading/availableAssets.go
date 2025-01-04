package futureCopyTrading

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"myapp/config/futureCopyTrading"
	"myapp/function"
	"myapp/future"
)

// OpenAfterAsset 开仓后资产
func OpenAfterAsset(assetDec decimal.Decimal, position []*GenDanPositionDetailsInfo, symbol []*future.SymbolInfo) (error, decimal.Decimal) {
	Asset := decimal.Zero
	openLongOrderMargin := decimal.Zero
	openShortOrderMargin := decimal.Zero
	for _, positions := range position {
		// 查找对应的交易对详情
		var Symbols *future.SymbolInfo
		for _, symbols := range symbol {
			if positions.ContractName == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}
		imr := future.One.Div(positions.Leverage)

		if positions.PositionType == 1 {
			openLongOrderMargin = openLongOrderMargin.Add(entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(imr.Add(Symbols.Tfr.Mul(future.Two))))
			longOrderMargin := entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(imr.Add(Symbols.Tfr.Mul(future.Two)))
			openLongPositionMargin := entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(imr.Add(Symbols.Tfr))
			fmt.Println("币对名:", positions.ContractName, "多仓仓位ID:", positions.PositionID, "仓位保证金:", openLongPositionMargin, "多仓总保证金:", openLongOrderMargin, "当前仓位保证金:", longOrderMargin)
		} else if positions.PositionType == 2 {
			mmr := decimal.NewFromFloat(0.008)
			openShortOrderMargin = openShortOrderMargin.Add(entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul((imr.Add(Symbols.Tfr)).Mul(future.One.Add(Symbols.Tfr)).Add(Symbols.Tfr.Mul(future.One.Sub(mmr)))))
			longOrderMargin := entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul((imr.Add(Symbols.Tfr)).Mul(future.One.Add(Symbols.Tfr)).Add(Symbols.Tfr.Mul(future.One.Sub(mmr))))
			openShortPositionMargin := entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(imr).Mul(future.One.Add(Symbols.Tfr)).Add(entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(Symbols.Tfr).Mul(future.One.Sub(mmr)))
			fmt.Println("币对名:", positions.ContractName, "空仓仓位ID:", positions.PositionID, "仓位保证金:", openShortPositionMargin, "空仓总保证金:", openShortOrderMargin, "维持保证金率:", mmr, "当前仓位保证金:", longOrderMargin)
		}
		Asset = assetDec.Sub(openLongOrderMargin).Sub(openShortOrderMargin)
	}

	return nil, Asset
}

// CloseAfterAsset 平仓后资产
func CloseAfterAsset(assetDec decimal.Decimal, position []*PositionInfo, symbol []*future.SymbolInfo) (error, decimal.Decimal) {
	Asset := decimal.Zero
	for _, positions := range position {
		// 查找对应的交易对详情
		var Symbols *future.SymbolInfo
		for _, symbols := range symbol {
			if positions.ContractName == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}
		closePriceDec, err := decimal.NewFromString(positions.ClosePrice)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}
		shareProfitDec, err := decimal.NewFromString(positions.WithHolding)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}
		imr := future.One.Div(positions.Leverage)

		if positions.PositionType == 1 {
			// 平仓后资产为：asset + 平仓后返还仓位保证金 + 平仓盈亏 - 平仓手续费
			positionMargin := entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(imr.Add(Symbols.Tfr))
			closePNL := closePriceDec.Sub(entryPriceDec).Mul(positions.CloseVolume).Mul(Symbols.Fv)
			openFee := entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr)
			closeFee := closePriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr)
			Asset = assetDec.Add(positionMargin).Add(closePNL).Sub(closeFee).Sub(shareProfitDec)
			positionPNL := closePNL.Sub(openFee).Sub(closeFee)
			fmt.Println("仓位保证金:", positionMargin, "平仓盈亏:", closePNL, "仓位盈亏:", positionPNL, "平仓手续费:", closeFee)
		} else if positions.PositionType == 2 {
			mmr := decimal.NewFromFloat(0.008)
			// 平仓后资产为：asset + 平仓后返还仓位保证金 + 平仓盈亏 - 平仓手续费
			positionMargin := entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(imr).Mul(future.One.Add(Symbols.Tfr)).Add(entryPriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr).Mul(future.One.Sub(mmr)))
			closePNL := entryPriceDec.Sub(closePriceDec).Mul(positions.CloseVolume).Mul(Symbols.Fv)
			closeFee := closePriceDec.Mul(positions.CloseVolume).Mul(Symbols.Fv).Mul(Symbols.Tfr)
			Asset = assetDec.Add(positionMargin).Add(closePNL).Sub(closeFee).Sub(shareProfitDec)
			fmt.Println("仓位保证金:", positionMargin, "平仓盈亏:", closePNL, "平仓手续费:", closeFee, "维持保证金率:", mmr)
		}
	}

	return nil, Asset
}

func AdjustLeverageAfterAsset(assetDec decimal.Decimal, leverage float64, position []*GenDanPositionDetailsInfo, symbol []*future.SymbolInfo) (error, decimal.Decimal) {
	Asset := decimal.Zero
	adjustLeverageAddLongPositionMargin := decimal.Zero
	adjustLeverageAddShortPositionMargin := decimal.Zero
	leverageDec := decimal.NewFromFloat(leverage)
	imr := future.One.Div(leverageDec)
	for _, positions := range position {
		// 查找对应的交易对详情
		var Symbols *future.SymbolInfo
		for _, symbols := range symbol {
			if positions.ContractName == symbols.Symbol {
				Symbols = symbols
				break
			}
		}

		entryPriceDec, err := decimal.NewFromString(positions.EntryPrice)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}

		MarginDec, err := decimal.NewFromString(positions.Margin)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}

		if positions.PositionType == 1 {
			adjustLeverageAfterPositionMargin := entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(imr.Add(Symbols.Tfr))
			// 调低杠杆后资产变动 = 当前资产 - (计算调整杠杆后的仓位保证金 - 当前仓位保证金)
			// 调低杠杆后需要添加的保证金 = 计算调整杠杆后的仓位保证金 - 当前仓位保证金
			adjustLeverageAddLongPositionMargin = adjustLeverageAddLongPositionMargin.Add(adjustLeverageAfterPositionMargin.Sub(MarginDec))
			adjustLeverageAfterAddPositionMargin := adjustLeverageAfterPositionMargin.Sub(MarginDec)
			fmt.Println("币对名:", positions.ContractName, "仓位ID:", positions.PositionID, "多仓总添加保证金:", adjustLeverageAddLongPositionMargin, "当前仓位添加保证金:", adjustLeverageAfterAddPositionMargin, "当前仓位保证金:", adjustLeverageAfterPositionMargin, "起始保证金:", MarginDec)
		} else if positions.PositionType == 2 {
			mmr := decimal.NewFromFloat(0.03)
			adjustLeverageAfterPositionMargin := entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(imr).Mul(future.One.Add(Symbols.Tfr)).Add(entryPriceDec.Mul(positions.Size).Mul(Symbols.Fv).Mul(Symbols.Tfr).Mul(future.One.Sub(mmr)))
			adjustLeverageAddShortPositionMargin = adjustLeverageAddShortPositionMargin.Add(adjustLeverageAfterPositionMargin.Sub(MarginDec))
			adjustLeverageAfterAddPositionMargin := adjustLeverageAfterPositionMargin.Sub(MarginDec)
			fmt.Println("币对名:", positions.ContractName, "仓位ID:", positions.PositionID, "空仓总添加保证金:", adjustLeverageAddShortPositionMargin, "当前仓位添加保证金:", adjustLeverageAfterAddPositionMargin, "当前仓位保证金:", adjustLeverageAfterPositionMargin, "起始保证金:", MarginDec)
		}
		Asset = assetDec.Sub(adjustLeverageAddLongPositionMargin).Sub(adjustLeverageAddShortPositionMargin)
	}

	return nil, Asset
}

type CopyTradingAssetResponse struct {
	Code      int                          `json:"code"`
	Message   string                       `json:"msg"`
	Data      CopyTradingAssetResponseInfo `json:"data"`
	ExtraInfo any                          `json:"extra"`
}

type CopyTradingAssetResponseInfo struct {
	AvatarURL string `json:"avatar"`
	Balance   string `json:"balance"`
	Nickname  string `json:"nickname"`
	IsTrader  int    `json:"trader"`
	Asset     string `json:"asset"`
}

func CopyTradingAsset() (error, decimal.Decimal) {
	url := futureCopyTrading.CopyTradingAssetUrl
	responseTest, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
	}

	var copyTradingAssetResponse CopyTradingAssetResponse
	if err := json.Unmarshal(responseTest, &copyTradingAssetResponse); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
	}

	if copyTradingAssetResponse.Code == 200 {
		assetDec, err := decimal.NewFromString(copyTradingAssetResponse.Data.Asset)
		if err != nil {
			return errors.New("金额转换错误:" + err.Error()), decimal.Zero
		}
		return nil, assetDec
	}

	return nil, decimal.Zero
}
