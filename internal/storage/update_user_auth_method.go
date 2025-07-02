package storage

import (
	"context"

	"github.com/tmazitov/auth_service.git/internal/staff"
)

func (s *Storage) UpdateUserAuthMethod(ctx context.Context, auth *staff.UserAuth, method *staff.UserAuthMethod) (int, error) {

	var (
		err error
	)

	// Start a new transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	// Rollback if function returns with error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert into user_auths table
	_, err = tx.NewInsert().
		Model(auth).
		On("CONFLICT (email) DO UPDATE SET last_auth_at=current_timestamp").
		Returning("id").
		Exec(ctx)
	if err != nil {
		return 0, err
	}

	count, err := tx.NewSelect().
		Model((*staff.UserAuthMethod)(nil)).
		Where("auth_method_id = ?", method.AuthMethodId).
		Where("user_id = ?", auth.Id).
		Count(ctx)

	if count == 0 {
		method.UserId = auth.Id
		_, err = tx.NewInsert().
			Model(method).
			Exec(ctx)
		if err != nil {
			return 0, err
		}
	} else {
		// Update last_auth_at in user_auth_methods table
		_, err = tx.NewUpdate().
			Model(method).
			Set("last_auth_at = CURRENT_TIMESTAMP").
			Where("user_id = ?", auth.Id).
			Where("auth_method_id = ?", method.AuthMethodId).
			Exec(ctx)
	}
	if err != nil {
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return auth.Id, nil
}
