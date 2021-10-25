package store

import (
	"context"
	"fmt"
)

type Auto struct {
	Uuid    uint32 `json:"uuid" db:"id"`
	Brand   string `json:"brand,omitempty" db:"brand,omitempty"`
	Model   string `json:"model,omitempty"  db:"model,omitempty"`
	Price   uint64 `json:"price,omitempty"  db:"price,omitempty"`
	Status  string `json:"status,omitempty"  db:"status,omitempty"`
	Mileage uint64 `json:"mileage,omitempty"  db:"mileage,omitempty"`
}

func (db *Db) GetAuto(
	ctx context.Context,
) ([]*Auto, error) {
	rows, err := db.QueryContext(ctx, `SELECT * FROM auto;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	autoArray := make([]*Auto, 0)
	for rows.Next() {
		var a Auto
		if err := rows.Scan(
			&a.Uuid,
			&a.Brand,
			&a.Model,
			&a.Price,
			&a.Status,
			&a.Mileage,
		); err != nil {
			return nil, err
		}
		autoArray = append(autoArray, &a)
	}

	return autoArray, nil
}

func (db *Db) AddAuto(
	ctx context.Context,
	req *Auto,
) (*Auto, error) {

	var auto Auto
	row := db.QueryRowContext(ctx, `
		INSERT INTO auto(
			brand, model, price, status, mileage            
		) VALUES ($1, $2, $3, $4, $5);
	`,
		req.Brand,
		req.Model,
		req.Price,
		req.Status,
		req.Mileage,
	)
	if err := row.Scan(
		&auto.Uuid,
		&auto.Brand,
		&auto.Model,
		&auto.Price,
		&auto.Status,
		&auto.Mileage,
	); err != nil {
		return nil, err
	}

	return &auto, nil
}

func (db *Db) UpdateAuto(
	ctx context.Context,
	req *Auto,
) (*Auto, error) {

	result, err := db.ExecContext(ctx, `
		UPDATE auto
		SET 
		    brand=$2,
		    model=$3,
		    price=$4,
		    status=$5,
		    mileage=$6
		WHERE id=$1;
		`,
		req.Uuid,
		req.Brand,
		req.Model,
		req.Price,
		req.Status,
		req.Mileage,
	)
	if err != nil {
		return nil, err
	}

	var auto Auto
	fmt.Println(result.RowsAffected())
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	fmt.Println(id)
	if err := db.QueryRowContext(
		ctx,
		`SELECT * FROM auto WHERE id = $1`,
		id,
	).Scan(
		&auto.Uuid,
		&auto.Brand,
		&auto.Model,
		&auto.Price,
		&auto.Status,
		&auto.Mileage,
	); err != nil {
		return nil, err
	}

	return &auto, nil
}

func (db *Db) DeleteAuto(
	ctx context.Context,
	uuid uint,
) error {

	result, err := db.ExecContext(ctx, `
		DELETE FROM auto
		WHERE id = $1;
		`,
		uuid,
	)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("No such auto with uuid: %d", uuid)
	}
	return nil
}
