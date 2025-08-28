This project is currently in progress!

Current Abstract: In Progress

Please check out my research journal entries detailing my experience on this project from start to finish [here!](/Journal.md)

To view the various lambda-functions referenced, please navigate to the [lambda-functions](https://github.com/djbertolo/aws-go-lambda-security-research/tree/main/lambda-functions) directory.

Below is a list of the Lambda functions that were created throughout this research project, list of concepts covered by the function, and a goal of the function.

|Lambda Function|Concepts Covered|Goal|
|---------------|----------------|-----------|
|Hello-World|Cross-compiling go for linux, Lambda's custom runtime, Creating Lambda Execution Role with IAM, Deploying code with the AWS CLI| Compile a Go binary, package it, and deploy it as a Lambda function manually. This is the foundational skill for all subsequent exercises.
|Processing-Input|Using Go structs for request/response bodies, Updating a function's code via the AWS CLI|Pass data into the function and receive a structured JSON response.
|Triggering-Gateway|Manually creating an HTTP API in API Gateway, Creating an integration between the API and Lambda function, Lambda permissions for API Gateway invocation| Create a public HTTP endpoint that executes the Lambda function.
