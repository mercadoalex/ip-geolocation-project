# Infrastructure Setup for IP Geolocation Project

This README provides instructions for setting up the infrastructure required to deploy the IP Geolocation project using Terraform on Vultr cloud.

## Prerequisites

Before you begin, ensure you have the following:

- A Vultr account with API access.
- Terraform installed on your local machine.
- Basic knowledge of Terraform and cloud infrastructure concepts.

## Getting Started

1. **Clone the Repository**

   Start by cloning the repository to your local machine:

   ```
   git clone https://github.com/yourusername/ip-geolocation-project.git
   cd ip-geolocation-project/infra/vultr
   ```

2. **Configure Terraform**

   Create a file named `terraform.tfvars` in the `infra/vultr` directory to store your Vultr API key and any other necessary variables:

   ```
   vultr_api_key = "your_vultr_api_key"
   region = "ewr"  # Example: New Jersey
   plan = "vc2-1c-1gb"  # Example: 1 CPU, 1 GB RAM
   ```

3. **Initialize Terraform**

   Run the following command to initialize Terraform and download the necessary providers:

   ```
   terraform init
   ```

4. **Plan the Deployment**

   Use the following command to see what resources will be created:

   ```
   terraform plan
   ```

5. **Apply the Configuration**

   Deploy the infrastructure by running:

   ```
   terraform apply
   ```

   Confirm the action when prompted. Terraform will create the specified resources on Vultr.

## Managing Infrastructure

- To update your infrastructure, modify the `main.tf` file as needed and run `terraform apply` again.
- To destroy the infrastructure, use the command:

   ```
   terraform destroy
   ```

   Confirm the action when prompted.

## Additional Resources

For more information on using Terraform with Vultr, refer to the [Vultr Terraform Provider Documentation](https://registry.terraform.io/providers/vultr/vultr/latest/docs).

For details on the application itself, please refer to the documentation in the `docs` directory.

## Conclusion

This README serves as a guide to set up the infrastructure for the IP Geolocation project. Ensure you follow the steps carefully and modify configurations as per your requirements. Happy coding!