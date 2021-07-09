terraform {
  required_providers {
    sbercloud = {
      source = "sbercloud-terraform/sbercloud"
      version = "1.3.0"
    }
  
  }
}


# Configure the SberCloud Provider
provider "sbercloud" {
  region     = "ru-moscow-1"
  access_key = "BWB8XOHMVCAPPIOPGDJP"
  secret_key = "rnnszkcwmwwKZ81GvMCVV3XVLIU1nolSH3O6wWbk"
  project_name = "ru-moscow-1_nkokunina"
}
