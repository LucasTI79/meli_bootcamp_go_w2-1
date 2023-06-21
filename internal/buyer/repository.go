package buyer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Repository encapsulates the storage of a buyer.
type IRepository interface {
	GetAll(ctx context.Context) []domain.Buyer
	Get(ctx context.Context, id int) *domain.Buyer
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, b domain.Buyer) int
	Update(ctx context.Context, b domain.Buyer)
	Delete(ctx context.Context, id int)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) []domain.Buyer {
	query := "SELECT * FROM buyers"
	rows, err := r.db.Query(query)
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

func (r *repository) Get(ctx context.Context, id int) *domain.Buyer {
	query := "SELECT * FROM buyers WHERE id = ?;"
	row := r.db.QueryRow(query, id)
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

func (r *repository) Exists(ctx context.Context, cardNumberID string) bool {
	query := "SELECT card_number_id FROM buyers WHERE card_number_id=?;"
	row := r.db.QueryRow(query, cardNumberID)
	err := row.Scan(&cardNumberID)
	return err == nil
}

func (r *repository) Save(ctx context.Context, b domain.Buyer) int {
	query := "INSERT INTO buyers(card_number_id,first_name,last_name) VALUES (?,?,?)"
	stmt, err := r.db.Prepare(query)
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

func (r *repository) Update(ctx context.Context, b domain.Buyer) {
	query := "UPDATE buyers SET card_number_id=?, first_name=?, last_name=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName, &b.ID)
	if err != nil {
		panic(err)
	}
}

func (r *repository) Delete(ctx context.Context, id int) {
	query := "DELETE FROM buyers WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}
