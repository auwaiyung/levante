package controllers

import "github.com/kataras/iris/mvc"

type CommentController struct {
	BaseController
}

func (this *CommentController) ReadHandle() mvc.Result{
	return ViewPageNotFound
}

func (this *CommentController) CreateHandle() mvc.Result{
	return ViewPageNotFound
}