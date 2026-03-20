# Layer 8 Types (L8Types)

A foundational library for Layer 8 distributed networking systems, providing Protocol Buffer-based type definitions, Go interfaces, and utilities for building distributed applications with service discovery, health monitoring, and secure communication.

## Recent Updates

### Latest Features (v1.7.0)
- **AI Agent Aggregation API**: SQL-style aggregation support for AI-powered analytics
  - `L8AggregateFunction` message for parsed aggregate functions (`count`, `sum`, `avg`, `min`, `max`)
  - `group_by` field on `L8Query` for grouping query results
  - `having` clause on `L8Query` for filtering on aggregate results
  - `aggregates` field for parsed SELECT-clause aggregate functions
  - Anthropic API key integration via `ANTHROPIC_API_KEY` environment variable
- **Two-Phase Authentication**: Enhanced login security with token hash verification
  - `token_hash` and `code` fields on `AuthUser` for two-phase auth flow
  - `token_hash` on `AuthToken` for secure session handoff
  - Updated `L8TFAVerify` to use `hash` field and return `token` on success
- **Service-to-Model Mapping**: Model name tracking per service area
  - `models` map on `L8ServiceAreas` linking area IDs to protobuf model names
  - `AddService()` now accepts an optional `modelName` parameter
  - `MergeServices()` propagates model mappings during health merges
- **Concurrency Safety**: Thread-safe service merging with mutex protection
  - `MergeServices()` function for safely merging service registrations into health records

### Previous Release (v1.6.0)
- **Time-Series Database (TSDB) Support**: Built-in time-series data collection and querying
  - New `ITSDBService` interface with `AddTSDB()` and `GetTSDB()` methods
  - `L8TimeSeriesPoint` message for timestamped data points
  - `L8TSDBQuery` for time-range queries with property filtering
  - `L8TSDBNotification` for real-time time-series change events
- **Data Import/Export Framework**: Comprehensive ETL and data migration system
  - `L8ImportTemplate` with column mappings and value transforms
  - 10 transform types: date format, enum mapping, unit conversion, trim, case conversion, default values, concatenation, split, and money formatting
  - AI-assisted column mapping with confidence scoring (`L8ImportAIMappingRequest/Response`)
  - Import execution with per-row error reporting
  - Template export/import for portable migration bundles
- **File Upload/Download**: Document management primitives
  - `L8FileUploadRequest/Response` with metadata, MIME types, and SHA-256 checksums
  - `L8FileDownloadRequest/Response` for file retrieval by storage path
- **CSV Export**: Data export for reporting and analysis
  - `L8CsvExportRequest/Response` with model-based export and suggested filenames
- **Service Groups**: Shared leader election across related services
  - `SetServiceGroup()` / `ServiceGroup()` on SLA for group-based coordination
  - Multiple services can share a single leader election (e.g., all Financial services)
  - `SystemServiceGroup` constant for system-level coordination
- **Performance Profiling (pprof)**: Integrated Go profiling data in health monitoring
  - `pprof_memory` and `pprof_cpu` fields on `L8Health` for runtime profiling
  - `pprof_collect` flag to control when profiling data is gathered
- **Fiscal Period Model**: Business period modeling
  - `L8Period` message with type (Yearly/Quarterly/Monthly), year, and value
  - `L8PeriodType` and `L8PeriodValue` enums covering months (January-December) and quarters (Q1-Q4)
- **Configurable Logs Directory**: Per-deployment log path configuration
  - `logs_directory` field in `L8SysConfig` for per-product log segregation

### Previous Release (v1.5.0)
- **Enhanced Security Framework**: TFA, Captcha, Registration, and Credential management
- **Advanced SLA Decorators**: Non-Unique Keys and Always Overwrite support
- **Local Communication**: Same-VNic messaging with `Local()` and `LocalRequest()`
- **Service Callbacks**: Continuation indication for fine-grained lifecycle control
- **Web Service Refactoring**: Improved plugin and bearer validation architecture
- **Apache 2.0 Licensing**: Proper copyright headers on all source files

### Previous Release (v1.4.0)
- **Map-Reduce Framework**: Distributed Map-Reduce capabilities for parallel data processing
  - New `IMapReduceService` interface for distributed computation
  - `MapReduce()` flag in query system for distributed aggregation
  - `Collect()` method for data collection and filtering across services
- **Leader Election**: Leader-based communication patterns
  - `Leader()` and `LeaderRequest()` methods in VNic for leader-only messaging
  - Automatic leader election and failover support
- **Remote VNet Support**: Enhanced multi-network connectivity
  - Added `Vnet()` field to web services for cross-network operations
  - Improved service discovery across network boundaries

### Previous Release (v1.3.0)
- **Enhanced Authentication System**: Token-based authentication with validation and activation mechanisms
- **Improved Security**: Added hash-based security functions and error handling for auth operations
- **Logging Enhancements**: Comprehensive logging system for debugging and monitoring
- **Sorting Capabilities**: Added flexible sorting mechanisms for query results
- **Token Management**: Secure token generation, validation, and lifecycle management

## Overview

Layer 8 Types serves as the core type system and interface library for Layer 8 networking applications, offering:

- **Protocol Buffer Schemas**: Comprehensive type definitions for distributed system components
- **Virtual Network Interface (VNic)**: Advanced networking abstractions with leader election and cross-network support
- **Service Discovery & Management**: Built-in service registration, discovery, and area-based routing
- **AI Agent Analytics**: SQL-style aggregation with GROUP BY, HAVING, and aggregate functions for AI-powered queries
- **Distributed Computing**: Map-Reduce framework for parallel data processing and aggregation
- **Time-Series Database (TSDB)**: Built-in time-series data collection and querying for metrics and monitoring
- **Data Import/Export**: ETL framework with templates, transforms, AI-assisted mapping, and CSV export
- **File Management**: Upload/download primitives with checksums and metadata
- **Health Monitoring**: Real-time system health tracking with pprof profiling and Unix `top`-style output
- **Transaction Management**: Distributed transaction support with state tracking
- **Notification System**: Property change and time-series notifications with event propagation
- **Security Framework**: Enhanced authentication with TFA, two-phase auth, captcha, token validation, and encryption

## Key Features

### Virtual Networking (VNic)
- **Multiple Communication Patterns**: Unicast, Multicast, Round-robin, Proximity-based, Leader selection, Local
- **Local Communication**: Same-VNic messaging for co-located services (v1.5.0)
- **Leader Election**: Automatic leader election with failover support (v1.4.0)
- **Remote VNet Support**: Cross-network service discovery and communication (v1.4.0)
- **Network Mode Support**: Native, Docker, and Kubernetes networking modes
- **Service API**: RESTful service interfaces (POST, PUT, PATCH, DELETE, GET)
- **Message Priorities**: 8-level priority system (P1-P8) for message handling
- **Transaction Support**: Distributed transaction state management

### Service Discovery & Management
- **Service Areas**: Logical service grouping and area-based routing
- **Service-to-Model Mapping**: Track protobuf model names per service area (v1.7.0)
- **Service Groups**: Shared leader election across related services (v1.6.0)
- **Health States**: Up, Down, Unreachable status tracking with statistics
- **Performance Profiling**: Integrated pprof memory and CPU profiling in health data (v1.6.0)
- **Replication**: Service replication with endpoint scoring and key-based routing
- **Dynamic Service Registration**: Runtime service addition and removal
- **Map-Reduce Framework**: Distributed computation and aggregation across services (v1.4.0)
- **Time-Series Database**: `ITSDBService` for collecting and querying time-series data (v1.6.0)

### AI Agent Analytics (v1.7.0)
- **Aggregate Functions**: `count(*)`, `sum(field)`, `avg(field)`, `min(field)`, `max(field)`
- **GROUP BY**: Group query results by one or more fields
- **HAVING Clause**: Filter on aggregate results using expression trees
- **Anthropic Integration**: Built-in support for `ANTHROPIC_API_KEY` environment variable

### Data Import/Export (v1.6.0)
- **Import Templates**: Column mappings with 10 transform types (date, enum, unit, trim, case, etc.)
- **AI-Assisted Mapping**: Automatic column-to-field mapping with confidence scores
- **Import Execution**: Batch import with per-row error reporting
- **CSV Export**: Model-based data export with suggested filenames
- **File Upload/Download**: Document storage with SHA-256 checksums and MIME type support
- **Template Portability**: Export/import template bundles across environments

### Type System & Reflection
- **Dynamic Type Registry**: Runtime type registration and introspection
- **RNode Reflection**: Advanced reflection system with decorators and table views
- **Serialization Framework**: Pluggable serialization with multiple format support
- **Query Engine**: Expression-based querying with comparators, conditions, and aggregation

### Security & Configuration
- **Token Authentication**: Secure token-based authentication with validation and activation
- **Two-Phase Authentication**: Token hash-based two-phase login flow (v1.7.0)
- **Two-Factor Authentication**: TFA setup and verification for enhanced security (v1.5.0)
- **Captcha Protection**: Bot protection with captcha challenge/response (v1.5.0)
- **User Registration**: Secure user registration workflow (v1.5.0)
- **Credential Management**: Secure credential fetching and handling (v1.5.0)
- **AES Encryption**: Built-in symmetric encryption for secure communication
- **Hash Functions**: Cryptographic hash support for data integrity and password security
- **System Configuration**: Comprehensive configuration management with VNet settings
- **Authentication Framework**: Enhanced AAA (Authentication, Authorization, Accounting) support
- **Access Control**: Resource-based security with permission management and scope views
- **Bearer Token Validation**: HTTP bearer token validation for web services (v1.5.0)

## Protocol Buffer Schemas

### Core Types
- **`api.proto`**: Query expressions, conditions, comparators, aggregate functions, time-series points, period model, two-phase auth, data import/export framework, file upload/download, and CSV export
- **`services.proto`**: Service discovery, areas, model mappings, replication indices, and transactions
- **`health.proto`**: Health monitoring, statistics, pprof profiling data, and system status tracking
- **`sysconfig.proto`**: System configuration, VNet settings, logs directory, and connection parameters
- **`notification.proto`**: Property change notifications, TSDB notifications, and event propagation
- **`web.proto`**: Web service definitions and plugin system
- **`reflect.proto`**: Type reflection, nodes, decorators, and table views
- **`system.proto`**: System messages and route tables

### Message Types
- **Query**: Advanced search with criteria, sorting, pagination, schema filtering, and aggregate functions (GROUP BY, HAVING)
- **TSDB**: Time-series data points and range-based queries (v1.6.0)
- **Import/Export**: Templates, column mappings, transforms, AI mapping, file upload/download, CSV export (v1.6.0)
- **Period**: Fiscal/business period modeling with yearly, quarterly, monthly support (v1.6.0)
- **Health**: Process health with statistics (CPU, memory, message counts, pprof data)
- **Services**: Service-to-area mappings and replication management
- **SysConfig**: Network configuration including VNet ports, UUIDs, logs directory, and keep-alive settings
- **NotificationSet**: Batched property change and TSDB notifications with sequencing

## Project Structure

```
l8types/
├── proto/                      # Protocol Buffer definitions
│   ├── api.proto              # Query system, TSDB, import/export, file management, periods
│   ├── services.proto         # Service discovery and management
│   ├── health.proto           # Health monitoring, statistics, and pprof profiling
│   ├── sysconfig.proto        # System configuration and logs directory
│   ├── notification.proto     # Event and TSDB notification system
│   ├── web.proto              # Web service interfaces
│   ├── reflect.proto          # Type reflection system
│   ├── system.proto           # System messages and route tables
│   ├── tests.proto            # Test-specific type definitions
│   └── make-bindings.sh       # Code generation script
├── go/                        # Go implementation
│   ├── ifs/                   # Core interfaces
│   │   ├── API.go             # Elements, Query, Property interfaces
│   │   ├── Registry.go        # Type registration and reflection
│   │   ├── Resources.go       # Resource management
│   │   ├── VNic.go            # Virtual network interface
│   │   ├── Message.go         # Message structure and operations
│   │   ├── MessageEnums.go    # Priority, Action, Transaction enums
│   │   ├── Security.go        # Security and authentication
│   │   ├── Services.go        # Service management
│   │   ├── Web.go             # Web service interfaces
│   │   ├── Env.go             # Environment variable management (AI keys)
│   │   └── ...                # Additional interfaces
│   ├── types/                 # Generated Protocol Buffer types
│   ├── nets/                  # Network protocol implementation
│   ├── aes/                   # AES encryption utilities
│   ├── sec/                   # Security provider loading and defaults
│   ├── tests/                 # Test suite
│   └── testtypes/             # Test-specific generated types
```

## Getting Started

### Prerequisites
- **Go 1.26.1+**: Core implementation language
- **Docker**: Required for Protocol Buffer code generation
- **Protocol Buffers**: For schema compilation (handled via Docker)
### Installation

```bash
# Clone the repository
git clone https://github.com/saichler/l8types.git
cd l8types

# Install Go dependencies
cd go && go mod download
```

### Generate Protocol Buffer Bindings

```bash
cd proto
./make-bindings.sh
```

This script uses Docker to generate Go bindings from all Protocol Buffer schemas and organizes them into the appropriate packages.

### Running Tests

```bash
cd go
go test ./...
```

## API Migration Guide

### Breaking Changes in v1.3.0

#### 1. Service Activation Refactor
The `Activate()` method has been significantly simplified:

**Before:**
```go
handler, err := service.Activate(serviceName, serviceArea, priority,
    resources, listener, args...)
```

**After:**
```go
sla := &ServiceLevelAgreement{}
sla.SetServiceName("my-service")
sla.SetServiceArea(1)
sla.SetArgs(args...)
handler, err := service.Activate(sla, vnic)
```

#### 2. Primary Keys API Change
**Before:**
```go
sla.SetPrimaryKeys([]string{"id", "name"})
```

**After:**
```go
sla.SetPrimaryKeys("id", "name")  // Variadic parameters
```

#### 3. Registry Enhancement
**New Method:**
```go
// Create a new instance of a registered type dynamically
newInstance := registry.NewOf(existingInstance)
```

#### 4. Web Service Definition Rename
**Before:**
```go
webServiceDef := sla.WebServiceDef()
sla.SetWebServiceDef(webService)
```

**After:**
```go
webService := sla.WebService()
sla.SetWebService(webService)
```

## Usage Examples

### Service Registration and VNic Setup

```go
import (
    "github.com/saichler/l8types/go/ifs"
    "github.com/saichler/l8types/go/types"
)

// Create system configuration
config := &types.SysConfig{
    LocalUuid: "node-123",
    VnetPort: 8080,
    MaxDataSize: 1024*1024,
    TxQueueSize: 1000,
    RxQueueSize: 1000,
}

// Add services to configuration (optional model name for service-to-model mapping)
ifs.AddService(config, "user-service", 1, "User")
ifs.AddService(config, "auth-service", 2)

// Create VNic for networking
vnic := NewVNic(config, resources)
vnic.Start()
```

### Service Level Agreement Configuration (v1.3.0+)

```go
// Create Service Level Agreement
sla := &ifs.ServiceLevelAgreement{}
sla.SetServiceName("user-service")
sla.SetServiceArea(1)
sla.SetStateful(true)
sla.SetTransactional(true)
sla.SetReplication(true)
sla.SetReplicationCount(3)
sla.SetPrimaryKeys("userId", "email")  // Variadic parameters
sla.SetArgs(dbConnection, cacheLayer)  // Variadic parameters

// Activate service with SLA
handler, err := services.Activate(sla, vnic)
if err != nil {
    log.Fatal("Failed to activate service:", err)
}

// Dynamic type creation using registry
registry := ifs.NewRegistry()
userType := &User{}
registry.Register(userType)

// Create new instances dynamically
newUser := registry.NewOf(userType).(*User)
```

### Message Communication Patterns

```go
// Unicast to specific destination
err := vnic.Unicast("node-456", "user-service", 1, ifs.POST, userData)

// Round-robin to service providers
destination, err := vnic.RoundRobin("auth-service", 2, ifs.GET, request)

// Request with response
response := vnic.Request("node-789", "data-service", 1, ifs.GET, query)

// Multicast to all service providers
err := vnic.Multicast("notification-service", 1, ifs.Notify, event)

// Leader-based communication (v1.4.0+)
err := vnic.Leader("consensus-service", 1, ifs.POST, proposal)
leaderResponse := vnic.LeaderRequest("state-service", 1, ifs.GET, stateQuery, timeout)

// Local communication - same VNic (v1.5.0+)
err := vnic.Local("local-service", 1, ifs.POST, localData)
localResponse := vnic.LocalRequest("local-service", 1, ifs.GET, query, timeout)
```

### Query System

```go
// Build complex query with expressions
query := &l8api.L8Query{
    RootType: "User",
    Properties: []string{"name", "email", "status"},
    Criteria: &l8api.L8Expression{
        Condition: &l8api.L8Condition{
            Comparator: &l8api.L8Comparator{
                Left: "status",
                Oper: "==",
                Right: "active",
            },
        },
        AndOr: "AND",
        Next: &l8api.L8Expression{
            Condition: &l8api.L8Condition{
                Comparator: &l8api.L8Comparator{
                    Left: "age",
                    Oper: ">",
                    Right: "18",
                },
            },
        },
    },
    SortBy: "name",
    Descending: false,
    Limit: 100,
    Page: 1,
    MapReduce: true, // Enable distributed Map-Reduce processing (v1.4.0+)
}

// Aggregation query with GROUP BY and HAVING (v1.7.0+)
aggQuery := &l8api.L8Query{
    RootType: "Order",
    GroupBy:  []string{"department"},
    Aggregates: []*l8api.L8AggregateFunction{
        {Function: "count", Field: "*", Alias: "orderCount"},
        {Function: "sum", Field: "amount", Alias: "totalAmount"},
    },
    Having: &l8api.L8Expression{
        Condition: &l8api.L8Condition{
            Comparator: &l8api.L8Comparator{
                Left: "count", Oper: ">", Right: "10",
            },
        },
    },
}
```

### Map-Reduce Operations (v1.4.0+)

```go
// Implement IMapReduceService for distributed computation
type MyMapReduceService struct {
    cache IDistributedCache
}

func (s *MyMapReduceService) Merge(results map[string]IElements) IElements {
    // Merge results from multiple nodes
    merged := NewElements()
    for nodeId, elements := range results {
        // Custom merging logic
        merged.Combine(elements)
    }
    return merged
}

// Use Collect for distributed data filtering and aggregation
results := cache.Collect(func(item interface{}) (bool, interface{}) {
    user := item.(*User)
    if user.Age > 18 && user.Status == "active" {
        return true, user.Summary() // Include in results with transformation
    }
    return false, nil // Exclude from results
})
```

### Health Monitoring

```go
// Create health information
health := &types.Health{
    AUuid: "node-123",
    Alias: "web-server-1",
    Status: types.HealthState_Up,
    Stats: &types.HealthStats{
        CpuUsage: 45.2,
        MemoryUsage: 512*1024*1024, // 512MB
        TxMsgCount: 1500,
        RxMsgCount: 1200,
    },
    StartTime: time.Now().Unix(),
}

// Format for display (Unix top-style output)
topData := &types.Top{
    Healths: map[string]*types.Health{
        "node-123": health,
    },
}
formatted := ifs.FormatTop(topData)
fmt.Println(formatted)
```

### Type Registration and Reflection

```go
// Register custom types
registry := NewRegistry()
registry.Register(&MyCustomType{})

// Create instances dynamically
info, err := registry.Info("MyCustomType")
if err == nil {
    instance, err := info.NewInstance()
}

// Use serialization
serializer := info.Serializer(ifs.PROTOBUF)
data, err := serializer.Serialize(instance)
```

### Time-Series Database (v1.6.0+)

```go
// Add time-series data points
notifications := []*l8notify.L8TSDBNotification{
    {
        PropertyId: "cpu.usage.node-123",
        Point: &l8api.L8TimeSeriesPoint{
            Timestamp: time.Now().Unix(),
            Value:     72.5,
        },
    },
}
tsdbService.AddTSDB(notifications)

// Query time-series data for a time range
points := tsdbService.GetTSDB("cpu.usage.node-123", startTime, endTime)
for _, point := range points {
    fmt.Printf("  %d: %.2f\n", point.Timestamp, point.Value)
}
```

### Service Groups (v1.6.0+)

```go
// Multiple services share a leader election via service group
sla := ifs.NewServiceLevelAgreement(handler, "fin-ledger", 40, true, callback)
sla.SetServiceGroup("Financial")  // All Financial services share one leader

sla2 := ifs.NewServiceLevelAgreement(handler2, "fin-budget", 40, true, callback2)
sla2.SetServiceGroup("Financial")  // Same group = shared election
```

### Two-Factor Authentication (v1.5.0+)

```go
// Setup TFA for a user
qrCode, secret, err := securityProvider.TFASetup(userId, vnic)
if err != nil {
    log.Fatal("TFA setup failed:", err)
}
// Display QR code to user for authenticator app

// Verify TFA code during login
err = securityProvider.TFAVerify(userId, totpCode, sessionId, vnic)
if err != nil {
    log.Fatal("TFA verification failed:", err)
}

// Get captcha for bot protection
captchaImage := securityProvider.Captcha()

// Register new user
err = securityProvider.Register(email, password, captchaResponse, vnic)
```

### Service Callbacks (v1.5.0+)

```go
// Implement IServiceCallback for lifecycle hooks
type MyCallback struct{}

func (c *MyCallback) Before(data interface{}, action ifs.Action, isLocal bool, vnic ifs.IVNic) (interface{}, bool, error) {
    // Pre-process data before service handles it
    // Return modified data, continue flag, and error
    processedData := preProcess(data)
    shouldContinue := true  // Set to false to stop processing
    return processedData, shouldContinue, nil
}

func (c *MyCallback) After(data interface{}, action ifs.Action, isLocal bool, vnic ifs.IVNic) (interface{}, bool, error) {
    // Post-process data after service handles it
    return postProcess(data), true, nil
}

// Use in SLA
sla := ifs.NewServiceLevelAgreement(handler, "my-service", 1, true, &MyCallback{})
```

## Architecture

### Network Layers
1. **Physical Network**: TCP/UDP transport layer
2. **VNet Protocol**: Custom protocol for message routing and service discovery
3. **Service Layer**: Business logic with service areas and discovery
4. **Application Layer**: User applications and services

### Message Flow
1. **Message Creation**: Applications create messages with service targets
2. **Routing**: VNic determines optimal destination based on service topology
3. **Serialization**: Messages are serialized using configured serializers
4. **Transport**: Messages are sent via network protocols
5. **Processing**: Receiving VNic deserializes and routes to appropriate handlers

### Service Discovery
- **Service Areas**: Logical groupings for service organization
- **Health Monitoring**: Continuous health checks with statistics
- **Load Balancing**: Multiple communication patterns for optimal distribution
- **Replication**: Service replication with intelligent endpoint selection

## Security

### Encryption
- **AES Symmetric Encryption**: Built-in encryption for sensitive data
- **Secure Message Transport**: Optional message-level encryption
- **Key Management**: Secure key distribution and rotation support

### Authentication & Authorization
- **AAA Framework**: Authentication, Authorization, and Accounting interfaces
- **Resource-Based Security**: Fine-grained access control
- **Service Authentication**: Mutual authentication between services

## Development

### Adding New Protocol Types
1. Define your schema in the appropriate `.proto` file
2. Run `./make-bindings.sh` to generate Go types
3. Implement required interfaces in `go/ifs/`
4. Add comprehensive tests

### Extending VNic Functionality
1. Implement new communication patterns in VNic interface
2. Add corresponding message routing logic
3. Update service discovery mechanisms as needed

### Custom Serialization
1. Implement `ISerializer` interface
2. Register serializer with type registry
3. Configure serialization mode in resources

## Testing

The project includes comprehensive test coverage for all major components:

```bash
# Quick test with automated coverage report
cd go
./test.sh  # Runs tests and opens coverage report automatically

# Or run tests manually
go test -v -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# View coverage report in browser
open coverage.html  # or cover.html
```

### Test Coverage
- **Current Coverage**: 95% overall test coverage
- **Coverage Reports Available**:
  - `go/coverage.html` - Detailed HTML coverage report
  - `go/cover.html` - Alternative coverage visualization
  - `go/coverage.out` - Raw coverage data

### Test Coverage Areas
- **Message Operations**: Marshal/unmarshal, cloning, edge cases
- **Encryption**: AES encryption/decryption with various key sizes
- **Network Protocol**: Read/write operations, protocol handling
- **Type Conversion**: String conversion functions and type safety
- **Authentication**: Token validation, activation, and error handling
- **Service Level Agreement**: SLA configuration and activation flows
- **Registry Operations**: Dynamic type creation and registration

## Dependencies

### Go Modules
- `github.com/google/uuid` v1.6.0: UUID generation for node identification
- `google.golang.org/protobuf` v1.36.11: Protocol Buffer runtime and code generation

### Build Tools
- **Docker**: Protocol Buffer code generation via containerized protoc
- **Go 1.25.4+**: Core language runtime and build tools

## Changelog

### Version 1.7.0 (Current)
- **AI Agent Aggregation API** - `L8AggregateFunction` message, `group_by`, `having`, and `aggregates` on `L8Query`
- **Two-Phase Authentication** - `token_hash` and `code` fields on `AuthUser`/`AuthToken`, updated TFA verify flow
- **Service-to-Model Mapping** - `models` map on `L8ServiceAreas`, `AddService()` accepts optional model name
- **Concurrency Safety** - Thread-safe `MergeServices()` for parallel service activation
- **Anthropic Integration** - `ANTHROPIC_API_KEY` environment variable support for AI agent features

### Version 1.6.0
- **Time-Series Database (TSDB)** - `ITSDBService` interface, `L8TimeSeriesPoint`, `L8TSDBQuery`, and TSDB notifications
- **Data Import/Export Framework** - Templates, column mappings, 10 transform types, AI-assisted mapping, execution with error reporting
- **File Upload/Download** - Document storage with metadata, MIME types, and SHA-256 checksums
- **CSV Export** - Model-based data export with suggested filenames
- **Service Groups** - Shared leader election across related services via `SetServiceGroup()`
- **Performance Profiling** - Integrated pprof memory/CPU data in health monitoring
- **Fiscal Period Model** - `L8Period` with yearly, quarterly, and monthly period support
- **Configurable Logs Directory** - Per-deployment log path in `L8SysConfig`

### Version 1.5.0
- **Enhanced Security Framework** - TFA, Captcha, Registration, and Credential management
- **Advanced SLA Decorators** - Non-Unique Keys and Always Overwrite support
- **Local Communication** - Same-VNic messaging with `Local()` and `LocalRequest()`
- **Service Callbacks** - Continuation indication for fine-grained control
- **Web Service Refactoring** - Improved architecture with plugin and bearer validation
- **Apache 2.0 Licensing** - Proper copyright headers on all source files

### Version 1.4.0
- **Map-Reduce Framework** - Distributed computation and data aggregation
- **Leader Election** - Leader-based communication patterns and automatic failover
- **Remote VNet Support** - Cross-network connectivity and service discovery
- **Protocol Enhancements** - Improved serialization utilities for services
- **Query System Enhancement** - Added MapReduce flag for distributed queries

### Version 1.3.0
- **Service Level Agreement (SLA) Framework** - Complete refactoring of service activation
- **Dynamic Type Creation** - Added `NewOf()` method to IRegistry
- **API Improvements** - Variadic parameters for better ergonomics
- **Code Quality** - 95% test coverage with comprehensive reporting
- **Documentation** - Interactive web documentation with visualizations

### Version 1.2.0
- Enhanced authentication system with token validation
- Improved security with hash-based functions
- Comprehensive logging system
- Flexible sorting mechanisms for queries
- Secure token lifecycle management

### Version 1.1.0
- Initial release with core networking features
- Protocol Buffer schemas
- Virtual Network Interface (VNic)
- Service discovery and health monitoring

## Contributing

We welcome contributions! Please:
1. Fork the repository
2. Create a feature branch
3. Run tests with `./test.sh` to ensure 95%+ coverage
4. Submit a pull request with clear description of changes

## License

This project is licensed under the terms specified in the LICENSE file.

---

**Note**: L8Types is designed for building distributed systems with sophisticated networking requirements. The "Layer 8" concept extends beyond traditional OSI model layers to provide application-level distributed system primitives.