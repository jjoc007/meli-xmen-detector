resource "aws_lambda_function" "poc_process_dna_lambda" {
  filename      =  "../ms/build/xmen-dna-process/main.zip"
  function_name = "poc_process_dna_lambda"
  role          = aws_iam_role.poc_lambda_role.arn
  handler       = "main"
  source_code_hash = filebase64sha256("../ms/build/xmen-dna-process/main.zip")
  runtime = "go1.x"
  timeout = 10

  environment {
    variables = {
      APP = local.app
    }
  }
}

resource "aws_lambda_function" "poc_get_stats_lambda" {
  filename      =  "../ms/build/xmen-stats-get/main.zip"
  function_name = "poc_get_stats_lambda"
  role          = aws_iam_role.poc_lambda_role.arn
  handler       = "main"
  source_code_hash = filebase64sha256("../ms/build/xmen-stats-get/main.zip")
  runtime = "go1.x"

  environment {
    variables = {
      APP = local.app
    }
  }
}
