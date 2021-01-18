package chouti

type Controller struct {
}

type NewsRouteConfig struct {
	LinkType string
	Title    string
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
			Title:    "热榜"},
		"news": {
			LinkType: "news",
			Title:    "42区"},
		"scoff": {
			LinkType: "scoff",
			Title:    "段子"},
		"pic": {
			LinkType: "pic",
			Title:    "图片"},
		"tec": {
			LinkType: "tec",
			Title:    "挨踢1024"},
		"ask": {
			LinkType: "ask",
			Title:    "你问我答"},
	}

	return Links
}
