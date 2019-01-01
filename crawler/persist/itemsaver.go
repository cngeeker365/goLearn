package persist

import (
	"awesomeProject/crawler/engine"
	"context"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
)

func ItemSaver() chan engine.Item{
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: go item #%d: %v\n", itemCount, item)
			itemCount++
			
			err:=save(item)
			if err !=nil {
				log.Printf("Item Saver: error saving item %v: %v\n", item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error{
	client, err := elastic.NewClient( elastic.SetSniff(false))//Must turn off sniff in docker
	if err != nil{
		return err
	}
	if item.Type ==""{
		return errors.New("must supply Type")
	}

	indexSvc := client.Index().
		Index("dating_profile").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)
	if item.Id != ""{
		indexSvc.Id(item.Id)
	}
	_, err = indexSvc.Do(context.Background())


	if err !=nil {
		return err
	}
	return nil
}
