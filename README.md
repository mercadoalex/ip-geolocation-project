# IP Geolocation Project

This project is designed to provide IP geolocation services using the MaxMind GeoLite2 database. It leverages the power of the GeoLite2 databases to accurately determine the geographical location of IP addresses. The application is built in Go and is structured to facilitate easy deployment and usage.

## Project Structure

```
ip-geolocation-project
├── src
│   ├── main.go          # Entry point of the application
│   ├── geo
│   │   └── geolite2.go  # Functions for interacting with the GeoLite2 database
│   └── utils
│       └── logger.go    # Logging utilities for the application
├── maxmind
│   └── GeoLite2-City.mmdb # MaxMind GeoLite2 City database file
├── infra
│   ├── vultr
│   │   └── main.tf      # Terraform configuration for Vultr cloud deployment
│   └── README.md        # Documentation for infrastructure setup
├── docs
│   ├── architecture.md   # Application architecture details
│   ├── usage.md          # Instructions on how to use the application
│   └── geolite2.md       # Explanation of the GeoLite2 database
├── README.md             # Overview of the project
├── go.mod                # Go module definition file
└── go.sum                # Checksums for module dependencies
```

## Getting Started

To get started with the IP Geolocation Project, follow these steps:

1. **Clone the Repository**: 
   ```
   git clone https://github.com/yourusername/ip-geolocation-project.git
   cd ip-geolocation-project
   ```

2. **Install Dependencies**: 
   Ensure you have Go installed, then run:
   ```
   go mod tidy
   ```

3. **Set Up the Database**: 
   Download the GeoLite2 City database from MaxMind and place it in the `maxmind` directory.

4. **Run the Application**: 
   Start the application by running:
   ```
   go run src/main.go
   ```

5. **Deploying with Terraform**: 
   Navigate to the `infra/vultr` directory and follow the instructions in `README.md` to deploy the application on Vultr cloud infrastructure.

## Documentation

- For detailed architecture, refer to `docs/architecture.md`.
- For usage instructions, see `docs/usage.md`.
- To understand the GeoLite2 database and its usage, check `docs/geolite2.md`.
- Infrastructure setup and deployment instructions can be found in `infra/README.md`.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author
Alejandro Mercado Peña