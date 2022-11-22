package bilibili

import (
	"context"
	"fmt"
	"rsshub/app/component"
	"rsshub/app/service"

	"github.com/gogf/gf/v2/encoding/gjson"
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

func getCookieMap(ctx context.Context) map[string]string {
	cookieMap := service.GetSiteCookies(ctx, "bilibili")
	return cookieMap
}

func getUsernameFromUserId(ctx context.Context, id string) string {
	redisKey := "BILI_USERNAME_FROM_ID_" + id
	var username string
	if value, err := component.GetRedis().Do(ctx, "GET", redisKey); err == nil {
		if value.String() != "" {
			username = value.String()
		}
	}

	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/space/acc/info?mid=%s&jsonp=jsonp", id)
	if resp := component.GetContent(ctx, apiUrl); resp != "" {
		jsonResp := gjson.New(resp)
		username = jsonResp.Get("data.name").String()
		component.GetRedis().Do(ctx, "SET", redisKey)
		component.GetRedis().Do(ctx, "EXPIRE", redisKey, 60*60*8)
	}

	return username
}

func getLiveIDFromShortID(ctx context.Context, id string) string {
	redisKey := "BILI_LIVE_ID_FROM_SHORT_ID_" + id
	var liveID string
	if value, err := component.GetRedis().Do(ctx, "GET", redisKey); err == nil {
		if value.String() != "" {
			liveID = value.String()
			return liveID
		}
	}

	apiUrl := "https://api.live.bilibili.com/room/v1/Room/room_init?id=" + id
	if resp := component.GetContent(ctx, apiUrl); resp != "" {
		respJson := gjson.New(resp)
		liveID = respJson.Get("data.room_id").String()
		component.GetRedis().Do(ctx, "SET", redisKey, liveID)
		component.GetRedis().Do(ctx, "EXPIRE", redisKey, 60*60*24)
	}

	return liveID
}

func getUsernameFromLiveID(ctx context.Context, id string) string {
	redisKey := "BILI_USERNAME_FROM_SHORT_ID_" + id
	var username string
	if value, err := component.GetRedis().Do(ctx, "GET", redisKey); err == nil {
		if value.String() != "" {
			username = value.String()
			return username
		}
	}

	apiUrl := "https://api.live.bilibili.com/live_user/v1/UserInfo/get_anchor_in_room?roomid=" + id
	if resp := component.GetContent(ctx, apiUrl); resp != "" {
		respJson := gjson.New(resp)
		username = respJson.Get("data.info.uname").String()
		component.GetRedis().Do(ctx, "SET", redisKey, username)
		component.GetRedis().Do(ctx, "EXPIRE", redisKey, 60*60*24)
	}
	return username
}
