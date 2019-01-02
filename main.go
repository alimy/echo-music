// Copyright © 2018 Micheal Li <alimy@gility.net>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/alimy/echo-music/cmd"
	_ "github.com/alimy/echo-music/module/serve/cmd"
	_ "github.com/alimy/echo-music/version"
	"github.com/unisx/logus"
)

func main() {
	defer logus.Sync()

	// Setup root cli command of application
	cmd.Setup(
		"echo-music",                        // command name
		"music infomation service",          // command short describe
		"music infomation provider service", // command long describe
	)

	// Execute start application
	cmd.Execute()
}