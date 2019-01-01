package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{}{
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: go item #%d: %v\n", itemCount, item)
			itemCount++
			
			_, err:=save(item)
			if err !=nil {
				log.Printf("Item Saver: error saving item %v: %v\n", item, err)
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error){
	client, err := elastic.NewClient( elastic.SetSniff(false))//Must turn off sniff in docker
	if err != nil{
		return "", err
	}
	resp, err:= client.Index().Index("dating_profile").Type("dating_profile").BodyJson(item).Do(context.Background())
	if err !=nil {
		return "", err
	}
	return resp.Id, nil
}
