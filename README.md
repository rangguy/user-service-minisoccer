<h3>User Service</h3>

## 🧩 Related Microservices

The User Service is part of the **MiniSoccer Platform**, which is built using a **microservices architecture**. Each service is responsible for a specific domain while communicating with others through APIs, message brokers, and payment gateway.

Below are the other related microservices:

---

### 🔗 Field Service (field-service-minisoccer)
**Repository:** https://github.com/rangguy/field-service-minisoccer

**Description:**
Responsible for managing all data related to mini soccer fields, including:
- Field details and metadata
- Availability schedules
- Pricing configuration
- Integration with the Order Service for slot availability checks

This service provides APIs to check field details and availability before a user creates an order.

---

### 🔗 Payment Service (payment-service-minisoccer)
**Repository:** https://github.com/rangguy/payment-service-minisoccer

**Description:**
Handles all payment processes within the platform, such as:
- Transaction initialization
- Integration with Midtrans
- Handling payment status callbacks/webhooks
- Storing payment history

This service updates the Order Service once a transaction is confirmed.

---

### 🔗 Order Service (order-service-minisoccer)
**Repository:** https://github.com/rangguy/order-service-minisoccer

**Description:**
Acts as the central service for processing bookings and orders, including:
- Creating new orders
- Validating field availability (via Field Service)
- Linking orders with users (via User Service)
- Sending payment instructions to the Payment Service
- Updating order status based on payment results

Order Service acts as the main coordinator between the User, Field, and Payment Services.

---


<h3>Description</h3>

<p>This repository will be used to manage user and auth</p>

<h3>Directory Structure</h3>

```
user-service
    L cmd                            → Contains the main entry point or initial configuration of the application
    L common                         → Stores common functions used throughout the application
    L config                         → Contains application configurations such as environment variables and other settings
    L constants                      → Stores global constant values used across the application
    L controllers                    → Manages control logic for handling HTTP requests
    L database                       → Contains files related to database management
        L seeders                    → Scripts for populating initial (seed) data into the database
    L domain                         → The application's domain module containing core domain elements
        L dto                        → Data Transfer Objects, used to define the structure of transferred data
        L models                     → Object models representing the application's or database's data structure
    L middlewares                    → Contains middleware for processing requests/responses before or after reaching the controller
    L repositories                   → Contains data access logic for interacting with the database
    L routes                         → Contains API route definitions
    L services                       → Stores the application's core business logic
```

## How to setup

```
- Clone this repository
- go mod tidy
- copy .env.example to .env (if you want to run with consul)
- copy .config.json.example to .config.json
```

## How to run

```bash
make watch-prepare (only for the first time or when you add new dependency)
make watch
```

## How to run with docker

```bash
docker-compose up -d --build --force-recreate
```

## How to build
```bash
make build
```
