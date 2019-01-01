package persist

import "log"

func ItemSaver() chan interface{}{
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: go item #%d: %v\n", itemCount, item)
			itemCount++
		}
	}()
	return out
}
