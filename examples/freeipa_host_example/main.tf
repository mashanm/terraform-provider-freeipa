terraform {
  required_providers {
    freeipa = {
      source = "hashicorp.com/mashanm/freeipa"
    }
  }
}

provider "freeipa" {
  host     = "duba-shp-doma01.corp.trimbletl.com"
  username = "terraform"
  password = "password"
  realm = "CORP.TRIMBLETL.COM"
}
