# Flipcoin Backend

This repository contains the backend service for a prediction market application, built with Go. It interacts with an Ethereum smart contract to manage prediction rounds, fetch data, and broadcast real-time updates to clients via WebSockets.

## Table of Contents

- [Features](#features)
- [How It's Built](#how-its-built)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [WebSocket Events](#websocket-events)
- [Configuration](#configuration)

## Features

-   **Smart Contract Interaction:** Communicates with an Ethereum prediction market smart contract.
-   **Round Management:** Automatically manages the lifecycle of prediction rounds (start, lock, close).
-   **Real-time Updates:** Uses WebSockets to broadcast round events and price updates to connected clients.
-   **REST API:** Provides endpoints to fetch current market price and details of specific rounds.
-   **Configuration:** Easily configurable through environment variables.

## How It's Built

This project is built using a modern Go stack, focusing on performance and real-time communication.

-   **[Go](https://go.dev/):** The core programming language used for the backend logic.
-   **[Fiber](https://gofiber.io/):** An Express.js-inspired web framework for Go, used to build the REST API. It's known for its high performance.
-   **[Go-Ethereum (Geth)](https://geth.ethereum.org/docs/developers/go-bindings):** The official Go implementation of the Ethereum protocol, used to interact with the Ethereum blockchain and the deployed smart contract.
-   **[WebSockets](https://pkg.go.dev/github.com/gofiber/contrib/websocket):** Implemented using Fiber's WebSocket library to provide real-time, bidirectional communication between the server and clients.
-   **[Godotenv](https://github.com/joho/godotenv):** A Go library to load environment variables from a `.env` file, simplifying configuration management.

## Project Structure
```
├── .gitignore
├── common
│   └── errors.go
├── config
│   ├── config.go
│   └── contract.go
├── contracts
│   └── PredictionMarket.go
├── go.mod
├── go.sum
├── main.go
├── public
│   ├── init.go
│   ├── manager.go
│   ├── routes.go
│   ├── types.go
│   └── utils.go
└── ws
    └── ws.go
```

-   **`/common`**: Contains common error types used across the application.
-   **`/config`**: Handles application and smart contract configuration, including loading from environment variables.
-   **`/contracts`**: Contains the Go bindings for the Solidity smart contract.
-   **`/public`**: Defines the public-facing REST API routes and the round management logic.
-   **`/ws`**: Implements the WebSocket server for real-time communication.
-   **`main.go`**: The entry point for the application.

## Getting Started

### Prerequisites

-   Go (version 1.23 or later)
-   An Ethereum node RPC URL
-   A deployed instance of the PredictionMarket smart contract

### Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/Rx-11/flipcoin-backend.git
    cd flipcoin-backend
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Create a `.env` file:**
    Create a `.env` file in the root of the project and add the necessary environment variables. See the [Configuration](#configuration) section for more details.

## Usage

To run the backend server, execute the following command from the root directory:

```sh
go run main.go
```

The server will start on http://localhost:3000.

## API Endpoints

-   **`GET /api/price`**
    Fetches the latest price from the oracle via the smart contract.

    **Response:**
    ```json
    {
      "price": "1234567890"
    }
    ```

-   **`GET /api/round`**
    Fetches details for a specific prediction round by its epoch number.

    **Query Parameters:**
    - `epoch` (integer): The epoch number of the round to fetch.

    **Example Request:**
    ```
    GET http://localhost:3000/api/round?epoch=100
    ```

    **Response:**
    ```json
    {
        "epoch": "100",
        "startTimestamp": "1672531200",
        "lockTimestamp": "1672531500",
        "closeTimestamp": "1672531800",
        "lockPrice": "1234500000",
        "closePrice": "0",
        "totalAmount": "5000000000000000000",
        "bullAmount": "3000000000000000000",
        "bearAmount": "2000000000000000000"
    }
    ```

## WebSocket Events

The server broadcasts the following events to all connected clients.

**WebSocket URL:** `ws://localhost:3000/ws/updates`

-   **`round_start`**: Sent when a new prediction round starts.
    ```json
    {
        "type": "round_start",
        "payload": {
            "epoch": "101",
            "transactionHash": "0xabcde12345..."
        }
    }
    ```
-   **`round_locked`**: Sent when a prediction round is locked.
    ```json
    {
        "type": "round_locked",
        "payload": {
            "epoch": "100",
            "transactionHash": "0xfghij67890..."
        }
    }
    ```
-   **`round_close`**: Sent when a prediction round is closed.
    ```json
    {
        "type": "round_close",
        "payload": {
            "epoch": "100",
            "transactionHash": "0xklmno54321..."
        }
    }
    ```

## Configuration

The following environment variables need to be set in a `.env` file in the project root:

| Variable | Description | Example |
| :--- | :--- | :--- |
| `RPC_URL` | The URL of the Ethereum node's RPC endpoint. | `https://sepolia.infura.io/v3/your-api-key` |
| `CONTRACT_ADDRESS`| The address of the deployed PredictionMarket contract. | `0x123AbC456dEf789...` |
| `ORACLE_ADDRESS` | The address of the price oracle. | `0x456DeF789aBc123...` |
| `OWNER_ADDRESS` | The address of the contract owner. | `0x789gHi123JkL456...` |
| `OPERATOR_ADDRESS`| The address of the operator. | `0xAbC123dEf456gHi...` |
| `PRIVATE_KEY_HEX` | The private key of the account used to send transactions. | `a1b2c3d4e5f6a1b2c3d4e5f6...` |
