package models

type Addon struct {
	ID          int64  `json:"id"`
	AddonName   string `json:"addon_name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
