package bing

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"rsshub/app/dao"
	"rsshub/lib"
)

func (ctl *Controller) GetDailyImage(req *ghttp.Request) {

	if value, err := g.Redis().DoVar("GET", "BING_DAILY_IMG"); err == nil {
		if value.String() != "" {
			_ = req.Response.WriteXmlExit(value.String())
		}
	}

	baseUrl := "http://www.bing.com"
	apiUrl := baseUrl + "/HPImageArchive.aspx?format=js&idx=0&n=1"
	header := getHeaders()
	rssData := dao.RSSFeed{}
	rssData.Title = "Bing每日壁纸"
	rssData.Link = "https://cn.bing.com/"
	rssData.Tag = []string{"壁纸", "图片"}
	rssData.ImageUrl = "https://cn.bing.com/sa/simg/favicon-2x.ico"
	if statusResp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		respJson := gjson.New(statusResp.ReadAllString())
		imageListJson := respJson.GetJsons("images")
		rssItems := make([]dao.RSSItem, 0)
		for _, imageJson := range imageListJson {
			rssItem := dao.RSSItem{
				Title: imageJson.GetString("copyright"),
				Link:  imageJson.GetString("copyrightlink"),
			}
			var description string
			description = description + fmt.Sprintf("<img src=\"%s%s\">", baseUrl, imageJson.GetString("url"))
			rssItem.Description = description
			rssItems = append(rssItems, rssItem)
		}
		rssData.Items = rssItems
	}

	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	g.Redis().DoVar("SET", "BING_DAILY_IMG", rssStr)
	g.Redis().DoVar("EXPIRE", "BING_DAILY_IMG", 60*60*6)
	_ = req.Response.WriteXmlExit(rssStr)
}
