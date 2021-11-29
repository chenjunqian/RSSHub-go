package ifan

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetFlash(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "IFAN_DAILY"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
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
	if resp, err := g.Client().SetHeaderMap(getHeaders()).Get(apiUrl); err == nil {
		respJson := gjson.New(resp.ReadAllString())
		itemJsonList := respJson.GetJsons("objects")
		rssItems := make([]dao.RSSItem, 0)
		for _, itemJson := range itemJsonList {
			rssItem := dao.RSSItem{}
			title := itemJson.GetString("post_title")
			author := itemJson.GetString("created_by.name")
			link := itemJson.GetString("post_url")
			time := itemJson.GetString("created_at")
			imageLink := itemJson.GetString("post_cover_image")
			content := parseCommonDetail(link)
			description := lib.GenerateDescription(imageLink, content)

			rssItem.Title = title
			rssItem.Created = time
			rssItem.Description = description
			rssItem.Link = link
			rssItem.Author = author
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "IFAN_DAILY", rssStr)
	g.Redis().DoVar("EXPIRE", "IFAN_DAILY", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}
