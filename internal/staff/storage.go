package staff

import "context"

type IStorage interface {
	UpdateUserAuthMethod(ctx context.Context, auth *UserAuth, method *UserAuthMethod) (int, error)
}
