package userRepo

import (
	"context"
	"database/sql"
	"fmt"

	userDomain "test-majoo/src/domain/user"
)

func (m *userRepo) GetListUser(ctx context.Context, offset, limit int, search, sort string) (res []userDomain.User, count int, err error) {
	query := `SELECT u.id, u.username, u.email, u.name, r.name	
	FROM users u
	JOIN roles r ON u.role_id = r.id
	WHERE u.is_active = 1`

	if search != "" {
		query += ` AND (username LIKE ?)`
	}

	switch sort {
	case "name asc":
		query += ` ORDER BY name ASC `
	case "name desc":
		query += ` ORDER BY name DESC `
	default:
		query += ` ORDER BY id desc `
	}

	query += `LIMIT ? OFFSET ?`

	var rows *sql.Rows = nil
	if search != "" {
		rows, err = m.Conn.QueryContext(ctx, query, "%"+search+"%", limit, offset)
	} else {
		rows, err = m.Conn.QueryContext(ctx, query, limit, offset)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		d := userDomain.User{}
		err = rows.Scan(&d.ID, &d.Username, &d.Email, &d.Name, &d.Role)
		if err != nil {
			fmt.Println(err)
			return
		}
		res = append(res, d)
	}

	// Query row count
	query = `SELECT COUNT(*)
	FROM users 	
	WHERE is_active = 1`
	if search != "" {
		query += ` AND (username LIKE ?)`
	}

	if search != "" {
		err = m.Conn.QueryRow(query, "%"+search+"%").Scan(&count)
	} else {
		err = m.Conn.QueryRow(query).Scan(&count)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *userRepo) Login(ctx context.Context, b userDomain.Login) (data userDomain.UserLogin, err error) {
	query := `SELECT u.id, u.username, u.password, u.role_id 	
	FROM users u 	
	where u.username = ?`
	err = m.Conn.QueryRow(query, b.Username).Scan(&data.ID, &data.Username, &data.Password,
		&data.RoleId)

	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return
}

func (m *userRepo) GetUserById(ctx context.Context, id int) (res userDomain.User, err error) {
	query := `SELECT id, username, email, name, role_id FROM users WHERE id = ?`
	err = m.Conn.QueryRow(query, id).Scan(&res.ID, &res.Username, &res.Email, &res.Name, &res.Role)
	if err != nil {
		return
	}
	return
}

func (m *userRepo) Create(ctx context.Context, b userDomain.SetUser) (err error) {
	query := `INSERT INTO users(username,email,name,password,role_id) 
	VALUES(?,?,?,?,?)`
	stmt, err := m.db.PrepareContext(ctx, query)
	fmt.Println(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.ExecContext(ctx,
		b.Username,
		b.Name,
		b.Email,
		b.Password,
		b.Role,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *userRepo) Update(ctx context.Context, b userDomain.SetUser, id int) (err error) {
	query := `UPDATE users 	
	SET username = ?, email = ?, role_id = ?
	WHERE id = ?`

	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, b.Username, b.Email, b.Role, id)
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

func (m *userRepo) Delete(ctx context.Context, id int) (err error) {
	query := `UPDATE users 	
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
