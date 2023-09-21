package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"os"
)

func main() {
	var profile string
	flag.StringVar(&profile, "p", "default", "profile name")
	flag.Parse()

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, // Must be set to enable
		Profile:           profile,                    // input "-p profile_name"
	})
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	client := sts.New(sess)

	identity, err := client.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Printf(
		"Account: %s\n, Arn: %s\n, UserId: %s\n",
		aws.StringValue(identity.Account),
		aws.StringValue(identity.UserId),
		aws.StringValue(identity.Arn),
	)
}
