package main

import (
	"github.com/song-without-words/huobi_Golang/cmd/accountclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/accountwebsocketclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/algoorderclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/commonclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/crossmarginclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/etfclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/isolatedmarginclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/marketclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/marketwebsocketclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/orderclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/orderwebsocketclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/stablecoinclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/subuserclientexample"
	"github.com/song-without-words/huobi_Golang/cmd/walletclientexample"
	"github.com/song-without-words/huobi_Golang/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	algoorderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	subuserclientexample.RunAllExamples()
	stablecoinclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}
