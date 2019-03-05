package main

import (
	"awesomeProject/block_chain/core"
	"fmt"
)

func main()  {
	blockChain := core.NewBlockChain()
	blockChain.SendData("Send 1 BTC to Jacky")
	blockChain.SendData("Send 2 EOS to Jacky")
	fmt.Println(blockChain.String())
}
