#############################
# DynamoDB
#############################
resource "aws_dynamodb_table" "main-table" {
  name           = var.table-name
  billing_mode   = "PROVISIONED"
  read_capacity  = 2
  write_capacity = 2
  hash_key       = "orderId"
  range_key      = "customerId"

  stream_enabled   = true
  stream_view_type = "NEW_IMAGE"

  attribute {
    name = "orderId"
    type = "S"
  }

  attribute {
    name = "customerId"
    type = "S"
  }

  attribute {
    name = "shipped"
    type = "S"
  }


  local_secondary_index {
    name            = "lsi-orderId-customerId"
    range_key       = "customerId"
    projection_type = "ALL"
  }

    global_secondary_index {
    name               = "gsi-shipped"
    hash_key           = "orderId"
    range_key          = "shipped"
    write_capacity     = 1
    read_capacity      = 1
    projection_type    = "ALL"
  }


  tags = {
    Name        = "${var.table-name}-${var.environment}"
    Environment = var.environment
  }
}

#############################
# Execut AWS CLI script
#############################
# resource "null_resource" "init-db" {
#   triggers = {
#     new = timestamp()
#   }
#   provisioner "local-exec" {
#     command = <<EOT
#       aws dynamodb batch-write-item --request-items file://../../static/mock-data.json --endpoint-url http://localhost:4566
#     EOT
#   }
#   depends_on = [aws_dynamodb_table.main-table]
# }
