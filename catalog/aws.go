package catalog

// awsDisplayNames maps AWS service identifiers (sas.Technology.Service
// when Provider is "aws") to their human-friendly product names.
var awsDisplayNames = map[string]string{
	"lambda":          "AWS Lambda",
	"ecs":             "Amazon ECS",
	"eks":             "Amazon EKS",
	"ec2":             "Amazon EC2",
	"rds":             "Amazon RDS",
	"s3":              "Amazon S3",
	"dynamodb":        "Amazon DynamoDB",
	"sqs":             "Amazon SQS",
	"sns":             "Amazon SNS",
	"api-gateway":     "Amazon API Gateway",
	"cloudfront":      "Amazon CloudFront",
	"elb":             "Elastic Load Balancing",
	"alb":             "Application Load Balancer",
	"secrets-manager": "AWS Secrets Manager",
}
