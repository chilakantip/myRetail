package main

import (
	"fmt"
	"strings"

	"github.com/chilakantip/my_retail/mg_persist"
)

type product struct {
	ProductID    int `json:"product_id"`
	CurrentPrise struct {
		Value        float64 `json:"value"`
		CurrencyCode string  `json:"currency_code"`
	} `json:"current_prise"`

	ProductDetails *mg_persist.ProductDetails `json:"product_details"`
}

func (p *product) validate() error {
	if strings.TrimSpace(p.ProductDetails.Name) == "" {
		return fmt.Errorf("product name is empty")
	}
	if strings.TrimSpace(p.ProductDetails.Description) == "" {
		return fmt.Errorf("product description is empty")
	}
	if strings.TrimSpace(p.ProductDetails.Type) == "" {
		return fmt.Errorf("product type is empty")
	}

	if p.CurrentPrise.Value <= 0.0 {
		return fmt.Errorf("product price is invalid")
	}
	if strings.TrimSpace(p.CurrentPrise.CurrencyCode) == "" {
		return fmt.Errorf("product prise currency is empty")
	}
	return nil
}
