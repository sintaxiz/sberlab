terraform {
  required_providers {
    sbercloud = {
      source = "sbercloud-terraform/sbercloud"
    }
  }
}


# Configure the SberCloud Provider
provider "sbercloud" {
  auth_url = "https://iam.ru-moscow-1.hc.sbercloud.ru/v3"
  region     = "ru-moscow-1"
  access_key = "CDFFLLBF1BBN1CGYPFAH"
  secret_key = "EfwZDgHSrGpYqjBls3AModiJQVDDpEqbuq1K78ih"
  project_name = "ru-moscow-1_nkokunina"
}
