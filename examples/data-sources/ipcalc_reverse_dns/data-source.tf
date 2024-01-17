# IPv4
data "ipcalc_reverse_dns" "ipv4" {
  ip_address = "203.0.113.1"
}

# IPv6
data "ipcalc_reverse_dns" "ipv6" {
  ip_address = "2001:db8::567:89ab"
}
