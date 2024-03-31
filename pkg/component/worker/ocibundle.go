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

package worker

import (
	"context"
	"github.com/iscas-fork/k0s/internal/pkg/dir"
	"github.com/iscas-fork/k0s/pkg/component/manager"
	"github.com/iscas-fork/k0s/pkg/component/prober"
	"github.com/iscas-fork/k0s/pkg/config"
	"github.com/iscas-fork/k0s/pkg/constant"
	"github.com/sirupsen/logrus"
)

// OCIBundleReconciler tries to import OCI bundle into the running containerd instance
type OCIBundleReconciler struct {
	k0sVars *config.CfgVars
	log     *logrus.Entry
	*prober.EventEmitter
}

var _ manager.Component = (*OCIBundleReconciler)(nil)

// NewOCIBundleReconciler builds new reconciler
func NewOCIBundleReconciler(vars *config.CfgVars) *OCIBundleReconciler {
	return &OCIBundleReconciler{
		k0sVars:      vars,
		log:          logrus.WithField("component", "OCIBundleReconciler"),
		EventEmitter: prober.NewEventEmitter(),
	}
}

func (a *OCIBundleReconciler) Init(_ context.Context) error {
	return dir.Init(a.k0sVars.OCIBundleDir, constant.ManifestsDirMode)
}

func (a *OCIBundleReconciler) Start(ctx context.Context) error {
	return nil
}

func (a *OCIBundleReconciler) Stop() error {
	return nil
}
