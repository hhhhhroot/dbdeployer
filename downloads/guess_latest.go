// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2019 Giuseppe Maxia
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
package downloads

type GuessInfo struct {
	Version string
	Url     string
}

var (
	DownloadUrlList = map[string]string{
		"5.7": "https://dev.mysql.com/get/Downloads/MySQL-5.7",
		"8.0": "https://dev.mysql.com/get/Downloads/MySQL-8.0",
	}
	FileNameTemplates = map[string]string{
		"linux":  "mysql-{{.Version}}-linux-x86_64-minimal.{{.Ext}}",
		"darwin": "mysql-{{.Version}}-macos10.14-x86_64.{{.Ext}}",
	}
	Extensions = map[string]map[string]string{
		"linux": {
			"5.7": "tar.gz",
			"8.0": "tar.xz",
		},
		"darwin": {
			"5.7": "tar.gz",
			"8.0": "tar.gz",
		},
	}
)
