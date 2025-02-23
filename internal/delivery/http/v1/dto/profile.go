package dto

import "time"

type Profile struct {
	Bio       string     `json:"bio"`
	AvatarURL string     `json:"avatar_url"`
	Website   string     `json:"website"`
	Birthday  *time.Time `json:"birthday"`
}
