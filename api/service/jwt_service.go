package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(accountCode string, propertyId string, subjectId string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	signingKey string
	issuer     string
}

type myClaims struct {
	AccountCode string `json:"account_code"`
	PropertyId  string `json:"property_id"`
	jwt.StandardClaims
}

func NewJwtService(signingKey string) JwtService {
	return &jwtService{
		signingKey: signingKey,
		issuer:     "ShiftingApi",
	}
}

func (s *jwtService) GenerateToken(accountCode string, propertyId string, subjectId string) (string, error) {
	claims := &myClaims{
		AccountCode: accountCode,
		PropertyId:  propertyId,
		StandardClaims: jwt.StandardClaims{
			Subject:  subjectId,
			Issuer:   s.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.signingKey))
	if err != nil {
		return "", fmt.Errorf("GenerateToken: failed signing token: %s", err.Error())
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("ValidateToken: invalid token: %v", t.Header["alg"])
		}

		return []byte(s.signingKey), nil
	})
}
