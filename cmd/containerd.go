/*
Copyright 2016 The Kubernetes Authors.

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

package cmd

import (
	"os" //nolint:staticcheck // Global math/rand seed is deprecated, but still used by external dependencies

	"github.com/containerd/containerd/cmd/ctr/app"
	"github.com/urfave/cli"

	cunionUtil "github.com/hyunseok95/cunion/pkg/util"
)

var pluginCmds = []cli.Command{}

func Containerd() {
	app := app.New()
	app.Commands = append(app.Commands, pluginCmds...)

	if err := app.Run(os.Args); err != nil {
		cunionUtil.CheckErr(err)
	}
}