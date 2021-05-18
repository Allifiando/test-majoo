package productRepo

import (
	"context"
	"database/sql"
	"fmt"

	productDomain "test-majoo/src/domain/product"
)

func (m *productRepo) GetListMerchantByUserId(ctx context.Context, offset, limit int, search, sort string, userId int64) (res []productDomain.Merchant, count int, err error) {
	query := `SELECT m.id, u.name, m.company_name, m.trademark, m.is_active	
	FROM merchants m
	JOIN users u ON m.user_id = u.id
	WHERE user_id = ?`

	if search != "" {
		query += ` AND (company_name LIKE ? OR trademark LIKE ?)`
	}

	switch sort {
	case "name asc":
		query += ` ORDER BY company_name ASC `
	case "name desc":
		query += ` ORDER BY company_name DESC `
	default:
		query += ` ORDER BY m.id desc `
	}

	query += `LIMIT ? OFFSET ?`

	var rows *sql.Rows = nil
	if search != "" {
		rows, err = m.Conn.QueryContext(ctx, query, userId, "%"+search+"%", "%"+search+"%", limit, offset)
	} else {
		rows, err = m.Conn.QueryContext(ctx, query, userId, limit, offset)
	}

	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		d := productDomain.Merchant{}
		err = rows.Scan(&d.ID, &d.User, &d.CompanyName, &d.Trademark, &d.IsActive)
		if err != nil {
			return
		}
		res = append(res, d)
	}

	// Query row count
	query = `SELECT COUNT(*)
	FROM merchants 	
	WHERE user_id = ?`
	if search != "" {
		query += ` AND (company_name LIKE ? OR trademark LIKE ?)`
	}

	if search != "" {
		err = m.Conn.QueryRow(query, userId, "%"+search+"%", "%"+search+"%").Scan(&count)
	} else {
		err = m.Conn.QueryRow(query, userId).Scan(&count)
	}
	if err != nil {
		return
	}
	return
}

func (m *productRepo) GetMerchantById(ctx context.Context, id int64) (data productDomain.Merchant, err error) {
	query := `SELECT m.id, u.name, m.company_name, m.trademark, m.is_active	
	FROM merchants m
	JOIN users u ON m.user_id = u.id
	WHERE m.id = ?`
	err = m.Conn.QueryRow(query, id).Scan(&data.ID, &data.User, &data.CompanyName,
		&data.Trademark, &data.IsActive)

	if err != nil {
		return data, err
	}
	return
}

func (m *productRepo) GetMerchantByName(ctx context.Context, name string) (data productDomain.Merchant, err error) {
	query := `SELECT m.id, u.name, m.company_name, m.trademark, m.is_active	
	FROM merchants m
	JOIN users u ON m.user_id = u.id
	WHERE m.company_name = ?`
	err = m.Conn.QueryRow(query, name).Scan(&data.ID, &data.User, &data.CompanyName,
		&data.Trademark, &data.IsActive)

	if err != nil {
		return data, err
	}
	return
}

func (m *productRepo) GetListOutletByMerchantId(ctx context.Context, offset, limit int, search, sort string, merchantId int64) (res []productDomain.Outlet, count int, err error) {
	query := `SELECT o.id, m.id, o.name, o.location, o.is_active	
	FROM outlets o
	JOIN merchants m ON o.merchant_id = m.id
	WHERE o.merchant_id = ?`

	if search != "" {
		query += ` AND (name LIKE ?)`
	}

	switch sort {
	case "name asc":
		query += ` ORDER BY name ASC `
	case "name desc":
		query += ` ORDER BY name DESC `
	default:
		query += ` ORDER BY o.id desc `
	}

	query += `LIMIT ? OFFSET ?`

	var rows *sql.Rows = nil
	if search != "" {
		rows, err = m.Conn.QueryContext(ctx, query, merchantId, "%"+search+"%", limit, offset)
	} else {
		rows, err = m.Conn.QueryContext(ctx, query, merchantId, limit, offset)
	}

	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		d := productDomain.Outlet{}
		err = rows.Scan(&d.ID, &d.MerchantId, &d.Name, &d.Location, &d.IsActive)
		if err != nil {
			return
		}
		res = append(res, d)
	}

	// Query row count
	query = `SELECT COUNT(*)
	FROM outlets 	
	WHERE merchant_id = ?`
	if search != "" {
		query += ` AND (name LIKE ?)`
	}

	if search != "" {
		err = m.Conn.QueryRow(query, merchantId, "%"+search+"%").Scan(&count)
	} else {
		err = m.Conn.QueryRow(query, merchantId).Scan(&count)
	}
	if err != nil {
		return
	}
	return
}

func (m *productRepo) GetOutletById(ctx context.Context, id int64) (data productDomain.Outlet, err error) {
	query := `SELECT o.id, m.id, o.name, o.location, o.is_active	
	FROM outlets o
	JOIN merchants m ON o.merchant_id = m.id	
	WHERE o.id = ?`
	err = m.Conn.QueryRow(query, id).Scan(&data.ID, &data.MerchantId, &data.Name,
		&data.Location, &data.IsActive)

	if err != nil {
		return data, err
	}
	return
}

func (m *productRepo) GetListProductByOutletId(ctx context.Context, offset, limit int, search, sort string, outletId int64) (res []productDomain.Product, count int, err error) {
	query := `SELECT p.id, o.id, p.name, p.price, p.qty, p.filename, p.is_active	
	FROM products p
	JOIN outlets o ON p.outlet_id = o.id
	WHERE p.outlet_id = ?`

	if search != "" {
		query += ` AND (name LIKE ?)`
	}

	switch sort {
	case "name asc":
		query += ` ORDER BY name ASC `
	case "name desc":
		query += ` ORDER BY name DESC `
	default:
		query += ` ORDER BY p.id desc `
	}

	query += `LIMIT ? OFFSET ?`

	var rows *sql.Rows = nil
	if search != "" {
		rows, err = m.Conn.QueryContext(ctx, query, outletId, "%"+search+"%", limit, offset)
	} else {
		rows, err = m.Conn.QueryContext(ctx, query, outletId, limit, offset)
	}

	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		d := productDomain.Product{}
		err = rows.Scan(&d.ID, &d.OutletId, &d.Name, &d.Price, &d.Qty, &d.Filename, &d.IsActive)
		if err != nil {
			return
		}
		res = append(res, d)
	}

	// Query row count
	query = `SELECT COUNT(*)
	FROM products 	
	WHERE outlet_id = ?`
	if search != "" {
		query += ` AND (name LIKE ?)`
	}

	if search != "" {
		err = m.Conn.QueryRow(query, outletId, "%"+search+"%").Scan(&count)
	} else {
		err = m.Conn.QueryRow(query, outletId).Scan(&count)
	}
	if err != nil {
		return
	}
	return
}

func (m *productRepo) GetProductById(ctx context.Context, id int64) (data productDomain.Product, err error) {
	query := `SELECT p.id, o.id, p.name, p.price, p.qty, p.filename, p.is_active	
	FROM products p
	JOIN outlets o ON p.outlet_id = o.id
	WHERE p.id = ?`
	err = m.Conn.QueryRow(query, id).Scan(&data.ID, &data.OutletId, &data.Name,
		&data.Price, &data.Qty, &data.Filename, &data.IsActive)

	if err != nil {
		return data, err
	}
	return
}

func (m *productRepo) CreateMerchant(ctx context.Context, b productDomain.SetMerchant) (err error) {
	query := `INSERT INTO merchants(user_id,company_name,trademark,is_active) 
	VALUES(?,?,?,?)`
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx,
		b.UserId,
		b.CompanyName,
		b.Trademark,
		b.IsActive,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *productRepo) CreateOutlet(ctx context.Context, b productDomain.SetOutlet) (err error) {
	query := `INSERT INTO outlets(merchant_id,name,location,is_active) 
	VALUES(?,?,?,?)`
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx,
		b.MerchantId,
		b.Name,
		b.Location,
		b.IsActive,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *productRepo) CreateProduct(ctx context.Context, b productDomain.SetProduct) (err error) {
	query := `INSERT INTO products(outlet_id,name,price,qty,filename,is_active) 
	VALUES(?,?,?,?,?,?)`
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx,
		b.OutletId,
		b.Name,
		b.Price,
		b.Qty,
		b.Filename,
		b.IsActive,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *productRepo) UpdateProduct(ctx context.Context, b productDomain.SetProduct, id int64) (err error) {
	query := `UPDATE products 	
	SET outlet_id = ?, name = ?, price = ?, qty = ?, filename = ?, is_active = ?
	WHERE id = ?`

	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, b.OutletId, b.Name, b.Price, b.Qty, b.Filename, b.IsActive, id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		// no affected data
		return
	}
	return
}

func (m *productRepo) DeleteProduct(ctx context.Context, id int64) (err error) {
	query := `UPDATE products 	
	SET is_active = 0 
	WHERE id = ?`

	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		// no affected data
		return
	}
	return
}
