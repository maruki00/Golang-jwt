package internal

import "context"

type Login struct {
}

func (l *Login) LoginAction(ctx context.Context, login, password string) {

	db, err := GetDB().ExecContext(ctx, "", 12344)
}
