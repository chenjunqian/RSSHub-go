package chaping

type controller struct {
}

var Controller = &controller{}

type NewsRouteConfig struct {
	Caty  string
	Title string
	Tags  []string
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
			Tags:  []string{"游戏"},
			Title: "游戏"},
		"techNews": {
			Caty:  "3",
			Tags:  []string{"科技"},
			Title: "科技新鲜事"},
		"techFun": {
			Caty:  "5",
			Tags:  []string{"科技"},
			Title: "趣味科技"},
		"debugTime": {
			Caty:  "6",
			Tags:  []string{"科技", "IT"},
			Title: "DEBUG TIME"},
		"internetFun": {
			Caty:  "6",
			Tags:  []string{"互联网"},
			Title: "互联网槽点"},
		"car": {
			Caty:  "9",
			Tags:  []string{"汽车"},
			Title: "公里每小时"},
	}

	return Links
}
