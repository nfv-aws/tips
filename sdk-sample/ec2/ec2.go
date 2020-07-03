package ec2

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

var (
	name          = "sdk-test"
	securitygroup = "default"
)

func ec2NewSession(sess *session.Session) *ec2.EC2 {
	svc := ec2.New(sess)
	return svc
}

// EC2インスタンスの作成
func Ec2CreateInstance(sess *session.Session) {
	svc := ec2NewSession(sess)

	// 作成するインスタンスのパラメータの指定
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		// ここで好みのパラメータに設定する
		ImageId:      aws.String("ami-0a1c2ec61571737db"),
		InstanceType: aws.String("t2.micro"),
		// ↓ではなくSecurityGroupIDを利用したID名指定でもどちらでも可能
		SecurityGroups: []*string{&securitygroup},
		// 作成するインスタンスの下限と上限
		MinCount: aws.Int64(1),
		MaxCount: aws.Int64(1),
		// タグでNameつけるときはこの処理を追加
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("instance"),
				Tags: []*ec2.Tag{
					{Key: aws.String("Name"), Value: &name},
				},
			},
		},
	})

	if err != nil {
		fmt.Println("Could not create instance", err)
		return
	}

	fmt.Println("Created instance", *runResult.Instances[0].InstanceId)

}

// EC2インスタンス情報の表示
func Ec2DescribeInstances(sess *session.Session) {
	svc := ec2NewSession(sess)

	res, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tag Name", "Instance Id", "Instance Type", "AZ", "Status"})

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			var tag_name string
			for _, t := range i.Tags {
				if *t.Key == "Name" {
					tag_name = *t.Value
				}
			}
			table.Append([]string{
				tag_name,
				*i.InstanceId,
				*i.InstanceType,
				*i.Placement.AvailabilityZone,
				*i.State.Name,
			})
		}
	}
	table.Render()
}
