package bilibili

import (
	"fmt"
	"rsshub/app/component"
	"rsshub/app/service"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
)

type Controller struct {
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["authority"] = "t.bilibili.com/"
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func getCookieMap() map[string]string {
	cookieMap := service.GetSiteCookies("bilibili")
	return cookieMap
}

func getUsernameFromUserId(id string) string {
	redisKey := "BILI_USERNAME_FROM_ID_" + id
	var username string
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			username = value.String()
		}
	}

	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/space/acc/info?mid=%s&jsonp=jsonp", id)
	if resp := component.GetContent(apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		username = jsonResp.GetString("data.name")
		g.Redis().DoVar("SET", redisKey)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*8)
	}

	return username
}

func getLiveIDFromShortID(id string) string {
	redisKey := "BILI_LIVE_ID_FROM_SHORT_ID_" + id
	var liveID string
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			liveID = value.String()
			return liveID
		}
	}

	apiUrl := "https://api.live.bilibili.com/room/v1/Room/room_init?id=" + id
	if resp := component.GetContent(apiUrl); resp != "" {
		respJson := gjson.New(resp)
		liveID = respJson.GetString("data.room_id")
		g.Redis().DoVar("SET", redisKey, liveID)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*24)
	}

	return liveID
}

func getUsernameFromLiveID(id string) string {
	redisKey := "BILI_USERNAME_FROM_SHORT_ID_" + id
	var username string
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			username = value.String()
			return username
		}
	}

	apiUrl := "https://api.live.bilibili.com/live_user/v1/UserInfo/get_anchor_in_room?roomid=" + id
	if resp := component.GetContent(apiUrl); resp != "" {
		respJson := gjson.New(resp)
		username = respJson.GetString("data.info.uname")
		g.Redis().DoVar("SET", redisKey, username)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*24)
	}
	return username
}
