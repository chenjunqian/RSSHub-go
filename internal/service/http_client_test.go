package service

import (
	"context"
	"testing"
)



func TestGetContent(t *testing.T) {
	resp := GetContent(context.TODO(), "https://www.baidu.com/")
	if resp == "" {
		t.Fatal("Get content failed")
	}
}