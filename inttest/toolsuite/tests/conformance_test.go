/*
Copyright 2022 k0s authors

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

package tests

import (
	"flag"
	"fmt"
	"testing"

	ts "github.com/iscas-fork/k0s/inttest/toolsuite"
	tsops "github.com/iscas-fork/k0s/inttest/toolsuite/operations"

	"github.com/stretchr/testify/suite"
)

type ConformanceConfig struct {
	SonobuoyVersion   string
	KubernetesVersion string
}

type ConformanceSuite struct {
	ts.ToolSuite
}

var config ConformanceConfig

func init() {
	flag.StringVar(&config.SonobuoyVersion, "sonobuoy-version", "", "The sonobuoy version to use")
	flag.StringVar(&config.KubernetesVersion, "conformance-kubernetes-version", "", "The kubernetes version of the conformance tests to run")
}

// TestConformanceSuite runs the Sonobuoy conformance tests for a specific k8s version.
func TestConformanceSuite(t *testing.T) {
	if config.SonobuoyVersion == "" {
		t.Fatal("--sonobuoy-version is a required parameter")
	}
	if config.KubernetesVersion == "" {
		t.Fatal("--conformance-kubernetes-version is a required parameter")
	}

	suite.Run(t, &ConformanceSuite{
		ts.ToolSuite{
			Operation: tsops.SonobuoyOperation(
				tsops.SonobuoyConfig{
					Version: config.SonobuoyVersion,
					Parameters: []string{
						"--mode=certified-conformance",
						"--plugin-env=e2e.E2E_EXTRA_ARGS=\"--ginkgo.v\"",
						fmt.Sprintf("--kubernetes-version=%s", config.KubernetesVersion),
					},
				},
			),
		},
	})
}
