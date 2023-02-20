package router

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	routerCatRootDir = "resource/public/html/router_catagory/"
	routerCatagoryList []CatagoryDirInfo
)

func InitRouterTemplateInfo(ctx context.Context) {
	routerCatagoryList = getAllCatagoryList(ctx)
}

func GetRouterCatagoryList() []CatagoryDirInfo {
	return routerCatagoryList
}

func getAllCatagoryList(ctx context.Context) []CatagoryDirInfo {
	var (
		catDirs     []*gres.File
		catDirInfos []CatagoryDirInfo
	)
	catDirs = gres.ScanDir(routerCatRootDir, "*", false)
	if len(catDirs) > 0 {
		for _, catDir := range catDirs {
			var (
				catDirInfo     CatagoryDirInfo
				catDirName     string
				catName        string
				catMetaContent []byte
				catMetaJson    *gjson.Json
				subCatDirList  []SubCatagoryDirInfo
			)
			catDirName = splitCatDirectoryName(catDir.Name())
			catMetaContent = gres.GetContent(routerCatRootDir + catDirName + "/meta.json")
			catMetaJson = gjson.New(catMetaContent)
			catName = catMetaJson.Get("name").String()
			subCatDirList = getSubCatInfoList(routerCatRootDir + catDirName)
			catDirInfo = CatagoryDirInfo{
				Name:            catName,
				DirName:         catDirName,
				SubCatagoryList: subCatDirList,
			}
			catDirInfos = append(catDirInfos, catDirInfo)
		}
	}

	return catDirInfos
}

func getSubCatInfoList(dir string) []SubCatagoryDirInfo {
	var (
		subCatDirs    []*gres.File
		subCatDirList []SubCatagoryDirInfo
	)
	subCatDirs = gres.ScanDir(dir, "*", false)
	if len(subCatDirs) > 0 {
		for _, subCatDir := range subCatDirs {
			if !subCatDir.FileInfo().IsDir() {
				continue
			}
			var (
				subCatDirInfo  = SubCatagoryDirInfo{}
				subCatName     string
				subCatDirName  string
				catMetaContent []byte
				catMetaJson    *gjson.Json
				metaJsonPath   string
			)
			subCatDirName = subCatDir.Name()
			metaJsonPath = subCatDirName + "/meta.json"
			catMetaContent = gres.GetContent(metaJsonPath)
			catMetaJson = gjson.New(catMetaContent)
			subCatName = catMetaJson.Get("name").String()
			subCatDirInfo.Name = subCatName
			subCatDirInfo.SubCatagoryDirName = splitCatDirectoryName(subCatDirName)
			subCatDirList = append(subCatDirList, subCatDirInfo)
		}
	}

	return subCatDirList
}

func splitCatDirectoryName(dirName string) string {
	var (
		dirArr []string
	)
	dirArr = gstr.Split(dirName, gfile.Separator)
	if len(dirArr) > 0 {
		return dirArr[len(dirArr)-1]
	}
	return ""
}
