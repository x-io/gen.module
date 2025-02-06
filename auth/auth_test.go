package auth

import (
	"fmt"
	"testing"
)

func TestUrlMatch(t *testing.T) {
	config := getConfig("../../cmd/web/acl/factory.json")

	// assert.Equal(t, false, isAuth(config, "PUT", "/api/shop/depot-order/202206241636319592613250", "depot-order:5"))
	// assert.Equal(t, true, isAuth(config, "PUT", "/api/shop/depot-order/202206241636319592613250", "depot-order:2"))

	// assert.Equal(t, false, isAuth(config, "PUT", "/api/shop/depot-order/202206241636319592613250/submit", "depot-order:3"))
	// assert.Equal(t, true, isAuth(config, "PUT", "/api/shop/depot-order/202206241636319592613250/submit", "depot-order:4"))

	// assert.Equal(t, false, isAuth(config, "PUT", "/api/shop/depot-order/submit", "depot-order:3"))
	// assert.Equal(t, true, isAuth(config, "PUT", "/api/shop/depot-order/submit", "depot-order:4"))

	getRole("order:ff", "order")

	fmt.Println(isAuth(config, "GET", "/api/factory/order", "order:ff"))

	//
	// order depot:ff,make:0,depot-order:ff,product:ff,staff:ff,depot-deploy:ff,role:ff [/api/factory/order order ]

}
