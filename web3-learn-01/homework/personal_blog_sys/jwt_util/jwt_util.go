package jwt_util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 固定的密钥，通过文件持久化存储
var secretKey *ecdsa.PrivateKey

func init() {
	// 密钥文件路径
	var err error
	secretKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Println("生成密钥失败 ", err)
		panic(err)
	}
}

// GenerateJWT 生成JWT
func GenerateJWT(userID uint, username string, email string) (string, error) {
	expireTime := time.Now().Add(3 * time.Hour)

	// 创建 Claims
	claims := CustomerClaims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 过期时间
			Issuer:    "web3-learn",                   // 签发人
		},
	}

	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("生成token异常 ", err)
		return "", err
	}

	return signedToken, nil
}

type CustomerClaims struct {
	UserID               uint   `json:"user_id"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	jwt.RegisteredClaims        // 标准 JWT 字段
}

// ValidateJWT 验证JWT并返回其中的声明
func ValidateJWT(tokenString string) (*CustomerClaims, error) {
	// 解析JWT
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &secretKey.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomerClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GetUserIDFromToken 从JWT中提取用户ID（便捷方法）
func GetUserIDFromToken(tokenString string) (uint, error) {
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
