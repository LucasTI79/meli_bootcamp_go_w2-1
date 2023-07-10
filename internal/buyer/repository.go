package buyer

import (
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

const (
	GetAllQuery = "SELECT id, card_number_id, first_name, last_name FROM buyers"
	GetQuery    = "SELECT id, card_number_id, first_name, last_name FROM buyers WHERE id = ?;"
	ExistsQuery = "SELECT card_number_id FROM buyers WHERE card_number_id=?;"
	InsertQuery = "INSERT INTO buyers(card_number_id,first_name,last_name) VALUES (?,?,?)"
	UpdateQuery = "UPDATE buyers SET card_number_id=?, first_name=?, last_name=?  WHERE id=?"
	DeleteQuery = "DELETE FROM buyers WHERE id = ?"

	CountPurchasesByAllBuyers = `SELECT b.id, b.card_number_id, b.first_name, b.last_name, count(po.id) "purchase_orders_count"
		FROM buyers b
		JOIN purchase_orders po ON b.id = po.buyer_id
		GROUP BY b.id`

	CountPurchasesByBuyer = `SELECT b.id, b.card_number_id, b.first_name, b.last_name, count(po.id) "purchase_orders_count"
		FROM buyers b
		JOIN purchase_orders po ON b.id = po.buyer_id
		WHERE b.id=?
		GROUP BY b.id`
)

type Repository interface {
	GetAll() []domain.Buyer
	Get(id int) *domain.Buyer
	Exists(cardNumberID string) bool
	Save(b domain.Buyer) int
	Update(b domain.Buyer)
	Delete(id int)
	CountPurchasesByAllBuyers() []domain.PurchasesByBuyerReport
	CountPurchasesByBuyer(id int) *domain.PurchasesByBuyerReport
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() []domain.Buyer {
	rows, err := r.db.Query(GetAllQuery)
	if err != nil {
		panic(err)
	}

	buyers := make([]domain.Buyer, 0)

	for rows.Next() {
		b := domain.Buyer{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
		buyers = append(buyers, b)
	}

	return buyers
}

func (r *repository) Get(id int) *domain.Buyer {
	row := r.db.QueryRow(GetQuery, id)
	b := domain.Buyer{}
	err := row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &b
}

func (r *repository) Exists(cardNumberID string) bool {
	row := r.db.QueryRow(ExistsQuery, cardNumberID)
	err := row.Scan(&cardNumberID)
	return err == nil
}

func (r *repository) Save(b domain.Buyer) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Update(b domain.Buyer) {
	stmt, err := r.db.Prepare(UpdateQuery)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName, &b.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(id int) {
	stmt, err := r.db.Prepare(DeleteQuery)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}

func (r *repository) CountPurchasesByAllBuyers() []domain.PurchasesByBuyerReport {
	rows, err := r.db.Query(CountPurchasesByAllBuyers)
	if err != nil {
		panic(err)
	}

	purchasesByBuyer := make([]domain.PurchasesByBuyerReport, 0)

	for rows.Next() {
		b := domain.PurchasesByBuyerReport{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName, &b.PurchasesCount)
		purchasesByBuyer = append(purchasesByBuyer, b)
	}

	return purchasesByBuyer
}

func (r *repository) CountPurchasesByBuyer(id int) *domain.PurchasesByBuyerReport {
	rows := r.db.QueryRow(CountPurchasesByBuyer, id)
	b := domain.PurchasesByBuyerReport{}
	err := rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName, &b.PurchasesCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &b
}
