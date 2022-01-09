package juejin

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetRecommand(req *ghttp.Request) {
	var (
		cacheKey string
		apiUrl   string
	)
	cacheKey = "JUEJIN_RECOMMAND_HOT"
	if value, err := g.Redis().DoVar("GET", cacheKey); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl = "https://api.juejin.cn/recommend_api/v1/article/recommend_all_feed?aid=2608&uuid=7007464640411485710"
	rssData := dao.RSSFeed{
		Title:       "掘金推荐 - 最热",
		Link:        apiUrl,
		Description: "掘金是一个帮助开发者成长的社区,是给开发者用的 Hacker News,给设计师用的 Designer News,和给产品经理用的 Medium。掘金的技术文章由稀土上聚集的技术大牛和极客共同编辑为你筛选出最优质的干货,其中包括：Android、iOS、前端、后端等方面的内容。用户每天都可以在这里找到技术世界的头条内容。与此同时,掘金内还有沸点、掘金翻译计划、线下活动、专栏文章等内容。即使你是 GitHub、StackOverflow、开源中国的用户,我们相信你也可以在这里有所收获。",
		ImageUrl:    "https://lf3-cdn-tos.bytescm.com/obj/static/xitu_juejin_web//static/favicons/favicon-32x32.png",
	}
	var payload = map[string]interface{}{
		"limit":       20,
		"sort_type":   3,
		"client_type": 2608,
		"cursor":      0,
	}
	if resp := component.PostContentByMobile(apiUrl, payload); resp != "" {
		rssItems := parseRecommand(resp)
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", cacheKey, rssStr)
	g.Redis().DoVar("EXPIRE", cacheKey, 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}

func parseRecommand(respString string) (items []dao.RSSItem) {
	var (
		jsonResp     *gjson.Json
		jsonItemList []*gjson.Json
	)
	jsonResp = gjson.New(respString)
	jsonItemList = jsonResp.GetJsons("data")
	for _, jsonItem := range jsonItemList {
		var (
			item         dao.RSSItem
			jsonItemInfo *gjson.Json
		)

		jsonItemInfo = jsonItem.GetJson("item_info")
		if jsonItem.GetInt("item_type") == 2 {
			var articleInfo *gjson.Json
			var articleBaseUrl = "https://juejin.im/post/"
			articleInfo = jsonItemInfo.GetJson("article_info")
			item.Title = articleInfo.GetString("title")
			item.Thumbnail = articleInfo.GetString("cover_image")
			item.Author = articleInfo.GetString("author_user_info.user_name")
			item.Created = articleInfo.GetString("ctime")
			item.Link = articleBaseUrl + articleInfo.GetString("article_id")
			item.Description = feed.GenerateDescription(articleInfo.GetString("brief_content"))
			items = append(items, item)
		} else if jsonItem.GetInt("item_type") == 14 {
			item.Title = jsonItemInfo.GetString("title")
			item.Thumbnail = jsonItemInfo.GetString("picture")
			item.Author = jsonItemInfo.GetString("author_name")
			item.Created = jsonItemInfo.GetString("ctime")
			item.Link = jsonItemInfo.GetString("url")
			item.Description = feed.GenerateDescription(jsonItemInfo.GetString("brief"))
			items = append(items, item)
		}
	}

	return
}
