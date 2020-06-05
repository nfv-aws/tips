package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/nfv-aws/tips/sdk-sample/ec2"
	"github.com/nfv-aws/wcafe-api-controller/config"
)

func sessInit() *session.Session {
	config.Configure()
	aws_region := config.C.SQS.Region
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(aws_region)}))
	return sess
}

func main() {
	sess := sessInit()
	ec2.Ec2CreateInstance(sess)
	ec2.Ec2DescribeInstances(sess)
}
