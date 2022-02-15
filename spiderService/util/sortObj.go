package util

import "Spider/spiderService/model"

type BybitArticleSlice []model.BybitArticle

func (a BybitArticleSlice) Len() int           { return len(a) }
func (a BybitArticleSlice) Less(i, j int) bool { return i > j }
func (a BybitArticleSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type BybitNewlyArticleSlice []model.BybitNewlyArticle

func (a BybitNewlyArticleSlice) Len() int           { return len(a) }
func (a BybitNewlyArticleSlice) Less(i, j int) bool { return i > j }
func (a BybitNewlyArticleSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type SlateArticleSlice []model.SlateArticle

func (a SlateArticleSlice) Len() int           { return len(a) }
func (a SlateArticleSlice) Less(i, j int) bool { return i > j }
func (a SlateArticleSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
