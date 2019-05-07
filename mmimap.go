package mimir

var (
	paymentCardConfigurations = map[string]paymentCardConfiguration{
		AmericanExpress:         americanExpressConfiguration,
		ChinaUnionPay:           chinaUnionPayConfiguration,
		DinnerClubInternational: dinnerClubInternationalConfiguration,
		DinnerClubCarteBlanche:  dinnerClubCarteBlancheConfiguration,
		DinnerClub:              dinnerClubConfiguration,
		Discover:                discoverConfiguration,
		Rupay:                   rupayConfiguration,
		InterPayment:            interPaymentConfiguration,
		InstaPayment:            instaPaymentConfiguration,
		JCB:                     jcbConfiguration,
		Maestro:                 maestroConfiguration,
		Dankort:                 dankortConfiguration,
		MIR:                     mirConfiguration,
		Mastercard:              mastercardConfiguration,
		Visa:                    visaConfiguration,
		UATP:                    uatpConfiguration,
	}

	americanExpressConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: AmericanExpress,
		Pattern:                     "^3[47][0-9]{13}$",
		CardSecurityCodeLength:      4,
		Examples: []string{
			"345752218692713",
			"377932215823195",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			15: "#### ###### #####",
		},
	}

	chinaUnionPayConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: ChinaUnionPay,
		Pattern:                     "^62[0-9]{14,17}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"6225410287201336",
			"62254102872013369",
			"622541028720133691",
			"6225410287201336918",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
			17: "#### #### #### #####",  // May not be accurate
			18: "#### #### #### ######", // May not be accurate
			19: "###### #############",
		},
	}

	dinnerClubInternationalConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: DinnerClubInternational,
		Pattern:                     "^3(0[0-59][0-9]{11}|[689][0-9]{12})$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"30037174022893",
			"30150312694976",
			"30219652520297",
			"30329187642019",
			"30454821880310",
			"30533713636318",
			"30902796777002",
			"38248931200191",
			"36892904829528",
			"39134092552753",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			14: "#### ###### ####",
		},
	}

	dinnerClubCarteBlancheConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: DinnerClubInternational,
		Pattern:                     "^30[0-5][0-9]{11}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"30037174022893",
			"30150312694976",
			"30219652520297",
			"30329187642019",
			"30454821880310",
			"30533713636318",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			14: "#### ###### ####",
		},
	}

	dinnerClubConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: DinnerClub,
		Pattern:                     "^5[45][0-9]{14}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"5420185851973859",
			"5590722767852282",
			"5574933443281743",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
		},
	}

	discoverConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: Discover,
		Pattern:                     "^6(011[0-9]{12}|22(12[6-9]|1[3-9][0-9]|[2-8][0-9]{2}|9[01][0-9]|92[0-5])[0-9]{10}|4[4-9][0-9]{13}|5[0-9]{14})$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"6011594786717005",
			"6011744921235898",
			"6011625899698235",
			"6011286921785324",
			"6011153184210640",
			"6011297183021415",
			"6442051015812940",
			"6453224731777921",
			"6467015443039624",
			"6477329463179212",
			"6486392641029856",
			"6491623729118200",
			"6549231111740110",
			"6221261529352302",
			"6221309814267758",
			"6229124911624542",
			"6229251960959246",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
		},
	}

	rupayConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: Rupay,
		Pattern:                     "^6(0[0-9]{14}|52[12][0-9]{12})$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"6032810923281217",
			"6521054127370108",
			"6522201132570169",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
		},
	}

	interPaymentConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: InterPayment,
		Pattern:                     "^636[0-9]{13,16}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"6364657271602637",
			"63646572716026378",
			"636465727160263783",
			"6364657271602637834",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
			17: "#### #### ##### ####",   // May not be accurate
			18: "#### #### ###### ####",  // May not be accurate
			19: "#### #### ####### ####", // May not be accurate
		},
	}

	instaPaymentConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: InstaPayment,
		Pattern:                     "^63[7-9][0-9]{13}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"6372766830370287",
			"6385122111231563",
			"6398826308675201",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
		},
	}

	jcbConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: JCB,
		Pattern:                     "^(352[89]|35[3-8][0-9])[0-9]{12,15}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"3528242311613520",
			"3529824509910443",
			"3589107644145290",
			"3540363124918023",
			"35403631249180234",
			"354036312491802342",
			"3540363124918023426",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
			17: "#### #### ##### ####",   // May not be accurate
			18: "#### #### ###### ####",  // May not be accurate
			19: "#### #### ####### ####", // May not be accurate
		},
	}

	maestroConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: Maestro,
		Pattern:                     "^((5(0|[6-8])|6[0-9])[0-9]{10,17})$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"505792268730",
			"5057922687301",
			"50579226873018",
			"505792268730183",
			"5057922687301834",
			"50579226873018342",
			"505792268730183426",
			"5057922687301834267",
			"5624731905154527",
			"5733843990180891",
			"5845992036759237",
			"5893571811944676",
			"5020450000555203",
			"5020369747104747",
			"6344712140017130",
			"6326519127091547",
			"6721562980126811",
			"6025059196982890",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			12: "#### #### ####", // May not be accurate
			13: "#### #### #####",
			14: "#### ###### ####", // May not be accurate
			15: "#### ###### #####",
			16: "#### #### #### ####",
			17: "#### #### #### #### #",
			18: "#### #### #### #### ##", // May not be accurate
			19: "#### #### #### #### ###",
		},
	}

	dankortConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: Dankort,
		Pattern:                     "^(5019|4571)[0-9]{12}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"4571151420387179",
			"5019492784413939",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
		},
	}

	mirConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: MIR,
		Pattern:                     "^220[0-4][0-9]{12}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"2200954085571885",
			"2201462642892074",
			"2202082047710553",
			"2203306833086392",
			"2204383503972804",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			15: "#### ###### #####",
		},
	}

	mastercardConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: Mastercard,
		Pattern:                     "^(5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12})$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"5122877833713855",
			"5132292671710369",
			"5122346267602737",
			"5211291116423210",
			"5222544024043232",
			"5372962217654946",
			"5481620704517038",
			"5536637164877434",
			"2222591215481074",
			"2719776701366414",
			"2720473203321611",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			16: "#### #### #### ####",
		},
	}

	visaConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: Visa,
		Pattern:                     "^4[0-9]{12,18}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"4287429749178",
			"42874297491787",
			"428742974917875",
			"4287429749178759",
			"42874297491787591",
			"428742974917875918",
			"4287429749178759183",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			13: "#### #### #####",   // May not be accurate
			14: "#### ###### ####",  // May not be accurate
			15: "#### ###### #####", // May not be accurate
			16: "#### #### #### ####",
			17: "#### #### #### #### #",   // May not be accurate
			18: "#### #### #### #### ##",  // May not be accurate
			19: "#### #### #### #### ###", // May not be accurate
		},
	}

	uatpConfiguration = paymentCardConfiguration{
		MajorIndustryIdentifierName: UATP,
		Pattern:                     "^1[0-9]{14}$",
		CardSecurityCodeLength:      3,
		Examples: []string{
			"100070352000006",
			"165291225432273",
		},
		ExpirationDateFormat: "01/2006",
		Structure: map[int]string{
			15: "#### ##### ######",
		},
	}
)

// GetPaymentCardConfiguration returns informations about the payment cards from a given issuer name
// If the countryCode does not exist in the list, returns an error `ErrCountryCodeDoesNotExist`
func GetPaymentCardConfiguration(issuer string) (*paymentCardConfiguration, error) {
	if conf, ok := paymentCardConfigurations[issuer]; ok {
		return &conf, nil
	}
	return nil, ErrIssuerDoesNotExist
}
