// Copyright 2023-2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package version

import (
	"context"

	interfaces "github.com/tenyx-corp/deepgram-go-sdk/pkg/client/interfaces"
)

/*
GetPrerecordedAPI is a function which controls the versioning of the live transcription API and provides
mechanism for:

- overriding the host endpoint
- overriding the version used
- overriding the endpoint path
- additional arguments to the query string/parameters

The return value is the complete URL endpoint to be used for the live transcription
*/
func GetPrerecordedAPI(ctx context.Context, host, version, path string, options *interfaces.PreRecordedTranscriptionOptions, args ...interface{}) (string, error) {
	return getAPIURL(ctx, "prerecorded", host, version, path, options, args...)
}
