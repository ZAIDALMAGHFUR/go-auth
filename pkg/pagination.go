package pkg

import (
	"fmt"
	"math"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Link struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}

type LaravelPaginatedResponse struct {
	CurrentPage   int         `json:"current_page"`
	Data          interface{} `json:"data"`
	FirstPageURL  string      `json:"first_page_url"`
	From          int         `json:"from"`
	LastPage      int         `json:"last_page"`
	LastPageURL   string      `json:"last_page_url"`
	Links         []Link      `json:"links"`
	NextPageURL   *string     `json:"next_page_url"`
	Path          string      `json:"path"`
	PerPage       int         `json:"per_page"`
	PrevPageURL   *string     `json:"prev_page_url"`
	To            int         `json:"to"`
	Total         int         `json:"total"`
}

// Struct-style response builder
func BuildPagination(ctx *fiber.Ctx, data interface{}, total, page, limit int) LaravelPaginatedResponse {
	path := strings.Split(ctx.OriginalURL(), "?")[0]
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	firstPageURL := fmt.Sprintf("%s?page=1", path)
	lastPageURL := fmt.Sprintf("%s?page=%d", path, totalPages)

	var nextPageURL *string
	if page < totalPages {
		url := fmt.Sprintf("%s?page=%d", path, page+1)
		nextPageURL = &url
	}

	var prevPageURL *string
	if page > 1 {
		url := fmt.Sprintf("%s?page=%d", path, page-1)
		prevPageURL = &url
	}

	// Build Links ala Laravel
	links := []Link{
		{URL: prevPageURL, Label: "« Previous", Active: false},
	}
	for i := 1; i <= totalPages; i++ {
		url := fmt.Sprintf("%s?page=%d", path, i)
		links = append(links, Link{
			URL:    &url,
			Label:  fmt.Sprintf("%d", i),
			Active: i == page,
		})
	}
	links = append(links, Link{URL: nextPageURL, Label: "Next »", Active: false})

	from := (page-1)*limit + 1
	to := from + limit - 1
	if to > total {
		to = total
	}
	if total == 0 {
		from = 0
		to = 0
	}

	return LaravelPaginatedResponse{
		CurrentPage:  page,
		Data:         data,
		FirstPageURL: firstPageURL,
		From:         from,
		LastPage:     totalPages,
		LastPageURL:  lastPageURL,
		Links:        links,
		NextPageURL:  nextPageURL,
		Path:         path,
		PerPage:      limit,
		PrevPageURL:  prevPageURL,
		To:           to,
		Total:        total,
	}
}

// Map-style pagination builder (tanpa links)
func Build(ctx *fiber.Ctx, data interface{}, total, page, limit int) map[string]interface{} {
	offset := (page - 1) * limit
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	baseURL := ctx.BaseURL() + ctx.Path()
	firstPageURL := fmt.Sprintf("%s?page=1", baseURL)
	lastPageURL := fmt.Sprintf("%s?page=%d", baseURL, totalPages)

	var nextPageURL interface{} = nil
	var prevPageURL interface{} = nil

	if page > 1 {
		prevPageURL = fmt.Sprintf("%s?page=%d", baseURL, page-1)
	}
	if page < totalPages {
		nextPageURL = fmt.Sprintf("%s?page=%d", baseURL, page+1)
	}

	from := offset + 1
	to := offset + limit
	if to > total {
		to = total
	}
	if total == 0 {
		from = 0
		to = 0
	}

	return map[string]interface{}{
		"current_page":   page,
		"data":           data,
		"per_page":       limit,
		"total":          total,
		"from":           from,
		"to":             to,
		"last_page":      totalPages,
		"path":           baseURL,
		"first_page_url": firstPageURL,
		"last_page_url":  lastPageURL,
		"next_page_url":  nextPageURL,
		"prev_page_url":  prevPageURL,
	}
}
