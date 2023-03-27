package middleware

import (
	"section22_buku/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(bukuId int, judul string) (string, error) {
	claims := jwt.MapClaims{}
	claims["bukuId"] = bukuId
	claims["judul"] = judul
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExportTokenBukuId(e echo.Context) int {
	buku := e.Get("buku").(*jwt.Token)
	if buku.Valid {
		claims := buku.Claims.(jwt.MapClaims)
		bukuId := claims["bukuId"].(float64)
		return int(bukuId)
	}

	return 0

}
