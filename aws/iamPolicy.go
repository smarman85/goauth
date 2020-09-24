package main

import (
  "encoding/json"
  "fmt"
  "os"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/iam"
)

type PolicyDocument struct {
  Version string
  Statement []StatementEntry
}

type StatementEntry struct {
  Effect string
  Action []string
  Resource string
}

func main() {
  // create session
  sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-east-1")},
  )
  if err != nil {
    logErrorf("Error creating session: %v", err)
  }

  // create service
  svc := iam.New(sess)
  fmt.Printf("%T", svc)

  // Create policy
  policy := PolicyDocument{
    Version: "2012-10-17",
    Statement: []StatementEntry{
      StatementEntry{
        Effect: "Allow",
        Action: []string{
          "logs:CreateLogGroup",
        },
        Resource: "iam:::*:logs:*",
      },
      StatementEntry{
        Effect: "Allow",
        Action: []string{
          "dynamodb:DeleteItem",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:Scan",
          "dynamodb:UpdateItem",
        },
        Resource: "aws:::dynamodbb:*",
      },
    },
  }

  // marshal the policy to json and pass CreatePolicy
  b, err := json.Marshal(&policy)
  if err != nil {
    logErrorf("Error marshalling poliyc", err)
    return
  }
  fmt.Println(string(b))

  //result, err := svc.CreatePolicy(&iam.CreatePolicyInput{
  //  PolicyDocument: aws.String(string(b)),
  //  PolicyName:     aws.String("myDynamodbPolicy"),
  //})

  //if err != nil {
  //  fmt.Println("Error", err)
  //  return
  //}

  //fmt.Println("New policy", result)

}

func logErrorf(msg string, args ...interface{}) {
  fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

