package bilibili

import (
	"context"
	"fmt"

	"rsshub/internal/dao"
	"rsshub/internal/service"
	"rsshub/internal/service/feed"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *Controller) GetLinvArea(req *ghttp.Request) {
	var ctx context.Context = context.Background()

	areaId := req.Get("areaId").String()
	order := req.Get("order").String()

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
	if resp := service.GetContent(ctx, apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		dataJsonList := jsonResp.GetJsons("data")

		for _, dataJson := range dataJsonList {
			itemList := dataJson.GetJsons("list")
			for _, item := range itemList {
				if item.Get("id").String() == areaId {
					parentTitle = dataJson.Get("name").String()
					parentId = dataJson.Get("id").String()
					areaTitle = item.Get("name").String()
					switch parentId {
					case "1":
						areaLink = fmt.Sprintf("https://live.bilibili.com/pages/area/ent-all#%s/%s", item.Get("cate_id"), areaId)
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
		if areaResp := service.GetContent(ctx, areaApiUrl); resp != "" {
			areaJsonResp := gjson.New(areaResp)
			dataJsonList := areaJsonResp.GetJsons("data")

			for _, dataJson := range dataJsonList {
				rssItem := dao.RSSItem{}
				rssItem.Title = dataJson.Get("uname").String() + " " + dataJson.Get("title").String()
				rssItem.Description = rssItem.Title
				rssItem.Created = time.Now().String()
				rssItem.Link = "https://live.bilibili.com/" + dataJson.Get("roomid").String()
				rssItems = append(rssItems, rssItem)
			}
		}

		rssData.Items = rssItems

	}

	rssStr := feed.GenerateRSS(rssData, req.Router.Uri)
	req.Response.WriteXmlExit(rssStr)
}
