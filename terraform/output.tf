# Public IPs for disk
output "disk_public_ips" {
  value = [for i in aws_instance.disk : i.public_ip]
}

# Private IPs for disk
output "disk_private_ips" {
  value = [for i in aws_instance.disk : i.private_ip]
}

# Public IPs for disk-migrate
output "disk_migrate_public_ips" {
  value = [for i in aws_instance.disk-migrate : i.public_ip]
}

# Private IPs for disk-migrate
output "disk_migrate_private_ips" {
  value = [for i in aws_instance.disk-migrate : i.private_ip]
}
