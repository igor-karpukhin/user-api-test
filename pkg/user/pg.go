package user

import (
	"database/sql"

	"go.uber.org/zap"
)

type PostgresUserDao struct {
	conn *sql.DB
	log  *zap.Logger
}

func NewPostgresUserDao(conn *sql.DB, log *zap.Logger) *PostgresUserDao {
	return &PostgresUserDao{
		conn: conn,
		log:  log,
	}
}

func (p *PostgresUserDao) ListUsers() ([]*User, error) {
	rows, err := p.conn.Query("SELECT id, name, birthday, age, hobbies FROM task.users")
	if err != nil {
		p.log.Error("unable to get a list of users", zap.Error(err))
		return nil, err
	}

	result := []*User{}
	for rows.Next() {
		u := &User{}

		err = rows.Scan(&u.ID, &u.Name, &u.Birthday, &u.Age, &u.Hobbies)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func (p *PostgresUserDao) GetUser(ID uint64) (*User, error) {
	rows, err := p.conn.Query("SELECT name, birthday, age, hobbies FROM task.users WHERE id = $1", ID)
	if err != nil {
		p.log.Error("unable to get user", zap.Uint64("id", ID), zap.Error(err))
		return nil, err
	}
	rows.Next()

	u := &User{
		ID: ID,
	}

	err = rows.Scan(&u.Name, &u.Birthday, &u.Age, &u.Hobbies)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (p *PostgresUserDao) DeleteUser(ID uint64) error {
	_, err := p.conn.Exec("DELETE FROM task.users WHERE id = $1", ID)
	return err
}

func (p *PostgresUserDao) CreateUser(user *User) (uint64, error) {
	res, err := p.conn.Exec("INSERT INTO task.users (name, birthday, age, hobbies) VALUES ($1, $2, $3, $4)",
		user.Name, user.Birthday, user.Age, user.Hobbies)
	if err != nil {
		p.log.Error("unable to insert user", zap.Uint64("id", user.ID), zap.Error(err))
		return 0, err
	}
	lId, err := res.LastInsertId()
	return uint64(lId), err
}

func (p *PostgresUserDao) UpdateUser(user *User) error {
	_, err := p.conn.Exec("UPDATE task.users SET name = $1, birthday = $2, age = $3, hobbies = $4 WHERE id = $5",
		user.Name, user.Birthday, user.Age, user.Hobbies, user.ID)
	if err != nil {
		p.log.Error("unable to update user", zap.Uint64("id", user.ID), zap.Error(err))
	}
	return nil
}
