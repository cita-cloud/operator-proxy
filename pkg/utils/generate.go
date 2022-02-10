package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"github.com/tjfoc/gmsm/sm3"
)

func GenerateChainId(name string) string {
	h := sm3.New()
	h.Write([]byte(name))
	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func GenerateAccountOrNodeName(name string) string {
	s := uuid.New().String()
	return name + "-" + s[len(s)-12:]
}

func GenerateAccountPassword() (string, error) {
	generate, err := password.Generate(16, 4, 4, false, false)
	return generate, err
}
