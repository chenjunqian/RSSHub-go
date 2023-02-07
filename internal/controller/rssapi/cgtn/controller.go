package cgtn

import (
	"context"
	"rsshub/internal/service"

	"github.com/anaskhan96/soup"
)

type controller struct {
}

var Controller = &controller{}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func getMainContent(ctx context.Context, url string) (content string) {
	if resp := service.GetContent(ctx,url); resp != "" {
		docs := soup.HTMLParse(resp)
		contentElem := docs.Find("div", "id", "cmsMainContent")
		if contentElem.Error == nil {
			content = contentElem.HTML()
		}
	}
	return
}
