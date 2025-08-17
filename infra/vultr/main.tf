provider "vultr" {
  api_key = var.vultr_api_key
}

resource "vultr_instance" "geo_instance" {
  label              = "geo-instance"
  region             = var.region
  plan               = var.plan
  os_id              = var.os_id
  enable_ipv6        = true
  firewall_group_id  = var.firewall_group_id

  tags = ["ip-geolocation"]

  user_data = <<-EOF
              #!/bin/bash
              apt-get update
              apt-get install -y golang-go
              # Additional setup commands for your application
              EOF
}

resource "vultr_firewall_group" "geo_firewall" {
  firewall_group = "geo-firewall"
}

resource "vultr_firewall_rule" "allow_http" {
  firewall_group_id = vultr_firewall_group.geo_firewall.id
  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "80"
    cidr      = "0.0.0.0/0"
  }
}

resource "vultr_firewall_rule" "allow_https" {
  firewall_group_id = vultr_firewall_group.geo_firewall.id
  rule {
    direction = "in"
    protocol  = "tcp"
    port      = "443"
    cidr      = "0.0.0.0/0"
  }
}

output "instance_ip" {
  value = vultr_instance.geo_instance.main_ip
}