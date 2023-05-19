package request

import (
	// jwt "github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt"
	// uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID          uint
	Address     string
	AuthorityId uint
}
