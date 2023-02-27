package router

type CatagoryDirInfo struct {
	DirName         string
	Name            string
	SubCatagoryList []SubCatagoryDirInfo
	CollapseOpen    bool
}

type SubCatagoryDirInfo struct {
	SubCatagoryDirName string
	Name               string
	SubCatagoryHtml    string
}

type SubCatagoryMetaInfo struct {
	Name    string
	Routers []SubCatagoryRouterDetail
}

type SubCatagoryRouterDetail struct {
	Name   string
	Router string
	Link   string
}
