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
	"strings"
	"testing"

	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/types/l8health"
	"github.com/saichler/l8types/go/types/l8services"
	"github.com/saichler/l8types/go/types/l8sysconfig"
)

func TestAddService(t *testing.T) {
	// Test with nil config - should not panic
	ifs.AddService(nil, "test-service", 1)

	// Test with empty config
	config := &l8sysconfig.L8SysConfig{}
	ifs.AddService(config, "test-service", 1)

	if config.LocalUuid == "" {
		t.Error("Expected LocalUuid to be set")
	}
	if config.Services == nil {
		t.Error("Expected Services to be initialized")
	}
	if config.Services.ServiceToAreas == nil {
		t.Error("Expected ServiceToAreas to be initialized")
	}

	// Check that service was added
	areas, ok := config.Services.ServiceToAreas["test-service"]
	if !ok {
		t.Error("Expected 'test-service' to be in ServiceToAreas")
	}
	if areas.Areas == nil {
		t.Error("Expected Areas to be initialized")
	}
	if !areas.Areas[1] {
		t.Error("Expected area 1 to be true for 'test-service'")
	}

	// Test adding another area to same service
	ifs.AddService(config, "test-service", 2)
	if !config.Services.ServiceToAreas["test-service"].Areas[2] {
		t.Error("Expected area 2 to be true for 'test-service'")
	}
	if len(config.Services.ServiceToAreas["test-service"].Areas) != 2 {
		t.Errorf("Expected 2 areas, got %d", len(config.Services.ServiceToAreas["test-service"].Areas))
	}

	// Test adding a different service
	ifs.AddService(config, "another-service", 3)
	if !config.Services.ServiceToAreas["another-service"].Areas[3] {
		t.Error("Expected area 3 to be true for 'another-service'")
	}
	if len(config.Services.ServiceToAreas) != 2 {
		t.Errorf("Expected 2 services, got %d", len(config.Services.ServiceToAreas))
	}
}

func TestRemoveService(t *testing.T) {
	// Test with nil services - should not panic
	ifs.RemoveService(nil, "test-service", 1)

	// Test with empty services
	services := &l8services.L8Services{}
	ifs.RemoveService(services, "test-service", 1)

	// Test with initialized but empty ServiceToAreas
	services.ServiceToAreas = make(map[string]*l8services.L8ServiceAreas)
	ifs.RemoveService(services, "test-service", 1)

	// Test with actual service
	services.ServiceToAreas["test-service"] = &l8services.L8ServiceAreas{
		Areas: make(map[int32]bool),
	}
	services.ServiceToAreas["test-service"].Areas[1] = true
	services.ServiceToAreas["test-service"].Areas[2] = true

	// Remove one area
	ifs.RemoveService(services, "test-service", 1)
	if services.ServiceToAreas["test-service"].Areas[1] {
		t.Error("Expected area 1 to be removed")
	}
	if !services.ServiceToAreas["test-service"].Areas[2] {
		t.Error("Expected area 2 to still exist")
	}
	if len(services.ServiceToAreas) != 1 {
		t.Error("Expected service to still exist since it has remaining areas")
	}

	// Remove last area - should remove service entirely
	ifs.RemoveService(services, "test-service", 2)
	if _, ok := services.ServiceToAreas["test-service"]; ok {
		t.Error("Expected 'test-service' to be removed entirely when last area is removed")
	}

	// Test removing non-existent service
	ifs.RemoveService(services, "non-existent", 1)
}

func TestNewUuid(t *testing.T) {
	// Test that NewUuid generates a valid UUID
	uuid1 := ifs.NewUuid()
	if uuid1 == "" {
		t.Error("Expected NewUuid to return a non-empty string")
	}

	// Check UUID format (should have dashes in the right places)
	if !strings.Contains(uuid1, "-") {
		t.Error("Expected UUID to contain dashes")
	}

	// Test that two calls generate different UUIDs
	uuid2 := ifs.NewUuid()
	if uuid1 == uuid2 {
		t.Error("Expected NewUuid to generate unique UUIDs")
	}

	// Test UUID length (standard UUID is 36 characters)
	if len(uuid1) != 36 {
		t.Errorf("Expected UUID length to be 36, got %d", len(uuid1))
	}

	// Test UUID format more specifically (8-4-4-4-12)
	parts := strings.Split(uuid1, "-")
	if len(parts) != 5 {
		t.Errorf("Expected UUID to have 5 parts separated by dashes, got %d", len(parts))
	}
	if len(parts) == 5 {
		if len(parts[0]) != 8 || len(parts[1]) != 4 || len(parts[2]) != 4 ||
		   len(parts[3]) != 4 || len(parts[4]) != 12 {
			t.Errorf("UUID format incorrect: %s", uuid1)
		}
	}
}

func TestMergeServices_NilTargetServices(t *testing.T) {
	hp := &l8health.L8Health{}
	src := &l8services.L8Services{
		ServiceToAreas: map[string]*l8services.L8ServiceAreas{
			"svcA": {Areas: map[int32]bool{1: true}},
		},
	}

	ifs.MergeServices(hp, src)

	if hp.Services != src {
		t.Error("Expected hp.Services to be assigned to src when initially nil")
	}
}

func TestMergeServices_NewServiceName(t *testing.T) {
	hp := &l8health.L8Health{
		Services: &l8services.L8Services{
			ServiceToAreas: map[string]*l8services.L8ServiceAreas{
				"existing": {Areas: map[int32]bool{1: true}},
			},
		},
	}
	src := &l8services.L8Services{
		ServiceToAreas: map[string]*l8services.L8ServiceAreas{
			"newSvc": {Areas: map[int32]bool{5: true}},
		},
	}

	ifs.MergeServices(hp, src)

	if _, ok := hp.Services.ServiceToAreas["existing"]; !ok {
		t.Error("Expected 'existing' service to be preserved")
	}
	got, ok := hp.Services.ServiceToAreas["newSvc"]
	if !ok {
		t.Fatal("Expected 'newSvc' to be added")
	}
	if !got.Areas[5] {
		t.Error("Expected 'newSvc' area 5 to be true")
	}
}

func TestMergeServices_MergeAreasIntoExistingService(t *testing.T) {
	hp := &l8health.L8Health{
		Services: &l8services.L8Services{
			ServiceToAreas: map[string]*l8services.L8ServiceAreas{
				"svcA": {Areas: map[int32]bool{1: true}},
			},
		},
	}
	src := &l8services.L8Services{
		ServiceToAreas: map[string]*l8services.L8ServiceAreas{
			"svcA": {Areas: map[int32]bool{2: true, 3: true}},
		},
	}

	ifs.MergeServices(hp, src)

	areas := hp.Services.ServiceToAreas["svcA"].Areas
	if !areas[1] || !areas[2] || !areas[3] {
		t.Errorf("Expected areas {1,2,3} all present, got %v", areas)
	}
}

func TestMergeServices_MergeModelsIntoExistingService(t *testing.T) {
	hp := &l8health.L8Health{
		Services: &l8services.L8Services{
			ServiceToAreas: map[string]*l8services.L8ServiceAreas{
				"svcA": {
					Areas:  map[int32]bool{1: true},
					Models: map[int32]string{1: "ModelA"},
				},
			},
		},
	}
	src := &l8services.L8Services{
		ServiceToAreas: map[string]*l8services.L8ServiceAreas{
			"svcA": {
				Areas:  map[int32]bool{2: true},
				Models: map[int32]string{2: "ModelB"},
			},
		},
	}

	ifs.MergeServices(hp, src)

	models := hp.Services.ServiceToAreas["svcA"].Models
	if models[1] != "ModelA" || models[2] != "ModelB" {
		t.Errorf("Expected models {1:ModelA, 2:ModelB}, got %v", models)
	}
}
