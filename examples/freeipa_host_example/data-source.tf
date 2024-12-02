data "freeipa_host" "example" {
  fqdn = "duba-nfws-otfa01.corp.trimbletl.com"
}

output "fqdn" {
  value = data.freeipa_host.example
}
