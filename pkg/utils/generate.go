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
