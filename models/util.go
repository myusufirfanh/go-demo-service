package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserJWT struct {
	ID          int64        `json:"id"`
	FullName    string       `json:"fullName"`
	Lang        string       `json:"lang"`
	Email       string       `json:"email"`
	PhoneNumber string       `json:"phoneNumber"`
	Agent       UserJWTAgent `json:"agent"`
	Role        string       `json:"role"`
	jwt.StandardClaims
}

// UserJWTAgent will return object of user agent belongs to in jwt token
type UserJWTAgent struct {
	ID           int64       `json:"id"`
	UserID       int64       `json:"userId"`
	ReferralCode string      `json:"referralCode"`
	ParentID     int64       `json:"parentId"`
	RootID       int64       `json:"rootId"`
	RootLevel    int64       `json:"rootLevel"`
	Type         string      `json:"type"`
	Kind         string      `json:"kind"`
	Description  string      `json:"description"`
	Additional   interface{} `json:"additional"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

// QoalaResponsePattern return object Qoala general response pattern
type QoalaResponsePattern struct {
	Status  string                    `json:"status"`
	Data    interface{}               `json:"data"`
	Message string                    `json:"message"`
	Code    int                       `json:"code"`
	Meta    *QoalaResponsePatternMeta `json:"meta,omitempty"`
}

// QoalaResponsePatternMeta return object for Qoala reposne meta
type QoalaResponsePatternMeta struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Total int `json:"total"`
	Count int `json:"count"`
}
