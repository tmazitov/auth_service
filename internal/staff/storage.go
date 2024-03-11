package staff

import "context"

type IStorage interface {
	AddUserAuthMethod(ctx context.Context, auth *UserAuth, method *UserAuthMethod) error
}
