package dianshangbao

type Controller struct {
}

type NewsRouteConfig struct {
	ChannelId string
	Title     string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func getNewsLinks() map[string]NewsRouteConfig {
	Links := map[string]NewsRouteConfig{
		"lingshou": {
			ChannelId: "lingshou",
			Title:     "零售"},
		"wuliu": {
			ChannelId: "wuliu",
			Title:     "物流"},
		"O2O": {
			ChannelId: "O2O",
			Title:     "O2O"},
		"zhifu": {
			ChannelId: "zhifu",
			Title:     "支付"},
		"B2B": {
			ChannelId: "B2B",
			Title:     "B2B"},
		"renwu": {
			ChannelId: "renwu",
			Title:     "人物"},
		"kuajing": {
			ChannelId: "kuajing",
			Title:     "跨境电商"},
		"guancha": {
			ChannelId: "guancha",
			Title:     "行业观察"},
	}

	return Links
}
