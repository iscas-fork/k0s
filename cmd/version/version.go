/*
Copyright 2021 k0s authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/iscas-fork/k0s/pkg/build"

	"github.com/spf13/cobra"
)

var (
	all   bool
	isJsn bool
)

func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the k0s version",

		Run: func(cmd *cobra.Command, args []string) {
			info := versionInfo{
				Version:    build.Version,
				Kubernetes: build.KubernetesVersion,
				Kine:       build.KineVersion,
			}

			info.Print(cmd.OutOrStdout())
		},
	}

	// append flags
	cmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "use to print all k0s version info")
	cmd.PersistentFlags().BoolVarP(&isJsn, "json", "j", false, "use to print all k0s version info in json")
	return cmd
}

type versionInfo struct {
	Version    string `json:"k0s,omitempty"`
	Kubernetes string `json:"kubernetes,omitempty"`
	Kine       string `json:"kine,omitempty"`
}

func (v versionInfo) Print(w io.Writer) {
	if all {
		fmt.Fprintln(w, "k0s :", v.Version)
		fmt.Fprintln(w, "kubernetes :", v.Kubernetes)
		fmt.Fprintln(w, "kine :", v.Kine)
	} else if isJsn {
		jsn, _ := json.MarshalIndent(v, "", "   ")
		fmt.Fprintln(w, string(jsn))
	} else {
		fmt.Fprintln(w, v.Version)
	}
}
