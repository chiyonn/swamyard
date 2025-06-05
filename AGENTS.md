# 📘 Project Guide: swarmyard

This document provides guidance and structure for working with the **swarmyard** trading system project.
It explains the project layout, conventions, and collaboration standards inspired by the `AGENTS.md` format.

---

## 🏗️ Project Structure

```
swarmyard/
├── api/
│   └── proto/
│       └── tradeexecutor/
│           ├── trade_executor.proto
│           ├── tradeexecutor.pb.go
│           └── tradeexecutor_grpc.pb.go
├── apps/
│   ├── bot-sma/
│   │   ├── main.go
│   │   └── Dockerfile
│   ├── executor/
│   │   ├── main.go
│   │   └── Dockerfile
│   ├── aggregator/
│   │   ├── main.go
│   │   └── Dockerfile
│   └── dashboard/
│       ├── package.json
│       ├── vite.config.ts
│       └── src/
│           ├── main.tsx
│           ├── components/
│           └── pages/
├── pkg/
│   ├── botcore/               # Botの共通ランタイム (run loop, state)
│   │   ├── runner.go
│   │   └── types.go
│   ├── strategy/              # 戦略ごとのビジネスロジック
│   │   ├── sma.go
│   │   ├── rsi.go
│   │   └── bollinger.go
│   ├── config/                # 設定読み込み共通処理
│   │   └── config.go
│   ├── logger/                # ログラッパー
│   │   └── logger.go
│   └── model/                 # 共通構造体
│       └── model.go
├── deploy/
│   ├── docker-compose.yml
│   └── .env
├── go.mod
├── go.sum
├── README.md
└── PROJECT_GUIDE.md
```

---

## 🧭 Coding Conventions

### Go Backend Guidelines
- Use idiomatic Go with modules (`go mod`)
- Follow `pkg/` layout for reusable libraries
- Group services under `apps/`, one folder per microservice
- gRPC communication should follow the structure in `api/proto`
- Each service should include a `main.go`, `handler/`, and optionally `internal/`

### React Frontend (Vite)
- Written in TypeScript
- Component-based structure in `src/components`
- Pages in `src/pages`
- Style using Tailwind CSS
- Follow PascalCase for component files (e.g., `BotStatusCard.tsx`)

---

## 🧪 Testing

### Go
- Use `testing` package with table-driven tests
- Place test files next to source with `_test.go` suffix

### Frontend
- Use `vitest` or `jest` for unit tests
- Keep tests colocated with components
- Use `npm test` and `npm run coverage` to verify test quality

---

## 🛠️ Tooling & Linting

### Go
- Format code with `gofmt`
- Lint with `golangci-lint`

### Frontend
```bash
npm run lint        # ESLint check
npm run type-check  # TypeScript validation
npm run build       # Vite build check
```

---

## 🔄 Pull Request Standards

- Describe changes clearly
- Scope each PR to one feature or bug fix
- Add screenshots for UI changes
- Reference issues where applicable
- Ensure all tests and builds pass

---

## 📡 Communication Between Services

- gRPC is used for Bot → Executor
- REST and WebSocket for Dashboard → Aggregator
- MySQL/MariaDB for persistent storage (trades, KPIs)

---

## 🧭 Next Steps

- Define gRPC handlers for `PlaceOrder`, `PauseBot`, `ResumeBot`
- Implement bot runner and config manager
- Create initial UI layout and connect to KPI data feed
