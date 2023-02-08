package weibo

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/cache"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetSearchHot(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := cache.GetCache(ctx, "WEIBO_HOT"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://m.weibo.cn/api/container/getIndex?containerid=106003type%3D25%26t%3D3%26disable_hot%3D1%26filter_type%3Drealtimehot&title=%E5%BE%AE%E5%8D%9A%E7%83%AD%E6%90%9C&extparam=filter_type%3Drealtimehot%26mi_cid%3D100103%26pos%3D0_0%26c_type%3D30%26display_time%3D1540538388&luicode=10000011&lfid=231583"
	headers := getHeaders()
	headers["Referer"] = "https://s.weibo.com/top/summary?cate=realtimehot"
	headers["MWeibo-Pwa"] = "1"
	headers["X-Requested-With"] = "XMLHttpRequest"
	rssData := dao.RSSFeed{
		Title:       "微博热搜榜",
		Link:        "https://s.weibo.com/top/summary?cate=realtimehot",
		Tag:         []string{"时事"},
		Description: "实时热点，每10分钟更新一次",
		ImageUrl:    "https://h5.sinaimg.cn/m/weibo-lite/icon-default-192.png",
	}
	if resp, err := service.GetHttpClient().SetHeaderMap(headers).Get(ctx, apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		dataJsonList := respJson.GetJsons("data.cards.0.card_group")
		rssItems := make([]dao.RSSItem, 0)
		for _, dataJson := range dataJsonList {
			rssItem := dao.RSSItem{}
			rssItem.Title = dataJson.Get("desc").String()
			rssItem.Link = "https://m.weibo.cn/search?containerid=100103type%3D1%26q%3D" + dataJson.Get("desc").String()
			rssItem.Content = feed.GenerateContent(dataJson.Get("desc").String())
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	cache.SetCache(ctx,"WEIBO_HOT", rssStr)
	req.Response.WriteXmlExit(rssStr)
}
