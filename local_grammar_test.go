package fmr

import (
	"testing"

	"github.com/liuzl/ling"
	"zliu.org/goutil"
)

func TestLocalGrammar(t *testing.T) {
	tests := []string{
		`天津，liang@zliu.org是我的邮箱，https://crawler.club是爬虫主页`,
		`关于FMR的介绍在这里：https://zliu.org/project/fmr/,好的`,
		`柏乡县是一个历史悠久的