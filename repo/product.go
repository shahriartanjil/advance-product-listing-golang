package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"ing_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	// productList []*Product  // in memory
	db *sqlx.DB
}

// constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
		INSERT INTO products (
		title,
		description,
		price,
		ing_url
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)
		RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
func (r *productRepo) Get(id int) (*Product, error) {
	var prd Product

	query := `
	SELECT
	 id,
	  title,
	   description,
	    price,
		 ing_url
		  from products
		  where id = $1
		  `
	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil

}
func (r *productRepo) List() ([]*Product, error) {
	var prdList []*Product

	query := `
	SELECT
	 id,
	  title,
	   description,
	    price,
		 ing_url
		  from products
		  `
	err := r.db.Select(&prdList, query)
	if err != nil {
		return nil, err
	}
	return prdList, nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	query := `
	UPDATE products
	SET title=$1, description=$2, price=$3, ing_url=$4 
	WHERE id = $5
	`
	result, err := r.db.Exec(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {

		return nil, sql.ErrNoRows
	}

	return &p, nil
}
func (r *productRepo) Delete(id int) error {
	query := `
		DELETE FROM products WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
