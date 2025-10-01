package order

import (
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error)
	GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error)
}

type Order struct {
	ID         string
	CreatedAt  time.Time
	TotalPrice string
	AccountID  string
	Products   []OrderedProduct
}

type OrderedProduct struct {
	ID          string
	Name        string
	Description string
	Price       string
	Quantity    string
}

type orderService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &orderService{r}
}

func (s *orderService) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {
	o := &Order{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().UTC(),
		AccountID: accountID,
		Products:  products,
	}
	o.TotalPrice = "0.0"
	for _, p := range products {
		pPrice, _ := strconv.Atoi(p.Price)
		pQuantity, _ := strconv.Atoi(p.Quantity)
		oTotalPrice, _ := strconv.Atoi(o.TotalPrice)
		oTotalPrice += pPrice * pQuantity
		o.TotalPrice = strconv.Itoa(oTotalPrice)
	}
	err := s.repository.PutOrder(ctx, *o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (s *orderService) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	return s.repository.GetOrdersForAcccount(ctx, accountID)
}
