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

package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

// var needs to be used instead of const as ldflags is used to fill this
// information in the release process
var (
	ClientVersion = "unknown"
	Goos          = "unknown"
	Goarch        = "unknown"
	GitCommit     = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	BuildDate     = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// version contains all the information related to the CLI version
type version struct {
	ClientVersion string `json:"clientVersion"`
	GitCommit     string `json:"gitCommit"`
	BuildDate     string `json:"buildDate"`
	GoOs          string `json:"goOs"`
	GoArch        string `json:"goArch"`
}

// versionString returns the CLI version
func versionString() string {
	return fmt.Sprintf("Version: %#v", version{
		ClientVersion,
		GitCommit,
		BuildDate,
		Goos,
		Goarch,
	})
}

func NewVersionCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "version",
		Short: "Print the cco version",
		Run:   versionCommandFunc,
	}
	return cc
}

func versionCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Println(versionString())
}
