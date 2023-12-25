package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ishanshre/GoRestApiExample/internals/models"
)

func (h *handler) HomeHandlerHtmx(c *gin.Context) {
	authors, err := h.repo.GetAllAuthors()
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"error": "Error getting authors",
		})
		return
	}
	data := make(map[string]interface{})
	data["authors"] = authors
	c.HTML(http.StatusOK, "index.tmpl", &models.TemplateData{
		Data: data,
	})
}

func (h *handler) AddAuthorHandlerHtmx(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"error": "error parsing form",
		})
		return
	}
	first_name := c.Request.Form.Get("first_name")
	last_name := c.Request.Form.Get("last_name")
	email := c.Request.Form.Get("email_address")
	author := models.Author{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
	}
	author_data, err := h.repo.CreateAuthor(author)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.tmpl", gin.H{
			"error": "error in creating author",
		})
		return
	}
	data := make(map[string]interface{})
	data["author"] = author_data

	c.HTML(http.StatusOK, "author", author_data)
}

func (h *handler) DeleteAuthorHandlerHtmx(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", nil)
		return
	}
	if err := h.repo.DeleteAuthorByID(id); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", nil)
		log.Println(err.Error())
		return
	}
	authors, _ := h.repo.GetAllAuthors()
	data := make(map[string]interface{})
	data["authors"] = authors
	c.HTML(http.StatusNoContent, "index.tmpl", &models.TemplateData{
		Data: data,
	})
}
