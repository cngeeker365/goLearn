package persist

import (
	"awesomeProject/crawler/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T)  {
	expected:= model.Profile{
		Info: []string{"未婚","29岁","天秤座(09.23-10.22)","180cm","66kg","工作地:阿坝茂县",
			"月收入:1.2-2万","其他职业","高中及以下","羌族","籍贯:四川阿坝","体型:运动员型",
			"社交场合会抽烟","社交场合会喝酒","已购房","已买车","没有小孩","是否想要孩子:想要孩子",
			"何时结婚:时机成熟就结婚","微微一笑"},
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	//TODO: Try to start up elastic search using docker go client.

	client,err:= elastic.NewClient(elastic.SetSniff(false))
	if err!=nil{
		panic(err)
	}
	resp, err:=client.Get().Index("dating_profile").Type("dating_profile").Id(id).Do(context.Background())
	if err!=nil{
		panic(err)
	}

	var actual model.Profile
	err=json.Unmarshal([]byte(*resp.Source), &actual)
	if err!=nil{
		panic(err)
	}

	t.Logf(" got %s\n actual %s\n", actual, expected)
}
