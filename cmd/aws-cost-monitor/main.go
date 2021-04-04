package main

import (
	"fmt"
	"github.com/cheuklau/aws-cost-monitor/ec2"
)

func main() {
	fmt.Println("Inside main")
	ec2.FindAllEc2s()
}