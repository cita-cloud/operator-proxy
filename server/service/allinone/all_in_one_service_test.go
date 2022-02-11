/*
 * Copyright Rivtower Technologies LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package allinone

import (
	"fmt"
	"github.com/cita-cloud/operator-proxy/pkg/utils"
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
	s := uuid.New().String()
	t.Log(s)
	t.Log(s[len(s)-12:])
	adminPwd, _ := password.Generate(16, 4, 4, false, false)
	t.Log(adminPwd)

	data := "test"
	h := sm3.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	fmt.Printf("digest value is: %x\n", sum)
}

func Test_generateChainId(t *testing.T) {
	a := utils.GenerateChainId("hello")
	b := fmt.Sprintf("%x", a)
	fmt.Printf("哈希结果为：%x", a)
	fmt.Printf("哈希结果为：%s", b)
	t.Log(a)
	t.Log(string(a))
}
