// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package web

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/helpers"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/resources"
	"github.com/subosito/gotenv"
)

var (
	webName               = "testthis2341now"
	resourceGroupLocation = "southcentralus"
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
func ExampleCreateAppServicePlan() {
	helpers.SetResourceGroupName("createorupdate")
	ctx := context.Background()
	defer resources.Cleanup(ctx)
	_, err := resources.CreateGroup(ctx, "whateves")
	if err != nil {
		helpers.PrintAndLog(err.Error())
	}
	_, err = CreateAppServicePlan(ctx, webName)
	if err != nil {
		helpers.PrintAndLog("Failed App Service Plan creation.")
		helpers.PrintAndLog(err.Error())
	} else {
		helpers.PrintAndLog("Created App Service Plan")
	}

	// Output:
	// created Created App Service Plan
}
func ExampleCreateWebSite() {
	helpers.SetResourceGroupName("createorupdate")
	ctx := context.Background()
	defer resources.Cleanup(ctx)
	_, err := resources.CreateGroup(ctx, "whateves")
	if err != nil {
		helpers.PrintAndLog(err.Error())
	}
	_, err = CreateWebSite(ctx, webName)
	if err != nil {
		helpers.PrintAndLog(err.Error())
	}
	helpers.PrintAndLog("created Website")

	// Output:
	// created Website
}
