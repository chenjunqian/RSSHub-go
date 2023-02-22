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
}
