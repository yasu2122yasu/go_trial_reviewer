package controller

import (
	"app/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	books := model.GetAll()
	c.HTML(200, "index.html", gin.H{"books": books})
}

// func Show(c *gin.Context) {
// ここに課題3-1を書いてください。
// }

// 投稿作成画面を表示するための関数です。
func GetCreate(c *gin.Context) {
	c.HTML(200, "create.html", gin.H{})
}

// func PostCreate(c *gin.Context) {
// ここに課題3-2を書いてください。
// }

// 編集画面を表示するための関数です。
func GetEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.GetOne(id)
	c.HTML(200, "edit.html", gin.H{"book": book})
}

// func PostEdit(c *gin.Context) {
// ここに課題3-3を書いてください。
// }

// 削除画面を表示するための関数です。
func GetDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.GetOne(id)
	c.HTML(200, "delete.html", gin.H{"book": book})
}

func PostDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Printf("ここに削除用のidを入れる%d", id)
	book := model.GetOne(id)
	book.Delete()
	c.Redirect(301, "/")
}
