package product

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofrs/uuid/v5"
)

type Single struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price,string"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type All []Single

// UnmarshalJSON handles both empty strings and float64 for the Price field
func (s *Single) UnmarshalJSON(data []byte) error {
	// Create a shadow struct to avoid recursion in UnmarshalJSON
	type Alias Single
	aux := &struct {
		Price string `json:"price"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	// Unmarshal into the shadow struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Handle the Price field
	if aux.Price == "" {
		s.Price = 0 // Default to 0 for an empty string
	} else {
		price, err := strconv.ParseFloat(aux.Price, 64)
		if err != nil {
			return err
		}
		s.Price = price
	}

	return nil
}
