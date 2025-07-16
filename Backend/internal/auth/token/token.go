package token
import (
    "github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
)
type Claims struct {
	User_id int64
	jwt.RegisteredClaims
}
func GetVerificationToken(user_id int64)(string, error){
	var jwtKey = []byte("your_jwt_secret")
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{
		User_id: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func GetToken(userID int64)(string, error){
	var jwtKey = []byte("your_jwt_secret")
	expirationTime := time.Now().AddDate(0, 6, 0)
	claims := &Claims{
		User_id : userID,
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
func ExtractIDFromToken(tokenString string, secret string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.User_id, nil
	} else {
		return 0, fmt.Errorf("invalid token")
	}
}
