output "poc_endpoint_api_gateway" {
  value = aws_api_gateway_deployment.poc_api_rest_dev_development.invoke_url
}