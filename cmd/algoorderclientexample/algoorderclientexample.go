package algoorderclientexample

import (
	"github.com/song-without-words/huobi_Golang/config"
	"github.com/song-without-words/huobi_Golang/logging/applogger"
	"github.com/song-without-words/huobi_Golang/pkg/client"
	"github.com/song-without-words/huobi_Golang/pkg/model"
	"github.com/song-without-words/huobi_Golang/pkg/model/algoorder"
)

func RunAllExamples() {
	placeOrder()
	getOpenOrders()
	getSpecificOrder()
	cancelOder()
	getHistoryOrders()
	cancelAllAfter()
}

func placeOrder() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	request := algoorder.PlaceOrderRequest{
		AccountId:     11136102,
		Symbol:        "htusdt",
		OrderPrice:    "4.4",
		OrderSide:     "buy",
		OrderSize:     "2",
		TimeInForce:   "gtc",
		OrderType:     "limit",
		ClientOrderId: "huobi1901",
		StopPrice:     "4",
	}
	resp, err := client.PlaceOrder(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			applogger.Info("Place algo order successfully, client order id: %s", resp.Data.ClientOrderId)
		} else {
			applogger.Error("Place algo order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
}

func cancelOder() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	request := algoorder.CancelOrdersRequest{
		ClientOrderIds: []string{"huobi1901"},
	}
	resp, err := client.CancelOrder(&request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data.Accepted != nil {
				for _, id := range resp.Data.Accepted {
					applogger.Info("Cancelled client order id success: %s", id)
				}
			}
			if resp.Data.Rejected != nil {
				for _, id := range resp.Data.Rejected {
					applogger.Error("Cancelled client order id error: %s", id)
				}
			}
		} else {
			applogger.Error("Cancel algo order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
}

func getOpenOrders() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	request := new(model.GetRequest).Init()
	request.AddParam("accountId", config.AccountId)

	resp, err := client.GetOpenOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data != nil {
				applogger.Info("There are total %d open orders", len(resp.Data))
				for _, o := range resp.Data {
					applogger.Info("Open orders, cid: %s, symbol: %s, status: %s", o.ClientOrderId, o.Symbol, o.OrderStatus)
				}
			}
		} else {
			applogger.Error("Get open order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
}

func getHistoryOrders() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	request := new(model.GetRequest).Init()
	request.AddParam("symbol", "htusdt")
	request.AddParam("orderStatus", "canceled")

	resp, err := client.GetHistoryOrders(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data != nil {
				applogger.Info("There are total %d history orders", len(resp.Data))
				for _, o := range resp.Data {
					applogger.Info("history orders, cid: %s, symbol: %s, status: %s", o.ClientOrderId, o.Symbol, o.OrderStatus)
				}
			}
		} else {
			applogger.Error("Get history order error, code: %d, message: %s", resp.Code, resp.Message)
		}
	}
}

func getSpecificOrder() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	request := new(model.GetRequest).Init()
	request.AddParam("clientOrderId", "huobi1901")

	resp, err := client.GetSpecificOrder(request)
	if err != nil {
		applogger.Error(err.Error())
	} else {
		if resp.Code == 200 {
			if resp.Data != nil {
				o := resp.Data
				applogger.Info("Get order, cid: %s, symbol: %s, status: %s", o.ClientOrderId, o.Symbol, o.OrderStatus)
			}
		} else {
			applogger.Error("Get order error, code: %s, message: %s", resp.Code, resp.Message)
		}
	}
}

func cancelAllAfter() {
	client := new(client.AlgoOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	request := algoorder.CancelAllAfterRequest{Timeout: 10}
	resp, err := client.CancelAllAfter(request)
	if err != nil {
		applogger.Error("cancelAllAfter error: %s", err)
	} else {
		applogger.Info("cancelAllAfter, %v", resp.Data)
	}
}
