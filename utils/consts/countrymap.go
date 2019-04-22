package consts

import (
	"github.com/vdlbk/mimir/models"
)

var (
	CountriesConfiguration = map[string]models.Configuration{
		AD.String(): andorraConfiguration,
		AE.String(): unitedArabEmiratesConfiguration,
		AL.String(): albaniaConfiguration,
		AT.String(): austriaConfiguration,
		AZ.String(): azerbaijanConfiguration,
		BA.String(): bosniaHerzegovinaConfiguration,
		BE.String(): belgiumConfiguration,
		BG.String(): bulgariaConfiguration,
		BH.String(): bahrainConfiguration,
		BR.String(): brazilConfiguration,
		BY.String(): republicBelarusConfiguration,
		CH.String(): switzerlandConfiguration,
		CR.String(): costaRicaConfiguration,
		CY.String(): cyprusConfiguration,
		CZ.String(): czechRepublicConfiguration,
		DE.String(): germanyConfiguration,
		DK.String(): denmarkConfiguration,
		DO.String(): dominicanRepublicConfiguration,
		EE.String(): estoniaConfiguration,
		ES.String(): spainConfiguration,
		FI.String(): finlandConfiguration,
		FO.String(): faroeIslandsConfiguration,
		FR.String(): franceConfiguration,
		GB.String(): unitedKingdomConfiguration,
		GE.String(): georgiaConfiguration,
		GI.String(): gibraltarConfiguration,
		GL.String(): greenlandConfiguration,
		GR.String(): greeceConfiguration,
		GT.String(): guatemalaConfiguration,
		HR.String(): croatiaConfiguration,
		HU.String(): hungaryConfiguration,
		IE.String(): irelandConfiguration,
		IL.String(): israelConfiguration,
		IQ.String(): iraqConfiguration,
		IS.String(): icelandConfiguration,
		IT.String(): italyConfiguration,
		JO.String(): jordanConfiguration,
		KW.String(): kuwaitConfiguration,
		KZ.String(): kazakhstanConfiguration,
		LB.String(): lebanonConfiguration,
		LC.String(): saintLuciaConfiguration,
		LI.String(): liechtensteinConfiguration,
		LT.String(): lithuaniaConfiguration,
		LU.String(): luxembourgConfiguration,
		LV.String(): latviaConfiguration,
		MC.String(): monacoConfiguration,
		MD.String(): moldovaConfiguration,
		ME.String(): montenegroConfiguration,
		MK.String(): macedoniaConfiguration,
		MR.String(): mauritaniaConfiguration,
		MT.String(): maltaConfiguration,
		MU.String(): mauritiusConfiguration,
		NL.String(): netherlandsConfiguration,
		NO.String(): norwayConfiguration,
		PK.String(): pakistanConfiguration,
		PL.String(): polandConfiguration,
		PS.String(): palestineConfiguration,
		PT.String(): portugalConfiguration,
		QA.String(): qatarConfiguration,
		RO.String(): romaniaConfiguration,
		RS.String(): serbiaConfiguration,
		SA.String(): saudiArabiaConfiguration,
		SC.String(): seychellesConfiguration,
		SE.String(): swedenConfiguration,
		SI.String(): sloveniaConfiguration,
		SK.String(): slovakiaConfiguration,
		SM.String(): sanMarinoConfiguration,
		ST.String(): saoTomePrincipeConfiguration,
		SV.String(): elSalvadorConfiguration,
		TL.String(): timorLesteConfiguration,
		TN.String(): tunisiaConfiguration,
		TR.String(): turkeyConfiguration,
		UA.String(): ukraineConfiguration,
		VA.String(): vaticanCityStateConfiguration,
		VG.String(): virginIslandsConfiguration,
		XK.String(): kosovoConfiguration,
	}

	andorraConfiguration = models.Configuration{
		CountryName:             "Andorra",
		CountryCode:             AD,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "2030200359100100",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "AD1200012030200359100100",
			PrintFormat: "AD12 0001 2030 2003 5910 0100",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "00012030200359100100",
		},
	}

	unitedArabEmiratesConfiguration = models.Configuration{
		CountryName:             "The United Arab Emirates",
		CountryCode:             AE,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "1234567890123456",
		IBANDefinition: models.Definition{
			Length:      23,
			Example:     "AE070331234567890123456",
			PrintFormat: "AE07 0331 2345 6789 0123 456",
		},
		BBANDefinition: models.Definition{
			Length:  19,
			Example: "0331234567890123456",
		},
	}

	albaniaConfiguration = models.Configuration{
		CountryName:             "Albania",
		CountryCode:             AL,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0000000235698741",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "AL47212110090000000235698741",
			PrintFormat: "AL47 2121 1009 0000 0002 3569 8741",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "212110090000000235698741",
		},
	}

	austriaConfiguration = models.Configuration{
		CountryName:             "Austria",
		CountryCode:             AT,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "BLZ 19043 Kto 234573201",
		IBANDefinition: models.Definition{
			Length:      20,
			Example:     "AT611904300234573201",
			PrintFormat: "AT61 1904 3002 3457 3201",
		},
		BBANDefinition: models.Definition{
			Length:  16,
			Example: "1904300234573201",
		},
	}

	azerbaijanConfiguration = models.Configuration{
		CountryName:             "Azerbaijan",
		CountryCode:             AZ,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "NABZ00000000137010001944",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "AZ21NABZ00000000137010001944",
			PrintFormat: "AZ21 NABZ 0000 0000 1370 1000 1944",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "NABZ00000000137010001944",
		},
	}

	bosniaHerzegovinaConfiguration = models.Configuration{
		CountryName:             "Bosnia and Herzegovina",
		CountryCode:             BA,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "199-044-00012002-79",
		IBANDefinition: models.Definition{
			Length:      20,
			Example:     "BA391290079401028494",
			PrintFormat: "BA39 1290 0794 0102 8494",
		},
		BBANDefinition: models.Definition{
			Length:  16,
			Example: "1990440001200279",
		},
	}

	belgiumConfiguration = models.Configuration{
		CountryName:             "Belgium",
		CountryCode:             BE,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "539-0075470-34",
		IBANDefinition: models.Definition{
			Length:      16,
			Example:     "BE68539007547034",
			PrintFormat: "BE68 5390 0754 7034",
		},
		BBANDefinition: models.Definition{
			Length:  12,
			Example: "539007547034",
		},
	}

	bulgariaConfiguration = models.Configuration{
		CountryName:             "Bulgaria",
		CountryCode:             BG,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "BNBG 9661 1020 3456 78",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "BG80BNBG96611020345678",
			PrintFormat: "BG80 BNBG 9661 1020 3456 78",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "BNBG96611020345678",
		},
	}

	bahrainConfiguration = models.Configuration{
		CountryName:             "Bahrain",
		CountryCode:             BH,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "00001299123456",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "BH67BMAG00001299123456",
			PrintFormat: "BH67 BMAG 0000 1299 1234 56",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "BMAG00001299123456",
		},
	}

	brazilConfiguration = models.Configuration{
		CountryName:             "Brazil",
		CountryCode:             BR,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0009795493C1",
		IBANDefinition: models.Definition{
			Length:      29,
			Example:     "BR1800360305000010009795493C1",
			PrintFormat: "BR18 0036 0305 0000 1000 9795 493C 1",
		},
		BBANDefinition: models.Definition{
			Length:  25,
			Example: "00360305000010009795493P1",
		},
	}

	republicBelarusConfiguration = models.Configuration{
		CountryName:             "Republic of Belarus",
		CountryCode:             BY,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "3600 0000 0000 0Z00 AB00",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "BY13NBRB3600900000002Z00AB00",
			PrintFormat: "BY13 NBRB 3600 9000 0000 2Z00 AB00",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "NBRB3600000000000Z00AB00",
		},
	}

	switzerlandConfiguration = models.Configuration{
		CountryName:             "Switzerland",
		CountryCode:             CH,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "762 1162-3852.957",
		IBANDefinition: models.Definition{
			Length:      21,
			Example:     "CH9300762011623852957",
			PrintFormat: "CH93 0076 2011 6238 5295 7",
		},
		BBANDefinition: models.Definition{
			Length:  17,
			Example: "00762011623852957",
		},
	}

	costaRicaConfiguration = models.Configuration{
		CountryName:             "Costa Rica",
		CountryCode:             CR,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "02001026284066",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "CR05015202001026284066",
			PrintFormat: "CR05 0152 0200 1026 2840 66",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "015202001026284066",
		},
	}

	cyprusConfiguration = models.Configuration{
		CountryName:             "Cyprus",
		CountryCode:             CY,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0000001200527600",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "CY17002001280000001200527600",
			PrintFormat: "CY17 0020 0128 0000 0012 0052 7600",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "002001280000001200527600",
		},
	}

	czechRepublicConfiguration = models.Configuration{
		CountryName:             "Czech Republic",
		CountryCode:             CZ,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "19-2000145399/0800",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "CZ6508000000192000145399",
			PrintFormat: "CZ65 0800 0000 1920 0014 5399",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "08000000192000145399",
		},
	}

	germanyConfiguration = models.Configuration{
		CountryName:             "Germany",
		CountryCode:             DE,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "532013000",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "DE89370400440532013000",
			PrintFormat: "DE89 3704 0044 0532 0130 00",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "370400440532013000",
		},
	}

	denmarkConfiguration = models.Configuration{
		CountryName:             "Denmark",
		CountryCode:             DK,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0040 0440116243",
		IBANDefinition: models.Definition{
			Length:      18,
			Example:     "DK5000400440116243",
			PrintFormat: "DK50 0040 0440 1162 43",
		},
		BBANDefinition: models.Definition{
			Length:  14,
			Example: "00400440116243",
		},
	}

	dominicanRepublicConfiguration = models.Configuration{
		CountryName:             "Dominican Republic",
		CountryCode:             DO,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "00000001212453611324",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "DO28BAGR00000001212453611324",
			PrintFormat: "DO28 BAGR 0000 0001 2124 5361 1324",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "BAGR00000001212453611324",
		},
	}

	estoniaConfiguration = models.Configuration{
		CountryName:             "Estonia",
		CountryCode:             EE,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "221020145685",
		IBANDefinition: models.Definition{
			Length:      20,
			Example:     "EE382200221020145685",
			PrintFormat: "EE38 2200 2210 2014 5685",
		},
		BBANDefinition: models.Definition{
			Length:  16,
			Example: "2200221020145685",
		},
	}

	spainConfiguration = models.Configuration{
		CountryName:             "Spain",
		CountryCode:             ES,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "2100 0418 45 0200051332",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "ES9121000418450200051332",
			PrintFormat: "ES91 2100 0418 4502 0005 1332",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "21000418450200051332",
		},
	}

	finlandConfiguration = models.Configuration{
		CountryName:             "Finland",
		CountryCode:             FI,
		IncludedCountryCode:     []models.CountryCode{AX},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{AX},
		AccountNumberExample:    "1234 5600 0007 85",
		IBANDefinition: models.Definition{
			Length:      18,
			Example:     "FI2112345600000785",
			PrintFormat: "FI21 1234 5600 0007 85",
		},
		BBANDefinition: models.Definition{
			Length:  14,
			Example: "12345600000785",
		},
	}

	faroeIslandsConfiguration = models.Configuration{
		CountryName:             "Faroe Islands",
		CountryCode:             FO,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "6460 0001631634",
		IBANDefinition: models.Definition{
			Length:      18,
			Example:     "FO6264600001631634",
			PrintFormat: "FO62 6460 0001 6316 34",
		},
		BBANDefinition: models.Definition{
			Length:  14,
			Example: "64600001631634",
		},
	}

	franceConfiguration = models.Configuration{
		CountryName:             "France",
		CountryCode:             FR,
		IncludedCountryCode:     []models.CountryCode{GF, GP, MQ, RE, PF, TF, YT, NC, BL, MF, PM, WF},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{GF, GP, MQ, YT, RE, PM, BL, MF},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: models.Definition{
			Length:      27,
			Example:     "FR1420041010050500013M02606",
			PrintFormat: "FR14 2004 1010 0505 0001 3M02 606",
		},
		BBANDefinition: models.Definition{
			Length:  23,
			Example: "20041010050500013M02606",
		},
	}

	unitedKingdomConfiguration = models.Configuration{
		CountryName:             "United Kingdom",
		CountryCode:             GB,
		IncludedCountryCode:     []models.CountryCode{IM, JE, GG},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "60-16-13 31926819",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "GB29NWBK60161331926819",
			PrintFormat: "GB29 NWBK 6016 1331 9268 19",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "NWBK60161331926819",
		},
	}

	georgiaConfiguration = models.Configuration{
		CountryName:             "Georgia",
		CountryCode:             GE,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0000000101904917",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "GE29NB0000000101904917",
			PrintFormat: "GE29 NB00 0000 0101 9049 17",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "NB0000000101904917",
		},
	}

	gibraltarConfiguration = models.Configuration{
		CountryName:             "Gibraltar",
		CountryCode:             GI,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0000 00007099 453",
		IBANDefinition: models.Definition{
			Length:      23,
			Example:     "GI75NWBK000000007099453",
			PrintFormat: "GI75 NWBK 0000 0000 7099 453",
		},
		BBANDefinition: models.Definition{
			Length:  19,
			Example: "NWBK000000007099453",
		},
	}

	greenlandConfiguration = models.Configuration{
		CountryName:             "Greenland",
		CountryCode:             GL,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "6471 0001000206",
		IBANDefinition: models.Definition{
			Length:      18,
			Example:     "GL8964710001000206",
			PrintFormat: "GL89 6471 0001 0002 06",
		},
		BBANDefinition: models.Definition{
			Length:  14,
			Example: "64710001000206",
		},
	}

	greeceConfiguration = models.Configuration{
		CountryName:             "Greece",
		CountryCode:             GR,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "01250000000012300695",
		IBANDefinition: models.Definition{
			Length:      27,
			Example:     "GR1601101250000000012300695",
			PrintFormat: "GR16 0110 1250 0000 0001 2300 695",
		},
		BBANDefinition: models.Definition{
			Length:  23,
			Example: "01101250000000012300695",
		},
	}

	guatemalaConfiguration = models.Configuration{
		CountryName:             "Guatemala",
		CountryCode:             GT,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "01020000001210029690",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "GT82TRAJ01020000001210029690",
			PrintFormat: "GT82 TRAJ 0102 0000 0012 1002 9690",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "TRAJ01020000001210029690",
		},
	}

	croatiaConfiguration = models.Configuration{
		CountryName:             "Croatia",
		CountryCode:             HR,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "1001005-1863000160",
		IBANDefinition: models.Definition{
			Length:      21,
			Example:     "HR1210010051863000160",
			PrintFormat: "HR12 1001 0051 8630 0016 0",
		},
		BBANDefinition: models.Definition{
			Length:  17,
			Example: "10010051863000160",
		},
	}

	hungaryConfiguration = models.Configuration{
		CountryName:             "Hungary",
		CountryCode:             HU,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "11773016-11111018-000000009",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "HU42117730161111101800000000",
			PrintFormat: "HU42 1177 3016 1111 1018 0000 0000",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "117730161111101800000000",
		},
	}

	irelandConfiguration = models.Configuration{
		CountryName:             "Ireland",
		CountryCode:             IE,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "93-11-52 12345678",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "IE29AIBK93115212345678",
			PrintFormat: "IE29 AIBK 9311 5212 3456 78",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "AIBK93115212345678",
		},
	}

	israelConfiguration = models.Configuration{
		CountryName:             "Israel",
		CountryCode:             IL,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "10-800-99999999",
		IBANDefinition: models.Definition{
			Length:      23,
			Example:     "IL620108000000099999999",
			PrintFormat: "IL62 0108 0000 0009 9999 999",
		},
		BBANDefinition: models.Definition{
			Length:  19,
			Example: "0108000000099999999",
		},
	}

	iraqConfiguration = models.Configuration{
		CountryName:             "Iraq",
		CountryCode:             IQ,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "123456789012",
		IBANDefinition: models.Definition{
			Length:      23,
			Example:     "IQ98NBIQ850123456789012",
			PrintFormat: "IQ98 NBIQ 8501 2345 6789 012",
		},
		BBANDefinition: models.Definition{
			Length:  19,
			Example: "NBIQ850123456789012",
		},
	}

	icelandConfiguration = models.Configuration{
		CountryName:             "Iceland",
		CountryCode:             IS,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0159-26-007654-551073-0339",
		IBANDefinition: models.Definition{
			Length:      26,
			Example:     "IS140159260076545510730339",
			PrintFormat: "IS14 0159 2600 7654 5510 7303 39",
		},
		BBANDefinition: models.Definition{
			Length:  22,
			Example: "0159260076545510730339",
		},
	}

	italyConfiguration = models.Configuration{
		CountryName:             "Italy",
		CountryCode:             IT,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "X 05428 11101 000000123456",
		IBANDefinition: models.Definition{
			Length:      27,
			Example:     "IT60X0542811101000000123456",
			PrintFormat: "IT60 X054 2811 1010 0000 0123 456",
		},
		BBANDefinition: models.Definition{
			Length:  23,
			Example: "X0542811101000000123456",
		},
	}

	jordanConfiguration = models.Configuration{
		CountryName:             "Jordan",
		CountryCode:             JO,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0001310000302",
		IBANDefinition: models.Definition{
			Length:      30,
			Example:     "JO94CBJO0010000000000131000302",
			PrintFormat: "JO94 CBJO 0010 0000 0000 0131 0003 02",
		},
		BBANDefinition: models.Definition{
			Length:  26,
			Example: "CBJO0010000000000131000302",
		},
	}

	kuwaitConfiguration = models.Configuration{
		CountryName:             "Kuwait",
		CountryCode:             KW,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "1234560101",
		IBANDefinition: models.Definition{
			Length:      30,
			Example:     "KW81CBKU0000000000001234560101",
			PrintFormat: "KW81 CBKU 0000 0000 0000 1234 5601 01",
		},
		BBANDefinition: models.Definition{
			Length:  26,
			Example: "CBKU0000000000001234560101",
		},
	}

	kazakhstanConfiguration = models.Configuration{
		CountryName:             "Kazakhstan",
		CountryCode:             KZ,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "KZ86 125K ZT50 0410 0100",
		IBANDefinition: models.Definition{
			Length:      20,
			Example:     "KZ86125KZT5004100100",
			PrintFormat: "KZ86 125K ZT50 0410 0100",
		},
		BBANDefinition: models.Definition{
			Length:  16,
			Example: "125KZT5004100100",
		},
	}

	lebanonConfiguration = models.Configuration{
		CountryName:             "Lebanon",
		CountryCode:             LB,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "01 001 901229114",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "LB62099900000001001901229114",
			PrintFormat: "LB62 0999 0000 0001 0019 0122 9114",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "099900000001001901229114",
		},
	}

	saintLuciaConfiguration = models.Configuration{
		CountryName:             "Saint Lucia",
		CountryCode:             LC,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0001 0001 0012 0012 0002 3015",
		IBANDefinition: models.Definition{
			Length:      32,
			Example:     "LC55HEMM000100010012001200023015",
			PrintFormat: "LC55 HEMM 0001 0001 0012 0012 0002 3015",
		},
		BBANDefinition: models.Definition{
			Length:  28,
			Example: "HEMM000100010012001200023015",
		},
	}

	liechtensteinConfiguration = models.Configuration{
		CountryName:             "Liechtenstein",
		CountryCode:             LI,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "8810 2324013AA",
		IBANDefinition: models.Definition{
			Length:      21,
			Example:     "LI21088100002324013AA",
			PrintFormat: "LI21 0881 0000 2324 013A A",
		},
		BBANDefinition: models.Definition{
			Length:  17,
			Example: "088100002324013AA",
		},
	}

	lithuaniaConfiguration = models.Configuration{
		CountryName:             "Lithuania",
		CountryCode:             LT,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "11101001000",
		IBANDefinition: models.Definition{
			Length:      20,
			Example:     "LT121000011101001000",
			PrintFormat: "LT12 1000 0111 0100 1000",
		},
		BBANDefinition: models.Definition{
			Length:  16,
			Example: "1000011101001000",
		},
	}

	luxembourgConfiguration = models.Configuration{
		CountryName:             "Luxembourg",
		CountryCode:             LU,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "9400644750000",
		IBANDefinition: models.Definition{
			Length:      20,
			Example:     "LU280019400644750000",
			PrintFormat: "LU28 0019 4006 4475 0000",
		},
		BBANDefinition: models.Definition{
			Length:  16,
			Example: "0019400644750000",
		},
	}

	latviaConfiguration = models.Configuration{
		CountryName:             "Latvia",
		CountryCode:             LV,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "LV80 BANK 0000 4351 9500 1",
		IBANDefinition: models.Definition{
			Length:      21,
			Example:     "LV80BANK0000435195001",
			PrintFormat: "LV80 BANK 0000 4351 9500 1",
		},
		BBANDefinition: models.Definition{
			Length:  17,
			Example: "BANK0000435195001",
		},
	}

	monacoConfiguration = models.Configuration{
		CountryName:             "Monaco",
		CountryCode:             MC,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0011111000h",
		IBANDefinition: models.Definition{
			Length:      27,
			Example:     "MC5811222000010123456789030",
			PrintFormat: "MC58 1122 2000 0101 2345 6789 030",
		},
		BBANDefinition: models.Definition{
			Length:  23,
			Example: "11222000010123456789030",
		},
	}

	moldovaConfiguration = models.Configuration{
		CountryName:             "Moldova",
		CountryCode:             MD,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "000225100013104168",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "MD24AG000225100013104168",
			PrintFormat: "MD24 AG00 0225 1000 1310 4168",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "AG000225100013104168",
		},
	}

	montenegroConfiguration = models.Configuration{
		CountryName:             "Montenegro",
		CountryCode:             ME,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "505 0000123456789 51",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "ME25505000012345678951",
			PrintFormat: "ME25 5050 0001 2345 6789 51",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "505000012345678951",
		},
	}

	macedoniaConfiguration = models.Configuration{
		CountryName:             "Macedonia",
		CountryCode:             MK,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "MK07 300 0000000424 25",
		IBANDefinition: models.Definition{
			Length:      19,
			Example:     "MK07250120000058984",
			PrintFormat: "MK07 2501 2000 0058 984",
		},
		BBANDefinition: models.Definition{
			Length:  15,
			Example: "250120000058984",
		},
	}

	mauritaniaConfiguration = models.Configuration{
		CountryName:             "Mauritania",
		CountryCode:             MR,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "00020 00101 00001234567 53",
		IBANDefinition: models.Definition{
			Length:      27,
			Example:     "MR1300020001010000123456753",
			PrintFormat: "MR13 0002 0001 0100 0012 3456 753",
		},
		BBANDefinition: models.Definition{
			Length:  23,
			Example: "00020001010000123456753",
		},
	}

	maltaConfiguration = models.Configuration{
		CountryName:             "Malta",
		CountryCode:             MT,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "12345MTLCAST001S",
		IBANDefinition: models.Definition{
			Length:      31,
			Example:     "MT84MALT011000012345MTLCAST001S",
			PrintFormat: "MT84 MALT 0110 0001 2345 MTLC AST0 01S",
		},
		BBANDefinition: models.Definition{
			Length:  27,
			Example: "MALT011000012345MTLCAST001S",
		},
	}

	mauritiusConfiguration = models.Configuration{
		CountryName:             "Mauritius",
		CountryCode:             MU,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "MU17 BOMM 0101 1010 3030 0200 000M UR",
		IBANDefinition: models.Definition{
			Length:      30,
			Example:     "MU17BOMM0101101030300200000MUR",
			PrintFormat: "MU17 BOMM 0101 1010 3030 0200 000M UR",
		},
		BBANDefinition: models.Definition{
			Length:  26,
			Example: "BOMM0101101030300200000MUR",
		},
	}

	netherlandsConfiguration = models.Configuration{
		CountryName:             "The Netherlands",
		CountryCode:             NL,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "041 71 64 300",
		IBANDefinition: models.Definition{
			Length:      18,
			Example:     "NL91ABNA0417164300",
			PrintFormat: "NL91 ABNA 0417 1643 00",
		},
		BBANDefinition: models.Definition{
			Length:  14,
			Example: "ABNA0417164300",
		},
	}

	norwayConfiguration = models.Configuration{
		CountryName:             "Norway",
		CountryCode:             NO,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "8601 11 17947",
		IBANDefinition: models.Definition{
			Length:      15,
			Example:     "NO9386011117947",
			PrintFormat: "NO93 8601 1117 947",
		},
		BBANDefinition: models.Definition{
			Length:  11,
			Example: "86011117947",
		},
	}

	pakistanConfiguration = models.Configuration{
		CountryName:             "Pakistan",
		CountryCode:             PK,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "00260101036360",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "PK36SCBL0000001123456702",
			PrintFormat: "PK36 SCBL 0000 0011 2345 6702",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "SCBL0000001123456702",
		},
	}

	polandConfiguration = models.Configuration{
		CountryName:             "Poland",
		CountryCode:             PL,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "61 1090 1014 0000 0712 1981 2874",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "PL61109010140000071219812874",
			PrintFormat: "PL61 1090 1014 0000 0712 1981 2874",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "109010140000071219812874",
		},
	}

	palestineConfiguration = models.Configuration{
		CountryName:             "State of Palestine",
		CountryCode:             PS,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "400123456702",
		IBANDefinition: models.Definition{
			Length:      29,
			Example:     "PS92PALS000000000400123456702",
			PrintFormat: "PS92 PALS 0000 0000 0400 1234 5670 2",
		},
		BBANDefinition: models.Definition{
			Length:  25,
			Example: "PALS000000000400123456702",
		},
	}

	portugalConfiguration = models.Configuration{
		CountryName:             "Portugal",
		CountryCode:             PT,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{}, // Azores, Madeira does not have proper Country Code
		AccountNumberExample:    "0002.0123.12345678901.54",
		IBANDefinition: models.Definition{
			Length:      25,
			Example:     "PT50000201231234567890154",
			PrintFormat: "PT50 0002 0123 1234 5678 9015 4",
		},
		BBANDefinition: models.Definition{
			Length:  21,
			Example: "000201231234567890154",
		},
	}

	qatarConfiguration = models.Configuration{
		CountryName:             "Qatar",
		CountryCode:             QA,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "00001234567890ABCDEFG",
		IBANDefinition: models.Definition{
			Length:      29,
			Example:     "QA58DOHB00001234567890ABCDEFG",
			PrintFormat: "QA58 DOHB 0000 1234 5678 90AB CDEF G",
		},
		BBANDefinition: models.Definition{
			Length:  25,
			Example: "DOHB00001234567890ABCDEFG",
		},
	}

	romaniaConfiguration = models.Configuration{
		CountryName:             "Romania",
		CountryCode:             RO,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "RO49 AAAA 1B31 0075 9384 0000",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "RO49AAAA1B31007593840000",
			PrintFormat: "RO49 AAAA 1B31 0075 9384 0000",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "AAAA1B31007593840000",
		},
	}

	serbiaConfiguration = models.Configuration{
		CountryName:             "Serbia",
		CountryCode:             RS,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "260-0056010016113-79",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "RS35260005601001611379",
			PrintFormat: "RS35 2600 0560 1001 6113 79",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "260005601001611379",
		},
	}

	saudiArabiaConfiguration = models.Configuration{
		CountryName:             "Saudi Arabia",
		CountryCode:             SA,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "608010167519",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "SA0380000000608010167519",
			PrintFormat: "SA03 8000 0000 6080 1016 7519",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "80000000608010167519",
		},
	}

	seychellesConfiguration = models.Configuration{
		CountryName:             "Seychelles",
		CountryCode:             SC,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "0000000000001497",
		IBANDefinition: models.Definition{
			Length:      31,
			Example:     "SC18SSCB11010000000000001497USD",
			PrintFormat: "SC18 SSCB 1101 0000 0000 0000 1497 USD",
		},
		BBANDefinition: models.Definition{
			Length:  27,
			Example: "SSCB11010000000000001497USD",
		},
	}

	swedenConfiguration = models.Configuration{
		CountryName:             "Sweden",
		CountryCode:             SE,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "1234 12 3456 1",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "SE4550000000058398257466",
			PrintFormat: "SE45 5000 0000 0583 9825 7466",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "50000000058398257466",
		},
	}

	sloveniaConfiguration = models.Configuration{
		CountryName:             "Slovenia",
		CountryCode:             SI,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "2633 0001 2039 086",
		IBANDefinition: models.Definition{
			Length:      19,
			Example:     "SI56263300012039086",
			PrintFormat: "SI56 2633 0001 2039 086",
		},
		BBANDefinition: models.Definition{
			Length:  15,
			Example: "263300012039086",
		},
	}

	slovakiaConfiguration = models.Configuration{
		CountryName:             "Slovakia",
		CountryCode:             SK,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "19-8742637541/1200",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "SK3112000000198742637541",
			PrintFormat: "SK31 1200 0000 1987 4263 7541",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "12000000198742637541",
		},
	}

	sanMarinoConfiguration = models.Configuration{
		CountryName:             "San Marino",
		CountryCode:             SM,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "U0322509800000000270100",
		IBANDefinition: models.Definition{
			Length:      27,
			Example:     "SM86U0322509800000000270100",
			PrintFormat: "SM86 U032 2509 8000 0000 0270 100",
		},
		BBANDefinition: models.Definition{
			Length:  23,
			Example: "U0322509800000000270100",
		},
	}

	saoTomePrincipeConfiguration = models.Configuration{
		CountryName:             "Sao Tome and Principe",
		CountryCode:             ST,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "01921942101.12",
		IBANDefinition: models.Definition{
			Length:      25,
			Example:     "ST32000200010192194210112",
			PrintFormat: "ST32 0002 0001 0192 1942 1011 2",

			// Given the Swift IBAN Registryn the following IBAN should be valid but they're not. The correct checksum is 32, not 68.
			//Example:     "ST68000200010192194210112",
			//PrintFormat: "ST68 0002 0001 0192 1942 1011 2",
		},
		BBANDefinition: models.Definition{
			Length:  21,
			Example: "000200010192194210112",
		},
	}

	elSalvadorConfiguration = models.Configuration{
		CountryName:             "El Salvador",
		CountryCode:             SV,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "00000000000000700025",
		IBANDefinition: models.Definition{
			Length:      28,
			Example:     "SV62CENR00000000000000700025",
			PrintFormat: "SV 62 CENR 00000000000000700025",
		},
		BBANDefinition: models.Definition{
			Length:  24,
			Example: "CENR00000000000000700025",
		},
	}

	timorLesteConfiguration = models.Configuration{
		CountryName:             "Timor-Leste",
		CountryCode:             TL,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "008 00123456789101 57",
		IBANDefinition: models.Definition{
			Length:      23,
			Example:     "TL380080012345678910157",
			PrintFormat: "TL38 0080 0123 4567 8910 157",
		},
		BBANDefinition: models.Definition{
			Length:  19,
			Example: "0080012345678910157",
		},
	}

	tunisiaConfiguration = models.Configuration{
		CountryName:             "Tunisia",
		CountryCode:             TN,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "10 006 0351835984788 31",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "TN5910006035183598478831",
			PrintFormat: "TN59 1000 6035 1835 9847 8831",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "10006035183598478831",
		},
	}

	turkeyConfiguration = models.Configuration{
		CountryName:             "Turkey",
		CountryCode:             TR,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "00519786457841326",
		IBANDefinition: models.Definition{
			Length:      26,
			Example:     "TR330006100519786457841326",
			PrintFormat: "TR33 0006 1005 1978 6457 8413 26",
		},
		BBANDefinition: models.Definition{
			Length:  22,
			Example: "0006100519786457841326",
		},
	}

	ukraineConfiguration = models.Configuration{
		CountryName:             "Ukraine",
		CountryCode:             UA,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "26007233566001",
		IBANDefinition: models.Definition{
			Length:      29,
			Example:     "UA213223130000026007233566001",
			PrintFormat: "UA21 3223 1300 0002 6007 2335 6600 1",
		},
		BBANDefinition: models.Definition{
			Length:  25,
			Example: "3223130000026007233566001",
		},
	}

	vaticanCityStateConfiguration = models.Configuration{
		CountryName:             "Vatican City State",
		CountryCode:             VA,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "123000012345678",
		IBANDefinition: models.Definition{
			Length:      22,
			Example:     "VA59001123000012345678",
			PrintFormat: "VA59 001 1230 0001 2345 678",
		},
		BBANDefinition: models.Definition{
			Length:  18,
			Example: "001123000012345678",
		},
	}

	virginIslandsConfiguration = models.Configuration{
		CountryName:             "Virgin Islands",
		CountryCode:             VG,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "00000 12 345 678 901",
		IBANDefinition: models.Definition{
			Length:      24,
			Example:     "VG96VPVG0000012345678901",
			PrintFormat: "VG96 VPVG 0000 0123 4567 8901",
		},
		BBANDefinition: models.Definition{
			Length:  20,
			Example: "VPVG0000012345678901",
		},
	}

	kosovoConfiguration = models.Configuration{
		CountryName:             "Kosovo",
		CountryCode:             XK,
		IncludedCountryCode:     []models.CountryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []models.CountryCode{},
		AccountNumberExample:    "1212 0123456789 06",
		IBANDefinition: models.Definition{
			Length:      20,
			Example:     "XK051212012345678906",
			PrintFormat: "XK05 1212 0123 4567 8906",
		},
		BBANDefinition: models.Definition{
			Length:  16,
			Example: "1212012345678906",
		},
	}
)
