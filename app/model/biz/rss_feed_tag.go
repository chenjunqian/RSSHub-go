package biz

type RespFeedTagData struct {
	Name  string `orm:"name"       json:"name"`
	Count string `orm:"count"       json:"count"`
}
