# ⚡ RunicNexus

RunicNexus is a highly scalable, distributed real-time message routing framework designed for Massive Multiplayer Online Role-Playing Games (MMORPGs) and heavy real-time applications. 

It acts as a **Stateless Gateway and Event Fabric**. It holds persistent player connections at the edge, normalizes network traffic, and routes high-frequency data (like physics and spatial movement) to external C++ simulation workers, while handling transactional data (like inventory and chat) via native Go plugins.

## 🏗️ Core Architecture

RunicNexus operates on a strict separation of concerns, acting as the "Traffic Cop" of your game architecture:

1. **The Edge Gateway (Golang):** Handles massive concurrent TCP/WebSocket connections. It is completely stateless, meaning it can be scaled horizontally infinitely behind a Load Balancer.
2. **The Nervous System (NATS):** An internal, high-speed Pub/Sub broker that handles server clustering, sharding, and real-time state broadcasting across all RunicNexus nodes.
3. **The Simulation Workers (C++ / gRPC):** CPU-heavy tasks (like spatial chunk rendering, collision, and physics) are offloaded to dedicated C++ workers via persistent gRPC streams.
4. **The Modules (Go Plugins):** Game-specific business logic (Quests, Inventory, Chat) is implemented by the user through a standardized `runic.Module` interface.

## 📂 Project Structure (Monorepo)

To separate the framework's internal mechanics from the user's game implementation, RunicNexus follows a strict directory layout:

```text
runic-nexus/
├── cmd/
│   └── test_server/           # User implementation (Entry point for testing the framework)
├── proto/                     # Language-agnostic contracts (Protobuf definitions)
├── internal/                  # The Engine Room (Private framework mechanics)
│   ├── broker/                # NATS connection and clustering logic
│   ├── network/               # Raw TCP/WebSocket listeners
│   └── protocol/              # Protobuf serialization and buffer management
├── pkg/                       # The Public API (What users import to build their game)
│   ├── runic/                 # The core Engine and Module interfaces
│   └── modules/               # Built-in framework modules (e.g., grpcproxy for C++)
├── deploy/                    # Kubernetes manifests and Dockerfiles (K3s/Helm)
```

## 🚀 Getting Started (Local Development)
*(Note: The framework is currently in active prototyping).*
**Prerequisites:**
* Go 1.21+
* Docker & Docker Compose (for the local NATS cluster)
* Protocol Buffers Compiler (`protoc`)

**1. Boot the NATS Cluster**
```bash
docker run -p 4222:4222 -p 8222:8222 nats:latest -m 8222
```

**2. Run the Test Server**
```bash
go run ./cmd/test_server/main.go
```

## 🛠️ Roadmap
- [x] Architectural Blueprint and Monorepo Setup
- [ ] Implement Protobuf `Envelope` definition (`proto/`)
- [ ] Build NATS Broker integration (`internal/broker/`)
- [ ] Create the Core Module Registry (`pkg/runic/`)
- [ ] Implement Raw TCP/WebSocket Listeners (`internal/network/`)
- [ ] Build the C++ gRPC Bridge Module (`pkg/modules/grpcproxy/`)
- [ ] Kubernetes Helm Charts for Oracle Cloud/AWS deployment
