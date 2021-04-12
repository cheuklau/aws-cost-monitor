package main

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/cheuklau/aws-cost-monitor/pkg/ec2"
	"github.com/cheuklau/aws-cost-monitor/pkg/input"
)

func main() {

	// Read input
	var rules input.RulesConfig
	rules = input.Run("examples/sample.yml")
	// Print input in pretty JSON
	rulesJson, err := json.MarshalIndent(rules, "", "  ")
	fmt.Printf("Retrieved input:\n%s\n",string(rulesJson))

	// Create AWS session
	sess, err := session.NewSession()
	if err != nil {
		log.Fatal("Unable to resolve AWS credentials. Error:", err)
	}

	// Perform EC2 scan
	ec2.Run(sess)
}