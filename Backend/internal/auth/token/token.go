package token
import (
    "github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
	"github.com/google/uuid"
	"backend/internal/env"
)
type Claims struct {
    UserID int64
	SessionID string
	jwt.RegisteredClaims
}
type Setupclaims struct {
	UserID int64
	SetupToken bool
	jwt.RegisteredClaims
}
func GenerateSessionID()string{
    return uuid.NewString() 
}
func GetVerificationToken(userID int64,session string)(string, error){
	var jwtKey = []byte(env.Env.JwtKey)
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{
		UserID: userID,
		SessionID: session,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func GetSetupToken(userID int64)(string, error){
	var jwtKey = []byte(env.Env.JwtKey)
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Setupclaims{
		UserID: userID,
		SetupToken: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func GetToken(userID int64, session string)(string, error){
	var jwtKey = []byte(env.Env.JwtKey)
	expirationTime := time.Now().AddDate(0, 6, 0)
	claims := &Claims{
		UserID : userID,
		SessionID: session,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func VerifyToken(requestToken string, secret string)(bool,error){
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
func ExtractFromToken(tokenString string, secret string) (int64,string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0,"",err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID,claims.SessionID ,nil
	} else {
		return 0,"", fmt.Errorf("invalid token")
	}
}
func ExtractFromSetupToken(tokenString string, secret string) (int64,bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Setupclaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0,false,err
	}

	if claims, ok := token.Claims.(*Setupclaims); ok && token.Valid {
		return claims.UserID,claims.SetupToken,nil
	} else {
		return 0, false,fmt.Errorf("invalid token")
	}
}
