package client

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error){
	client, err:= rpcsupport.NewClient(host)
	if err!=nil{
		panic(err)
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: go item #%d: %v\n", itemCount, item)
			itemCount++

			//Call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err !=nil {
				log.Printf("Item Saver: error saving item %v: %v\n", item, err)
			}
		}
	}()
	return out, nil
}