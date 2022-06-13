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
	Authorize       Authorize  `json:"authorize"`
}

type Authorize struct {
	CompanyId   int         `json:"company"`
	RoleId      int         `json:"role_id"`
	Role        string      `json:"role"`
	Permissions interface{} `json:"permission"`
}
