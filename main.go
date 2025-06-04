package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/ShivangSrivastava/vani/analysis"
	"github.com/ShivangSrivastava/vani/data"
	"github.com/ShivangSrivastava/vani/handler"
	"github.com/ShivangSrivastava/vani/logger"
	"github.com/ShivangSrivastava/vani/rpc"
)

func main() {
	// Check for --version
	version := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *version {
		fmt.Printf("%v %v", data.GetData().ServerName, data.GetData().Version)
		return
	}
	// Initialise logger with new log file, and add init message
	logger.Init("/home/shivang/personal/vani/vani.log")
	logger.Info("Started vani")

	// Takes input as stdio from lsp client
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	// state save all info of file and there text data
	state := analysis.NewState()

	for scanner.Scan() {
		// take one message at a time, based on split function said
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Error(fmt.Sprintf("%v", err))
			continue
		}

		// call handle and give it method and content, let it do remaining work
		handler.Handle(state, method, contents)
	}
}
