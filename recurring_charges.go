package goshopify

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

const billingBasePath = "admin/recurring_application_charges"

// Theme represents a Shopify theme
type RecurringApplicationCharge struct {
	CappedAmount    decimal.Decimal `json:"capped_amount"`
	ConfirmationUrl string          `json:"confirmation_url"`
	ReturnUrl       string          `json:"return_url"`
	Name            string          `json:"name"`
	Price           decimal.Decimal `json:"price"`
	CreatedAt       *time.Time      `json:"created_at"`
	UpdatedAt       *time.Time      `json:"updated_at"`
	Id              int             `json:"id"`
	Status          string          `json:"status"`
	Terms           string          `json:"terms"`
	Test            string          `json:"test,omitempty"`
}

// Represents the result from the themes/X.json endpoint
type RecurringApplicationChargeResource struct {
	RecurringApplicationCharge *RecurringApplicationCharge `json:"recurring_application_charge"`
}

// ThemeService handles communication with the theme related methods of
// the Shopify API.
type RecurringApplicationChargeService struct {
	client *Client
}

// Create a new webhook
func (s *RecurringApplicationChargeService) Create(charge RecurringApplicationCharge) (*RecurringApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", billingBasePath)
	wrappedData := RecurringApplicationChargeResource{RecurringApplicationCharge: &charge}
	resource := new(RecurringApplicationChargeResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.RecurringApplicationCharge, err
}
