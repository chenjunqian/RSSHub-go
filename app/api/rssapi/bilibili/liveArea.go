package bilibili

import (
	"fmt"
	"rsshub/app/dao"
	"rsshub/lib"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func (ctl *Controller) GetLinvArea(req *ghttp.Request) {

	areaId := req.GetString("areaId")
	order := req.GetString("order")

	var orderTitle string

	switch order {
	case "live_time":
		orderTitle = "最新开播"
		break
	case "online":
		orderTitle = "人气直播"
		break
	}

	rssData := dao.RSSFeed{}
	var parentTitle string
	var parentId string
	var areaTitle string
	var areaLink string
	apiUrl := "https://api.live.bilibili.com/room/v1/Area/getList"
	header := getHeaders()
	header["Referer"] = "https://link.bilibili.com/p/center/index"
	if resp, err := g.Client().SetHeaderMap(header).Get(apiUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		dataJsonList := jsonResp.GetJsons("data")

		for _, dataJson := range dataJsonList {
			itemList := dataJson.GetJsons("list")
			for _, item := range itemList {
				if item.GetString("id") == areaId {
					parentTitle = dataJson.GetString("name")
					parentId = dataJson.GetString("id")
					areaTitle = item.GetString("name")
					switch parentId {
					case "1":
						areaLink = fmt.Sprintf("https://live.bilibili.com/pages/area/ent-all#%s/%s", item.GetString("cate_id"), areaId)
						break
					case "2":
					case "3":
						areaLink = fmt.Sprintf("https://live.bilibili.com/p/eden/area-tags#/%s/%s", parentId, areaId)
						break
					case "4":
						areaLink = "https://live.bilibili.com/pages/area/draw"
						break
					}
				}
			}
		}

		rssData.Title = fmt.Sprintf("哔哩哔哩直播-%s·%s分区-%s", parentTitle, areaTitle, orderTitle)
		rssData.Link = areaLink
		rssData.Description = rssData.Title
		rssData.ImageUrl = "https://www.bilibili.com/favicon.ico"

		rssItems := make([]dao.RSSItem, 0)
		areaApiUrl := fmt.Sprintf("https://api.live.bilibili.com/room/v1/area/getRoomList?area_id=%s&sort_type=%s&page_size=30&page_no=1", areaId, order)
		if areaResp, err := g.Client().SetHeaderMap(header).Get(areaApiUrl); err == nil {
			areaJsonResp := gjson.New(areaResp.ReadAllString())
			dataJsonList := areaJsonResp.GetJsons("data")

			for _, dataJson := range dataJsonList {
				rssItem := dao.RSSItem{}
				rssItem.Title = dataJson.GetString("uname") + " " + dataJson.GetString("title")
				rssItem.Description = rssItem.Title
				rssItem.Created = time.Now().String()
				rssItem.Link = "https://live.bilibili.com/" + dataJson.GetString("roomid")
				rssItems = append(rssItems, rssItem)
			}
		}

		rssData.Items = rssItems

	}

	rssStr := lib.GenerateRSS(rssData, req.Router.Uri)
	_ = req.Response.WriteXmlExit(rssStr)
}
