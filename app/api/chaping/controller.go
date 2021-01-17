package chaping

type Controller struct {
}

type NewsRouteConfig struct {
	Caty  string
	Title string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func getNewsLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"game": {
			Caty:  "1",
			Title: "游戏"},
		"techNews": {
			Caty:  "3",
			Title: "科技新鲜事"},
		"techFun": {
			Caty:  "5",
			Title: "趣味科技"},
		"debugTime": {
			Caty:  "6",
			Title: "DEBUG TIME"},
		"internetFun": {
			Caty:  "6",
			Title: "互联网槽点"},
		"car": {
			Caty:  "9",
			Title: "公里每小时"},
	}

	return Links
}
