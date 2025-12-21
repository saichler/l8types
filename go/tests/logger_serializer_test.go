/*
© 2025 Sharon Aicler (saichler@gmail.com)

Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tests

import (
	"testing"

	"github.com/saichler/l8types/go/ifs"
)

func TestLogLevelString(t *testing.T) {
	tests := []struct {
		level    ifs.LogLevel
		expected string
	}{
		{ifs.Trace_Level, "(Trace)"},
		{ifs.Debug_Level, "(Debug)"},
		{ifs.Info_Level, "(Info) "},
		{ifs.Warning_Level, "(Warn )"},
		{ifs.Error_Level, "(Error)"},
		{ifs.LogLevel(99), ""},
	}

	for _, test := range tests {
		result := test.level.String()
		if result != test.expected {
			t.Errorf("For level %d, expected '%s', got '%s'", test.level, test.expected, result)
		}
	}
}

func TestIsNil(t *testing.T) {
	// Test nil value
	var nilPtr *string
	if !ifs.IsNil(nilPtr) {
		t.Error("Expected IsNil to return true for nil pointer")
	}

	// Test non-nil value
	str := "test"
	if ifs.IsNil(&str) {
		t.Error("Expected IsNil to return false for non-nil pointer")
	}

	// Test nil interface
	var nilInterface interface{}
	if !ifs.IsNil(nilInterface) {
		t.Error("Expected IsNil to return true for nil interface")
	}

	// Test nil slice
	var nilSlice []string
	if !ifs.IsNil(nilSlice) {
		t.Error("Expected IsNil to return true for nil slice")
	}

	// Test non-nil slice
	nonNilSlice := []string{"test"}
	if ifs.IsNil(nonNilSlice) {
		t.Error("Expected IsNil to return false for non-nil slice")
	}

	// Test nil map
	var nilMap map[string]string
	if !ifs.IsNil(nilMap) {
		t.Error("Expected IsNil to return true for nil map")
	}

	// Test non-nil map
	nonNilMap := make(map[string]string)
	if ifs.IsNil(nonNilMap) {
		t.Error("Expected IsNil to return false for non-nil map")
	}
}

func TestSerializerModeConstants(t *testing.T) {
	// Test that serializer mode constants are defined correctly
	if ifs.BINARY != 1 {
		t.Errorf("Expected BINARY to be 1, got %d", ifs.BINARY)
	}
	if ifs.JSON != 2 {
		t.Errorf("Expected JSON to be 2, got %d", ifs.JSON)
	}
	if ifs.STRING != 3 {
		t.Errorf("Expected STRING to be 3, got %d", ifs.STRING)
	}
}
