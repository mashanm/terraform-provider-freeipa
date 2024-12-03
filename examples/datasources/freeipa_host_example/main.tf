terraform {
  required_providers {
    freeipa = {
      source = "registry.terraform.io/mashanm/freeipa"
      version = "0.0.4"
    }
  }
}

provider "freeipa" {
  host     = "duba-shp-doma01.corp.trimbletl.com"
  username = "terraform"
  password = "password"
  realm = "CORP.TRIMBLETL.COM"
}
