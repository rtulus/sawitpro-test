package repository

import (
	"context"
	"time"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) AddNewUser(ctx context.Context, input AddNewUserInput) (err error) {

	sql := `INSERT INTO users (
		full_name, 
		phone_number, 
		passwd, 
		salt, 
		created_on
	) VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		$5
	)`

	_, err = r.Db.ExecContext(ctx, sql, input.FullName, input.PhoneNumber, input.HashedPassword, input.Salt, time.Now())
	return err
}

func (r *Repository) SelectUserByPhoneNumber(ctx context.Context, input SelectUserByPhoneNumberInput) (output Users, err error) {

	sql := `SELECT 
		user_id,
		full_name,
		passwd,
		salt
	FROM
		users
	WHERE
		phone_number = $1
	`

	var user Users
	err = r.Db.QueryRowContext(ctx, sql, input.PhoneNumber).Scan(
		&user.UserID,
		&user.FullName,
		&user.HashedPassword,
		&user.Salt,
	)
	if err != nil {
		return user, err
	}
	user.PhoneNumber = input.PhoneNumber

	return user, nil
}

func (r *Repository) IncrementUserSuccessfulLogin(ctx context.Context, input IncrementUserSuccessfulLoginInput) (err error) {

	sql := `UPDATE users 
	SET
		successful_login = successful_login + 1
	WHERE
		user_id = $1
	`

	_, err = r.Db.ExecContext(ctx, sql, input.UserID)
	return err
}

func (r *Repository) UpdateFullName(ctx context.Context, input UpdateFullNameInput) (err error) {

	sql := `UPDATE users 
	SET
		full_name = $1
	WHERE
		user_id = $2
	`

	_, err = r.Db.ExecContext(ctx, sql, input.FullName, input.UserID)
	return err
}

func (r *Repository) UpdatePhoneNumber(ctx context.Context, input UpdatePhoneNumberInput) (err error) {

	sql := `UPDATE users 
	SET
		phone_number = $1
	WHERE
		user_id = $2
	`

	_, err = r.Db.ExecContext(ctx, sql, input.PhoneNumber, input.UserID)
	return err
}
