package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"bytes"
)

// ANSI escape code for blue color
const BlueColor = "\033[34m"

// ANSI escape code for resetting color
const ResetColor = "\033[0m"

func main() {
	// Define a flag for the JWT token
	tokenFlag := flag.String("t", "", "JWT token")
	flag.Parse()

	// Check if the token flag is provided
	if *tokenFlag == "" {
		fmt.Println("Please provide a JWT token using the -t flag")
		return
	}

	// Split the token into its three parts: header, payload, and signature
	parts := strings.Split(*tokenFlag, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid JWT token format")
		return
	}

	headerJSON, _ := base64.RawURLEncoding.DecodeString(parts[0])
	payloadJSON, _ := base64.RawURLEncoding.DecodeString(parts[1])
	signature := parts[2]

	fmt.Print("-------------------------------------")

	// Print labels in blue color
	fmt.Println(BlueColor)
	fmt.Print("Header:")
	fmt.Print(ResetColor) // Reset color after the label
	fmt.Println(" ", string(headerJSON))

	fmt.Print(BlueColor)
	fmt.Print("Payload:")
	fmt.Print(ResetColor) // Reset color after the label
	stringifiedPayload, err := stringifyJSON(payloadJSON)
	if err != nil {
		fmt.Println("Error decoding payload:", err)
		return
	}
	fmt.Println(" ", stringifiedPayload)

	fmt.Print(BlueColor)
	fmt.Print("Signature:")
	fmt.Print(ResetColor) // Reset color after the label
	fmt.Println(" ", signature)
	fmt.Print("-------------------------------------")
	
}

func stringifyJSON(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
