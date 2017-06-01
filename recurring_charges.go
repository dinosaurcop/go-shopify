package goshopify

import (
	"fmt"
	"time"
)

const billingBasePath = "admin/recurring_application_charges"

// Theme represents a Shopify theme
type RecurringApplicationCharge struct {
	CappedAmount    string     `json:"capped_amount"`
	ConfirmationUrl string     `json:"confirmation_url"`
	ReturnUrl       string     `json:"return_url"`
	Name            string     `json:"name"`
	Price           string     `json:"price"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	Id              int        `json:"id"`
	Status          string     `json:"status"`
	Test            bool       `json:"test"`
}

type RecurringApplicationChargeRequest struct {
	RecurringApplicationCharge
	Price        float64 `json:"price"`
	CappedAmount float64 `json:"capped_amount"`
}

// Represents the result from the themes/X.json endpoint
type RecurringApplicationChargeResource struct {
	RecurringApplicationCharge *RecurringApplicationChargeRequest `json:"recurring_application_charge"`
}

// ThemeService handles communication with the theme related methods of
// the Shopify API.
type RecurringApplicationChargeService struct {
	client *Client
}

// Create a new webhook
func (s *RecurringApplicationChargeService) Create(charge RecurringApplicationChargeRequest) (*RecurringApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", billingBasePath)
	wrappedData := RecurringApplicationChargeResource{RecurringApplicationCharge: &charge}
	resource := new(RecurringApplicationChargeResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.RecurringApplicationCharge, err
}
