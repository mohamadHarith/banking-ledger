# Banking Ledger Service

## 🚀 Introduction
A backend service built as part of an assessment for Goland Developer position at Innoscripta. The service manages bank accounts and transactions with high scalibility and reliability.


## 🛠️ Features
- Create accounts with initial balances
- Deposit and withdraw funds
- View transaction history
- Reliable and scalable transactions logging using message queues
- ACID compliance for financial operations

## 🧱 Technology Stack
- **Golang** for backend services
- **MySQL** for account balances
- **MongoDB** for transaction logs
- **RabbitMQ** for message queue
- **Docker Compose** for container orchestration

## 🏗️ Architecture Overview
![Bank System drawio](https://github.com/user-attachments/assets/fafb563c-8459-441c-8104-bc9ff9d46b54)

## 🧩 Service Explanations (`services/`):
- **api-gateway/**: Manages API endpoints and routes requests to the appropriate services.
- **accounts/**: Handles account creation, balance updates, and account management.
- **transactions/**: Processes deposits, withdrawals, and records transaction history.
- **shared/**: Contains shared utilities, configurations, and common modules.

## 📦 Installation & Usage
### Prerequisites
- Docker & Docker Compose installed

### Run the Service
```bash
docker-compose up --build
```

### REST API Documentation (`services/api-gateway`)
- **Create Account:** `POST /accounts`
  Request: `{ "name": "John Doe", "initial_balance": 1000 }`
  Response: `{ "account_id": "12345" }`

- **Deposit Funds:** `POST /transactions/deposit`
  Request: `{ "account_id": "12345", "amount": 500 }`
  Response: `{ "balance": 1500 }`

- **Withdraw Funds:** `POST /transactions/withdraw`
  Request: `{ "account_id": "12345", "amount": 200 }`
  Response: `{ "balance": 1300 }`

- **Transaction History:** `GET /transactions/{account_id}`
  Response: `[{"type": "deposit", "amount": 500}, {"type": "withdraw", "amount": 200}]`

## 🧪 Testing
```bash
go test ./...
```

## 📂 Project Structure
```
├── docker-compose.yml
├── services/
│   ├── api-gateway/    (API routes and request handling)
│   ├── accounts/       (Account management logic)
│   ├── transactions/   (Transaction processing and logging)
├── shared/             (Common utilities and configurations)
└── README.md
```

## ✅ Improvements
- Add rate limiting to prevent abuse
- Implement retry mechanisms in the queue
- Secure endpoints with authentication tokens

## 📝 Best Practices Followed
- **SOLID Principles:** Proper separation of concerns and dependency injection.
- **DRY and KISS:** Simple, reusable functions with minimal redundancy.
- **Clean Architecture:** Clear separation of business logic and framework dependencies.
