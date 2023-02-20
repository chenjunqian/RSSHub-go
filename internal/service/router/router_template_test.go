package router

import (
	_ "rsshub/internal/packed"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestInitRouterTemplateInfo(t *testing.T) {

}

func TestGetAllCatagoryList(t *testing.T) {
	var (
		ctx            = gctx.New()
		catDirNameList []CatagoryDirInfo
	)
	catDirNameList = GetAllCatagoryList(ctx)
	if len(catDirNameList) == 0 {
		t.Fatal("Get catagory directory list failed")
	}
}

func Test_splitCatDirName(t *testing.T) {
	var (
		catDirName        = "test_dir"
		fullCatDirName    = routerCatRootDir + catDirName
		splitCatDirResult string
	)

	splitCatDirResult = splitCatDirectoryName(fullCatDirName)
	if splitCatDirResult != catDirName {
		t.Fatal("split catagory directory failed.")
	}
}

func Test_getSubCatInfoList(t *testing.T) {
	var (
		catDirName = "new_media"	
		subCatDirInfoList []SubCatagoryDirInfo
	)
	subCatDirInfoList = getSubCatInfoList(routerCatRootDir + catDirName)
	if len(subCatDirInfoList) == 0 {
		t.Fatal("Get sub catagory info failed")
	}
}
