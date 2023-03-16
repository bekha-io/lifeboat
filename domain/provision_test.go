package domain

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	m.Run()
}

func TestGetProvisionAll_Mixing(t *testing.T) {
	d1 := GetProvisionAll()
	d2 := GetProvisionAll()

	fmt.Println(d1[0])
	fmt.Println(d2[0])

	if reflect.DeepEqual(d1, d2) {
		t.FailNow()
	}
}
