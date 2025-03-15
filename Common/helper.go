package common

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func HttpRequest(url, method string, data, header []byte) ([]byte, error) {
	var err error

	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		log.Fatalln("error:", err)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if len(header) > 0 {
		headerMap := new(map[string]interface{})
		err = json.Unmarshal(header, headerMap)
		if err != nil {
			log.Fatalln("error:", err)
		}
		for k, v := range *headerMap {
			if k == "" || v == "" {
				continue
			}
			request.Header.Set(k, v.(string))
		}
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln("error:", err)
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error:", err)
	}
	return respBytes, nil
}
func HttpPost(url string, data []byte, header ...byte) ([]byte, error) {
	return HttpRequest(url, "POST", data, header)

}

func HttpGet(url string, header ...byte) ([]byte, error) {
	return HttpRequest(url, "GET", []byte{}, header)

}

func IF(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))                // 将字符串转换为[]byte进行哈希计算
	return fmt.Sprintf("%x", hasher.Sum(nil)) // 返回十六进制字符串
}

func RFC3339ToNormalTime(rfc3339 string) string {
	if len(rfc3339) < 19 || rfc3339 == "" || !strings.Contains(rfc3339, "T") {
		return rfc3339
	}
	return strings.Split(rfc3339, "T")[0] + " " + strings.Split(rfc3339, "T")[1][:8]
}

type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`  //	ID
	UserName string `json:"username"` // 用户名
	Role     int    `json:"role"`     // 权限

}

type CustomerClaims struct {
	User JwtPayLoad `json:"user"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func GenerateToken(secretKey string, iat, seconds, userid int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["UserId"] = userid
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func GenerateToken1(user JwtPayLoad, secretKey string, expires int64) (string, error) {
	claims := CustomerClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateToken 验证 JWT Token
func ValidateToken(signedToken string, accessSecret string, expires int64) (*CustomerClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomerClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
