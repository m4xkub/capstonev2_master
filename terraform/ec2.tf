resource "aws_instance" "disk" {
  count                       = var.enable_cluster_1 ? 3 : 0
  ami                         = var.ec2Ami
  instance_type               = var.instanceType
  key_name                    = var.keyName
  subnet_id                   = aws_subnet.private-subnet.id
  vpc_security_group_ids      = [aws_security_group.security-group.id]
  associate_public_ip_address = true

  tags = {
    Name = "disk-${count.index + 1}"
  }
}

resource "aws_instance" "disk-migrate" {
  count                       = var.enable_cluster_2 ? 3 : 0
  ami                         = var.ec2Ami
  instance_type               = var.instanceType
  key_name                    = var.keyName
  subnet_id                   = aws_subnet.private-subnet.id
  vpc_security_group_ids      = [aws_security_group.security-group.id]
  associate_public_ip_address = true

  tags = {
    Name = "disk-migrate-${count.index + 1}"
  }
}

