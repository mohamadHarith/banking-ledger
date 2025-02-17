# Banking Ledger Service

## 🚀 Introduction
A backend service built as part of an assessment for Goland Developer position at Innoscripta. The service manages bank accounts and transactions with high scalibility and reliability.


## 🛠️ Features
- Create accounts with initial balances
- Deposit and withdraw funds
- Retrieval of detailed transaction logs/ ledger
- Reliable and scalable transactions logging using message queues
- ACID compliance for financial operations
- Horizontally scalled service
- Secure APIs using JWT authentication

## 🧱 Technology Stack
- **Golang** for backend services
- **MySQL** for account balances
- **MongoDB** for transaction logs
- **RabbitMQ** for message queue
- **Docker Compose** for container orchestration
- **gRPC** for RPC calls 

## 🏗️ Architecture Overview

![Bank System-Page-2 drawio (2)](https://github.com/user-attachments/assets/c6f4c440-332a-41a3-ac12-7ac9ad4ffc91)

This architecture diagram represents a banking ledger system with multiple microservices communicating via RabbitMQ and an API Gateway. Below is an overview of the components:

**API Gateway**:
- Acts as the single entry point for client requests.
- Subscribes to `transaction_balance_queue` from RabbitMQ.
- Caches account balances in Redis for fast retrieval.
- Load balances the request to `transaction-processor-1` and `transaction-processor-2`

**Authentication Service**: 
- Creates user accounts.
- Manages user authentication.

**Transaction Processor Services (1 and 2)**:
- Process financial transactions.
- Publish transaction logs (`transaction.log`) and account balances (`transaction.balance`) to RabbitMQ via `transaction_exchange`.
- Store transaction results in MySQL.

**Transaction Logger Service**:  
- Subscribes to `transaction_log_queue` from RabbitMQ.  
- Stores transaction logs in MongoDB.  

**Redis**: Caches account balances received from the API Gateway to reduce database load. 

**RabbitMQ (Message Broker)**: Handles asynchronous communication.  
 


## 🧩 Service Explanations (`services/`):
| Service Name                       | Description                                         | Port |
|------------------------------------|-----------------------------------------------------|------|
| **api-gateway/**                   | Manages API endpoints and routes requests           | 5001 |
| **authentication-service/**        | Auhtenticates user                                  | 5002 |
| **transaction-logger-service/**    | Stores and receives transaction logs (ledger)       | 5003 |
| **transaction-processor-service/** | Processes account creation deposits,and withdrawals | 5004 |


## 📦 Installation & Usage
### Prerequisites
- Docker & Docker Compose installed

### Run the Service
```bash
docker-compose up --build
```

Note: Normally we do not commit `.env` file to the repository. It should be confidential and should be handled using tools like secrets manager. However, for the sake of simiplicity for this task, we have commited the environment variables file in this repo.

### REST APIs
- **Create User:**
To create user account
Sample request:
```console
curl --location 'localhost:5001/user' \
--header 'Content-Type: application/json' \
--data '{
    "fullName":"Mohamad Harith Bin Habib Rahman",
    "userName":"harith97",
    "password":"123456$"
}'
```
Sample response:
```json
{
    "errorCode": 200,
    "message": "success"
}
```

- **Login:**
To login and obtain access token
Sample request:
```console
curl --location 'localhost:5001/login' \
--header 'Content-Type: application/json' \
--data '{
    "userName":"harith97",
    "password":"123456$"
}'
```

Sample response:
```json
{
    "errorCode": 200,
    "item": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3ODAxODUsImlhdCI6MTczOTc4MDE3MCwic3ViIjoiYTM3YzczYzAtNTI4Mi00OTk1LTgxZDQtYTQ4ODNiMmUzYzU5In0.kqJDTMluX3-cBXGnVs4Vc83rRnxFd1HBU7vCBJ7HDIs"
    },
    "message": "success"
}
```

- **Account:**
To create ledger account with initial balance  
Sample request:
```console
curl --location 'localhost:5001/account' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3ODI5NjYsImlhdCI6MTczOTc4MjA2Niwic3ViIjoiYTM3YzczYzAtNTI4Mi00OTk1LTgxZDQtYTQ4ODNiMmUzYzU5In0.NxOKLY7_Ox_IQxxAUKcauhKiK2bCxLfASdFtJbWG5tg' \
--data '{
    "initialBalance":100
}'
```

Sample response:
```json
{
    "errorCode": 200,
    "item": {
        "Id": "011af361-298a-4623-8251-9f5072f0becf",
        "UserId": "a37c73c0-5282-4995-81d4-a4883b2e3c59",
        "Balance": 100,
        "CreatedAt": "2025-02-17T08:53:21.389983Z",
        "UpdatedAt": "2025-02-17T08:53:21.389983Z"
    },
    "message": "success"
}
```

- **Deposit:**
To deposit funds to an account  
Sample request:
```console
curl --location 'localhost:5001/deposit' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3ODM3MjAsImlhdCI6MTczOTc4MjgyMCwic3ViIjoiYTM3YzczYzAtNTI4Mi00OTk1LTgxZDQtYTQ4ODNiMmUzYzU5In0.P0lKKQdFn9KlpD68r8_6ofQyIkXajrHiDudSnTXiBZ8' \
--data '{
    "accountId":"011af361-298a-4623-8251-9f5072f0becf",
    "amount":50,
    "description":"Savings"
}'
```

Sample response:
```
{
    "errorCode": 200,
    "message": "success"
}
```

- **Withdraw:**
To withdraw funds from an account  
Sample request:
```console
curl --location 'localhost:5001/withdraw' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3ODM3MjAsImlhdCI6MTczOTc4MjgyMCwic3ViIjoiYTM3YzczYzAtNTI4Mi00OTk1LTgxZDQtYTQ4ODNiMmUzYzU5In0.P0lKKQdFn9KlpD68r8_6ofQyIkXajrHiDudSnTXiBZ8' \
--data '{
    "accountId":"011af361-298a-4623-8251-9f5072f0becf",
    "amount":20,
    "description":"Foodpanda"
}'
```

Sample response:
```json
{
    "errorCode": 200,
    "message": "success"
}
```

- **Balance:**
To get balance for an account  
Sample request:
```console
curl --location --request GET 'localhost:5001/balance' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3ODM3MjAsImlhdCI6MTczOTc4MjgyMCwic3ViIjoiYTM3YzczYzAtNTI4Mi00OTk1LTgxZDQtYTQ4ODNiMmUzYzU5In0.P0lKKQdFn9KlpD68r8_6ofQyIkXajrHiDudSnTXiBZ8' \
--data '{
    "accountId":"011af361-298a-4623-8251-9f5072f0becf"
}'
```

Sample response:
```json
{
    "errorCode": 200,
    "item": {
        "balance": 51
    },
    "message": "success"
}
```

- **Transactions:**
To get transaction logs/ ledger for an account  
Sample request:  
```console
curl --location 'localhost:5001/transactions' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3ODcwODMsImlhdCI6MTczOTc4NjE4Mywic3ViIjoiYTM3YzczYzAtNTI4Mi00OTk1LTgxZDQtYTQ4ODNiMmUzYzU5In0.LL7-1ip2pZ6XtrbDZ3qOC8ELJpEAtyO8Uz8PUAs7YOw' \
--data '{
    "accountId":"011af361-298a-4623-8251-9f5072f0becf",
    "page":1
}'
```

Sample response:  
```json
{
    "errorCode": 200,
    "item": [
        {
            "id": "e3021802-5149-4c59-905c-8866ecc00921",
            "accountId": "011af361-298a-4623-8251-9f5072f0becf",
            "userId": "a37c73c0-5282-4995-81d4-a4883b2e3c59",
            "amount": 20,
            "balance": 51,
            "description": "Savings",
            "createdAt": "2025-02-17T09:08:59.554Z"
        },
        {
            "id": "e5772b0b-5970-45e2-a4f6-8161ef16352a",
            "accountId": "011af361-298a-4623-8251-9f5072f0becf",
            "userId": "a37c73c0-5282-4995-81d4-a4883b2e3c59",
            "amount": -20,
            "balance": 31,
            "description": "Foodpanda",
            "createdAt": "2025-02-17T09:04:54.481Z"
        },
        {
            "id": "6fef4586-8684-47be-86a7-b5405ad92c9d",
            "accountId": "011af361-298a-4623-8251-9f5072f0becf",
            "userId": "a37c73c0-5282-4995-81d4-a4883b2e3c59",
            "amount": 50,
            "balance": 51,
            "description": "Savings",
            "createdAt": "2025-02-17T09:00:41.733Z"
        },
        {
            "id": "62b48e1f-ea86-45a2-8a89-b9ce9ca9570c",
            "accountId": "011af361-298a-4623-8251-9f5072f0becf",
            "userId": "a37c73c0-5282-4995-81d4-a4883b2e3c59",
            "amount": 1,
            "balance": 1,
            "description": "Account Creation Initial Balance",
            "createdAt": "2025-02-17T08:53:21.389Z"
        }
    ],
    "message": "success",
    "pagination": {
        "totalRecords": 4,
        "currentPage": 1,
        "totalPages": 1,
        "nextPage": 0,
        "prevPage": 0
    }
}
```

## ✅ Improvements
- Add rate limiting to prevent abuse
- Add distributed tracing and logging
- Add more robust tests

## 🏁 Conclusion
This project successfully meets the requirements of the assessment by showcasing:
- **Technical Proficiency:** Use of Golang, MySQL, MongoDB, RabbitMQ, and Docker.
- **Scalability and Reliability:** Through message queues and microservice architecture.
- **Code Quality:** Adherence to best practices and clean architecture principles.

