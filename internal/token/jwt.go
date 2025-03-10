package token

import (
	"github.com/go-chi/jwtauth/v5"
	"time"
)

type JWTTokenService struct {
	secretKey string
	jwtAuth   *jwtauth.JWTAuth
}

func NewJWTTokenService(secretKey string) *JWTTokenService {
	jwtAuth := jwtauth.New("HS256", []byte(secretKey), nil)
	return &JWTTokenService{secretKey: secretKey, jwtAuth: jwtAuth}
}

func (s *JWTTokenService) GenerateToken(email string, userID uint) (string, error) {
	_, tokenString, err := s.jwtAuth.Encode(map[string]interface{}{"user_id": userID, "email": email, "exp": time.Now().Add(time.Hour).Unix()})

	return tokenString, err
}

func (s *JWTTokenService) GetJWTAuth() *jwtauth.JWTAuth {
	return s.jwtAuth
}
