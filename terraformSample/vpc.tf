# variable subnet {
#   type        = string
#   default     = "trf-test-subnet"
#   description = "name of subnet"
# }


# Create a VPC
resource "sbercloud_vpc" "vpc_v1" {
  name = "tf-test-vpc"
  cidr = "192.168.0.0/16"
}


# Create subnet for a VPC
resource "sbercloud_vpc_subnet" "subnet_v1" {
  name       = "tf-test-subnet"
  cidr       = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"
  vpc_id     = sbercloud_vpc.vpc_v1.id
}
