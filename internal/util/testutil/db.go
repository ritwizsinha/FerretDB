// Copyright 2021 FerretDB Inc.
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

package testutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// DatabaseName returns a stable database name for that test.
func DatabaseName(tb testing.TB) string {
	tb.Helper()

	// database names are always lowercase
	name := strings.ToLower(tb.Name())

	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "$", "_")

	require.Less(tb, len(name), 64)
	return name
}

// CollectionName returns a stable collection name for that test.
func CollectionName(tb testing.TB) string {
	tb.Helper()

	// do not use strings.ToLower because collection names can contain uppercase letters
	name := tb.Name()

	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "$", "_")

	require.Less(tb, len(name), 255)
	return name
}
