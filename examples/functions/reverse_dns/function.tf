output "reverse_ipv4" {
  value = provider::ipcalc::reverse_dns("1.2.3.4")
}

output "reverse_ipv6" {
  value = provider::ipcalc::reverse_dns("2001:db8::567:89ab")
}
