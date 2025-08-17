# Architecture Overview

## Introduction
This document outlines the architecture of the IP Geolocation Project, which utilizes the MaxMind GeoLite2 database for IP geolocation lookups. The application is designed to be deployed on Vultr cloud infrastructure, leveraging Terraform for resource management.

## Components
The architecture consists of the following key components:

1. **Application Server**
   - The entry point of the application is located in `src/main.go`. This file initializes the server and sets up the necessary routes for handling incoming IP geolocation requests.

2. **GeoLite2 Database Interaction**
   - The `src/geo/geolite2.go` file contains functions that interact with the GeoLite2 database. It provides methods to query the database and retrieve geolocation information based on IP addresses.

3. **Logging Utilities**
   - The `src/utils/logger.go` file provides logging functionalities. It exports functions to log messages at various levels (info, warning, error), which aids in debugging and monitoring the application.

4. **MaxMind GeoLite2 Database**
   - The `maxmind/GeoLite2-City.mmdb` file is the MaxMind GeoLite2 City database used for IP geolocation lookups. This database contains the geolocation data that the application queries.

5. **Infrastructure as Code**
   - The `infra/vultr/main.tf` file contains the Terraform configuration for deploying the application on Vultr cloud infrastructure. It defines the necessary resources such as virtual machines, networking, and configurations required for deployment.

6. **Documentation**
   - The `docs` directory contains various documentation files:
     - `architecture.md`: This file outlines the architecture of the application.
     - `usage.md`: Provides instructions on how to use the application, including API endpoints and example requests.
     - `geolite2.md`: Explains the GeoLite2 database, its structure, and how to effectively use it within the application.

## Data Flow
1. **Client Request**: The client sends an IP geolocation request to the application server.
2. **Route Handling**: The server processes the request and invokes the appropriate handler defined in `main.go`.
3. **Database Query**: The handler calls functions in `geolite2.go` to query the GeoLite2 database using the provided IP address.
4. **Response Generation**: The application retrieves the geolocation data and formats it into a response.
5. **Logging**: Throughout the process, relevant events and errors are logged using the utilities defined in `logger.go`.
6. **Client Response**: The server sends the geolocation data back to the client.

## Deployment
The application is deployed using Terraform scripts defined in `infra/vultr/main.tf`. This setup ensures that the application is scalable and can be managed efficiently on the Vultr cloud infrastructure.

## Conclusion
This architecture provides a robust framework for IP geolocation services, leveraging the power of the GeoLite2 database and the flexibility of cloud infrastructure. The modular design allows for easy maintenance and scalability as the project evolves.