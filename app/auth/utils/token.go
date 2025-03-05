package utils

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
)

var JWTKey = []byte("your_secret_key")

type Claims struct {
    UserID uint
    Email  string
    jwt.RegisteredClaims
}

func GenerateToken(userID uint) (string, error) {
    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JWTKey)
}