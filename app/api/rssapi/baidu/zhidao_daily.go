package baidu

import (
	"rsshub/app/component"
	"rsshub/app/dao"
	"rsshub/app/service/feed"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/encoding/gcharset"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *controller) GetZhiDaoDaily(req *ghttp.Request) {
	if value, err := g.Redis().DoVar("GET", "BAIDU_ZHIDAO_DAILY"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	apiUrl := "https://zhidao.baidu.com/daily?fr=daohang"
	rssData := dao.RSSFeed{
		Title:       "百度知道日报",
		Link:        apiUrl,
		Description: "百度知道日报精选",
		Tag:         []string{"知识", "百科", "问答"},
		ImageUrl:    "https://www.baidu.com/favicon.ico",
	}
	if resp := component.GetContent(apiUrl); resp != "" {
		respString, _ := gcharset.Convert("UTF-8", "gbk", resp)
		docs := soup.HTMLParse(respString)
		itemList := docs.FindAll("li", "class", "clearfix")
		rssItems := make([]dao.RSSItem, 0)
		for _, item := range itemList {
			var (
				title      string
				contentDiv soup.Root
				content    string
				link       string
				imageLink  string
			)
			title = item.Find("img").Attrs()["title"]
			imageLink = item.Find("img").Attrs()["src"]
			contentDiv = item.Find("div", "class", "summer")
			if contentDiv.Error != nil {
				continue
			}
			link = contentDiv.Find("a").Attrs()["href"]
			link = "https://zhidao.baidu.com/" + link
			content = parseDetail(link)
			rssItem := dao.RSSItem{
				Title:     title,
				Link:      link,
				Content:   feed.GenerateContent(content),
				Thumbnail: imageLink,
			}
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BAIDU_ZHIDAO_DAILY", rssStr)
	g.Redis().DoVar("EXPIRE", "BAIDU_ZHIDAO_DAILY", 60*60*4)
	_ = req.Response.WriteXmlExit(rssStr)
}

func parseDetail(detailLink string) (detailData string) {
	var (
		resp string
	)
	if resp = component.GetContent(detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)
		respString, _ = gcharset.Convert("UTF-8", "gbk", resp)
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("div", "class", "detail")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request baidu article detail failed, link  %s \n", detailLink)
	}

	return
}
