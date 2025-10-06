package main

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/krishna102001/ecommerce-microservice/order"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateAccount(ctx context.Context, in AccountInput) (*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.PostAccount(ctx, in.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Account{
		ID:   a.ID,
		Name: a.Name,
	}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, in ProductInput) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	inPrice := strconv.FormatFloat(in.Price, 'f', 2, 2)
	p, err := r.server.catalogClient.PostProduct(ctx, in.Name, in.Description, inPrice)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	pPrice, _ := strconv.ParseFloat(p.Price, 2)
	return &Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       pPrice,
	}, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, in OrderInput) (*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var products []order.OrderedProduct
	for _, p := range in.Products {
		if p.Quantity <= 0 {
			return nil, ErrInvalidParameter
		}
		pQuantity := strconv.Itoa(p.Quantity)
		products = append(products, order.OrderedProduct{
			ID:       p.ID,
			Quantity: pQuantity,
		})
	}
	o, err := r.server.orderClient.PostOrder(ctx, in.AccountID, products)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	oTotalPrice, _ := strconv.ParseFloat(o.TotalPrice, 2)
	return &Order{
		ID:         o.ID,
		CreatedAt:  o.CreatedAt,
		TotalPrice: oTotalPrice,
	}, nil
}
