package controllers

import (
	"net/http"
	"uooobarry/gin-notes/models"
	"uooobarry/gin-notes/services"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
	CategoryService *services.CategoryService
}

func NewHomeController(service *services.CategoryService) *HomeController {
	return &HomeController{CategoryService: service}
}

func (ctrl *HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"message": "Welcome to the Home Page",
	})
}

func (ctrl *HomeController) CategoriesList(c *gin.Context) {
	categories, err := ctrl.CategoryService.CategoryRepo.List()
	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, c, err)
		return
	}

	c.HTML(http.StatusOK, "categories.html", gin.H{
		"categories": categories,
	})
}

func (ctrl *HomeController) NewCategory(c *gin.Context) {
	c.HTML(http.StatusOK, "new_category_input.html", gin.H{})
}

func (ctrl *HomeController) CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBind(&category); err != nil {
		RenderErrorPage(http.StatusBadRequest, c, err)
		return
	}
	if err := ctrl.CategoryService.CategoryRepo.Create(&category); err != nil {
		RenderErrorPage(http.StatusInternalServerError, c, err)
		return
	}

	categories, err := ctrl.CategoryService.CategoryRepo.List()
	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, c, err)
		return
	}
	c.HTML(http.StatusOK, "categories.html", gin.H{
		"categories": categories,
	})
}
