package cyzone

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
		"news": {
			ChannelId: "",
			Title:     "最新"},
		"kuailiyu": {
			ChannelId: "5",
			Title:     "快鲤鱼"},
		"chuangtou": {
			ChannelId: "14",
			Title:     "创投"},
		"kechuang": {
			ChannelId: "13",
			Title:     "科创板"},
		"qiche": {
			ChannelId: "8",
			Title:     "汽车"},
		"haiwai": {
			ChannelId: "10",
			Title:     "海外"},
		"xiaofei": {
			ChannelId: "9",
			Title:     "消费"},
		"keji": {
			ChannelId: "7",
			Title:     "科技"},
		"yiliao": {
			ChannelId: "27",
			Title:     "医疗"},
		"wenyu": {
			ChannelId: "11",
			Title:     "文娱"},
		"chengshi": {
			ChannelId: "16",
			Title:     "城市"},
		"zhengce": {
			ChannelId: "15",
			Title:     "政策"},
		"texie": {
			ChannelId: "6",
			Title:     "特写"},
		"ganhuo": {
			ChannelId: "12",
			Title:     "干货"},
		"kejigu": {
			ChannelId: "33",
			Title:     "科技股"},
	}

	return Links
}
