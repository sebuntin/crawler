package persist

import (
	"context"
	"engine"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver() chan interface{} {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	out := make(chan interface{})
	go func() {
		count := 0
		for {
			item := <-out
			count++
			log.Printf("Item Saver: got item #%d, %v", count, item)
			if book, ok := item.(engine.Item); ok {
				if err := elasticSave(book, "douban_book", client); err != nil {
					log.Printf("Item Saver: error saving item %v: %v\n", item, err)
					continue
				}
			}
		}
	}()
	return out
}

func elasticSave(item engine.Item, index string, esClient *elastic.Client)error {
	// 如果es跑在docker上，需要关掉sniff
	//opt := elastic.SetSniff(false)
	// 存数据
	indexService := esClient.Index().
		Index(index).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
