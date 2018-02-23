// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package servicebus

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/helpers"
	"github.com/subosito/gotenv"
)

var (
	resourceGroupNameSuffix = "servicebus"
	serviceBusNamespaceName = "azure-sdk-for-go-sample-namespace-" + strings.ToLower(helpers.GetRandomLetterSequence(4))
)

func TestMain(m *testing.M) {
	err := parseArgs()
	if err != nil {
		log.Fatalln("failed to parse args")
	}

	os.Exit(m.Run())
}

func parseArgs() error {
	gotenv.Load()
	err := helpers.ParseArgs()
	if err != nil {
		return fmt.Errorf("cannot parse args: %v", err)
	}

	return nil
}
