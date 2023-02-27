package router

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	routerCatRootDir   = "resource/template/router_meta/router_catagory/"
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
				subCatHtml     string
			)
			subCatDirName = subCatDir.Name()
			metaJsonPath = subCatDirName + "/meta.json"
			catMetaContent = gres.GetContent(metaJsonPath)
			catMetaJson = gjson.New(catMetaContent)
			subCatName = catMetaJson.Get("name").String()
			subCatHtml = string(gres.GetContent(subCatDirName + "/index.html"))
			subCatDirInfo.Name = subCatName
			subCatDirInfo.SubCatagoryDirName = splitCatDirectoryName(subCatDirName)
			subCatDirInfo.SubCatagoryHtml = subCatHtml
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

func getSubCatRouterMetaInfo(ctx context.Context, metaPath string) SubCatagoryMetaInfo {

	var (
		subMetaInfo *SubCatagoryMetaInfo
		subMetaStr  string
		subMetaJson *gjson.Json
	)

	subMetaStr = string(gres.GetContent(metaPath))
	subMetaJson = gjson.New(subMetaStr)
	subMetaInfo = new(SubCatagoryMetaInfo)
	if err := subMetaJson.Scan(subMetaInfo); err != nil {
		g.Log().Fatal(ctx, "Parse sub-catagory router meta info failed: ", err)
	}

	return *subMetaInfo
}

func genSubCatHtml(ctx context.Context, subCatRouterDetail []SubCatagoryRouterDetail) string {

	var (
		err     error
		cataTpl string
		htmlStr string
		tplView gview.View
	)

	cataTpl = string(gres.GetContent(routerCatRootDir + "cata_tpl.html"))
	if cataTpl == "" {
		g.Log().Fatal(ctx, "Failed to get catagorey template")
	}

	tplView = *gview.New()
	tplView.SetConfigWithMap(g.Map{
		"Data": g.Map{
			"routers": subCatRouterDetail,
		},
	})

	htmlStr, err = tplView.ParseContent(ctx, cataTpl)
	if err != nil {
		g.Log().Fatal(ctx, " Generate sub-catagory router html failed, error : ", err)
	}

	return htmlStr
}
