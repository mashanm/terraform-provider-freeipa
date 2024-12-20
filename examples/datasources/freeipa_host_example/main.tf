# Copyright (c) HashiCorp, Inc.

terraform {
  required_providers {
    freeipa = {
      source  = "registry.terraform.io/mashanm/freeipa"
      version = "0.0.8"
    }
  }
}

provider "freeipa" {
  host     = "duba-shp-doma01.corp.example.com"
  username = "terraform"
  password = "password"
  realm    = "corp.example.com"
}
