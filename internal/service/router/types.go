package router

type CatagoryDirInfo struct {
	DirName         string
	Name            string
	SubCatagoryList []SubCatagoryDirInfo
}

type SubCatagoryDirInfo struct {
	SubCatagoryDirName string
	Name               string
}
