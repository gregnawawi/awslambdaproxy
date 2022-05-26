package main

import (
	"fmt"
	"github.com/gregnawawi/awslambdaproxy/cmd/awslambdaproxy"
	"os"
)

func main() {
	if err := awslambdaproxy.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
