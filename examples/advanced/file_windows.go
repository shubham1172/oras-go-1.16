/*
Copyright The ORAS Authors.
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

package main

import (
	"strings"
	"unicode"
)

// parseFileRef parse file reference on windows.
// Windows systems does not allow ':' in the file path except for drive letter.
func parseFileRef(ref string, mediaType string) (string, string) {
	i := strings.Index(ref, ":")
	if i < 0 {
		return ref, mediaType
	}

	// In case it is C:\
	if i == 1 && len(ref) > 2 && ref[2] == '\\' && unicode.IsLetter(rune(ref[0])) {
		i = strings.Index(ref[3:], ":")
		if i < 0 {
			return ref, mediaType
		}
		i += 3
	}

	return ref[:i], ref[i+1:]
}
