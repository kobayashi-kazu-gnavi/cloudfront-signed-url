package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"time"
)

const (
	privateKeyString = `-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----
`
	keyPairID = "keyPairID"
	region    = "region"
	url       = "url"
)

func main() {
	lambda.Start(handler)
}

// handler lambda handler.
func handler(ctx context.Context) {
	// convert to rsa.PrivateKey
	block, _ := pem.Decode([]byte(privateKeyString))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("failed to parse private key: %v\n", err)
		return
	}

	// create signer
	signer := sign.NewURLSigner(keyPairID, privateKey)

	// create signed url
	signedURL, err := signer.Sign(url, time.Now().Add(24*time.Hour))
	if err != nil {
		fmt.Printf("failed to create signed cookie: %v\n", err)
		return
	}
	fmt.Printf("signed url: %s\n", signedURL)
}
