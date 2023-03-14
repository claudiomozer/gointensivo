package database

import (
	"database/sql"

	"github.com/claudiomozer/taxas/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	return err
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select COUNT(id) from orders").Scan(&total)
	return total, err
}
