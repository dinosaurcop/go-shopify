package goshopify

import (
	"fmt"
	"time"
)

const pagesBasePath = "admin/pages"

type MetaField struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueType string `json:"value_type"`
	Namespace string `json:"namespace"`
}
type Page struct {
	Author         string      `json:"author"`
	Title          string      `json:"title"`
	BodyHTML       string      `json:"body_html"`
	Handle         string      `json:"handle"`
	ID             int         `json:"id"`
	Metafield      []MetaField `json:"metafield"`
	ShopID         int         `json:"shop_id"`
	TemplateSuffix *string     `json:"template_suffix, omitempty"`
	CreatedAt      *time.Time  `json:"created_at"`
	UpdatedAt      *time.Time  `json:"updated_at"`
}

// Represents the result from the pages/X.json endpoint
type PageResource struct {
	Page *Page `json:"page"`
}

// Represents the result from the pages.json endpoint
type PagesResource struct {
	Pages []Page `json:"pages"`
}

// PageService handles communication with the page related methods of
// the Shopify API.
type PageService struct {
	client *Client
}

// List pages
func (s *PageService) List(options interface{}) ([]Page, error) {
	path := fmt.Sprintf("%s.json", pagesBasePath)
	resource := new(PagesResource)
	err := s.client.Get(path, resource, options)
	return resource.Pages, err
}

// Get individual page
func (s *PageService) Get(id int, options interface{}) (*Page, error) {
	path := fmt.Sprintf("%s/%d.json", pagesBasePath, id)
	resource := new(PageResource)
	err := s.client.Get(path, resource, options)
	return resource.Page, err
}

// Create a new page
func (s *PageService) Create(page Page) (*Page, error) {
	path := fmt.Sprintf("%s.json", pagesBasePath)
	wrappedData := PageResource{Page: &page}
	resource := new(PageResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Page, err
}

// Update an existing page.
func (s *PageService) Update(page Page) (*Page, error) {
	path := fmt.Sprintf("%s/%d.json", pagesBasePath, page.ID)
	wrappedData := PageResource{Page: &page}
	resource := new(PageResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Page, err
}
