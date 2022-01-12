package allinone

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"github.com/tjfoc/gmsm/sm3"
	"k8s.io/apimachinery/pkg/util/wait"
)

func TestAPIs(t *testing.T) {
	err := wait.Poll(2*time.Second, 10*time.Second, func() (done bool, err error) {

		return true, nil
	})
	t.Fatal(err)
}

func TestUuid(t *testing.T) {
	s:=uuid.New().String()
	t.Log(s)
	t.Log(s[len(s)-12:])
	adminPwd, _ := password.Generate(16, 4, 4, false, false)
	t.Log(adminPwd)

	data := "test"
	h := sm3.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	fmt.Printf("digest value is: %x\n",sum)
}
