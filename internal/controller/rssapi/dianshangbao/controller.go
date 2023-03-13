package dianshangbao

type controller struct {
}

var Controller = &controller{}

type NewsRouteConfig struct {
	ChannelId string
	Title     string
	Tags      []string
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
			Tags:      []string{"电商", "零售"},
			Title:     "零售"},
		"wuliu": {
			ChannelId: "wuliu",
			Tags:      []string{"电商", "物流"},
			Title:     "物流"},
		"O2O": {
			ChannelId: "O2O",
			Tags:      []string{"电商"},
			Title:     "O2O"},
		"zhifu": {
			ChannelId: "zhifu",
			Tags:      []string{"电商", "金融"},
			Title:     "支付"},
		"B2B": {
			ChannelId: "B2B",
			Tags:      []string{"电商"},
			Title:     "B2B"},
		"renwu": {
			ChannelId: "renwu",
			Tags:      []string{"电商", "其他"},
			Title:     "人物"},
		"kuajing": {
			ChannelId: "kuajing",
			Tags:      []string{"电商", "海外"},
			Title:     "跨境电商"},
		"guancha": {
			ChannelId: "guancha",
			Tags:      []string{"电商", "其他"},
			Title:     "行业观察"},
	}

	return Links
}
