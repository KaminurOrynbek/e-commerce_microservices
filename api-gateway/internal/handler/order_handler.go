package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type OrderHandler struct {
	serviceURL string
	client     *http.Client
}

func NewOrderHandler(serviceURL string) *OrderHandler {
	return &OrderHandler{
		serviceURL: serviceURL,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *OrderHandler) ProxyRequest(c *gin.Context) {
	targetURL, err := url.Parse(h.serviceURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("Error proxying request: %v", err),
		})
	}

	path := c.Param("path")
	c.Request.URL.Path = path

	proxy.ServeHTTP(c.Writer, c.Request)
}
