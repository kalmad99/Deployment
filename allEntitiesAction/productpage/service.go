package productpage

import "github.com/kalmad99/Deployment/entity"

type CategoryService interface {
	Categories() ([]entity.Category, []error)
	Category(id uint) (*entity.Category, []error)
	UpdateCategory(category *entity.Category) (*entity.Category, []error)
	DeleteCategory(id uint) (*entity.Category, []error)
	StoreCategory(category *entity.Category) (*entity.Category, []error)
	ItemsInCategory(category *entity.Category) ([]entity.Product, []error)
}

// ItemService specifies item related database operations
type ItemService interface {
	Items() ([]entity.Product, []error)
	Item(id uint) (*entity.Product, []error)
	UpdateItem(product *entity.Product) (*entity.Product, []error)
	DeleteItem(id uint) (*entity.Product, []error)
	StoreItem(product *entity.Product) (*entity.Product, []error)
	RateProduct(product *entity.Product) (*entity.Product, []error)
	SearchProduct(index string) ([]entity.Product, error)
	StoreItemCateg(product *entity.Product) []error
}
