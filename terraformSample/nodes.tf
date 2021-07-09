# Create nodes
data "sbercloud_availability_zones" "myaz" {}

data "sbercloud_compute_flavors" "myflavor" {
  availability_zone = data.sbercloud_availability_zones.myaz.names[0]
  performance_type  = "normal"
  cpu_core_count    = 2
  memory_size       = 4
}

data "sbercloud_images_image" "ubuntu" {
  name        = "Ubuntu 18.04 server 64bit"
  most_recent = true
}


# Node 1
resource "sbercloud_compute_instance" "tf-test-node1" {
  name              = "tf-test-node1"
  image_id          = data.sbercloud_images_image.ubuntu.id
  flavor_id         = data.sbercloud_compute_flavors.myflavor.ids[0]
  security_groups   = ["default"]
  availability_zone = "ru-moscow-1a"
  system_disk_type = "SAS"
  user_data        = "#!/bin/bash\necho \nHello, the time is now $(date -R)\n | tee /root/output.txt"
  key_pair          = "KeyPair-default"
  network {
    uuid = sbercloud_vpc_subnet.subnet_v1.id
  }
}

# Node 2
resource "sbercloud_compute_instance" "tf-test-node2" {
  name              = "tf-test-node2"
  image_id          = data.sbercloud_images_image.ubuntu.id
  flavor_id         = data.sbercloud_compute_flavors.myflavor.ids[0]
  security_groups   = ["default"]
  availability_zone = "ru-moscow-1a"
  system_disk_type = "SAS"
  key_pair          = "KeyPair-default"
  user_data         = file("${path.module}/init.sh")

  network {
    uuid = sbercloud_vpc_subnet.subnet_v1.id
  }
}

