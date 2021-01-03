package bilibili

import (
	"fmt"

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

func getUsernameFromUserId(id string) string {
	redisKey := "BILI_USERNAME_FROM_ID_" + id
	var username string
	if value, err := g.Redis().DoVar("GET", redisKey); err == nil {
		if value.String() != "" {
			username = value.String()
		}
	}

	apiUrl := fmt.Sprintf("https://api.bilibili.com/x/space/acc/info?mid=%s&jsonp=jsonp", id)
	headers := getHeaders()
	if resp, err := g.Client().SetHeaderMap(headers).Get(apiUrl); err == nil {
		jsonResp := gjson.New(resp.ReadAllString())
		username = jsonResp.GetString("data.name")
		g.Redis().DoVar("SET", redisKey)
		g.Redis().DoVar("EXPIRE", redisKey, 60*60*8)
	}

	return username
}
