package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/munsy/gobattlenet"
)

type cmdConfig struct {
	Key    string
	Token  string
	Region string
	Locale string
}

func (c cmdConfig) Settings() *battlenet.Settings {
	s := &battlenet.Settings{}

	s.Client = &http.Client{Timeout: (10 * time.Second)}

	switch c.Locale {
	case "en_US":
		s.Locale = battlenet.Locale.AmericanEnglish
		break
	case "pt_BR":
		s.Locale = battlenet.Locale.BrazilianPortuguese
		break
	case "en_GB":
		s.Locale = battlenet.Locale.BritishEnglish
		break
	case "es_ES":
		s.Locale = battlenet.Locale.CastilianSpanish
		break
	case "pt_PT":
		s.Locale = battlenet.Locale.EuropeanPortuguese
		break
	case "es_MX":
		s.Locale = battlenet.Locale.MexicanSpanish
		break
	case "zh_CN":
		s.Locale = battlenet.Locale.SimplifiedChinese
		break
	case "fr_FR":
		s.Locale = battlenet.Locale.StandardFrench
		break
	case "de_DE":
		s.Locale = battlenet.Locale.StandardGerman
		break
	case "it_IT":
		s.Locale = battlenet.Locale.StandardItalian
		break
	case "ko_KR":
		s.Locale = battlenet.Locale.StandardKorean
		break
	case "ru_RU":
		s.Locale = battlenet.Locale.StandardRussian
		break
	case "zh_TW":
		s.Locale = battlenet.Locale.TraditionalChinese
		break
	default:
		fmt.Println("Invalid type. Possible types are:")
		fmt.Println("\t- en_US (American English)")
		fmt.Println("\t- pt_BR (Brazilian Portuguese)")
		fmt.Println("\t- en_GB (British English)")
		fmt.Println("\t- es_ES (Castilian Spanish)")
		fmt.Println("\t- pt_PT (European Portuguese)")
		fmt.Println("\t- es_MX (Mexican Spanish)")
		fmt.Println("\t- zh_CN (Simplified Chinese)")
		fmt.Println("\t- fr_FR (Standard French)")
		fmt.Println("\t- de_DE (Standard German)")
		fmt.Println("\t- it_IT (Standard Italian)")
		fmt.Println("\t- ko_KR (Standard Korean)")
		fmt.Println("\t- ru_RU (Standard Russian)")
		fmt.Println("\t- zh_TW (Traditional Chinese)")
		os.Exit(1)
	}

	switch c.Region {
	case "us":
		s.Region = battlenet.Regions.US
		break
	case "eu":
		s.Region = battlenet.Regions.EU
		break
	case "kr":
		s.Region = battlenet.Regions.KR
		break
	case "tw":
		s.Region = battlenet.Regions.TW
		break
	case "sea":
		s.Region = battlenet.Regions.SEA
		break
	case "cn":
		s.Region = battlenet.Regions.CN
		break
	default:
		fmt.Println("Invalid type. Possible types are:")
		fmt.Println("\t- us (United States)")
		fmt.Println("\t- eu (Europe)")
		fmt.Println("\t- kr (Korea)")
		fmt.Println("\t- tw (Taiwan)")
		fmt.Println("\t- sea (Oceanic)")
		fmt.Println("\t- cn (China)")
		os.Exit(1)
	}

	return s
}
