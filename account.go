package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/munsy/gobattlenet"
)

var (
	accountCommand       = flag.NewFlagSet("account", flag.ExitOnError)
	accountQuotaFlag     = accountCommand.Bool("quota", false, "Display Battle.net API usage statistics for this period.")
	accountEndpointFlag  = accountCommand.Bool("endpoint", false, "Display endpoint that was used to access the given data.")
	accountUserAgentFlag = accountCommand.Bool("useragent", false, "Display user-agent that was used to make the given request.")
	accountLocaleFlag    = accountCommand.Bool("locale", false, "Display locale that was used to make the given request.")

	accountBattleIDFlag   = accountCommand.Bool("bid", false, "Display Battle.net ID and BattleTag.")
	accountSc2ProfileFlag = accountCommand.Bool("sc2", false, "Display details about a Starcraft II character.")
	accountWoWProfileFlag = accountCommand.Bool("wow", false, "Display details about a World of Warcraft profile.")
)

func parseAccountCommand() {
	checkAllAccountConfigs()

	client, err := battlenet.AccountClient(config.Settings(), *tokenFlag)

	if nil != err {
		panic(err)
	}

	if *accountBattleIDFlag == true {
		response, err := client.BattleID()

		if nil != err {
			panic(err)
		}
		if nil == response {
			fmt.Println("Nil response received.")
			os.Exit(1)
		}

		quota = response.Quota
		lastEndpoint = response.Endpoint

		printResult(response.Data)
	} else if *accountSc2ProfileFlag == true {
		response, err := client.Sc2OauthProfile()

		if nil != err {
			panic(err)
		}

		quota = response.Quota
		lastEndpoint = response.Endpoint

		printResult(response.Data)
	} else if *accountWoWProfileFlag == true {
		response, err := client.WoWOauthProfile()

		if nil != err {
			panic(err)
		}

		quota = response.Quota
		lastEndpoint = response.Endpoint

		printResult(response.Data)
	} else {
		accountCommand.PrintDefaults()
		os.Exit(1)
	}

	if *accountQuotaFlag {
		printQuota()
	}
	if *accountEndpointFlag {
		printLastEndpoint()
	}
	if *accountUserAgentFlag {
		fmt.Printf("User-Agent: %s\n", client.UserAgent())
	}
	if *accountLocaleFlag {
		fmt.Printf("Locale: %s\n", client.Locale())
	}
}
