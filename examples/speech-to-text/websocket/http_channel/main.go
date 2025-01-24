// Copyright 2023-2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package main

// streaming
import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"

	interfaces "github.com/tenyx-corp/deepgram-go-sdk/pkg/client/interfaces"
	client "github.com/tenyx-corp/deepgram-go-sdk/pkg/client/listen"
)

const (
	STREAM_URL = "http://stream.live.vc.bbcmedia.co.uk/bbc_world_service"
)

func main() {
	// init library
	client.InitWithDefault()
	// client.Init(client.InitLib{
	// 	LogLevel: client.LogLevelTrace, // LogLevelDefault, LogLevelFull, LogLevelDebug, LogLevelTrace
	// })

	// Go context
	ctx := context.Background()

	// print instructions
	fmt.Print("\n\nPress ENTER to exit!\n\n")

	// set the Transcription options
	transcriptOptions := &interfaces.LiveTranscriptionOptions{
		Language:  "en-US",
		Punctuate: true,
	}

	// create a Deepgram client
	dgClient, err := client.NewWSUsingChanForDemo(ctx, transcriptOptions)
	if err != nil {
		fmt.Println("ERROR creating LiveTranscription connection:", err)
		return
	}

	// get the HTTP stream
	httpClient := new(http.Client)

	res, err := httpClient.Get(STREAM_URL)
	if err != nil {
		fmt.Printf("httpClient.Get failed. Err: %v\n", err)
		return
	}
	fmt.Printf("Stream is up and running %s\n", reflect.TypeOf(res))

	// connect the websocket to Deepgram
	bConnected := dgClient.Connect()
	if !bConnected {
		fmt.Println("Client.Connect failed")
		os.Exit(1)
	}

	go func() {
		// feed the HTTP stream to the Deepgram client (this is a blocking call)
		dgClient.Stream(bufio.NewReader(res.Body))
	}()

	// wait for user input to exit
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	// close HTTP stream
	res.Body.Close()

	// close DG client
	dgClient.Stop()

	fmt.Printf("\n\nProgram exiting...\n")
}
