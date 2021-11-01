package biz

type RssFeedChannelData struct {
	Id          string
	Title       string
	ChannelDesc string
	ImageUrl    string
	Link        string
	RsshubLink  string
	Count       string
	Sub         int
}

type RssFeedChannelCatalogData struct {
	Id          string
	Title       string
	ChannelDesc string
	ImageUrl    string
	Link        string
	RsshubLink  string
	Count       string
	ItemList    []RssFeedItemData
	Sub         int
}
