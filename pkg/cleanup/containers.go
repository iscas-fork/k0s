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

package cleanup

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"k8s.io/mount-utils"
	"os"
	"strings"
)

type containers struct {
	Config *Config
}

// Name returns the name of the step
func (c *containers) Name() string {
	return "containers steps"
}

// Run removes all the pods and mounts and stops containers afterwards
// Run starts containerd if custom CRI is not configured
func (c *containers) Run() error {

	if err := c.stopAllContainers(); err != nil {
		logrus.Debugf("error stopping containers: %v", err)
	}

	return nil
}

func removeMount(path string) error {
	var msg []string

	mounter := mount.New("")
	procMounts, err := mounter.List()
	if err != nil {
		return err
	}
	for _, v := range procMounts {
		if strings.Contains(v.Path, path) {
			logrus.Debugf("Unmounting: %s", v.Path)
			if err = mounter.Unmount(v.Path); err != nil {
				msg = append(msg, err.Error())
			}

			logrus.Debugf("Removing: %s", v.Path)
			if err := os.RemoveAll(v.Path); err != nil {
				msg = append(msg, err.Error())
			}
		}
	}
	if len(msg) > 0 {
		return fmt.Errorf("%v", strings.Join(msg, "\n"))
	}
	return nil
}

func (c *containers) isCustomCriUsed() bool {
	return true
}

func (c *containers) stopAllContainers() error {
	return nil
}
