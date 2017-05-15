package goshopify

import (
	"fmt"
	"time"
)

const themesBasePath = "admin/themes"

// Theme represents a Shopify theme
type Theme struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Role        string     `json:"role"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Previewable bool       `json:"previewable"`
	Processing  bool       `json:"processing"`
}

// Represents the result from the themes/X.json endpoint
type ThemeResource struct {
	Theme *Theme `json:"theme"`
}

// Represents the result from the themes.json endpoint
type ThemesResource struct {
	Themes []Theme `json:"themes"`
}

// ThemeService handles communication with the theme related methods of
// the Shopify API.
type ThemeService struct {
	client *Client
}

// List themes
func (s *ThemeService) List(options interface{}) ([]Theme, error) {
	path := fmt.Sprintf("%s.json", themesBasePath)
	resource := new(ThemesResource)
	err := s.client.Get(path, resource, options)
	return resource.Themes, err
}

// Get individual theme
func (s *ThemeService) Get(id int, options interface{}) (*Theme, error) {
	path := fmt.Sprintf("%s/%d.json", themesBasePath, id)
	resource := new(ThemeResource)
	err := s.client.Get(path, resource, options)
	return resource.Theme, err
}
