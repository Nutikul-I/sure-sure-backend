# Payso-Check-Slip-Dashboard-Backend

The Payso-Check-Slip-Dashboard-Backend is a backend service designed to manage and process check slip transactions for the Payso platform. This service provides various functionalities, including handling merchant registrations, transaction overviews, and payment processing. It is built using Go and leverages SQL databases for data storage and retrieval.

## Key Features:
1. Merchant Management: Provides APIs to manage merchant registrations and retrieve merchant statistics.
2. Transaction Overview: Offers detailed transaction statistics and overviews, allowing for comprehensive reporting and analysis.
3. Payment Processing: Handles the creation, updating, and retrieval of payment postbacks, ensuring smooth transaction workflows.
4. Admin Dashboard: Implements use cases for the admin dashboard, including quota and merchant usage comparisons, and monthly transaction summaries.

## Project Structure:
1. internal/model/admin_dashboard: Contains the business logic and repository interfaces for the admin dashboard functionalities.
2. internal/model/payment: Includes the payment processing logic and repository interfaces.
3. pkg/util: Utility functions and helpers used across the project.

4. router: Defines the HTTP routes and handlers for the service.

## Environment Configuration:
The project uses environment variables defined in the .env file to configure database connections and logging settings. Key environment variables include:

1. DEV_DB_HOST, DEV_DB_PORT, DEV_DB_USER, DEV_DB_PASSWORD, DEV_DB_NAME: Database connection details.
2. LOG_FILE_PATH, LOG_LEVEL: Logging configuration.
3. JWT_SECRET_TOKEN: Secret token for JWT authentication.

## Docker and Docker Compose:
The project includes a Dockerfile and a docker-compose.yml file to facilitate containerization and orchestration. These files define the necessary steps to build and run the service in a Docker container, making it easy to deploy and manage.

## Getting Started:
To get started with the Payso-Check-Slip-Dashboard-Backend, clone the repository and follow the instructions in the README file to set up the development environment, configure the necessary environment variables, and run the service using Docker Compose.