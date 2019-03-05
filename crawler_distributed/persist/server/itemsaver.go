package main

import (
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/persist"
	"awesomeProject/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

var port = flag.Int("port",0,"the port for me to listen on")
func main(){
	flag.Parse()
	if *port==0{
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(ServeRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
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