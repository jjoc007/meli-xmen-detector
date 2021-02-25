resource "aws_api_gateway_rest_api" "poc_rest_api" {
  name        = "poc_rest_api_dna_crud"
  description = "API Gateway"
}

resource "aws_api_gateway_deployment" "poc_api_rest_dev_development" {
  depends_on = [
    aws_api_gateway_integration.poc_mutant_post_integration,
    aws_api_gateway_integration.poc_stats_get_integration,
  ]

  rest_api_id = aws_api_gateway_rest_api.poc_rest_api.id
  stage_name  = "dev"

  triggers = {
    redeployment = sha1(join(",", list(
    jsonencode(aws_api_gateway_integration.poc_stats_get_integration),
    jsonencode(aws_api_gateway_integration.poc_mutant_post_integration)
    )))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_resource" "poc_mutant_resource" {
  rest_api_id = aws_api_gateway_rest_api.poc_rest_api.id
  parent_id   = aws_api_gateway_rest_api.poc_rest_api.root_resource_id
  path_part   = "mutant"
}

resource "aws_api_gateway_method" "poc_mutant_post_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_mutant_resource.id
  http_method   = "POST"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_mutant_post_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_mutant_resource.id
  http_method             = aws_api_gateway_method.poc_mutant_post_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_process_dna_lambda.invoke_arn
}

resource "aws_api_gateway_resource" "poc_stats_resource" {
  rest_api_id = aws_api_gateway_rest_api.poc_rest_api.id
  parent_id   = aws_api_gateway_rest_api.poc_rest_api.root_resource_id
  path_part   = "stats"
}

resource "aws_api_gateway_method" "poc_stats_get_method" {
  rest_api_id   = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id   = aws_api_gateway_resource.poc_stats_resource.id
  http_method   = "GET"
  authorization = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "poc_stats_get_integration" {
  rest_api_id             = aws_api_gateway_rest_api.poc_rest_api.id
  resource_id             = aws_api_gateway_resource.poc_stats_resource.id
  http_method             = aws_api_gateway_method.poc_stats_get_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.poc_get_stats_lambda.invoke_arn
}