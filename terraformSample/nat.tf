resource "sbercloud_vpc_eip" "eipNAT" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "tf-test-ipForNat"
    size        = 8
    share_type  = "PER"
    charge_mode = "traffic"
  }
}


resource "sbercloud_nat_gateway" "nat_1" {
  name                = "tf-test-nat"
  description         = "test for terraform"
  spec                = "1"
  vpc_id              = sbercloud_vpc.vpc_v1.id
  subnet_id           = sbercloud_vpc_subnet.subnet_v1.id
}

resource "sbercloud_nat_dnat_rule" "dnat_1" {
  floating_ip_id        = sbercloud_vpc_eip.eipNAT.id
  nat_gateway_id        = sbercloud_nat_gateway.nat_1.id
  private_ip            = sbercloud_compute_instance.tf-test-node1.access_ip_v4
  protocol              = "tcp"
  internal_service_port = 22
  external_service_port = 2020
}


resource "sbercloud_nat_dnat_rule" "dnat_2" {
  floating_ip_id        = sbercloud_vpc_eip.eipNAT.id
  nat_gateway_id        = sbercloud_nat_gateway.nat_1.id
  private_ip            = sbercloud_compute_instance.tf-test-node2.access_ip_v4
  protocol              = "tcp"
  internal_service_port = 22
  external_service_port = 2021
}

# resource "sbercloud_nat_snat_rule" "snat_1" {
#   nat_gateway_id = sbercloud_nat_gateway.nat_1.id
#   subnet_id     = sbercloud_vpc_subnet.subnet_v1.subnet_id
#   floating_ip_id = "0a166fc5-a904-42fb-b1ef-cf18afeeddca"
# }
