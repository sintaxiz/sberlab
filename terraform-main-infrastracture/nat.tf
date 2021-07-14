variable "natName" {
  description = "Name of NAT"
}

variable "serverPort" {
  description = "Port on which server is listening"
}

resource "sbercloud_vpc_eip" "nat_eip" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name = "eip-for-${var.natName}"
    size = 5
    share_type = "PER"
    charge_mode = "bandwidth"
  }
}

resource "sbercloud_nat_gateway" "nat_01" {
  name = var.natName
  description = "NAT Gateway"
  spec = "1"
  vpc_id = sbercloud_vpc.vpc_01.id
  subnet_id = sbercloud_vpc_subnet.subnet_01.id
}

resource "sbercloud_nat_snat_rule" "snat_01" {
  nat_gateway_id = sbercloud_nat_gateway.nat_01.id
  subnet_id = sbercloud_vpc_subnet.subnet_01.id
  floating_ip_id = sbercloud_vpc_eip.nat_eip.id
}

resource "sbercloud_nat_dnat_rule" "dnat_01" {
  floating_ip_id = sbercloud_vpc_eip.nat_eip.id
  nat_gateway_id = sbercloud_nat_gateway.nat_01.id
  private_ip = sbercloud_compute_instance.ecs_master.access_ip_v4
  protocol = "tcp"
  internal_service_port = 22
  external_service_port = 22
}

resource "sbercloud_nat_dnat_rule" "dnat_02" {
  floating_ip_id = sbercloud_vpc_eip.nat_eip.id
  nat_gateway_id = sbercloud_nat_gateway.nat_01.id
  private_ip = sbercloud_compute_instance.ecs_master.access_ip_v4
  protocol = "tcp"
  internal_service_port = var.serverPort
  external_service_port = 80
}
resource "sbercloud_nat_dnat_rule" "dnat_for_kuberctl" {
  floating_ip_id = sbercloud_vpc_eip.nat_eip.id
  nat_gateway_id = sbercloud_nat_gateway.nat_01.id
  private_ip = sbercloud_compute_instance.ecs_master.access_ip_v4
  protocol = "tcp"
  internal_service_port = 6443
  external_service_port = 6443
}

resource "sbercloud_nat_dnat_rule" "dnat_for_kuberapi" {
  floating_ip_id = sbercloud_vpc_eip.nat_eip.id
  nat_gateway_id = sbercloud_nat_gateway.nat_01.id
  private_ip = sbercloud_compute_instance.ecs_master.access_ip_v4
  protocol = "tcp"
  internal_service_port = 11251
  external_service_port = 11251
}

resource "sbercloud_nat_dnat_rule" "dnat_for_nodeport_service" {
  floating_ip_id = sbercloud_vpc_eip.nat_eip.id
  nat_gateway_id = sbercloud_nat_gateway.nat_01.id
  private_ip = sbercloud_compute_instance.ecs_master.access_ip_v4
  protocol = "tcp"
  internal_service_port = 31234
  external_service_port = 31234
}

resource "sbercloud_nat_dnat_rule" "dnat_for_nodeport_backend" {
  floating_ip_id = sbercloud_vpc_eip.nat_eip.id
  nat_gateway_id = sbercloud_nat_gateway.nat_01.id
  private_ip = sbercloud_compute_instance.ecs_master.access_ip_v4
  protocol = "tcp"
  internal_service_port = 31333
  external_service_port = 8080
}

