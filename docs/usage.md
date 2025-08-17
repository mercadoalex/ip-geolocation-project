# Usage Instructions for IP Geolocation Project

## Overview
This document provides instructions on how to use the IP Geolocation application, including the available API endpoints and example requests.

## API Endpoints

### 1. Get Geolocation by IP Address
- **Endpoint:** `/api/geolocation`
- **Method:** `GET`
- **Query Parameters:**
  - `ip`: The IP address for which to retrieve geolocation information.

#### Example Request
```
GET /api/geolocation?ip=8.8.8.8
```

#### Example Response
```json
{
  "ip": "8.8.8.8",
  "country": "United States",
  "region": "California",
  "city": "Mountain View",
  "latitude": 37.386,
  "longitude": -122.0838
}
```

### 2. Health Check
- **Endpoint:** `/api/health`
- **Method:** `GET`

#### Example Request
```
GET /api/health
```

#### Example Response
```json
{
  "status": "healthy"
}
```

## Running the Application
1. Ensure you have Go installed on your machine.
2. Clone the repository:
   ```
   git clone https://github.com/yourusername/ip-geolocation-project.git
   ```
3. Navigate to the project directory:
   ```
   cd ip-geolocation-project
   ```
4. Install dependencies:
   ```
   go mod tidy
   ```
5. Run the application:
   ```
   go run src/main.go
   ```

## Testing the API
You can use tools like `curl` or Postman to test the API endpoints. Make sure the application is running before making requests.

## Conclusion
This document serves as a guide to using the IP Geolocation application. For further details on the architecture and the GeoLite2 database, please refer to the other documentation files in the `docs` directory.