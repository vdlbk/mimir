package mimir

var (
	// countriesConfiguration holds the mapping between the country code and the country specifications
	countriesConfiguration = map[string]configuration{
		ad.String(): andorraConfiguration,
		ae.String(): unitedArabEmiratesConfiguration,
		al.String(): albaniaConfiguration,
		at.String(): austriaConfiguration,
		az.String(): azerbaijanConfiguration,
		ba.String(): bosniaHerzegovinaConfiguration,
		be.String(): belgiumConfiguration,
		bg.String(): bulgariaConfiguration,
		bh.String(): bahrainConfiguration,
		br.String(): brazilConfiguration,
		by.String(): republicBelarusConfiguration,
		ch.String(): switzerlandConfiguration,
		cr.String(): costaRicaConfiguration,
		cy.String(): cyprusConfiguration,
		cz.String(): czechRepublicConfiguration,
		de.String(): germanyConfiguration,
		dk.String(): denmarkConfiguration,
		do.String(): dominicanRepublicConfiguration,
		ee.String(): estoniaConfiguration,
		es.String(): spainConfiguration,
		fi.String(): finlandConfiguration,
		fo.String(): faroeIslandsConfiguration,
		fr.String(): franceConfiguration,
		gb.String(): unitedKingdomConfiguration,
		ge.String(): georgiaConfiguration,
		gi.String(): gibraltarConfiguration,
		gl.String(): greenlandConfiguration,
		gr.String(): greeceConfiguration,
		gt.String(): guatemalaConfiguration,
		hr.String(): croatiaConfiguration,
		hu.String(): hungaryConfiguration,
		ie.String(): irelandConfiguration,
		il.String(): israelConfiguration,
		iq.String(): iraqConfiguration,
		is.String(): icelandConfiguration,
		it.String(): italyConfiguration,
		jo.String(): jordanConfiguration,
		kw.String(): kuwaitConfiguration,
		kz.String(): kazakhstanConfiguration,
		lb.String(): lebanonConfiguration,
		lc.String(): saintLuciaConfiguration,
		li.String(): liechtensteinConfiguration,
		lt.String(): lithuaniaConfiguration,
		lu.String(): luxembourgConfiguration,
		lv.String(): latviaConfiguration,
		mc.String(): monacoConfiguration,
		md.String(): moldovaConfiguration,
		me.String(): montenegroConfiguration,
		mk.String(): macedoniaConfiguration,
		mr.String(): mauritaniaConfiguration,
		mt.String(): maltaConfiguration,
		mu.String(): mauritiusConfiguration,
		nl.String(): netherlandsConfiguration,
		no.String(): norwayConfiguration,
		pk.String(): pakistanConfiguration,
		pl.String(): polandConfiguration,
		ps.String(): palestineConfiguration,
		pt.String(): portugalConfiguration,
		qa.String(): qatarConfiguration,
		ro.String(): romaniaConfiguration,
		rs.String(): serbiaConfiguration,
		sa.String(): saudiArabiaConfiguration,
		sc.String(): seychellesConfiguration,
		se.String(): swedenConfiguration,
		si.String(): sloveniaConfiguration,
		sk.String(): slovakiaConfiguration,
		sm.String(): sanMarinoConfiguration,
		st.String(): saoTomePrincipeConfiguration,
		sv.String(): elSalvadorConfiguration,
		tl.String(): timorLesteConfiguration,
		tn.String(): tunisiaConfiguration,
		tr.String(): turkeyConfiguration,
		ua.String(): ukraineConfiguration,
		va.String(): vaticanCityStateConfiguration,
		vg.String(): virginIslandsConfiguration,
		xk.String(): kosovoConfiguration,
	}

	andorraConfiguration = configuration{
		CountryName:             "Andorra",
		CountryCode:             ad,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "2030200359100100",
		IBANDefinition: definition{
			Length:      24,
			Example:     "AD1200012030200359100100",
			PrintFormat: "AD12 0001 2030 2003 5910 0100",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "00012030200359100100",
		},
	}

	unitedArabEmiratesConfiguration = configuration{
		CountryName:             "The United Arab Emirates",
		CountryCode:             ae,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "1234567890123456",
		IBANDefinition: definition{
			Length:      23,
			Example:     "AE070331234567890123456",
			PrintFormat: "AE07 0331 2345 6789 0123 456",
		},
		BBANDefinition: definition{
			Length:  19,
			Example: "0331234567890123456",
		},
	}

	albaniaConfiguration = configuration{
		CountryName:             "Albania",
		CountryCode:             al,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0000000235698741",
		IBANDefinition: definition{
			Length:      28,
			Example:     "AL47212110090000000235698741",
			PrintFormat: "AL47 2121 1009 0000 0002 3569 8741",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "212110090000000235698741",
		},
	}

	austriaConfiguration = configuration{
		CountryName:             "Austria",
		CountryCode:             at,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "BLZ 19043 Kto 234573201",
		IBANDefinition: definition{
			Length:      20,
			Example:     "AT611904300234573201",
			PrintFormat: "AT61 1904 3002 3457 3201",
		},
		BBANDefinition: definition{
			Length:  16,
			Example: "1904300234573201",
		},
	}

	azerbaijanConfiguration = configuration{
		CountryName:             "Azerbaijan",
		CountryCode:             az,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "NABZ00000000137010001944",
		IBANDefinition: definition{
			Length:      28,
			Example:     "AZ21NABZ00000000137010001944",
			PrintFormat: "AZ21 NABZ 0000 0000 1370 1000 1944",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "NABZ00000000137010001944",
		},
	}

	bosniaHerzegovinaConfiguration = configuration{
		CountryName:             "Bosnia and Herzegovina",
		CountryCode:             ba,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "199-044-00012002-79",
		IBANDefinition: definition{
			Length:      20,
			Example:     "BA391290079401028494",
			PrintFormat: "BA39 1290 0794 0102 8494",
		},
		BBANDefinition: definition{
			Length:  16,
			Example: "1990440001200279",
		},
	}

	belgiumConfiguration = configuration{
		CountryName:             "Belgium",
		CountryCode:             be,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "539-0075470-34",
		IBANDefinition: definition{
			Length:      16,
			Example:     "BE68539007547034",
			PrintFormat: "BE68 5390 0754 7034",
		},
		BBANDefinition: definition{
			Length:  12,
			Example: "539007547034",
		},
	}

	bulgariaConfiguration = configuration{
		CountryName:             "Bulgaria",
		CountryCode:             bg,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "BNBG 9661 1020 3456 78",
		IBANDefinition: definition{
			Length:      22,
			Example:     "BG80BNBG96611020345678",
			PrintFormat: "BG80 BNBG 9661 1020 3456 78",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "BNBG96611020345678",
		},
	}

	bahrainConfiguration = configuration{
		CountryName:             "Bahrain",
		CountryCode:             bh,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00001299123456",
		IBANDefinition: definition{
			Length:      22,
			Example:     "BH67BMAG00001299123456",
			PrintFormat: "BH67 BMAG 0000 1299 1234 56",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "BMAG00001299123456",
		},
	}

	brazilConfiguration = configuration{
		CountryName:             "Brazil",
		CountryCode:             br,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0009795493C1",
		IBANDefinition: definition{
			Length:      29,
			Example:     "BR1800360305000010009795493C1",
			PrintFormat: "BR18 0036 0305 0000 1000 9795 493C 1",
		},
		BBANDefinition: definition{
			Length:  25,
			Example: "00360305000010009795493P1",
		},
	}

	republicBelarusConfiguration = configuration{
		CountryName:             "Republic of Belarus",
		CountryCode:             by,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "3600 0000 0000 0Z00 AB00",
		IBANDefinition: definition{
			Length:      28,
			Example:     "BY13NBRB3600900000002Z00AB00",
			PrintFormat: "BY13 NBRB 3600 9000 0000 2Z00 AB00",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "NBRB3600000000000Z00AB00",
		},
	}

	switzerlandConfiguration = configuration{
		CountryName:             "Switzerland",
		CountryCode:             ch,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "762 1162-3852.957",
		IBANDefinition: definition{
			Length:      21,
			Example:     "CH9300762011623852957",
			PrintFormat: "CH93 0076 2011 6238 5295 7",
		},
		BBANDefinition: definition{
			Length:  17,
			Example: "00762011623852957",
		},
	}

	costaRicaConfiguration = configuration{
		CountryName:             "Costa Rica",
		CountryCode:             cr,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "02001026284066",
		IBANDefinition: definition{
			Length:      22,
			Example:     "CR05015202001026284066",
			PrintFormat: "CR05 0152 0200 1026 2840 66",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "015202001026284066",
		},
	}

	cyprusConfiguration = configuration{
		CountryName:             "Cyprus",
		CountryCode:             cy,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0000001200527600",
		IBANDefinition: definition{
			Length:      28,
			Example:     "CY17002001280000001200527600",
			PrintFormat: "CY17 0020 0128 0000 0012 0052 7600",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "002001280000001200527600",
		},
	}

	czechRepublicConfiguration = configuration{
		CountryName:             "Czech Republic",
		CountryCode:             cz,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "19-2000145399/0800",
		IBANDefinition: definition{
			Length:      24,
			Example:     "CZ6508000000192000145399",
			PrintFormat: "CZ65 0800 0000 1920 0014 5399",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "08000000192000145399",
		},
	}

	germanyConfiguration = configuration{
		CountryName:             "Germany",
		CountryCode:             de,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "532013000",
		IBANDefinition: definition{
			Length:      22,
			Example:     "DE89370400440532013000",
			PrintFormat: "DE89 3704 0044 0532 0130 00",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "370400440532013000",
		},
	}

	denmarkConfiguration = configuration{
		CountryName:             "Denmark",
		CountryCode:             dk,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0040 0440116243",
		IBANDefinition: definition{
			Length:      18,
			Example:     "DK5000400440116243",
			PrintFormat: "DK50 0040 0440 1162 43",
		},
		BBANDefinition: definition{
			Length:  14,
			Example: "00400440116243",
		},
	}

	dominicanRepublicConfiguration = configuration{
		CountryName:             "Dominican Republic",
		CountryCode:             do,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00000001212453611324",
		IBANDefinition: definition{
			Length:      28,
			Example:     "DO28BAGR00000001212453611324",
			PrintFormat: "DO28 BAGR 0000 0001 2124 5361 1324",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "BAGR00000001212453611324",
		},
	}

	estoniaConfiguration = configuration{
		CountryName:             "Estonia",
		CountryCode:             ee,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "221020145685",
		IBANDefinition: definition{
			Length:      20,
			Example:     "EE382200221020145685",
			PrintFormat: "EE38 2200 2210 2014 5685",
		},
		BBANDefinition: definition{
			Length:  16,
			Example: "2200221020145685",
		},
	}

	spainConfiguration = configuration{
		CountryName:             "Spain",
		CountryCode:             es,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "2100 0418 45 0200051332",
		IBANDefinition: definition{
			Length:      24,
			Example:     "ES9121000418450200051332",
			PrintFormat: "ES91 2100 0418 4502 0005 1332",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "21000418450200051332",
		},
	}

	finlandConfiguration = configuration{
		CountryName:             "Finland",
		CountryCode:             fi,
		IncludedCountryCode:     []countryCode{ax},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{ax},
		AccountNumberExample:    "1234 5600 0007 85",
		IBANDefinition: definition{
			Length:      18,
			Example:     "FI2112345600000785",
			PrintFormat: "FI21 1234 5600 0007 85",
		},
		BBANDefinition: definition{
			Length:  14,
			Example: "12345600000785",
		},
	}

	faroeIslandsConfiguration = configuration{
		CountryName:             "Faroe Islands",
		CountryCode:             fo,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "6460 0001631634",
		IBANDefinition: definition{
			Length:      18,
			Example:     "FO6264600001631634",
			PrintFormat: "FO62 6460 0001 6316 34",
		},
		BBANDefinition: definition{
			Length:  14,
			Example: "64600001631634",
		},
	}

	franceConfiguration = configuration{
		CountryName:             "France",
		CountryCode:             fr,
		IncludedCountryCode:     []countryCode{gf, gp, mq, re, pf, tf, yt, nc, bl, mf, pm, wf},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{gf, gp, mq, yt, re, pm, bl, mf},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "FR1420041010050500013M02606",
			PrintFormat: "FR14 2004 1010 0505 0001 3M02 606",
		},
		BBANDefinition: definition{
			Length:  23,
			Example: "20041010050500013M02606",
		},
	}

	unitedKingdomConfiguration = configuration{
		CountryName:             "United Kingdom",
		CountryCode:             gb,
		IncludedCountryCode:     []countryCode{im, je, gg},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "60-16-13 31926819",
		IBANDefinition: definition{
			Length:      22,
			Example:     "GB29NWBK60161331926819",
			PrintFormat: "GB29 NWBK 6016 1331 9268 19",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "NWBK60161331926819",
		},
	}

	georgiaConfiguration = configuration{
		CountryName:             "Georgia",
		CountryCode:             ge,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0000000101904917",
		IBANDefinition: definition{
			Length:      22,
			Example:     "GE29NB0000000101904917",
			PrintFormat: "GE29 NB00 0000 0101 9049 17",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "NB0000000101904917",
		},
	}

	gibraltarConfiguration = configuration{
		CountryName:             "Gibraltar",
		CountryCode:             gi,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0000 00007099 453",
		IBANDefinition: definition{
			Length:      23,
			Example:     "GI75NWBK000000007099453",
			PrintFormat: "GI75 NWBK 0000 0000 7099 453",
		},
		BBANDefinition: definition{
			Length:  19,
			Example: "NWBK000000007099453",
		},
	}

	greenlandConfiguration = configuration{
		CountryName:             "Greenland",
		CountryCode:             gl,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "6471 0001000206",
		IBANDefinition: definition{
			Length:      18,
			Example:     "GL8964710001000206",
			PrintFormat: "GL89 6471 0001 0002 06",
		},
		BBANDefinition: definition{
			Length:  14,
			Example: "64710001000206",
		},
	}

	greeceConfiguration = configuration{
		CountryName:             "Greece",
		CountryCode:             gr,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "01250000000012300695",
		IBANDefinition: definition{
			Length:      27,
			Example:     "GR1601101250000000012300695",
			PrintFormat: "GR16 0110 1250 0000 0001 2300 695",
		},
		BBANDefinition: definition{
			Length:  23,
			Example: "01101250000000012300695",
		},
	}

	guatemalaConfiguration = configuration{
		CountryName:             "Guatemala",
		CountryCode:             gt,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "01020000001210029690",
		IBANDefinition: definition{
			Length:      28,
			Example:     "GT82TRAJ01020000001210029690",
			PrintFormat: "GT82 TRAJ 0102 0000 0012 1002 9690",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "TRAJ01020000001210029690",
		},
	}

	croatiaConfiguration = configuration{
		CountryName:             "Croatia",
		CountryCode:             hr,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "1001005-1863000160",
		IBANDefinition: definition{
			Length:      21,
			Example:     "HR1210010051863000160",
			PrintFormat: "HR12 1001 0051 8630 0016 0",
		},
		BBANDefinition: definition{
			Length:  17,
			Example: "10010051863000160",
		},
	}

	hungaryConfiguration = configuration{
		CountryName:             "Hungary",
		CountryCode:             hu,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "11773016-11111018-000000009",
		IBANDefinition: definition{
			Length:      28,
			Example:     "HU42117730161111101800000000",
			PrintFormat: "HU42 1177 3016 1111 1018 0000 0000",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "117730161111101800000000",
		},
	}

	irelandConfiguration = configuration{
		CountryName:             "Ireland",
		CountryCode:             ie,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "93-11-52 12345678",
		IBANDefinition: definition{
			Length:      22,
			Example:     "IE29AIBK93115212345678",
			PrintFormat: "IE29 AIBK 9311 5212 3456 78",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "AIBK93115212345678",
		},
	}

	israelConfiguration = configuration{
		CountryName:             "Israel",
		CountryCode:             il,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "10-800-99999999",
		IBANDefinition: definition{
			Length:      23,
			Example:     "IL620108000000099999999",
			PrintFormat: "IL62 0108 0000 0009 9999 999",
		},
		BBANDefinition: definition{
			Length:  19,
			Example: "0108000000099999999",
		},
	}

	iraqConfiguration = configuration{
		CountryName:             "Iraq",
		CountryCode:             iq,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "123456789012",
		IBANDefinition: definition{
			Length:      23,
			Example:     "IQ98NBIQ850123456789012",
			PrintFormat: "IQ98 NBIQ 8501 2345 6789 012",
		},
		BBANDefinition: definition{
			Length:  19,
			Example: "NBIQ850123456789012",
		},
	}

	icelandConfiguration = configuration{
		CountryName:             "Iceland",
		CountryCode:             is,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0159-26-007654-551073-0339",
		IBANDefinition: definition{
			Length:      26,
			Example:     "IS140159260076545510730339",
			PrintFormat: "IS14 0159 2600 7654 5510 7303 39",
		},
		BBANDefinition: definition{
			Length:  22,
			Example: "0159260076545510730339",
		},
	}

	italyConfiguration = configuration{
		CountryName:             "Italy",
		CountryCode:             it,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "X 05428 11101 000000123456",
		IBANDefinition: definition{
			Length:      27,
			Example:     "IT60X0542811101000000123456",
			PrintFormat: "IT60 X054 2811 1010 0000 0123 456",
		},
		BBANDefinition: definition{
			Length:  23,
			Example: "X0542811101000000123456",
		},
	}

	jordanConfiguration = configuration{
		CountryName:             "Jordan",
		CountryCode:             jo,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0001310000302",
		IBANDefinition: definition{
			Length:      30,
			Example:     "JO94CBJO0010000000000131000302",
			PrintFormat: "JO94 CBJO 0010 0000 0000 0131 0003 02",
		},
		BBANDefinition: definition{
			Length:  26,
			Example: "CBJO0010000000000131000302",
		},
	}

	kuwaitConfiguration = configuration{
		CountryName:             "Kuwait",
		CountryCode:             kw,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "1234560101",
		IBANDefinition: definition{
			Length:      30,
			Example:     "KW81CBKU0000000000001234560101",
			PrintFormat: "KW81 CBKU 0000 0000 0000 1234 5601 01",
		},
		BBANDefinition: definition{
			Length:  26,
			Example: "CBKU0000000000001234560101",
		},
	}

	kazakhstanConfiguration = configuration{
		CountryName:             "Kazakhstan",
		CountryCode:             kz,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "KZ86 125K ZT50 0410 0100",
		IBANDefinition: definition{
			Length:      20,
			Example:     "KZ86125KZT5004100100",
			PrintFormat: "KZ86 125K ZT50 0410 0100",
		},
		BBANDefinition: definition{
			Length:  16,
			Example: "125KZT5004100100",
		},
	}

	lebanonConfiguration = configuration{
		CountryName:             "Lebanon",
		CountryCode:             lb,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "01 001 901229114",
		IBANDefinition: definition{
			Length:      28,
			Example:     "LB62099900000001001901229114",
			PrintFormat: "LB62 0999 0000 0001 0019 0122 9114",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "099900000001001901229114",
		},
	}

	saintLuciaConfiguration = configuration{
		CountryName:             "Saint Lucia",
		CountryCode:             lc,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0001 0001 0012 0012 0002 3015",
		IBANDefinition: definition{
			Length:      32,
			Example:     "LC55HEMM000100010012001200023015",
			PrintFormat: "LC55 HEMM 0001 0001 0012 0012 0002 3015",
		},
		BBANDefinition: definition{
			Length:  28,
			Example: "HEMM000100010012001200023015",
		},
	}

	liechtensteinConfiguration = configuration{
		CountryName:             "Liechtenstein",
		CountryCode:             li,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "8810 2324013AA",
		IBANDefinition: definition{
			Length:      21,
			Example:     "LI21088100002324013AA",
			PrintFormat: "LI21 0881 0000 2324 013A A",
		},
		BBANDefinition: definition{
			Length:  17,
			Example: "088100002324013AA",
		},
	}

	lithuaniaConfiguration = configuration{
		CountryName:             "Lithuania",
		CountryCode:             lt,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "11101001000",
		IBANDefinition: definition{
			Length:      20,
			Example:     "LT121000011101001000",
			PrintFormat: "LT12 1000 0111 0100 1000",
		},
		BBANDefinition: definition{
			Length:  16,
			Example: "1000011101001000",
		},
	}

	luxembourgConfiguration = configuration{
		CountryName:             "Luxembourg",
		CountryCode:             lu,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "9400644750000",
		IBANDefinition: definition{
			Length:      20,
			Example:     "LU280019400644750000",
			PrintFormat: "LU28 0019 4006 4475 0000",
		},
		BBANDefinition: definition{
			Length:  16,
			Example: "0019400644750000",
		},
	}

	latviaConfiguration = configuration{
		CountryName:             "Latvia",
		CountryCode:             lv,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "LV80 BANK 0000 4351 9500 1",
		IBANDefinition: definition{
			Length:      21,
			Example:     "LV80BANK0000435195001",
			PrintFormat: "LV80 BANK 0000 4351 9500 1",
		},
		BBANDefinition: definition{
			Length:  17,
			Example: "BANK0000435195001",
		},
	}

	monacoConfiguration = configuration{
		CountryName:             "Monaco",
		CountryCode:             mc,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0011111000h",
		IBANDefinition: definition{
			Length:      27,
			Example:     "MC5811222000010123456789030",
			PrintFormat: "MC58 1122 2000 0101 2345 6789 030",
		},
		BBANDefinition: definition{
			Length:  23,
			Example: "11222000010123456789030",
		},
	}

	moldovaConfiguration = configuration{
		CountryName:             "Moldova",
		CountryCode:             md,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "000225100013104168",
		IBANDefinition: definition{
			Length:      24,
			Example:     "MD24AG000225100013104168",
			PrintFormat: "MD24 AG00 0225 1000 1310 4168",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "AG000225100013104168",
		},
	}

	montenegroConfiguration = configuration{
		CountryName:             "Montenegro",
		CountryCode:             me,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "505 0000123456789 51",
		IBANDefinition: definition{
			Length:      22,
			Example:     "ME25505000012345678951",
			PrintFormat: "ME25 5050 0001 2345 6789 51",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "505000012345678951",
		},
	}

	macedoniaConfiguration = configuration{
		CountryName:             "Macedonia",
		CountryCode:             mk,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "MK07 300 0000000424 25",
		IBANDefinition: definition{
			Length:      19,
			Example:     "MK07250120000058984",
			PrintFormat: "MK07 2501 2000 0058 984",
		},
		BBANDefinition: definition{
			Length:  15,
			Example: "250120000058984",
		},
	}

	mauritaniaConfiguration = configuration{
		CountryName:             "Mauritania",
		CountryCode:             mr,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00020 00101 00001234567 53",
		IBANDefinition: definition{
			Length:      27,
			Example:     "MR1300020001010000123456753",
			PrintFormat: "MR13 0002 0001 0100 0012 3456 753",
		},
		BBANDefinition: definition{
			Length:  23,
			Example: "00020001010000123456753",
		},
	}

	maltaConfiguration = configuration{
		CountryName:             "Malta",
		CountryCode:             mt,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "12345MTLCAST001S",
		IBANDefinition: definition{
			Length:      31,
			Example:     "MT84MALT011000012345MTLCAST001S",
			PrintFormat: "MT84 MALT 0110 0001 2345 MTLC AST0 01S",
		},
		BBANDefinition: definition{
			Length:  27,
			Example: "MALT011000012345MTLCAST001S",
		},
	}

	mauritiusConfiguration = configuration{
		CountryName:             "Mauritius",
		CountryCode:             mu,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "MU17 BOMM 0101 1010 3030 0200 000M UR",
		IBANDefinition: definition{
			Length:      30,
			Example:     "MU17BOMM0101101030300200000MUR",
			PrintFormat: "MU17 BOMM 0101 1010 3030 0200 000M UR",
		},
		BBANDefinition: definition{
			Length:  26,
			Example: "BOMM0101101030300200000MUR",
		},
	}

	netherlandsConfiguration = configuration{
		CountryName:             "The Netherlands",
		CountryCode:             nl,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "041 71 64 300",
		IBANDefinition: definition{
			Length:      18,
			Example:     "NL91ABNA0417164300",
			PrintFormat: "NL91 ABNA 0417 1643 00",
		},
		BBANDefinition: definition{
			Length:  14,
			Example: "ABNA0417164300",
		},
	}

	norwayConfiguration = configuration{
		CountryName:             "Norway",
		CountryCode:             no,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "8601 11 17947",
		IBANDefinition: definition{
			Length:      15,
			Example:     "NO9386011117947",
			PrintFormat: "NO93 8601 1117 947",
		},
		BBANDefinition: definition{
			Length:  11,
			Example: "86011117947",
		},
	}

	pakistanConfiguration = configuration{
		CountryName:             "Pakistan",
		CountryCode:             pk,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00260101036360",
		IBANDefinition: definition{
			Length:      24,
			Example:     "PK36SCBL0000001123456702",
			PrintFormat: "PK36 SCBL 0000 0011 2345 6702",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "SCBL0000001123456702",
		},
	}

	polandConfiguration = configuration{
		CountryName:             "Poland",
		CountryCode:             pl,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "61 1090 1014 0000 0712 1981 2874",
		IBANDefinition: definition{
			Length:      28,
			Example:     "PL61109010140000071219812874",
			PrintFormat: "PL61 1090 1014 0000 0712 1981 2874",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "109010140000071219812874",
		},
	}

	palestineConfiguration = configuration{
		CountryName:             "State of Palestine",
		CountryCode:             ps,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "400123456702",
		IBANDefinition: definition{
			Length:      29,
			Example:     "PS92PALS000000000400123456702",
			PrintFormat: "PS92 PALS 0000 0000 0400 1234 5670 2",
		},
		BBANDefinition: definition{
			Length:  25,
			Example: "PALS000000000400123456702",
		},
	}

	portugalConfiguration = configuration{
		CountryName:             "Portugal",
		CountryCode:             pt,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{}, // Azores, Madeira does not have proper Country Code
		AccountNumberExample:    "0002.0123.12345678901.54",
		IBANDefinition: definition{
			Length:      25,
			Example:     "PT50000201231234567890154",
			PrintFormat: "PT50 0002 0123 1234 5678 9015 4",
		},
		BBANDefinition: definition{
			Length:  21,
			Example: "000201231234567890154",
		},
	}

	qatarConfiguration = configuration{
		CountryName:             "Qatar",
		CountryCode:             qa,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00001234567890ABCDEFG",
		IBANDefinition: definition{
			Length:      29,
			Example:     "QA58DOHB00001234567890ABCDEFG",
			PrintFormat: "QA58 DOHB 0000 1234 5678 90AB CDEF G",
		},
		BBANDefinition: definition{
			Length:  25,
			Example: "DOHB00001234567890ABCDEFG",
		},
	}

	romaniaConfiguration = configuration{
		CountryName:             "Romania",
		CountryCode:             ro,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "RO49 AAAA 1B31 0075 9384 0000",
		IBANDefinition: definition{
			Length:      24,
			Example:     "RO49AAAA1B31007593840000",
			PrintFormat: "RO49 AAAA 1B31 0075 9384 0000",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "AAAA1B31007593840000",
		},
	}

	serbiaConfiguration = configuration{
		CountryName:             "Serbia",
		CountryCode:             rs,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "260-0056010016113-79",
		IBANDefinition: definition{
			Length:      22,
			Example:     "RS35260005601001611379",
			PrintFormat: "RS35 2600 0560 1001 6113 79",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "260005601001611379",
		},
	}

	saudiArabiaConfiguration = configuration{
		CountryName:             "Saudi Arabia",
		CountryCode:             sa,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "608010167519",
		IBANDefinition: definition{
			Length:      24,
			Example:     "SA0380000000608010167519",
			PrintFormat: "SA03 8000 0000 6080 1016 7519",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "80000000608010167519",
		},
	}

	seychellesConfiguration = configuration{
		CountryName:             "Seychelles",
		CountryCode:             sc,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0000000000001497",
		IBANDefinition: definition{
			Length:      31,
			Example:     "SC18SSCB11010000000000001497USD",
			PrintFormat: "SC18 SSCB 1101 0000 0000 0000 1497 USD",
		},
		BBANDefinition: definition{
			Length:  27,
			Example: "SSCB11010000000000001497USD",
		},
	}

	swedenConfiguration = configuration{
		CountryName:             "Sweden",
		CountryCode:             se,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "1234 12 3456 1",
		IBANDefinition: definition{
			Length:      24,
			Example:     "SE4550000000058398257466",
			PrintFormat: "SE45 5000 0000 0583 9825 7466",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "50000000058398257466",
		},
	}

	sloveniaConfiguration = configuration{
		CountryName:             "Slovenia",
		CountryCode:             si,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "2633 0001 2039 086",
		IBANDefinition: definition{
			Length:      19,
			Example:     "SI56263300012039086",
			PrintFormat: "SI56 2633 0001 2039 086",
		},
		BBANDefinition: definition{
			Length:  15,
			Example: "263300012039086",
		},
	}

	slovakiaConfiguration = configuration{
		CountryName:             "Slovakia",
		CountryCode:             sk,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "19-8742637541/1200",
		IBANDefinition: definition{
			Length:      24,
			Example:     "SK3112000000198742637541",
			PrintFormat: "SK31 1200 0000 1987 4263 7541",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "12000000198742637541",
		},
	}

	sanMarinoConfiguration = configuration{
		CountryName:             "San Marino",
		CountryCode:             sm,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "U0322509800000000270100",
		IBANDefinition: definition{
			Length:      27,
			Example:     "SM86U0322509800000000270100",
			PrintFormat: "SM86 U032 2509 8000 0000 0270 100",
		},
		BBANDefinition: definition{
			Length:  23,
			Example: "U0322509800000000270100",
		},
	}

	saoTomePrincipeConfiguration = configuration{
		CountryName:             "Sao Tome and Principe",
		CountryCode:             st,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "01921942101.12",
		IBANDefinition: definition{
			Length:      25,
			Example:     "ST32000200010192194210112",
			PrintFormat: "ST32 0002 0001 0192 1942 1011 2",

			// Given the Swift IBAN Registryn the following IBAN should be valid but they're not. The correct checksum is 32, not 68.
			//Example:     "ST68000200010192194210112",
			//PrintFormat: "ST68 0002 0001 0192 1942 1011 2",
		},
		BBANDefinition: definition{
			Length:  21,
			Example: "000200010192194210112",
		},
	}

	elSalvadorConfiguration = configuration{
		CountryName:             "El Salvador",
		CountryCode:             sv,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00000000000000700025",
		IBANDefinition: definition{
			Length:      28,
			Example:     "SV62CENR00000000000000700025",
			PrintFormat: "SV 62 CENR 00000000000000700025",
		},
		BBANDefinition: definition{
			Length:  24,
			Example: "CENR00000000000000700025",
		},
	}

	timorLesteConfiguration = configuration{
		CountryName:             "Timor-Leste",
		CountryCode:             tl,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "008 00123456789101 57",
		IBANDefinition: definition{
			Length:      23,
			Example:     "TL380080012345678910157",
			PrintFormat: "TL38 0080 0123 4567 8910 157",
		},
		BBANDefinition: definition{
			Length:  19,
			Example: "0080012345678910157",
		},
	}

	tunisiaConfiguration = configuration{
		CountryName:             "Tunisia",
		CountryCode:             tn,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "10 006 0351835984788 31",
		IBANDefinition: definition{
			Length:      24,
			Example:     "TN5910006035183598478831",
			PrintFormat: "TN59 1000 6035 1835 9847 8831",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "10006035183598478831",
		},
	}

	turkeyConfiguration = configuration{
		CountryName:             "Turkey",
		CountryCode:             tr,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00519786457841326",
		IBANDefinition: definition{
			Length:      26,
			Example:     "TR330006100519786457841326",
			PrintFormat: "TR33 0006 1005 1978 6457 8413 26",
		},
		BBANDefinition: definition{
			Length:  22,
			Example: "0006100519786457841326",
		},
	}

	ukraineConfiguration = configuration{
		CountryName:             "Ukraine",
		CountryCode:             ua,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "26007233566001",
		IBANDefinition: definition{
			Length:      29,
			Example:     "UA213223130000026007233566001",
			PrintFormat: "UA21 3223 1300 0002 6007 2335 6600 1",
		},
		BBANDefinition: definition{
			Length:  25,
			Example: "3223130000026007233566001",
		},
	}

	vaticanCityStateConfiguration = configuration{
		CountryName:             "Vatican City State",
		CountryCode:             va,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "123000012345678",
		IBANDefinition: definition{
			Length:      22,
			Example:     "VA59001123000012345678",
			PrintFormat: "VA59 001 1230 0001 2345 678",
		},
		BBANDefinition: definition{
			Length:  18,
			Example: "001123000012345678",
		},
	}

	virginIslandsConfiguration = configuration{
		CountryName:             "Virgin Islands",
		CountryCode:             vg,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00000 12 345 678 901",
		IBANDefinition: definition{
			Length:      24,
			Example:     "VG96VPVG0000012345678901",
			PrintFormat: "VG96 VPVG 0000 0123 4567 8901",
		},
		BBANDefinition: definition{
			Length:  20,
			Example: "VPVG0000012345678901",
		},
	}

	kosovoConfiguration = configuration{
		CountryName:             "Kosovo",
		CountryCode:             xk,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "1212 0123456789 06",
		IBANDefinition: definition{
			Length:      20,
			Example:     "XK051212012345678906",
			PrintFormat: "XK05 1212 0123 4567 8906",
		},
		BBANDefinition: definition{
			Length:  16,
			Example: "1212012345678906",
		},
	}
)

// GetCountryConfiguration returns ISO 13616-Compliant IBAN Formats from a given country code
// If the countryCode does not exist in the list, returns an error `ErrCountryCodeDoesNotExist`
func GetCountryConfiguration(countryCode string) (*configuration, error) {
	if conf, ok := countriesConfiguration[countryCode]; ok {
		return &conf, nil
	}
	return nil, ErrCountryCodeDoesNotExist
}
