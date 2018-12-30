package parser

import (
	"awesomeProject/crawler/fetcher"
	"fmt"
	"testing"
)

var cont = `<div class="introduce-content" data-v-3e01facc data-v-0566ce42><div class="title" data-v-3e01facc><div class="color-section" data-v-3e01facc></div> <span data-v-3e01facc>个人资料（ID：1320662004）</span></div> <div class="basicInfo-section" data-v-3e01facc><div class="tag" data-v-3e01facc>未婚</div><div class="tag" data-v-3e01facc>29岁</div><div class="tag" data-v-3e01facc>天秤座(09.23-10.22)</div><div class="tag" data-v-3e01facc>180cm</div><div class="tag" data-v-3e01facc>66kg</div><div class="tag" data-v-3e01facc>工作地:阿坝茂县</div><div class="tag" data-v-3e01facc>月收入:1.2-2万</div><div class="tag" data-v-3e01facc>其他职业</div><div class="tag" data-v-3e01facc>高中及以下</div></div> <div class="detailInfo-section" data-v-3e01facc><div class="tag" data-v-3e01facc>羌族</div><div class="tag" data-v-3e01facc>体型:运动员型</div><div class="tag" data-v-3e01facc>社交场合会抽烟</div><div class="tag" data-v-3e01facc>社交场合会喝酒</div><div class="tag" data-v-3e01facc>已购房</div><div class="tag" data-v-3e01facc>已买车</div><div class="tag" data-v-3e01facc>没有小孩</div><div class="tag" data-v-3e01facc>是否想要孩子:想要孩子</div><div class="tag" data-v-3e01facc>何时结婚:时机成熟就结婚</div></div></div> <div class="introduce-content" data-v-24b9ac8d data-v-0566ce42><div class="title" data-v-24b9ac8d><div class="color-section" data-v-24b9ac8d></div> <span data-v-24b9ac8d>择偶条件</span></div> <div class="objectInfo-section" data-v-24b9ac8d><div class="tag" data-v-24b9ac8d>19-33岁</div><div class="tag" data-v-24b9ac8d>163-174cm</div><div class="tag" data-v-24b9ac8d>工作地:四川阿坝</div><div class="tag" data-v-24b9ac8d>中专</div><div class="tag" data-v-24b9ac8d>未婚</div><div class="tag" data-v-24b9ac8d>体型:苗条</div><div class="tag" data-v-24b9ac8d>可以喝酒</div><div class="tag" data-v-24b9ac8d>不要吸烟</div><div class="tag" data-v-24b9ac8d>没有小孩</div><div class="tag" data-v-24b9ac8d>是否想要孩子:想要孩子</div></div></div> <div class="za-footer" data-v-3900bb52 data-v-0566ce42>`
func TestParseProfile(t *testing.T) {
	contents, _ :=fetcher.Fetch("http://m.zhenai.com/u/1320662004")
	fmt.Printf("%s", contents)
	ParseProfile(contents,"微微一笑")
}


