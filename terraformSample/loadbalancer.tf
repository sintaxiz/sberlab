# # Load balancer
# resource "sbercloud_lb_loadbalancer" "lb_1" {
#   name = "tf-test-lb"
#   vip_subnet_id = sbercloud_vpc_subnet.subnet_v1.subnet_id
# }

# resource "sbercloud_vpc_eip" "eip_1" {
#   publicip {
#     type = "5_bgp"
#   }
#   bandwidth {
#     name        = "tf-test-ip"
#     size        = 8
#     share_type  = "PER"
#     charge_mode = "traffic"
#   }
# }

# resource "sbercloud_networking_eip_associate" "eip_1" {
#   public_ip = sbercloud_vpc_eip.eip_1.publicip[0].ip_address
#   port_id   = sbercloud_lb_loadbalancer.lb_1.vip_port_id
# }

# # Add listener
# resource "sbercloud_lb_listener" "listener_1" {
#   protocol        = "TCP"
#   protocol_port   = 80
#   loadbalancer_id = sbercloud_lb_loadbalancer.lb_1.id
# }

# # Add backend server group
# # Manages an ELB pool resource within SberCloud.
# resource "sbercloud_lb_pool" "pool_1" {
#   protocol    = "TCP"
#   lb_method   = "ROUND_ROBIN"
#   listener_id = sbercloud_lb_listener.listener_1.id
# }

# resource "sbercloud_lb_member" "member_1" {
#   address       = sbercloud_compute_instance.tf-test-node1.access_ip_v4
#   protocol_port = 8080
#   pool_id       = sbercloud_lb_pool.pool_1.id
#   subnet_id     = sbercloud_vpc_subnet.subnet_v1.subnet_id
# }

# resource "sbercloud_lb_member" "member_2" {
#   address       = sbercloud_compute_instance.tf-test-node2.access_ip_v4
#   protocol_port = 8080
#   pool_id       = sbercloud_lb_pool.pool_1.id
#   subnet_id     = sbercloud_vpc_subnet.subnet_v1.subnet_id
# }