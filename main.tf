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
    name            = "gsi-shipped"
    hash_key        = "orderId"
    range_key       = "shipped"
    write_capacity  = 1
    read_capacity   = 1
    projection_type = "ALL"
  }


  tags = {
    Name        = "${var.table-name}-${var.environment}"
    Environment = var.environment
  }
}
#
# REMBER ONLY UNCOMMENT ONE OF THE TWO STEPS BELOW. THE Go script requires compiling, instructions provided in the README
#
###################################
# Execut AWS CLI script  Option 1/2
###################################
# resource "null_resource" "init-db" {

// This will cause the upload script to only execute when the table changes id (recreate). 
#   triggers = {
#     new = aws_dynamodb_table.main-table.id
#   }
#   provisioner "local-exec" {
#     command = <<EOT
#       aws dynamodb batch-write-item --request-items file://static/formatted-data.json --endpoint-url ${var.dynamodb-addr}
#     EOT
#   }
#   depends_on = [aws_dynamodb_table.main-table]
# }


#####################################
# Execut Go binary script Option 2/2
#####################################
resource "null_resource" "init-db-go" {
  // This will cause the upload script to only execute when the table changes id (recreate). 
  triggers = {
    new = aws_dynamodb_table.main-table.id
  }
  provisioner "local-exec" {
    command = <<EOT
      ./upload-logic
    EOT
    environment = {
      TABLE_NAME = aws_dynamodb_table.main-table.id
      REGION = var.region
      DYNAMODB_ADDR = var.dynamodb-addr
      JSON_PATH = var.json-file-path
    }
  }
  depends_on = [aws_dynamodb_table.main-table]
}