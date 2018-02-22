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
	webName               = "az-samples-go-web-" + helpers.GetRandomLetterSequence(5)
	resourceGroupName     = "canttouchthis"
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

func ExampleCreateWebSite() {
	helpers.SetResourceGroupName(resourceGroupName)
	ctx := context.Background()
	defer resources.Cleanup(ctx)
	_, err := resources.CreateGroup(ctx, helpers.ResourceGroupName())
	if err != nil {
		helpers.PrintAndLog("Failed Web Site creation.")
		helpers.PrintAndLog(err.Error())
	}
	_, err = CreateAppServicePlan(ctx, webName)
	if err != nil {
		helpers.PrintAndLog("Failed App Service Plan creation.")
		helpers.PrintAndLog(err.Error())
	} else {
		helpers.PrintAndLog("Created App Service Plan")
		_, err = CreateWebSite(ctx, webName)
		if err != nil {
			helpers.PrintAndLog(err.Error())
		} else {
			helpers.PrintAndLog("Created Website")
		}

	}

	// Output:
	// Created App Service Plan
	// Created Website
}
