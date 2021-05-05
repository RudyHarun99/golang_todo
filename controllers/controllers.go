package controllers

import (
	"github.com/gin-gonic/gin"
	"GO14/finalProject/models"
	"gorm.io/gorm"
)

// GetTodos godoc
// @Summary Get all todos
// @Description get all todos
// @ID get-todos
// @Produce  json
// @Success 200 {array} models.Todos
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var todos []models.Todos
	db.Debug().Find(&todos)

	c.JSON(200, gin.H{
		"todos": todos,
	})
}

// CreateTodo godoc
// @Summary Create a todo
// @Description create a new todo
// @ID create-todo
// @Accept  json
// @Produce  json
// @Param requestbody body models.Todos true "request body json"
// @Success 200 {object} models.Todos
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	title := c.PostForm("Title")
	description := c.PostForm("Description")
	status := c.PostForm("Status")

	data := models.Todos{
		Title: title,
		Description: description,
		Status: status,
	}
	
	db.Debug().Create(&data)

	c.JSON(200, gin.H{
		"data": data,
	})
}

// GetTodo godoc
// @Summary Get a todo
// @Description get a todo
// @ID get-todo
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todos
// @Router /todos/{id} [get]
func GetTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var todo models.Todos
	db.Debug().Where("id = ?", id).Find(&todo)

	c.JSON(200, gin.H{
		"todo": todo,
	})
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description update a todo
// @ID update-todo
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param requestbody body models.Todos true "request body json"
// @Success 200 {object} models.Todos
// @Router /todos/{id} [put]
func EditTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	title := c.PostForm("Title")
	description := c.PostForm("Description")
	status := c.PostForm("Status")

	data := models.Todos{
		Title: title,
		Description: description,
		Status: status,
	}

	var todo models.Todos
	db.Debug().Model(&todo).Where("id = ?", id).Updates(&data)
	db.Debug().Where("id = ?", id).Find(&todo)

	c.JSON(200, gin.H{
		"message": todo,
	})
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description delete a todo
// @ID delete-todo
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todos
// @Router /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var todo models.Todos
	db.Debug().Where("id = ?", id).Delete(&todo)

	c.JSON(200, gin.H{
		"message": id,
	})
}