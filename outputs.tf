########################
# Main Table
########################
output "main-table-id" {
  value = aws_dynamodb_table.main-table.id
}

output "main-table-arn" {
  value = aws_dynamodb_table.main-table.arn
}

output "main-table-stream-arn" {
  value = aws_dynamodb_table.main-table.stream_arn
}
