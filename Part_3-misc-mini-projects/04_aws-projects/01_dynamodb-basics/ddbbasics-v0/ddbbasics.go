package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewSession returns a pointer to a new AWS Session
func NewSession() *session.Session {
	credentials.NewSharedCredentials()

	sessionConfig := aws.Config{
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{

		}),
	}
}

func main() {

}