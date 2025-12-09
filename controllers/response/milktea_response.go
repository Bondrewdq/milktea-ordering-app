package response

type MilkTeaResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
}

type ListMilkTeaResponse struct {
	Total  int               `json:"total"`
	Items  []MilkTeaResponse `json:"items"`
	Page   int               `json:"page"`
	PageSize int               `json:"page_size"`
}