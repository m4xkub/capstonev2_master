resource "aws_instance" "disk" {
  count                       = var.enable_cluster_1 ? 2 : 0
  ami                         = var.ec2Ami
  instance_type               = var.instanceType
  key_name                    = var.keyName
  subnet_id                   = aws_subnet.private-subnet.id
  vpc_security_group_ids      = [aws_security_group.security-group.id]
  associate_public_ip_address = true

  user_data = <<-EOF
    #!/bin/bash
    sudo apt update && sudo apt upgrade -y
    sudo apt install golang-go -y

    echo "postfix postfix/main_mailer_type select No configuration" | sudo debconf-set-selections
    sudo DEBIAN_FRONTEND=noninteractive apt install -y drbd-utils postfix

    git clone https://github.com/m4xkub/capstonev2_worker.git /home/ubuntu/capstonev2_worker
    chown -R ubuntu:ubuntu /home/ubuntu/capstonev2_worker


    sudo -u ubuntu bash -c "cd /home/ubuntu/capstonev2_worker && nohup go run main.go &"

  EOF

  tags = {
    Name = "disk-${count.index + 1}"
  }
}

resource "aws_instance" "disk-migrate" {
  count                       = var.enable_cluster_2 ? 2 : 0
  ami                         = var.ec2Ami
  instance_type               = var.instanceType
  key_name                    = var.keyName
  subnet_id                   = aws_subnet.private-subnet.id
  vpc_security_group_ids      = [aws_security_group.security-group.id]
  associate_public_ip_address = true
  user_data                   = <<-EOF
    #!/bin/bash
    sudo apt update && sudo apt upgrade -y
    sudo apt install golang-go -y

    echo "postfix postfix/main_mailer_type select No configuration" | sudo debconf-set-selections
    sudo DEBIAN_FRONTEND=noninteractive apt install -y drbd-utils postfix

    git clone https://github.com/m4xkub/capstonev2_worker.git /home/ubuntu/capstonev2_worker
    chown -R ubuntu:ubuntu /home/ubuntu/capstonev2_worker


    sudo -u ubuntu bash -c "cd /home/ubuntu/capstonev2_worker && nohup go run main.go &"

  EOF
  tags = {
    Name = "disk-migrate-${count.index + 1}"
  }
}

resource "aws_ebs_volume" "disk" {
  for_each          = { for idx, instance in aws_instance.disk : idx => instance }
  availability_zone = each.value.availability_zone
  size              = 10
  type              = "gp3"

  tags = {
    Name = "ebs-disk-${each.key}"
  }
}

resource "aws_volume_attachment" "disk" {
  for_each     = aws_ebs_volume.disk
  device_name  = "/dev/sdf"
  volume_id    = each.value.id
  instance_id  = aws_instance.disk[each.key].id
  force_detach = true
}

resource "aws_ebs_volume" "disk-migrate" {
  for_each          = { for idx, instance in aws_instance.disk-migrate : idx => instance }
  availability_zone = each.value.availability_zone
  size              = 10
  type              = "gp3"

  tags = {
    Name = "ebs-disk-migrate-${each.key}"
  }
}

resource "aws_volume_attachment" "disk-migrate" {
  for_each     = aws_ebs_volume.disk-migrate
  device_name  = "/dev/sdf"
  volume_id    = each.value.id
  instance_id  = aws_instance.disk-migrate[each.key].id
  force_detach = true
}
