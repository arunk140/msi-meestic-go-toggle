package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// Define a struct to represent the JSON response
type Response struct {
	Profiles []string `json:"profiles"`
	Active   int      `json:"active"`
}

func main() {

	meesticPath := "/run/meestic.sock"
	meesticConn, err := net.Dial("unix", meesticPath)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer meesticConn.Close()

	response := make([]byte, 1024) // Adjust the buffer size as needed
	n, err := meesticConn.Read(response)
	if err != nil {
		fmt.Printf("Failed to read response from server: %v\n", err)
		os.Exit(1)
	}

	var parsedResponse Response
	err = json.Unmarshal(response[:n], &parsedResponse)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v\n", err)
		os.Exit(1)
	}

	profileLen := len(parsedResponse.Profiles)
	fmt.Printf("Num Profiles: %d\n", profileLen)
	// Extract the 'active' index as a variable
	activeIndex := parsedResponse.Active
	fmt.Printf("Active index: %d\n", activeIndex)

	idx := (activeIndex + 1) % profileLen // Replace with your desired index value
	message := fmt.Sprintf("{\"apply\": %d}\n", idx)

	// Process the received response
	fmt.Printf("Received response from server: %s\n", string(response[:n]))
	fmt.Printf("Applying profile %d\n", idx)

	_, err = meesticConn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Failed to send message to server: %v\n", err)
		os.Exit(1)
	}
	// Read and process the response
	response = make([]byte, 1024) // Adjust the buffer size as needed
	n, err = meesticConn.Read(response)
	if err != nil {
		fmt.Printf("Failed to read response from server: %v\n", err)
		os.Exit(1)
	}

	// Process the received response
	fmt.Printf("Received response from server: %s\n", string(response[:n]))
}
