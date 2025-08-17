# L8Types

A foundational library for Layer 8 distributed networking systems, providing Protocol Buffer-based type definitions, Go interfaces, and utilities for building distributed applications with service discovery, health monitoring, and secure communication.

## Overview

L8Types serves as the core type system and interface library for Layer 8 networking applications, offering:

- **Protocol Buffer Schemas**: Comprehensive type definitions for distributed system components
- **Virtual Network Interface (VNic)**: Advanced networking abstractions for service communication
- **Service Discovery & Management**: Built-in service registration, discovery, and area-based routing
- **Health Monitoring**: Real-time system health tracking with Unix `top`-style output formatting
- **Transaction Management**: Distributed transaction support with state tracking
- **Notification System**: Property change notifications and event propagation
- **Security Framework**: Authentication, authorization, and encryption interfaces
- **Multi-Language Support**: Go implementation with Zig bindings

## Key Features

### Virtual Networking (VNic)
- **Multiple Communication Patterns**: Unicast, Multicast, Round-robin, Proximity-based, Leader selection
- **Network Mode Support**: Native, Docker, and Kubernetes networking modes
- **Service API**: RESTful service interfaces (POST, PUT, PATCH, DELETE, GET)
- **Message Priorities**: 8-level priority system (P1-P8) for message handling
- **Transaction Support**: Distributed transaction state management

### Service Discovery & Management
- **Service Areas**: Logical service grouping and area-based routing
- **Health States**: Up, Down, Unreachable status tracking with statistics
- **Replication**: Service replication with endpoint scoring and key-based routing
- **Dynamic Service Registration**: Runtime service addition and removal

### Type System & Reflection
- **Dynamic Type Registry**: Runtime type registration and introspection
- **RNode Reflection**: Advanced reflection system with decorators and table views
- **Serialization Framework**: Pluggable serialization with multiple format support
- **Query Engine**: Expression-based querying with comparators and conditions

### Security & Configuration
- **AES Encryption**: Built-in symmetric encryption for secure communication
- **System Configuration**: Comprehensive configuration management with VNet settings
- **Authentication Framework**: AAA (Authentication, Authorization, Accounting) support
- **Access Control**: Resource-based security with permission management

## Protocol Buffer Schemas

### Core Types
- **`api.proto`**: Query expressions, conditions, comparators, and routing tables
- **`services.proto`**: Service discovery, areas, replication indices, and transactions
- **`health.proto`**: Health monitoring, statistics, and system status tracking
- **`config.proto`**: System configuration, VNet settings, and connection parameters
- **`notification.proto`**: Property change notifications and event propagation
- **`web.proto`**: Web service definitions and plugin system
- **`reflect.proto`**: Type reflection, nodes, decorators, and table views

### Message Types
- **Query**: Advanced search with criteria, sorting, pagination, and schema filtering
- **Health**: Process health with statistics (CPU, memory, message counts)
- **Services**: Service-to-area mappings and replication management
- **SysConfig**: Network configuration including VNet ports, UUIDs, and keep-alive settings
- **NotificationSet**: Batched property change notifications with sequencing

## Project Structure

```
l8types/
├── proto/                      # Protocol Buffer definitions
│   ├── api.proto              # Query system and routing
│   ├── services.proto         # Service discovery and management
│   ├── health.proto           # Health monitoring and statistics
│   ├── config.proto           # System configuration
│   ├── notification.proto     # Event notification system
│   ├── web.proto              # Web service interfaces
│   ├── reflect.proto          # Type reflection system
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
│   │   └── ...                # Additional interfaces
│   ├── types/                 # Generated Protocol Buffer types
│   ├── nets/                  # Network protocol implementation
│   ├── aes/                   # AES encryption utilities
│   ├── tests/                 # Test suite
│   └── testtypes/             # Test-specific generated types
└── zig/                       # Zig implementation (experimental)
    ├── src/                   # Zig source files
    └── build.zig              # Build configuration
```

## Getting Started

### Prerequisites
- **Go 1.23.8+**: Core implementation language
- **Docker**: Required for Protocol Buffer code generation
- **Protocol Buffers**: For schema compilation (handled via Docker)
- **Zig** (optional): For Zig language bindings

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

// Add services to configuration
ifs.AddService(config, "user-service", 1)
ifs.AddService(config, "auth-service", 2)

// Create VNic for networking
vnic := NewVNic(config, resources)
vnic.Start()
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
```

### Query System

```go
// Build complex query with expressions
query := &types.Query{
    RootType: "User",
    Properties: []string{"name", "email", "status"},
    Criteria: &types.Expression{
        Condition: &types.Condition{
            Comparator: &types.Comparator{
                Left: "status",
                Oper: "==",
                Right: "active",
            },
        },
        AndOr: "AND",
        Next: &types.Expression{
            Condition: &types.Condition{
                Comparator: &types.Comparator{
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
}
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

## Dependencies

### Go Modules
- `github.com/google/uuid`: UUID generation for node identification
- `google.golang.org/protobuf`: Protocol Buffer runtime and code generation

### Build Tools
- **Docker**: Protocol Buffer code generation via containerized protoc
- **Go 1.23.8+**: Core language runtime and build tools

## License

This project is licensed under the terms specified in the LICENSE file.

---

**Note**: L8Types is designed for building distributed systems with sophisticated networking requirements. The "Layer 8" concept extends beyond traditional OSI model layers to provide application-level distributed system primitives.