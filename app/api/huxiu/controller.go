package huxiu

type Controller struct {
}

type LinkRouteConfig struct {
	ChannelId string
	Title     string
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func getChannelInfoLinks() map[string]LinkRouteConfig {
	Links := map[string]LinkRouteConfig{
		"auto": {
			ChannelId: "21",
			Title:     "车与出行"},
		"young": {
			ChannelId: "106",
			Title:     "年轻一代"},
		"consumer": {
			ChannelId: "103",
			Title:     "财经"},
		"tech": {
			ChannelId: "105",
			Title:     "科技前沿"},
		"finance": {
			ChannelId: "115",
			Title:     "财经 "},
		"entertainment": {
			ChannelId: "22",
			Title:     "娱乐淘金"},
		"medical": {
			ChannelId: "111",
			Title:     "医疗健康"},
		"culture": {
			ChannelId: "113",
			Title:     "文化教育"},
		"oversea": {
			ChannelId: "114",
			Title:     "出海"},
		"realestate": {
			ChannelId: "102",
			Title:     "金融地产"},
		"enterprise": {
			ChannelId: "110",
			Title:     "企业服务"},
		"startup": {
			ChannelId: "2",
			Title:     "创业"},
		"social": {
			ChannelId: "112",
			Title:     "社交"},
		"global": {
			ChannelId: "107",
			Title:     "全球热点"},
		"life": {
			ChannelId: "4",
			Title:     "生活腔调"},
	}

	return Links
}
