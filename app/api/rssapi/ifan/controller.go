package ifan

import (
	"rsshub/app/component"

	"github.com/anaskhan96/soup"
	"github.com/gogf/gf/frame/g"
)

type Controller struct {
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func parseCommonDetail(detailLink string) (detailData string) {
	var (
		resp string
	)
    if resp = component.GetContent(detailLink); resp != "" {
		var (
			docs        soup.Root
			articleElem soup.Root
			respString  string
		)

		respString = resp
		docs = soup.HTMLParse(respString)
		articleElem = docs.Find("article", "class", "o-single-content__body__content")
		detailData = articleElem.HTML()

	} else {
		g.Log().Errorf("Request ifan article detail failed, link  %s \n", detailLink)
	}

	return
}
