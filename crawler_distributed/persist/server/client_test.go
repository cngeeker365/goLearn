package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"awesomeProject/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T){
	const host = ":1234"
	const index = "test"
	//start ItemSaverServer
	go ServeRpc(host, index)

	//之所以在这里加sleep（偷懒），是因为如果不这样做，后续client调用时，rpc实际还未启动成功
	//在实际工作中，rpc启动后应该通过一定的手段通知client自己已经启动成功了
	time.Sleep(time.Second)

	//start ItemSaverClient
	client, err:= rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}

	//call save
	item:= engine.Item{
		Url:"http://m.zhenai.com/u/1320662004",
		Type:"zhenai",
		Id:"1320662004",
		Payload: model.Profile{
			Info: []string{"未婚","29岁","天秤座(09.23-10.22)","180cm","66kg","工作地:阿坝茂县",
				"月收入:1.2-2万","其他职业","高中及以下","羌族","籍贯:四川阿坝","体型:运动员型",
				"社交场合会抽烟","社交场合会喝酒","已购房","已买车","没有小孩","是否想要孩子:想要孩子",
				"何时结婚:时机成熟就结婚","微微一笑"}},
	}

	result:=""
	client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result!="ok"{
		t.Error("result: %s\n err: %s\n", result, err)
	}

}
