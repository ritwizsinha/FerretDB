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

package integration

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateFieldCompatInc(t *testing.T) {
	t.Parallel()

	testCases := map[string]updateCompatTestCase{
		"Int32": {
			update: bson.D{{"$inc", bson.D{{"v", int32(42)}}}},
			skip:   "https://github.com/FerretDB/FerretDB/issues/972",
		},
		"Int32Negative": {
			update: bson.D{{"$inc", bson.D{{"v", int32(-42)}}}},
			skip:   "https://github.com/FerretDB/FerretDB/issues/972",
		},
		"EmptyUpdatePath": {
			update: bson.D{{"$inc", bson.D{{}}}},
			skip:   "https://github.com/FerretDB/FerretDB/issues/673",
		},
		"DotNotationFieldExist": {
			update: bson.D{{"$inc", bson.D{{"v.foo", int32(1)}}}},
			skip:   "https://github.com/FerretDB/FerretDB/issues/972",
		},
		"DotNotationFieldNotExist": {
			update: bson.D{{"$inc", bson.D{{"foo.bar", int32(1)}}}},
			skip:   "https://github.com/FerretDB/FerretDB/issues/972",
		},
	}

	testUpdateCompat(t, testCases)
}

func TestUpdateFieldCompatUnset(t *testing.T) {
	t.Parallel()

	testCases := map[string]updateCompatTestCase{
		"Simple": {
			update:        bson.D{{"$unset", bson.D{{"v", ""}}}},
			skipForTigris: true,
		},
		"NotExistedField": {
			update: bson.D{{"$unset", bson.D{{"foo", ""}}}},
		},
		"NestedField": {
			update:        bson.D{{"$unset", bson.D{{"v", bson.D{{"array", ""}}}}}},
			skipForTigris: true,
		},
		"DotNotationDocumentFieldExist": {
			update: bson.D{{"$unset", bson.D{{"v.foo", ""}}}},
		},
		"DotNotationDocumentFieldNotExist": {
			update: bson.D{{"$unset", bson.D{{"foo.bar", ""}}}},
		},
		"DotNotationArrayFieldExist": {
			update: bson.D{{"$unset", bson.D{{"v.array.0", int32(1)}}}},
		},
		"DotNotationArrayFieldNotExist": {
			update: bson.D{{"$unset", bson.D{{"foo.0.baz", int32(1)}}}},
		},
		"DocumentDotNotationArrFieldNotExist": {
			update: bson.D{{"$unset", bson.D{{"v.0.foo", int32(1)}}}},
		},
	}

	testUpdateCompat(t, testCases)
}
