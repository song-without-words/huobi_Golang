package market

import (
	"github.com/shopspring/decimal"
	"github.com/song-without-words/huobi_Golang/pkg/model/base"
)

type SubscribeCandlestickResponse struct {
	base base.WebSocketResponseBase
	Tick *Tick
	Data []Tick
}
type Tick struct {
	Id     int64           `json:"id"`
	Amount decimal.Decimal `json:"amount"`
	Count  int             `json:"count"`
	Open   decimal.Decimal `json:"open"`
	Close  decimal.Decimal `json:"close"`
	Low    decimal.Decimal `json:"low"`
	High   decimal.Decimal `json:"high"`
	Vol    decimal.Decimal `json:"vol"`
}
