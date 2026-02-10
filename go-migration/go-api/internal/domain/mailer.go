package domain

import "context"

type Mailer interface {
	SendEmail(ctx context.Context, user string, sub string, body string) error
}
