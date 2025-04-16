package http

import (
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/domain"
	"github.com/KaminurOrynbek/e-commerce_microservices/inventory_service/internal/handler/http/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryHandler(repo domain.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{
		categoryRepo: repo,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	category := req.ToCategory()
	if err := h.categoryRepo.Create(c.Request.Context(), category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.FromCategory(category))
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	category, err := h.categoryRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category: " + err.Error()})
		return
	}
	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, dto.FromCategory(category))
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	categories, err := h.categoryRepo.List(c.Request.Context(), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories: " + err.Error()})
		return
	}

	response := dto.CategoryListResponse{
		Data: make([]dto.CategoryResponse, len(categories)),
		Meta: struct {
			Page  int `json:"page"`
			Limit int `json:"limit"`
		}{
			Page:  page,
			Limit: limit,
		},
	}

	for i, category := range categories {
		response.Data[i] = *dto.FromCategory(category)
	}

	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	existingCategory, err := h.categoryRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category: " + err.Error()})
		return
	}
	if existingCategory == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	existingCategory.Update(req.Name, req.Description)

	if err := h.categoryRepo.Update(c.Request.Context(), existingCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.FromCategory(existingCategory))
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID format"})
		return
	}

	existingCategory, err := h.categoryRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category: " + err.Error()})
		return
	}
	if existingCategory == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := h.categoryRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *CategoryHandler) RegisterRoutes(router *gin.Engine) {
	categories := router.Group("/api/categories")
	{
		categories.POST("", h.CreateCategory)
		categories.GET("/:id", h.GetCategory)
		categories.GET("", h.ListCategories)
		categories.PUT("/:id", h.UpdateCategory)
		categories.DELETE("/:id", h.DeleteCategory)
	}
}
