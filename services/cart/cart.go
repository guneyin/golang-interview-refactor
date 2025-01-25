package cart

import (
	"interview/repositories/cart"
)

type Service struct {
	repo cart.Repository
}

func NewService() *Service {
	return &Service{repo: cart.NewRepository()}
}

func (s *Service) Get(sessionID string) (map[string]any, error) {
	items, err := s.repo.GetCart(sessionID)
	if err != nil {
		return nil, err
	}

	dataItems := make([]map[string]any, len(items))
	for i, item := range items {
		dataItems[i] = make(map[string]any)
		dataItems[i]["ID"] = item.ID
		dataItems[i]["Quantity"] = item.Quantity
		dataItems[i]["Price"] = item.Price
		dataItems[i]["Product"] = item.ProductName
	}

	data := make(map[string]any)
	data["CartItems"] = dataItems
	return data, nil
}

func (s *Service) Add(sessionID, product string, qty uint) error {
	return s.repo.AddItem(sessionID, product, qty)
}

func (s *Service) Delete(sessionID string, id uint) error {
	return s.repo.DeleteItem(sessionID, id)
}
