package product_record

import (
	"database/sql"
	"errors"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
)

const (
	InsertQuery = "INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)"
	ExistsQuery = "SELECT product_id, last_update_date FROM product_records WHERE product_id=? AND last_update_date=?"
	GetQuery    = "SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records WHERE id=?"
)

type Repository interface {
	Save(productRecord domain.ProductRecord) int
	Exists(productId int, lastUpdateDate time.Time) bool
	Get(id int) *domain.ProductRecord
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(productRecord domain.ProductRecord) int {
	stmt, err := r.db.Prepare(InsertQuery)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(productRecord.LastUpdateDate, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductID)
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(id)
}

func (r *repository) Exists(productId int, lastUpdateDate time.Time) bool {
	row := r.db.QueryRow(ExistsQuery, productId, lastUpdateDate)
	lastUpdateDateString := helpers.ToFormattedDateTime(lastUpdateDate)
	err := row.Scan(&productId, &lastUpdateDateString)
	return err == nil
}

func (r *repository) Get(id int) *domain.ProductRecord {
	row := r.db.QueryRow(GetQuery, id)
	productRecord := domain.ProductRecord{}
	var lastUpdateDate string

	err := row.Scan(&productRecord.ID, &lastUpdateDate, &productRecord.PurchasePrice, &productRecord.SalePrice, &productRecord.ProductID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	productRecord.LastUpdateDate = helpers.ToDateTime(lastUpdateDate)

	return &productRecord
}
