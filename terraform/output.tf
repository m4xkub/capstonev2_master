output "disk_public_ips" {
  value = [for i in aws_instance.disk : i.public_ip]
}

output "disk_migrate_public_ips" {
  value = [for i in aws_instance.disk-migrate : i.public_ip]
}
