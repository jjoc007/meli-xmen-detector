resource "aws_lambda_permission" "lambda_permission_mutant_post_rest" {
  depends_on    = [aws_lambda_function.poc_process_dna_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_process_dna_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_mutant_post_method.http_method}${aws_api_gateway_resource.poc_mutant_resource.path}"
}

resource "aws_lambda_permission" "lambda_permission_stats_get_rest" {
  depends_on    = [aws_lambda_function.poc_get_stats_lambda]
  principal     = "apigateway.amazonaws.com"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.poc_get_stats_lambda.function_name
  source_arn = "${aws_api_gateway_rest_api.poc_rest_api.execution_arn}/*/${aws_api_gateway_method.poc_stats_get_method.http_method}${aws_api_gateway_resource.poc_stats_resource.path}"
}