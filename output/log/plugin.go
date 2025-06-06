// Copyright 2023 Blink Labs Software
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

package log

import (
	"github.com/jumbohurric/adder/internal/logging"
	"github.com/jumbohurric/adder/plugin"
)

var cmdlineOptions struct {
	level string
}

func init() {
	plugin.Register(
		plugin.PluginEntry{
			Type:               plugin.PluginTypeOutput,
			Name:               "log",
			Description:        "display events to the console using the logger",
			NewFromOptionsFunc: NewFromCmdlineOptions,
			Options: []plugin.PluginOption{
				{
					Name:         "level",
					Type:         plugin.PluginOptionTypeString,
					Description:  "specifies the log level to use",
					DefaultValue: "info",
					Dest:         &(cmdlineOptions.level),
				},
			},
		},
	)
}

func NewFromCmdlineOptions() plugin.Plugin {
	p := New(
		WithLogger(
			logging.GetLogger().With("plugin", "output.log"),
		),
		WithLevel(cmdlineOptions.level),
	)
	return p
}
