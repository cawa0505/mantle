// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/coreos/mantle/cli"
	"github.com/coreos/mantle/kola"
)

const (
	cliName        = "kola"
	cliDescription = "The CoreOS Superdeep Borehole"
	// http://en.wikipedia.org/wiki/Kola_Superdeep_Borehole
)

// main test harness
var cmdRun = &cli.Command{
	Name:        "run",
	Summary:     "Run run kola tests by category",
	Description: "run all kola tests (default) or related groups",
	Run:         kola.RunTests,
}

func init() {
	cli.Register(cmdRun)
}

func main() {
	cli.Run(cliName, cliDescription)
}