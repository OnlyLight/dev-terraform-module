output "main_vpc_id" {
  value       = aws_vpc.main.id
  description = "The main VPC id"
}

output "private_zone1_id" {
  value       = aws_subnet.private_zone1.id
  description = "The id private zone 1"
}

output "private_zone2_id" {
  value       = aws_subnet.private_zone2.id
  description = "The id private zone 2"
}

output "public_zone1_id" {
  value       = aws_subnet.public_zone1.id
  description = "The id public zone 1"
}

output "public_zone2_id" {
  value       = aws_subnet.public_zone2.id
  description = "The id public zone 2"
}
