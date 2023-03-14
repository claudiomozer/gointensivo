package usecase

import "github.com/claudiomozer/taxas/internal/entity"

type OrderInputDto struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDto struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func (c *CalculateFinalPrice) Execute(input OrderInputDto) (*OrderOutputDto, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)

	if err != nil {
		return &OrderOutputDto{}, nil
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return &OrderOutputDto{}, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return &OrderOutputDto{}, err
	}

	return &OrderOutputDto{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
