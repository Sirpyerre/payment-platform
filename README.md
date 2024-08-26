# Online Payment Platform

This project is an online payment platform that enables merchants to process payments, retrieve payment details, and process refunds via a simple API. The platform is built using Go and the Echo framework, with PostgreSQL as the database.

## Prerequisites

Before you begin, ensure you have the following installed on your local machine:

- **Go**: Version 1.18 or higher
- **PostgreSQL**: Version 13 or higher
- **WireMock**: As a bank simulator, follow the installation instructions from the official [WireMock website](https://wiremock.org/docs/download-and-installation/).

## Environment Variables

The application requires the following environment variables to be set:

```bash
ENVIRONMENT=
SECRET_KEY=
GO_ENV=
LOG_LEVEL=
PORT=
VERSION_API=
POSTGRES_SERVER=
POSTGRES_DATABASE=payment-platform
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_CONNECT_TIMEOUT=
POSTGRES_MAX_OPEN_CONNS=
POSTGRES_MAX_IDLE_CONNS=
POSTGRES_CONN_MAX_LIFETIME=
POSTGRES_QUERY_TIMEOUT=
PAYMENT_BANK_URL=
PAYMENT_BANK_TIMEOUT=
```

- **ENVIRONMENT**: The environment in which the application is running (e.g., development, production).
- **SECRET_KEY**: A secret key used for encryption and security.
- **GO_ENV**: Go environment (e.g., development, production).
- **LOG_LEVEL**: The log level for the application (e.g., debug, info, error).
- **PORT**: The port on which the API will run.
- **VERSION_API**: API version (e.g., v1).
- **POSTGRES_SERVER**: The PostgreSQL server address.
- **POSTGRES_DATABASE**: The name of the database (should be set to `payment-platform`).
- **POSTGRES_USER**: The PostgreSQL user.
- **POSTGRES_PASSWORD**: The PostgreSQL user's password.
- **POSTGRES_CONNECT_TIMEOUT**: Timeout in seconds for connecting to the PostgreSQL database.
- **POSTGRES_MAX_OPEN_CONNS**: Maximum number of open connections to the database.
- **POSTGRES_MAX_IDLE_CONNS**: Maximum number of idle connections in the connection pool.
- **POSTGRES_CONN_MAX_LIFETIME**: Maximum amount of time a connection may be reused.
- **POSTGRES_QUERY_TIMEOUT**: Timeout in seconds for queries to the PostgreSQL database.
- **PAYMENT_BANK_URL**: The URL of the WireMock bank simulator.
- **PAYMENT_BANK_TIMEOUT**: Timeout in seconds for requests to the bank simulator.

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-repository/online-payment-platform.git
   cd online-payment-platform
   ```

2. **Install Dependencies**:

   Make sure all dependencies are installed by running:

   ```bash
   go mod tidy
   ```

3. **Set Up PostgreSQL**:

   Create a PostgreSQL database named `payment-platform` and ensure your environment variables are correctly configured to connect to it.

   ```bash
   createdb payment-platform
   ```

4. **Run Database Migrations**:

   Apply database migrations to set up the necessary tables.

   ```bash
   go run cmd/migrate.go
   ```

5. **Set Up WireMock**:

   Follow the [WireMock installation guide](https://wiremock.org/docs/download-and-installation/) to set up the bank simulator. Ensure that the `PAYMENT_BANK_URL` environment variable points to your local WireMock instance.

## Running the Application

To start the application, run:

```bash
go run main.go
```

The application will start on the port specified by the `PORT` environment variable.

## API Endpoints

The following endpoints are available in the API:

1. **Process Payment**
    - **Method**: `POST`
    - **Endpoint**: `/api/v1/payment/process`
    - **Description**: Processes a payment through the online payment platform.

2. **Get Payment Details**
    - **Method**: `GET`
    - **Endpoint**: `/api/v1/payment/:id`
    - **Description**: Retrieves details of a previously made payment using a unique payment identifier.

3. **Process Refund**
    - **Method**: `POST`
    - **Endpoint**: `/api/v1/payment/:id/refund`
    - **Description**: Processes a refund for a specific transaction.

## Testing the API

After the application is up and running, you can use tools like Postman or cURL to interact with the API endpoints. Ensure that WireMock is running as your bank simulator to properly test payment processing and refunds.

## Contributing

If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are welcome.

## License

This project is open-source and available under the [MIT License](LICENSE).

---

This README provides clear instructions and explanations for setting up and running the project, making it easy for any developer to get started.