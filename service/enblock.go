package service

import (
	"app/pkg/util"
	"fmt"
)

func CreateTx(userId string, json map[string]string) {
	fmt.Println(userId, json)
	a := util.Sha256(userId)

	fmt.Println(a)
}
