package ifan

import (
	"context"
	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetFlash(req *ghttp.Request) {
	var ctx context.Context = context.Background()
	if value, err := service.GetRedis().Do(ctx,"GET", "IFAN_DAILY"); err == nil {
		if value.String() != "" {
			req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://sso.ifanr.com/api/v5/wp/article/?post_category=早报"
	rssData := dao.RSSFeed{
		Title:       "爱范-早报",
		Link:        "https://sso.ifanr.com/api/v5/wp/article/?post_category=早报",
		Tag:         []string{"媒体", "科技"},
		Description: "爱范每日早报",
		ImageUrl:    "https://images.ifanr.cn/wp-content/themes/ifanr-5.0-pc/static/images/favicon.ico",
	}
	if resp := service.GetContent(ctx,apiUrl); resp != "" {

		respJson := gjson.New(resp)
		itemJsonList := respJson.GetJsons("objects")
		rssItems := make([]dao.RSSItem, 0)
		for _, itemJson := range itemJsonList {
			rssItem := dao.RSSItem{}
			title := itemJson.Get("post_title").String()
			author := itemJson.Get("created_by.name").String()
			link := itemJson.Get("post_url").String()
			time := itemJson.Get("created_at").String()
			imageLink := itemJson.Get("post_cover_image").String()
			content := parseCommonDetail(ctx, link)
			content = feed.GenerateContent(content)

			rssItem.Title = title
			rssItem.Created = time
			rssItem.Content = content
			rssItem.Link = link
			rssItem.Author = author
			rssItem.Thumbnail = imageLink
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	service.GetRedis().Do(ctx,"SET", "IFAN_DAILY", rssStr)
	service.GetRedis().Do(ctx,"EXPIRE", "IFAN_DAILY", 60*60*4)
	req.Response.WriteXmlExit(rssStr)
}
