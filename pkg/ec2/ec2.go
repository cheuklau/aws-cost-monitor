package ec2

import (
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Ec2 struct {
	instanceType string
	// todo: add more info including tag for owner
}

func Run(sess *session.Session) {
	fmt.Println("Starting EC2 scan...")

	// Create EC2 session
	svc := ec2.New(sess)

	// Find all instances
	// Note: Do not need to pass a map by reference
	allEc2s := make(map[string]Ec2)
	FindAllEc2s(svc, allEc2s)
	for instanceId, instanceMeta := range allEc2s {
		fmt.Println(instanceId, instanceMeta.instanceType)
	}
}

func FindAllEc2s(svc *ec2.EC2, allEc2s map[string]Ec2) {
	result, err := svc.DescribeInstances(nil)
	if err != nil {
		log.Fatal("Unable to retrieve EC2 instances. Error:", err)
	}
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			// todo: add more info including tag for owner
			allEc2s[*instance.InstanceId] = Ec2{
				*instance.InstanceType,
			}
		}
	}
}