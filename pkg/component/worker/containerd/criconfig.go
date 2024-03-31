/*
Copyright 2023 k0s authors

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

package containerd

import (
	"runtime"
	"strings"

	"github.com/iscas-fork/k0s/pkg/apis/k0s/v1beta1"
	"github.com/sirupsen/logrus"
)

// TODO: move to K0sVars
const containerdCRIConfigPathPosix = "/run/k0s/containerd-cri.toml"
const containerdCRIConfigPathWindows = "C:\\var\\lib\\k0s\\run\\containerd-cri.toml"

type CRIConfigurer struct {
	loadPath       string
	pauseImage     string
	criRuntimePath string

	log *logrus.Entry
}

func NewConfigurer(pauseImage *v1beta1.ImageSpec, importsPath string) *CRIConfigurer {
	c := &CRIConfigurer{
		loadPath:   importsPath,
		pauseImage: pauseImage.URI(),
		log:        logrus.WithField("component", "containerd"),
	}
	if runtime.GOOS == "windows" {
		c.criRuntimePath = containerdCRIConfigPathWindows

	} else {
		c.criRuntimePath = containerdCRIConfigPathPosix
	}
	return c
}

func escapedPath(s string) string {
	// double escape for windows because containerd expects
	// double backslash in the configuration but golang templates
	// unescape double slash to a single slash
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(s, "\\", "\\\\")
	}
	return s
}

// We need to use custom struct so we can unmarshal the CRI plugin config only
type config struct {
	Version int
	Plugins map[string]interface{} `toml:"plugins"`
}
