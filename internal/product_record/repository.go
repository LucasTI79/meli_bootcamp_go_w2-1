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

	CountRecordsByAllProductsQuery = `SELECT p.id "product_id", p.description, count(pr.id) "records_count"
		FROM products p
		LEFT JOIN product_records pr ON p.id = pr.product_id
		GROUP BY p.id`

	CountRecordsByProductQuery = `SELECT p.id "product_id", p.description, count(pr.id) "records_count"
		FROM products p
		LEFT JOIN product_records pr ON p.id = pr.product_id
		WHERE p.id=?
		GROUP BY p.id`
)

type Repository interface {
	Save(productRecord domain.ProductRecord) int
	Exists(productId int, lastUpdateDate time.Time) bool
	Get(id int) *domain.ProductRecord
	CountRecordsByAllProducts() []domain.RecordsByProductReport
	CountRecordsByProduct(id int) *domain.RecordsByProductReport
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

func (r *repository) CountRecordsByAllProducts() []domain.RecordsByProductReport {
	rows, err := r.db.Query(CountRecordsByAllProductsQuery)
	if err != nil {
		panic(err)
	}

	recordsByProducts := make([]domain.RecordsByProductReport, 0)

	for rows.Next() {
		record := domain.RecordsByProductReport{}
		_ = rows.Scan(&record.ProductID, &record.Description, &record.RecordsCount)
		recordsByProducts = append(recordsByProducts, record)
	}

	return recordsByProducts
}

func (r *repository) CountRecordsByProduct(id int) *domain.RecordsByProductReport {
	rows := r.db.QueryRow(CountRecordsByProductQuery, id)
	record := domain.RecordsByProductReport{}
	err := rows.Scan(&record.ProductID, &record.Description, &record.RecordsCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		panic(err)
	}

	return &record
}
