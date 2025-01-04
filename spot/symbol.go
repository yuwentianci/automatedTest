package spot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"myapp/function"
)

// 定义结构体
type Response struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Result  []Result `json:"result"`
}

type Result struct {
	ID                  int           `json:"ID"`
	CreatedAt           string        `json:"CreatedAt"`
	UpdatedAt           string        `json:"UpdatedAt"`
	Icon                string        `json:"icon"`
	BaseAsset           string        `json:"base_asset"`
	BaseAssetName       string        `json:"base_asset_name"`
	BaseAssetPrecision  int           `json:"base_asset_precision"`
	QuoteAsset          string        `json:"quote_asset"`
	QuoteAssetName      string        `json:"quote_asset_name"`
	QuoteAssetPrecision int           `json:"quote_asset_precision"`
	Symbol              string        `json:"symbol"`
	TickSize            string        `json:"tick_size"`
	MinQuantity         string        `json:"min_quantity"`
	Status              int           `json:"status"`
	Recommend           int           `json:"recommend"`
	LimitTakerFee       float64       `json:"limit_taker_fee"`
	LimitMakerFee       float64       `json:"limit_maker_fee"`
	MarketTakerFee      float64       `json:"market_taker_fee"`
	SiteID              int           `json:"site_id"`
	DisplayOrder        int           `json:"display_order"`
	DisplaySearch       int           `json:"display_search"`
	AlertPercent        string        `json:"alert_percent"`
	AllowdealAt         int           `json:"allowdeal_at"`
	ForbiddendealAt     int           `json:"forbiddendeal_at"`
	PoQuoteAssetMin     int           `json:"po_quote_asset_min"`
	PoBaseAssetMin      int           `json:"po_base_asset_min"`
	Flag                string        `json:"flag"`
	FlagOrder           int           `json:"flag_order"`
	Gear                int           `json:"gear"`
	FullName            string        `json:"full_name"`
	AssetIcon           string        `json:"asset_icon"`
	ThreeURL            string        `json:"three_url"`
	ThreeIcon           string        `json:"three_icon"`
	ReminderTitle       string        `json:"reminder_title"`
	ReminderContent     string        `json:"reminder_content"`
	ReminderExist       bool          `json:"reminder_exist"`
	Chain               []interface{} `json:"chain"` // 由于 `chain` 是空的，我们使用 `interface{}` 处理
	Classes             []Class       `json:"classes"`
	Open                string        `json:"open"`
	Last                string        `json:"last"`
	High                string        `json:"high"`
	Low                 string        `json:"low"`
	Deal                string        `json:"deal"`
	Volume              string        `json:"volume"`
}

type Class struct {
	DisplayOrder  int    `json:"display_order"`
	SymbolID      int    `json:"symbol_id"`
	SymbolClassID int    `json:"symbol_class_id"`
	Shorthand     string `json:"shorthand"`
}

func Symbol() (error, decimal.Decimal) {
	url := "https://www.biconomy.com/api/v1/symbol"
	responseTest, err := function.GetDetails(url)
	if err != nil {
		fmt.Println(err)
		return err, decimal.Zero
	}

	var balanceData Response
	if err := json.Unmarshal(responseTest, &balanceData); err != nil {
		return errors.New("解析JSON响应时发生错误:" + err.Error()), decimal.Zero
	}

	// 创建一个新的 Excel 文件
	f := excelize.NewFile()
	// 创建一个新的工作表
	sheetName := "Symbols"
	index, _ := f.NewSheet(sheetName)

	// 写入表头
	headers := []string{
		"ID", "CreatedAt", "UpdatedAt", "Icon", "BaseAsset", "BaseAssetName",
		"BaseAssetPrecision", "QuoteAsset", "QuoteAssetName", "QuoteAssetPrecision",
		"Symbol", "TickSize", "MinQuantity", "Status", "Recommend", "LimitTakerFee",
		"LimitMakerFee", "MarketTakerFee", "SiteID", "DisplayOrder", "DisplaySearch",
		"AlertPercent", "AllowdealAt", "ForbiddendealAt", "PoQuoteAssetMin",
		"PoBaseAssetMin", "Flag", "FlagOrder", "Gear", "FullName", "AssetIcon",
		"ThreeURL", "ThreeIcon", "ReminderTitle", "ReminderContent", "ReminderExist",
		"Open", "Last", "High", "Low", "Deal", "Volume",
	}

	for colIndex, header := range headers {
		cell, _ := excelize.ColumnNumberToName(colIndex + 1) // Column numbers are 1-based
		cell = fmt.Sprintf("%s1", cell)
		f.SetCellValue(sheetName, cell, header)
	}

	// 写入数据
	for rowIndex, result := range balanceData.Result {
		row := rowIndex + 2 // start from the second row
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), result.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), result.CreatedAt)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), result.UpdatedAt)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), result.Icon)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), result.BaseAsset)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), result.BaseAssetName)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), result.BaseAssetPrecision)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), result.QuoteAsset)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), result.QuoteAssetName)
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", row), result.QuoteAssetPrecision)
		f.SetCellValue(sheetName, fmt.Sprintf("K%d", row), result.Symbol)
		f.SetCellValue(sheetName, fmt.Sprintf("L%d", row), result.TickSize)
		f.SetCellValue(sheetName, fmt.Sprintf("M%d", row), result.MinQuantity)
		f.SetCellValue(sheetName, fmt.Sprintf("N%d", row), result.Status)
		f.SetCellValue(sheetName, fmt.Sprintf("O%d", row), result.Recommend)
		f.SetCellValue(sheetName, fmt.Sprintf("P%d", row), result.LimitTakerFee)
		f.SetCellValue(sheetName, fmt.Sprintf("Q%d", row), result.LimitMakerFee)
		f.SetCellValue(sheetName, fmt.Sprintf("R%d", row), result.MarketTakerFee)
		f.SetCellValue(sheetName, fmt.Sprintf("S%d", row), result.SiteID)
		f.SetCellValue(sheetName, fmt.Sprintf("T%d", row), result.DisplayOrder)
		f.SetCellValue(sheetName, fmt.Sprintf("U%d", row), result.DisplaySearch)
		f.SetCellValue(sheetName, fmt.Sprintf("V%d", row), result.AlertPercent)
		f.SetCellValue(sheetName, fmt.Sprintf("W%d", row), result.AllowdealAt)
		f.SetCellValue(sheetName, fmt.Sprintf("X%d", row), result.ForbiddendealAt)
		f.SetCellValue(sheetName, fmt.Sprintf("Y%d", row), result.PoQuoteAssetMin)
		f.SetCellValue(sheetName, fmt.Sprintf("Z%d", row), result.PoBaseAssetMin)
		f.SetCellValue(sheetName, fmt.Sprintf("AA%d", row), result.Flag)
		f.SetCellValue(sheetName, fmt.Sprintf("AB%d", row), result.FlagOrder)
		f.SetCellValue(sheetName, fmt.Sprintf("AC%d", row), result.Gear)
		f.SetCellValue(sheetName, fmt.Sprintf("AD%d", row), result.FullName)
		f.SetCellValue(sheetName, fmt.Sprintf("AE%d", row), result.AssetIcon)
		f.SetCellValue(sheetName, fmt.Sprintf("AF%d", row), result.ThreeURL)
		f.SetCellValue(sheetName, fmt.Sprintf("AG%d", row), result.ThreeIcon)
		f.SetCellValue(sheetName, fmt.Sprintf("AH%d", row), result.ReminderTitle)
		f.SetCellValue(sheetName, fmt.Sprintf("AI%d", row), result.ReminderContent)
		f.SetCellValue(sheetName, fmt.Sprintf("AJ%d", row), result.ReminderExist)
		f.SetCellValue(sheetName, fmt.Sprintf("AK%d", row), result.Open)
		f.SetCellValue(sheetName, fmt.Sprintf("AL%d", row), result.Last)
		f.SetCellValue(sheetName, fmt.Sprintf("AM%d", row), result.High)
		f.SetCellValue(sheetName, fmt.Sprintf("AN%d", row), result.Low)
		f.SetCellValue(sheetName, fmt.Sprintf("AO%d", row), result.Deal)
		f.SetCellValue(sheetName, fmt.Sprintf("AP%d", row), result.Volume)
	}

	// 设置工作表为默认工作表
	f.SetActiveSheet(index)

	// 保存 Excel 文件
	filePath := "symbols.xlsx"
	if err := f.SaveAs(filePath); err != nil {
		return errors.New("保存 Excel 文件时发生错误:" + err.Error()), decimal.Zero
	}

	fmt.Println("Excel 文件已成功保存到:", filePath)
	return nil, decimal.Zero
}
