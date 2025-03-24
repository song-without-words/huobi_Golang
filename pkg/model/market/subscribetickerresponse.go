package market

import "github.com/song-without-words/huobi_Golang/pkg/model/base"

type SubscribeTickerResponse struct {
	base.WebSocketResponseBase
	Data *Ticker
	Tick *Ticker
}
