package client

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/song-without-words/huobi_Golang/internal"
	"github.com/song-without-words/huobi_Golang/internal/requestbuilder"
	"github.com/song-without-words/huobi_Golang/pkg/model"
	"github.com/song-without-words/huobi_Golang/pkg/model/etf"
)

// Responsible to operate ETF
type ETFClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *ETFClient) Init(accessKey string, secretKey string, host string, s string) *ETFClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, s)
	return p
}

// Return the basic information of ETF creation and redemption, as well as ETF constituents
func (p *ETFClient) GetSwapConfig(etfName string) (*etf.SwapConfig, error) {
	request := new(model.GetRequest).Init()

	request.AddParam("etf_name", etfName)

	url := p.privateUrlBuilder.Build("GET", "/etf/swap/config", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := etf.GetSwapConfigResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Swap in ETF
func (p *ETFClient) SwapIn(request etf.SwapRequest) (bool, error) {

	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return false, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/etf/swap/in", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return false, postErr
	}
	result := etf.SwapResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return false, jsonErr
	}

	if result.Code == 200 && result.Success == true {
		return result.Success, nil
	}
	return false, errors.New(postResp)
}

// Swap out ETF
func (p *ETFClient) SwapOut(request etf.SwapRequest) (bool, error) {

	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return false, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/etf/swap/out", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return false, postErr
	}
	result := etf.SwapResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return false, jsonErr
	}

	if result.Code == 200 && result.Success == true {
		return result.Success, nil
	}
	return false, errors.New(postResp)

}

// Get past creation and redemption.(up to 100 records)
func (p *ETFClient) GetSwapList(etfName string, offset int, limit int) ([]*etf.SwapList, error) {
	request := new(model.GetRequest).Init()

	request.AddParam("etf_name", etfName)
	request.AddParam("offset", strconv.Itoa(offset))
	request.AddParam("limit", strconv.Itoa(limit))

	url := p.privateUrlBuilder.Build("GET", "/etf/swap/list", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := etf.GetSwapListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}
