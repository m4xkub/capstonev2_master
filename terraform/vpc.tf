resource "aws_vpc" "storage-cluster-vpc" {
  cidr_block = "10.1.0.0/16"
  // fix cidr_block
  tags = {
    Name = "storage-cluster-vpc"
  }
}

resource "aws_subnet" "private-subnet" {
  vpc_id                  = aws_vpc.storage-cluster-vpc.id
  cidr_block              = "10.1.1.0/24"
  availability_zone       = var.az
  map_public_ip_on_launch = false

  tags = {
    Name = "private-subnet"
  }
}

resource "aws_subnet" "public-subnet" {
  vpc_id                  = aws_vpc.storage-cluster-vpc.id
  cidr_block              = "10.1.2.0/24"
  availability_zone       = var.az
  map_public_ip_on_launch = false

  tags = {
    Name = "public-subnet"
  }
}

# resource "aws_eip" "nat-gateway-elastic-ip" {
#   domain = "vpc"

#   tags = {
#     Name = "nat-gateway-elastic-ip"
#   }
# }

resource "aws_internet_gateway" "storage-cluster-internet-gateway" {
  vpc_id = aws_vpc.storage-cluster-vpc.id

  tags = {
    Name = "storage-cluster-internet-gateway"
  }
}

# resource "aws_nat_gateway" "nat-gateway" {
#   allocation_id = aws_eip.nat-gateway-elastic-ip.id
#   subnet_id     = aws_subnet.public-subnet.id

#   tags = {
#     Name = "nat-gateway"
#   }
# }

resource "aws_route_table" "private-route-table" {
  vpc_id = aws_vpc.storage-cluster-vpc.id

  #   route {
  #     cidr_block = "0.0.0.0/0"
  #     gateway_id = aws_nat_gateway.nat-gateway.id
  #   }

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.storage-cluster-internet-gateway.id
  }

  route {
    cidr_block = "10.1.0.0/16"
    gateway_id = "local"
  }

  tags = {
    Name = "private-route-table"
  }
}

resource "aws_route_table" "public-route-table" {
  vpc_id = aws_vpc.storage-cluster-vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.storage-cluster-internet-gateway.id
  }

  route {
    cidr_block = "10.1.0.0/16"
    gateway_id = "local"
  }

  tags = {
    Name = "public-route-table"
  }
}

resource "aws_route_table_association" "public-subnet-asso" {
  subnet_id      = aws_subnet.public-subnet.id
  route_table_id = aws_route_table.public-route-table.id
}

resource "aws_route_table_association" "private-subnet-asso" {
  subnet_id      = aws_subnet.private-subnet.id
  route_table_id = aws_route_table.private-route-table.id
}
