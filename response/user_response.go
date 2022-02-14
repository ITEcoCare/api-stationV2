package response

import "time"

type ResponseUserAuth struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Avatar          string     `json:"avatar"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Token           string     `json:"token"`
}
