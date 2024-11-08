package request

type GetCategories struct {
	Page  int
	Limit int
	Name  *string
}

type CreateCategory struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
