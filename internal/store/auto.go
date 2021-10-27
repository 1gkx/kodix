package store

import (
	"context"
	"fmt"
)

type Auto struct {
	Uuid    uint32 `json:"uuid" db:"id"`
	Brand   string `json:"brand,omitempty" db:"brand,omitempty"`
	Model   string `json:"model,omitempty"  db:"model,omitempty"`
	Price   uint32 `json:"price,omitempty"  db:"price,omitempty"`
	Status  string `json:"status,omitempty"  db:"status,omitempty"`
	Mileage uint32 `json:"mileage,omitempty"  db:"mileage,omitempty"`
}

func (db *Db) GetAuto(
	ctx context.Context,
) ([]*Auto, error) {

	rows, err := db.QueryContext(
		ctx,
		`SELECT
       		id, brand, model, price, status, mileage
    	FROM auto;`)
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

	result, err := db.ExecContext(ctx, `
		INSERT INTO auto(
			brand, model, price, status, mileage            
		) VALUES (?, ?, ?, ?, ?);
	`,
		req.Brand,
		req.Model,
		req.Price,
		req.Status,
		req.Mileage,
	)
	if err != nil {
		return nil, err
	}

	id, err := result. LastInsertId()
	if err != nil {
		return nil, err
	}

	var auto Auto
	if err := db.QueryRowContext(
		ctx,
		`SELECT
       		id, brand, model, price, status, mileage
    	FROM auto WHERE id = ?;`,
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

func (db *Db) UpdateAuto(
	ctx context.Context,
	req *Auto,
) (*Auto, error) {

	result, err := db.ExecContext(ctx, `
		UPDATE auto
			set brand=?, model=?, price=?, status=?, mileage=?
		WHERE id=?;
		`,
		req.Brand,
		req.Model,
		req.Price,
		req.Status,
		req.Mileage,
		req.Uuid,
	)
	if err != nil {
		return nil, err
	}
	if c, err := result.RowsAffected(); err != nil || c == 0 {
		return nil, fmt.Errorf("failed update auto with id: %d", req.Uuid)
	}

	var auto Auto
	if err := db.QueryRowContext(
		ctx,
		`SELECT * FROM auto WHERE id = $1`,
		req.Uuid,
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
	uuid uint32,
) error {

	result, err := db.ExecContext(ctx, `
		DELETE FROM auto
		WHERE id = ?;
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
