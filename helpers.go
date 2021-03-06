package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
)

func printResult(v interface{}) {
	fields := reflect.ValueOf(v).Elem()
	fieldTypes := fields.Type()

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		fmt.Printf("%s: %v\n", fieldTypes.Field(i).Name, field.Interface())
	}
}

func printQuota() {
	fmt.Printf("X-Plan-Qps-Allotted: %v\n", quota.QPSAllotted())
	fmt.Printf("X-Plan-Qps-Current: %v\n", quota.QPSCurrent())
	fmt.Printf("X-Plan-Quota-Allotted: %v\n", quota.QuotaAlloted())
	fmt.Printf("X-Plan-Quota-Current: %v\n", quota.QuotaCurrent())
	fmt.Printf("X-Plan-Quota-Reset: %v\n", quota.QuotaReset())
}

func printUserAgent() {
	fmt.Printf("User-Agent: %s\n", client.UserAgent())
}
func printLocale() {
	fmt.Printf("Locale: %s\n", client.Locale())
}

func printLastEndpoint() {
	fmt.Printf("Endpoint: %s\n", lastEndpoint)
}

func checkTokenFlag() {
	if *tokenFlag == "" {
		fmt.Println("Token not set.")
		os.Exit(1)
	}
}

func checkKeyFlag() {
	if *keyFlag == "" {
		fmt.Println("Key not set.")
		os.Exit(1)
	}
}

func checkRegionFlag() {
	if *regionFlag == "" {
		fmt.Println("Region not set.")
		os.Exit(1)
	}
}

func checkLocaleFlag() {
	if *localeFlag == "" {
		fmt.Println("Locale not set.")
		os.Exit(1)
	}
}

func checkAllGameConfigs() {
	checkKeyFlag()
	checkRegionFlag()
	checkLocaleFlag()
}

func checkAllAccountConfigs() {
	checkTokenFlag()
	checkRegionFlag()
	checkLocaleFlag()
}

func checkAllConfigs() {
	checkKeyFlag()
	checkTokenFlag()
	checkRegionFlag()
	checkLocaleFlag()
}

func writeTOML(key, token, region, locale string) error {
	var inputs = cmdConfig{
		key,
		token,
		region,
		locale,
	}

	var buffer bytes.Buffer

	encoder := toml.NewEncoder(&buffer)

	err := encoder.Encode(inputs)

	if nil != err {
		return err
	}

	return ioutil.WriteFile("config.toml", buffer.Bytes(), 0644)
}

func readTOML(filename string) {
	if _, err := toml.DecodeFile(filename, &config); err != nil {
	}
}
