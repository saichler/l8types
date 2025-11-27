package tests

import (
	"testing"

	"github.com/saichler/l8types/go/ifs"
	"github.com/saichler/l8types/go/types/l8web"
)

// Mock implementations for testing
type mockServiceHandler struct{}

func (m *mockServiceHandler) Activate(*ifs.ServiceLevelAgreement, ifs.IVNic) error { return nil }
func (m *mockServiceHandler) DeActivate() error                                    { return nil }
func (m *mockServiceHandler) Post(ifs.IElements, ifs.IVNic) ifs.IElements          { return nil }
func (m *mockServiceHandler) Put(ifs.IElements, ifs.IVNic) ifs.IElements           { return nil }
func (m *mockServiceHandler) Patch(ifs.IElements, ifs.IVNic) ifs.IElements         { return nil }
func (m *mockServiceHandler) Delete(ifs.IElements, ifs.IVNic) ifs.IElements        { return nil }
func (m *mockServiceHandler) Get(ifs.IElements, ifs.IVNic) ifs.IElements           { return nil }
func (m *mockServiceHandler) Failed(ifs.IElements, ifs.IVNic, *ifs.Message) ifs.IElements {
	return nil
}
func (m *mockServiceHandler) TransactionConfig() ifs.ITransactionConfig { return nil }
func (m *mockServiceHandler) WebService() ifs.IWebService               { return nil }

type mockSLAServiceCallback struct{}

func (m *mockSLAServiceCallback) BeforePost(interface{}, ifs.IVNic) interface{}   { return nil }
func (m *mockSLAServiceCallback) AfterPost(interface{}, ifs.IVNic) interface{}    { return nil }
func (m *mockSLAServiceCallback) BeforePut(interface{}, ifs.IVNic) interface{}    { return nil }
func (m *mockSLAServiceCallback) AfterPut(interface{}, ifs.IVNic) interface{}     { return nil }
func (m *mockSLAServiceCallback) BeforePatch(interface{}, ifs.IVNic) interface{}  { return nil }
func (m *mockSLAServiceCallback) AfterPatch(interface{}, ifs.IVNic) interface{}   { return nil }
func (m *mockSLAServiceCallback) BeforeDelete(interface{}, ifs.IVNic) interface{} { return nil }
func (m *mockSLAServiceCallback) AfterDelete(interface{}, ifs.IVNic) interface{}  { return nil }
func (m *mockSLAServiceCallback) BeforeGet(interface{}, ifs.IVNic) interface{}    { return nil }
func (m *mockSLAServiceCallback) AfterGet(interface{}, ifs.IVNic) interface{}     { return nil }

type mockSLAStorage struct{}

func (m *mockSLAStorage) Put(string, interface{}) error                               { return nil }
func (m *mockSLAStorage) Get(string) (interface{}, error)                             { return nil, nil }
func (m *mockSLAStorage) Delete(string) (interface{}, error)                          { return nil, nil }
func (m *mockSLAStorage) Collect(f func(interface{}) (bool, interface{})) map[string]interface{} {
	return nil
}
func (m *mockSLAStorage) CacheEnabled() bool { return false }

type mockSLAWebService struct{}

func (m *mockSLAWebService) Vnet() uint32                              { return 0 }
func (m *mockSLAWebService) ServiceName() string                       { return "" }
func (m *mockSLAWebService) ServiceArea() byte                         { return 0 }
func (m *mockSLAWebService) PostBody() string                          { return "" }
func (m *mockSLAWebService) PostResp() string                          { return "" }
func (m *mockSLAWebService) PutBody() string                           { return "" }
func (m *mockSLAWebService) PutResp() string                           { return "" }
func (m *mockSLAWebService) PatchBody() string                         { return "" }
func (m *mockSLAWebService) PatchResp() string                         { return "" }
func (m *mockSLAWebService) DeleteBody() string                        { return "" }
func (m *mockSLAWebService) DeleteResp() string                        { return "" }
func (m *mockSLAWebService) GetBody() string                           { return "" }
func (m *mockSLAWebService) GetResp() string                           { return "" }
func (m *mockSLAWebService) Serialize() *l8web.L8WebService            { return nil }
func (m *mockSLAWebService) DeSerialize(*l8web.L8WebService)           {}
func (m *mockSLAWebService) Plugin() string                            { return "" }

func TestNewServiceLevelAgreement(t *testing.T) {
	handler := &mockServiceHandler{}
	callback := &mockSLAServiceCallback{}
	serviceName := "testService"
	serviceArea := byte(1)
	stateful := true

	sla := ifs.NewServiceLevelAgreement(handler, serviceName, serviceArea, stateful, callback)

	if sla == nil {
		t.Fatal("NewServiceLevelAgreement returned nil")
	}

	if sla.ServiceHandlerInstance() == nil {
		t.Error("ServiceHandlerInstance not set correctly")
	}

	if sla.ServiceName() != serviceName {
		t.Errorf("Expected ServiceName %s, got %s", serviceName, sla.ServiceName())
	}

	if sla.ServiceArea() != serviceArea {
		t.Errorf("Expected ServiceArea %d, got %d", serviceArea, sla.ServiceArea())
	}

	if sla.Stateful() != stateful {
		t.Errorf("Expected Stateful %t, got %t", stateful, sla.Stateful())
	}

	if sla.Callback() == nil {
		t.Error("Callback not set correctly")
	}
}

func TestServiceLevelAgreement_ServiceItem(t *testing.T) {
	sla := createTestSLA()
	testItem := "test item"

	sla.SetServiceItem(testItem)

	if sla.ServiceItem() != testItem {
		t.Errorf("Expected ServiceItem %v, got %v", testItem, sla.ServiceItem())
	}
}

func TestServiceLevelAgreement_ServiceItemList(t *testing.T) {
	sla := createTestSLA()
	testList := []string{"item1", "item2"}

	sla.SetServiceItemList(testList)

	if sla.ServiceItemList() == nil {
		t.Error("ServiceItemList should not be nil")
	}
}

func TestServiceLevelAgreement_InitItems(t *testing.T) {
	sla := createTestSLA()
	initItems := []interface{}{"init1", "init2", "init3"}

	sla.SetInitItems(initItems)

	retrieved := sla.InitItems()
	if len(retrieved) != len(initItems) {
		t.Errorf("Expected %d InitItems, got %d", len(initItems), len(retrieved))
	}
}

func TestServiceLevelAgreement_PrimaryKeys(t *testing.T) {
	sla := createTestSLA()
	keys := []string{"id", "name", "timestamp"}

	sla.SetPrimaryKeys(keys...)

	retrieved := sla.PrimaryKeys()
	if len(retrieved) != len(keys) {
		t.Errorf("Expected %d PrimaryKeys, got %d", len(keys), len(retrieved))
	}

	for i, key := range keys {
		if retrieved[i] != key {
			t.Errorf("Expected PrimaryKey[%d] to be %s, got %s", i, key, retrieved[i])
		}
	}
}

func TestServiceLevelAgreement_Store(t *testing.T) {
	sla := createTestSLA()
	store := &mockSLAStorage{}

	sla.SetStore(store)

	if sla.Store() == nil {
		t.Error("Store not set correctly")
	}
}

func TestServiceLevelAgreement_Voter(t *testing.T) {
	sla := createTestSLA()

	if sla.Voter() != false {
		t.Error("Expected default Voter to be false")
	}

	sla.SetVoter(true)

	if sla.Voter() != true {
		t.Error("Expected Voter to be true after setting")
	}
}

func TestServiceLevelAgreement_Transactional(t *testing.T) {
	sla := createTestSLA()

	if sla.Transactional() != false {
		t.Error("Expected default Transactional to be false")
	}

	sla.SetTransactional(true)

	if sla.Transactional() != true {
		t.Error("Expected Transactional to be true after setting")
	}
}

func TestServiceLevelAgreement_Replication(t *testing.T) {
	sla := createTestSLA()

	if sla.Replication() != false {
		t.Error("Expected default Replication to be false")
	}

	sla.SetReplication(true)

	if sla.Replication() != true {
		t.Error("Expected Replication to be true after setting")
	}
}

func TestServiceLevelAgreement_ReplicationCount(t *testing.T) {
	sla := createTestSLA()

	if sla.ReplicationCount() != 0 {
		t.Error("Expected default ReplicationCount to be 0")
	}

	expectedCount := 3
	sla.SetReplicationCount(expectedCount)

	if sla.ReplicationCount() != expectedCount {
		t.Errorf("Expected ReplicationCount %d, got %d", expectedCount, sla.ReplicationCount())
	}
}

func TestServiceLevelAgreement_WebService(t *testing.T) {
	sla := createTestSLA()
	webService := &mockSLAWebService{}

	sla.SetWebService(webService)

	if sla.WebService() == nil {
		t.Error("WebService not set correctly")
	}
}

func TestServiceLevelAgreement_Args(t *testing.T) {
	sla := createTestSLA()
	args := []interface{}{"arg1", 123, true}

	sla.SetArgs(args...)

	retrieved := sla.Args()
	if len(retrieved) != len(args) {
		t.Errorf("Expected %d Args, got %d", len(args), len(retrieved))
	}

	for i, arg := range args {
		if retrieved[i] != arg {
			t.Errorf("Expected Arg[%d] to be %v, got %v", i, arg, retrieved[i])
		}
	}
}

func TestServiceLevelAgreement_MetadataFunc(t *testing.T) {
	sla := createTestSLA()

	// Initially should be nil
	if sla.MetadataFunc() != nil {
		t.Error("Expected MetadataFunc to be nil initially")
	}

	// Add a metadata function
	testFunc := func(data interface{}) (bool, string) {
		return true, "test"
	}

	sla.AddMetadataFunc("testFunc", testFunc)

	metadataFuncs := sla.MetadataFunc()
	if metadataFuncs == nil {
		t.Fatal("MetadataFunc should not be nil after adding function")
	}

	if len(metadataFuncs) != 1 {
		t.Errorf("Expected 1 MetadataFunc, got %d", len(metadataFuncs))
	}

	if _, exists := metadataFuncs["testFunc"]; !exists {
		t.Error("Expected testFunc to exist in MetadataFunc map")
	}

	// Test the function
	if fn, ok := metadataFuncs["testFunc"]; ok {
		result, message := fn(nil)
		if !result || message != "test" {
			t.Errorf("Expected function to return (true, 'test'), got (%t, '%s')", result, message)
		}
	}
}

func TestServiceLevelAgreement_AddMultipleMetadataFuncs(t *testing.T) {
	sla := createTestSLA()

	func1 := func(data interface{}) (bool, string) { return true, "func1" }
	func2 := func(data interface{}) (bool, string) { return false, "func2" }
	func3 := func(data interface{}) (bool, string) { return true, "func3" }

	sla.AddMetadataFunc("func1", func1)
	sla.AddMetadataFunc("func2", func2)
	sla.AddMetadataFunc("func3", func3)

	metadataFuncs := sla.MetadataFunc()
	if len(metadataFuncs) != 3 {
		t.Errorf("Expected 3 MetadataFuncs, got %d", len(metadataFuncs))
	}

	// Verify all functions exist
	for _, name := range []string{"func1", "func2", "func3"} {
		if _, exists := metadataFuncs[name]; !exists {
			t.Errorf("Expected %s to exist in MetadataFunc map", name)
		}
	}
}

// Helper function to create a test SLA instance
func createTestSLA() *ifs.ServiceLevelAgreement {
	return ifs.NewServiceLevelAgreement(
		&mockServiceHandler{},
		"testService",
		byte(1),
		true,
		&mockSLAServiceCallback{},
	)
}
