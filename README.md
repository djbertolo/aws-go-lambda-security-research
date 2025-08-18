Title: A Practical Investigation of Event Data Injection Vulnerabilities in Go-based AWS Lambda Functions

Author: [Your Name]

Keywords: Serverless Security, AWS Lambda, Golang (Go), Application Security, Vulnerability Research, Event Data Injection, Cloud-Native Security.

Project Motivation and Scope:

Serverless computing, particularly AWS Lambda, has become a dominant paradigm for building scalable, event-driven applications. Go is an increasingly popular language for developing high-performance Lambda functions due to its fast startup times and efficient concurrency. However, this combination introduces a unique attack surface that is not always addressed by traditional security models. This project focuses on investigating a critical class of serverless vulnerabilities: event data injection. The core research question is: How can untrusted data from AWS event sources (e.g., API Gateway, S3, SQS) be leveraged to manipulate the execution flow and compromise the security of a Go-based Lambda function? The scope of this investigation will be limited to developing and analyzing proof-of-concept vulnerabilities within a controlled AWS environment.

Methodology:

This research will be conducted through a hands-on, empirical approach involving four key phases:

Literature Review & Threat Modeling: Synthesize existing knowledge from sources like the OWASP Serverless Top 10 and academic papers to identify theoretical event data injection attack vectors relevant to Go applications on AWS Lambda.

Development of a Vulnerable Testbed: Create a collection of small, single-purpose AWS Lambda functions written in Go. Each function will be intentionally designed with a potential vulnerability (e.g., OS command injection, data leakage through improper error handling, SQL injection) that can be triggered by manipulated event data.

Reproducible Deployment: The entire serverless application, including the Lambda functions and their corresponding event triggers (e.g., API Gateway endpoints), will be defined and deployed using an Infrastructure as Code (IaC) framework like AWS SAM or Terraform. This ensures the experimental environment is version-controlled and easily reproducible.

Proof-of-Concept (PoC) Exploitation and Analysis: For each vulnerable function, develop and execute a PoC exploit. This will involve crafting malicious event payloads (e.g., JSON objects for API Gateway) and demonstrating the security impact. The root cause of each vulnerability will be analyzed in the context of Go's language features and the Lambda execution environment.

Expected Outcomes & Contributions:

The primary outcome of this project will be a public GitHub repository serving as both a tangible research artifact and an educational resource. This repository will contain:

The source code for all vulnerable Go-based Lambda functions.

The IaC templates required to deploy the testbed.

Detailed documentation and scripts for each PoC exploit.

A REPORT.md file summarizing the findings, analyzing each vulnerability class, and proposing specific mitigation strategies and secure coding best practices for developers.

This project will serve as a practical demonstration of research capability, providing a deep, hands-on understanding of application security in modern cloud-native environments. 