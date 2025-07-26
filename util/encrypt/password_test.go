package encrypt

import (
	"testing"
)

func TestPassword(t *testing.T) {
	password := "admin123"
	hash, err := HashPassword(password, DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hash)
}

func TestVerifyPassword(t *testing.T) {
	password := "123456"
	hash := "$2a$12$cqpWQxWnXRFsKfou4MzxQuerFzrPjVuTy/vjajnflfbJJBSMRj2ge"
	err := VerifyPassword(password, hash)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("密码正确")
	}
}
