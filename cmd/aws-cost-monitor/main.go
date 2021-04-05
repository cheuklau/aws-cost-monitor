package main

import (
	"log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/cheuklau/aws-cost-monitor/pkg/ec2"
)

func main() {
	// Create AWS session
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal("Unable to resolve AWS credentials. Error:", err)
	}

	// Perform EC2 scan
	ec2.Run(sess)
}