package pg_persist

import (
	"github.com/pkg/errors"
)

var (
	ErrNoRecords      = errors.New("no_records")
	ErrNoRowsAffected = errors.New("no_rows_affected")
)

const (
	insProd = `INSERT INTO public.products(product_cost, currency_code, updated_on) VALUES ($1, $2, current_timestamp) RETURNING product_id`
	selProd = `SELECT  product_cost, currency_code FROM public.products where product_id=$1`
	updProd = `UPDATE public.products SET product_cost=$1, currency_code=$2,updated_on=current_timestamp WHERE product_id=$3`
	delProd = `DELETE FROM public.products WHERE product_id=$1`
)

func AddProduct(prise float64, code string) (id int64, err error) {
	prepare, err := Db.Prepare(insProd)
	rows, err := prepare.Query(prise, code)
	defer rows.Close()
	if err != nil {
		return 0, errors.Wrap(err, "insert product failed")
	}
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, errors.Wrap(err, "insert product failed")
		}
	}

	return
}

func GetProduct(id int) (value float64, code string, err error) {
	prepare, err := Db.Prepare(selProd)
	rows, err := prepare.Query(id)
	defer rows.Close()
	if err != nil {
		err = errors.Wrap(err, "failed to get product info")
		return
	}
	if rows.Next() {
		err = rows.Scan(&value, &code)
		if err != nil {
			err = errors.Wrap(err, "failed to get product info")
			return
		}
	} else {
		err = ErrNoRecords
	}

	return
}

func UpdateProduct(id int, prise float64, code string) error {
	res, err := Db.Exec(updProd, prise, code, id)
	if err != nil {
		return errors.Wrap(err, "failed to update product")
	}
	if v, _ := res.RowsAffected(); v <= 0 {
		return ErrNoRowsAffected
	}

	return nil
}

func DeleteProduct(id int) error {
	res, err := Db.Exec(delProd, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete product")
	}
	if v, _ := res.RowsAffected(); v <= 0 {
		return ErrNoRowsAffected
	}

	return nil
}
