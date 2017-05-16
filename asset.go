package goshopify

import (
	"fmt"
	"time"
)

// Asset represents a Shopify asset
type Asset struct {
	ThemeID     int        `json:"theme_id"`
	Key         string     `json:"key"`
	Value       string     `json:"value"`
	PublicUrl   string     `json:"public_url"`
	Source      string     `json:"source"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	ContentType string     `json:"content_type"`
	Attachment  string     `json:"attachment"`
	Size        int        `json:"size"`
}

type AssetOptions struct {
	Key     string `url:"asset[key],omitempty"`
	ThemeID int    `url:"theme_id,omitempty"`
}

// Represents the result from the assets/X.json endpoint
type AssetResource struct {
	Asset *Asset `json:"asset"`
}

// Represents the result from the assets.json endpoint
type AssetsResource struct {
	Assets []Asset `json:"assets"`
}

// AssetService handles communication with the asset related methods of
// the Shopify API.
type AssetService struct {
	client *Client
}

// List assets
func (s *AssetService) List(themeID int, options interface{}) ([]Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", themesBasePath, themeID)
	resource := new(AssetsResource)
	err := s.client.Get(path, resource, options)
	return resource.Assets, err
}

// Get individual asset
func (s *AssetService) Get(themeID int, key string) (*Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", themesBasePath, themeID)
	resource := new(AssetResource)
	options := AssetOptions{Key: key, ThemeID: themeID}
	err := s.client.Get(path, resource, options)
	return resource.Asset, err
}

// Update an existing asset.
func (s *AssetService) Upsert(asset *Asset) (*Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", themesBasePath, asset.ThemeID)
	wrappedData := AssetResource{Asset: asset}
	resource := new(AssetResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Asset, err
}
