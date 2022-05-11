output "public_ip" {
  //value       = aws_instance.terragrunt_ubuntu.public_ip
  value       = aws_instance.terragrunt_Postgres_DB.public_ip
 }
