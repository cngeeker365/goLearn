package main

import (
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/persist"
	"awesomeProject/crawler_distributed/rpcsupport"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func main(){
	log.Fatal(ServeRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func ServeRpc(host, index string) error{
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(	host,
						 		&persist.ItemSaverService{
									Client: client,
									Index:	index,
								})
}