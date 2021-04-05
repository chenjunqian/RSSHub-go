package chouti

type Controller struct {
}

type NewsRouteConfig struct {
	LinkType string
	Title    string
	Tags     []string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func getNewsLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"hot": {
			LinkType: "hot",
			Tags:     []string{"社区", "门户", "其他"},
			Title:    "热榜"},
		"news": {
			LinkType: "news",
			Tags:     []string{"社区", "门户", "其他"},
			Title:    "42区"},
		"scoff": {
			LinkType: "scoff",
			Tags:     []string{"社区", "搞笑"},
			Title:    "段子"},
		"pic": {
			LinkType: "pic",
			Tags:     []string{"图片"},
			Title:    "图片"},
		"tec": {
			LinkType: "tec",
			Tags:     []string{"科技", "IT"},
			Title:    "挨踢1024"},
		"ask": {
			LinkType: "ask",
			Tags:     []string{"社区", "门户", "问答"},
			Title:    "你问我答"},
	}

	return Links
}
