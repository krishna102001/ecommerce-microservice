package main

import (
	"context"
	"log"
	"strconv"
	"time"
)

type accountResolver struct {
	server *Server
}

func (r *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	orderList, err := r.server.orderClient.GetOrdersForAccount(ctx, obj.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var orders []*Order
	for _, o := range orderList {
		var products []*OrderedProduct
		for _, p := range o.Products {
			pQuantity, _ := strconv.Atoi(p.Quantity)
			pPrice, _ := strconv.ParseFloat(p.Price, 2)
			products = append(products, &OrderedProduct{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Quantity:    pQuantity,
				Price:       pPrice,
			})
		}
		oTotalPrice, _ := strconv.ParseFloat(o.TotalPrice, 2)
		orders = append(orders, &Order{
			ID:         o.ID,
			CreatedAt:  o.CreatedAt,
			TotalPrice: oTotalPrice,
			Products:   products,
		})
	}
	return orders, nil
}
