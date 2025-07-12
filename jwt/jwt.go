package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/x-io/gen"
	"github.com/x-io/gen.module/dbs"
	"github.com/x-io/gen/errors"
)

// MapClaims MapClaims
type MapClaims = jwt.MapClaims

// Get ...
func Get(c *gen.Context, name string) string {
	data := c.Data("JWT")

	if v, ok := data.(jwt.MapClaims); ok {
		if vv, ok := v[name]; ok {
			return vv.(string)
		}
	}
	return ""
}

// GetOperator ...
func GetOperator(c *gen.Context) *dbs.Operator {
	data := c.Data("JWT")

	if v, ok := data.(jwt.MapClaims); ok {
		var d dbs.Operator

		if vv, ok := v["id"]; ok {
			d.ID = dbs.String(vv.(string))
		}

		if vv, ok := v["name"]; ok {
			d.Name = dbs.String(vv.(string))
		}

		if vv, ok := v["depot"]; ok {
			d.Depot = dbs.String(vv.(string))
		}
		return &d
	}
	return nil
}

// New New
func New(claim MapClaims, tokenExp time.Duration, key interface{}) (string, error) {
	claim["exp"] = time.Now().Add(tokenExp).Unix()

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(key)
}

// Refresh Refresh 只有在token过期前10分钟，过期后30天内才可以刷新
func Refresh(token string, key interface{}, refresh func(claim MapClaims) (string, error)) (string, error) {
	data, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Always check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key for validation
		return key, nil
	})

	if data.Valid {
		claim := data.Claims.(jwt.MapClaims)
		t := time.Unix(int64(claim["exp"].(float64)), 0)
		if time.Since(t).Minutes() > -10 {
			return refresh(claim)
		}
		//	fmt.Println(time.Since(t).Minutes())
		return "", nil
	}

	ve := err.(*jwt.ValidationError)
	if ve.Errors == jwt.ValidationErrorExpired {
		claim := data.Claims.(jwt.MapClaims)

		t := time.Unix(int64(claim["exp"].(float64)), 0)
		if time.Since(t).Hours()/24 < 30 {
			return refresh(claim)
		}
	}

	return "", errors.HTTP(http.StatusUnauthorized)
}
