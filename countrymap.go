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

		gf.String(): frenchGuianaConfiguration,
		gp.String(): guadaloupeConfiguration,
		mq.String(): martiniqueConfiguration,
		re.String(): reunionConfiguration,
		pf.String(): frenchPolynesiaConfiguration,
		tf.String(): frenchSouthernTerritoriesConfiguration,
		yt.String(): mayotteConfiguration,
		nc.String(): newCaledoniaConfiguration,
		bl.String(): saintBarthelemyConfiguration,
		mf.String(): collectivityOfSaintMartinConfiguration,
		pm.String(): saintPierreAndMiquelonConfiguration,
		wf.String(): wallisAndFutunaIslandsConfiguration,
		ax.String(): alandIslandsConfiguration,


		// Experimental IBAN Countries
		dz.String(): algeriaConfiguration,
		ao.String(): angolaConfiguration,
		bj.String(): beninConfiguration,
		bf.String(): burkinaFasoConfiguration,
		bi.String(): burundiConfiguration,
		cm.String(): cameroonConfiguration,
		cv.String(): capeVerdeConfiguration,
		cf.String(): centralAfricanRepublicConfiguration,
		td.String(): chadConfiguration,
		km.String(): comorosConfiguration,
		cg.String(): congoConfiguration,
		dj.String(): djiboutiConfiguration,
		eg.String(): egyptConfiguration,
		ga.String(): gabonConfiguration,
		gw.String(): guineaBissauConfiguration,
		hn.String(): hondurasConfiguration,
		ir.String(): iranConfiguration,
		ci.String(): ivoryCoastConfiguration,
		mg.String(): madagascarConfiguration,
		ml.String(): maliConfiguration,
		ma.String(): moroccoConfiguration,
		mz.String(): mozambiqueConfiguration,
		ni.String(): nicaraguaConfiguration,
		ne.String(): nigerConfiguration,
		sn.String(): senegalConfiguration,
		tg.String(): togoConfiguration,
	}

	andorraConfiguration = configuration{
		CountryName:             "Andorra",
		CountryCode:             ad,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "2030200359100100",
		IBANDefinition: definition{
			Length:      24,
			Example:     "AD1200012030200359100100",
			PrintFormat: "AD12 0001 2030 2003 5910 0100",
			Structure:   "cckkbbbbssssaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "00012030200359100100",
			Structure: "bbbbssssaaaaaaaaaaaa",
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
			Structure:   "cckkbbbaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    19,
			Example:   "0331234567890123456",
			Structure: "bbbaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbssssxaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "212110090000000235698741",
			Structure: "bbbssssxaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    16,
			Example:   "1904300234573201",
			Structure: "bbbbbaaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "NABZ00000000137010001944",
			Structure: "wwwwaaaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbsssaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    16,
			Example:   "1990440001200279",
			Structure: "bbbsssaaaaaaaaxx",
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
			Structure:   "cckkbbbaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    12,
			Example:   "539007547034",
			Structure: "bbbaaaaaaaxx",
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
			Structure:   "cckkwwwwssssttaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "BNBG96611020345678",
			Structure: "wwwwssssttaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "BMAG00001299123456",
			Structure: "wwwwaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbbbbsssssaaaaaaaaaatn",
		},
		BBANDefinition: definition{
			Length:    25,
			Example:   "00360305000010009795493P1",
			Structure: "bbbbbbbbsssssaaaaaaaaaatn",
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
			Structure:   "cckkbbbbssssaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "NBRB3600000000000Z00AB00",
			Structure: "bbbbssssaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    17,
			Example:   "00762011623852957",
			Structure: "bbbbbaaaaaaaaaaaa",
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
			Structure:   "cckkobbbaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "015202001026284066",
			Structure: "obbbaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbsssssaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "002001280000001200527600",
			Structure: "bbbsssssaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbssssssaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "08000000192000145399",
			Structure: "bbbbssssssaaaaaaaaaa",
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
			Structure:   "cckkbbbbbbbbaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "370400440532013000",
			Structure: "bbbbbbbbaaaaaaaaaa",
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
			Structure:   "cckkbbbbaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    14,
			Example:   "00400440116243",
			Structure: "bbbbaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "BAGR00000001212453611324",
			Structure: "wwwwaaaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbssaaaaaaaaaaax",
		},
		BBANDefinition: definition{
			Length:    16,
			Example:   "2200221020145685",
			Structure: "bbssaaaaaaaaaaax",
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
			Structure:   "cckkbbbbssssxxaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "21000418450200051332",
			Structure: "bbbbssssxxaaaaaaaaaa",
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
			Structure:   "cckkbbbbbbaaaaaaax",
		},
		BBANDefinition: definition{
			Length:    14,
			Example:   "12345600000785",
			Structure: "bbbbbbaaaaaaax",
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
			Structure:   "cckkbbbbaaaaaaaaax",
		},
		BBANDefinition: definition{
			Length:    14,
			Example:   "64600001631634",
			Structure: "bbbbaaaaaaaaax",
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
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
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
			Structure:   "cckkwwwwssssssaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "NWBK60161331926819",
			Structure: "wwwwssssssaaaaaaaa",
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
			Structure:   "cckkwwaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "NB0000000101904917",
			Structure: "wwaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    19,
			Example:   "NWBK000000007099453",
			Structure: "wwwwaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    14,
			Example:   "64710001000206",
			Structure: "bbbbaaaaaaaaaa",
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
			Structure:   "cckkbbbssssaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "01101250000000012300695",
			Structure: "bbbssssaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwmmttaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "TRAJ01020000001210029690",
			Structure: "wwwwmmttaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbbbaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    17,
			Example:   "10010051863000160",
			Structure: "bbbbbbbaaaaaaaaaa",
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
			Structure:   "cckkbbbsssssaaaaaaaaaaaaaaax",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "117730161111101800000000",
			Structure: "bbbsssssaaaaaaaaaaaaaaax",
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
			Structure:   "cckkwwwwssssssaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "AIBK93115212345678",
			Structure: "wwwwssssssaaaaaaaa",
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
			Structure:   "cckkbbbsssaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    19,
			Example:   "0108000000099999999",
			Structure: "bbbsssaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwsssaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    19,
			Example:   "NBIQ850123456789012",
			Structure: "wwwwsssaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbssaaaaaaiiiiiiiiii",
		},
		BBANDefinition: definition{
			Length:    22,
			Example:   "0159260076545510730339",
			Structure: "bbbbssaaaaaaiiiiiiiiii",
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
			Structure:   "cckkxbbbbbsssssaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "X0542811101000000123456",
			Structure: "xbbbbbsssssaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwssssaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    26,
			Example:   "CBJO0010000000000131000302",
			Structure: "wwwwssssaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    26,
			Example:   "CBKU0000000000001234560101",
			Structure: "wwwwaaaaaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    16,
			Example:   "125KZT5004100100",
			Structure: "bbbaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "099900000001001901229114",
			Structure: "bbbbaaaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    28,
			Example:   "HEMM000100010012001200023015",
			Structure: "wwwwaaaaaaaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    17,
			Example:   "088100002324013AA",
			Structure: "bbbbbaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    16,
			Example:   "1000011101001000",
			Structure: "bbbbbaaaaaaaaaaa",
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
			Structure:   "cckkbbbaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    16,
			Example:   "0019400644750000",
			Structure: "bbbaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    17,
			Example:   "BANK0000435195001",
			Structure: "wwwwaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "11222000010123456789030",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
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
			Structure:   "cckkwwaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "AG000225100013104168",
			Structure: "wwaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbaaaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "505000012345678951",
			Structure: "bbbaaaaaaaaaaaaaxx",
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
			Structure:   "cckkbbbaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    15,
			Example:   "250120000058984",
			Structure: "bbbaaaaaaaaaaxx",
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
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "00020001010000123456753",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
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
			Structure:   "cckkwwwwsssssaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    27,
			Example:   "MALT011000012345MTLCAST001S",
			Structure: "wwwwsssssaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbbssaaaaaaaaaaaaooommm",
		},
		BBANDefinition: definition{
			Length:    26,
			Example:   "BOMM0101101030300200000MUR",
			Structure: "bbbbbbssaaaaaaaaaaaaooommm",
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
			Structure:   "cckkwwwwaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    14,
			Example:   "ABNA0417164300",
			Structure: "wwwwaaaaaaaaaa",
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
			Structure:   "cckkbbbbaaaaaax",
		},
		BBANDefinition: definition{
			Length:    11,
			Example:   "86011117947",
			Structure: "bbbbaaaaaax",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "SCBL0000001123456702",
			Structure: "wwwwaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbssssxaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "109010140000071219812874",
			Structure: "bbbssssxaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwoooooooooaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    25,
			Example:   "PALS000000000400123456702",
			Structure: "wwwwoooooooooaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    21,
			Example:   "000201231234567890154",
			Structure: "bbbbssssaaaaaaaaaaaxx",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    25,
			Example:   "DOHB00001234567890ABCDEFG",
			Structure: "wwwwaaaaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "AAAA1B31007593840000",
			Structure: "wwwwaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbaaaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "260005601001611379",
			Structure: "bbbaaaaaaaaaaaaaxx",
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
			Structure:   "cckkbbaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "80000000608010167519",
			Structure: "bbaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwssssaaaaaaaaaaaaaaaammm",
		},
		BBANDefinition: definition{
			Length:    27,
			Example:   "SSCB11010000000000001497USD",
			Structure: "wwwwssssaaaaaaaaaaaaaaaammm",
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
			Structure:   "cckkbbbaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "50000000058398257466",
			Structure: "bbbaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbsssaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    15,
			Example:   "263300012039086",
			Structure: "bbsssaaaaaaaaxx",
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
			Structure:   "cckkbbbbssssssaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "12000000198742637541",
			Structure: "bbbbssssssaaaaaaaaaa",
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
			Structure:   "cckkxbbbbbsssssaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "U0322509800000000270100",
			Structure: "xbbbbbsssssaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbssssaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    21,
			Example:   "000200010192194210112",
			Structure: "bbbbssssaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "CENR00000000000000700025",
			Structure: "wwwwaaaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    19,
			Example:   "0080012345678910157",
			Structure: "bbbaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbsssaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "10006035183598478831",
			Structure: "bbsssaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbboaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    22,
			Example:   "0006100519786457841326",
			Structure: "bbbbboaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbbbbaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    25,
			Example:   "3223130000026007233566001",
			Structure: "bbbbbbaaaaaaaaaaaaaaaaaaa",
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
			Structure:   "cckkbbbaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    18,
			Example:   "001123000012345678",
			Structure: "bbbaaaaaaaaaaaaaaa",
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
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "VPVG0000012345678901",
			Structure: "wwwwaaaaaaaaaaaaaaaa",
		},
	}

	kosovoConfiguration = configuration{
		CountryName:             "Kosovo",
		CountryCode:             xk,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "1212 0123456789 06",
		IBANDefinition: definition{
			Length:      20,
			Example:     "XK051212012345678906",
			PrintFormat: "XK05 1212 0123 4567 8906",
			Structure:   "cckkbbbbaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    16,
			Example:   "1212012345678906",
			Structure: "bbbbaaaaaaaaaaaa",
		},
	}

	frenchGuianaConfiguration = configuration{
		CountryName:             "French Guiana",
		CountryCode:             gf,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "GF4120041010050500013M02606",
			PrintFormat: "GF41 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	guadaloupeConfiguration = configuration{
		CountryName:             "Guadaloupe",
		CountryCode:             gp,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "GP1120041010050500013M02606",
			PrintFormat: "GP11 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	martiniqueConfiguration = configuration{
		CountryName:             "Martinique",
		CountryCode:             mq,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "MQ5120041010050500013M02606",
			PrintFormat: "MQ51 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	reunionConfiguration = configuration{
		CountryName:             "Reunion",
		CountryCode:             re,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "RE4220041010050500013M02606",
			PrintFormat: "RE42 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	frenchPolynesiaConfiguration = configuration{
		CountryName:             "French Polynesia",
		CountryCode:             pf,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "PF5720041010050500013M02606",
			PrintFormat: "PF57 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	frenchSouthernTerritoriesConfiguration = configuration{
		CountryName:             "French Southern Territories",
		CountryCode:             tf,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "TF2120041010050500013M02606",
			PrintFormat: "TF21 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	mayotteConfiguration = configuration{
		CountryName:             "Mayotte",
		CountryCode:             yt,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "YT3120041010050500013M02606",
			PrintFormat: "YT31 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	newCaledoniaConfiguration = configuration{
		CountryName:             "New Caledonia",
		CountryCode:             nc,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "NC8420041010050500013M02606",
			PrintFormat: "NC84 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	saintBarthelemyConfiguration = configuration{
		CountryName:             "Saint Barthélemy",
		CountryCode:             bl,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "BL6820041010050500013M02606",
			PrintFormat: "BL68 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	collectivityOfSaintMartinConfiguration = configuration{
		CountryName:             "Saint Martin",
		CountryCode:             mf,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "MF8420041010050500013M02606",
			PrintFormat: "MF84 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	saintPierreAndMiquelonConfiguration = configuration{
		CountryName:             "Kosovo",
		CountryCode:             pm,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "PM3620041010050500013M02606",
			PrintFormat: "PM36 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	wallisAndFutunaIslandsConfiguration = configuration{
		CountryName:             "Wallis and Futuna",
		CountryCode:             wf,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20041 01005 0500013M026 06",
		IBANDefinition: definition{
			Length:      27,
			Example:     "WF9120041010050500013M02606",
			PrintFormat: "WF91 2004 1010 0505 0001 3M02 606",
			Structure:   "cckkbbbbbsssssaaaaaaaaaaaxx",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20041010050500013M02606",
			Structure: "bbbbbsssssaaaaaaaaaaaxx",
		},
	}

	alandIslandsConfiguration = configuration{
		CountryName:             "Åland Islands",
		CountryCode:             ax,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         true,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "1234 5600 0007 85",
		IBANDefinition: definition{
			Length:      18,
			Example:     "AX2112345600000785",
			PrintFormat: "AX21 1234 5600 0007 85",
			Structure:   "cckkbbbbbbaaaaaaax",
		},
		BBANDefinition: definition{
			Length:    14,
			Example:   "12345600000785",
			Structure: "bbbbbbaaaaaaax",
		},
	}

	algeriaConfiguration = configuration{
		CountryName:             "Algeria",
		CountryCode:             dz,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "22738346840171372597",
		IBANDefinition: definition{
			Length:      24,
			Example:     "DZ8022738346840171372597",
			PrintFormat: "DZ80 2273 8346 8401 7137 2597",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    20,
			Example:   "22738346840171372597",
			Structure: "aaaaaaaaaaaaaaaaaaaa",
		},
	}

	angolaConfiguration = configuration{
		CountryName:             "Angola",
		CountryCode:             ao,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "226781323520060113196",
		IBANDefinition: definition{
			Length:      25,
			Example:     "AO27226781323520060113196",
			PrintFormat: "AO27 2267 8132 3520 0601 1319 6",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    21,
			Example:   "226781323520060113196",
			Structure: "aaaaaaaaaaaaaaaaaaaaa",
		},
	}

	beninConfiguration = configuration{
		CountryName:             "Benin",
		CountryCode:             bj,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "99932227891101178429964",
		IBANDefinition: definition{
			Length:      28,
			Example:     "BJ73B99932227891101178429964",
			PrintFormat: "BJ73 B999 3222 7891 1011 7842 9964",
			Structure:   "cckkbaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "B99932227891101178429964",
			Structure: "baaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	burkinaFasoConfiguration = configuration{
		CountryName:             "Burkina Faso",
		CountryCode:             bf,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "30134020015400945000643",
		IBANDefinition: definition{
			Length:      27,
			Example:     "BF1030134020015400945000643",
			PrintFormat: "BF10 3013 4020 0154 0094 5000 643",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "30134020015400945000643",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	burundiConfiguration = configuration{
		CountryName:             "Burundi",
		CountryCode:             bi,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "201011067444",
		IBANDefinition: definition{
			Length:      16,
			Example:     "BI43201011067444",
			PrintFormat: "BI43 2010 1106 7444",
			Structure:   "cckkaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    12,
			Example:   "201011067444",
			Structure: "aaaaaaaaaaaa",
		},
	}

	cameroonConfiguration = configuration{
		CountryName:             "Cameroon",
		CountryCode:             cm,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "10003001000500000605306",
		IBANDefinition: definition{
			Length:      27,
			Example:     "CM2110003001000500000605306",
			PrintFormat: "CM21 1000 3001 0005 0000 0605 306",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "10003001000500000605306",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	capeVerdeConfiguration = configuration{
		CountryName:             "Cape Verde",
		CountryCode:             cv,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "000300004547069110176",
		IBANDefinition: definition{
			Length:      25,
			Example:     "CV64000300004547069110176",
			PrintFormat: "CV64 0003 0000 4547 0691 1017 6",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    21,
			Example:   "000300004547069110176",
			Structure: "aaaaaaaaaaaaaaaaaaaaa",
		},
	}

	centralAfricanRepublicConfiguration = configuration{
		CountryName:             "Central African Republic",
		CountryCode:             cf,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "20001000010120069700160",
		IBANDefinition: definition{
			Length:      27,
			Example:     "CF4220001000010120069700160",
			PrintFormat: "CF42 2000 1000 0101 2006 9700 160",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "20001000010120069700160",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	chadConfiguration = configuration{
		CountryName:             "Chad",
		CountryCode:             td,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00020000556026733000006",
		IBANDefinition: definition{
			Length:      27,
			Example:     "TD7200020000556026733000006",
			PrintFormat: "TD72 0002 0000 5560 2673 3000 006",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "00020000556026733000006",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	comorosConfiguration = configuration{
		CountryName:             "Comoros",
		CountryCode:             km,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "40002000055602673300064",
		IBANDefinition: definition{
			Length:      27,
			Example:     "KM4640002000055602673300064",
			PrintFormat: "KM46 4000 2000 0556 0267 3300 064",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "40002000055602673300064",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	congoConfiguration = configuration{
		CountryName:             "Congo",
		CountryCode:             cg,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "30011000202151234567890",
		IBANDefinition: definition{
			Length:      27,
			Example:     "CG5230011000202151234567890",
			PrintFormat: "CG52 3001 1000 2021 5123 4567 890",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "30011000202151234567890",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	djiboutiConfiguration = configuration{
		CountryName:             "Djibouti",
		CountryCode:             dj,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "10002010010409943020008",
		IBANDefinition: definition{
			Length:      27,
			Example:     "DJ2110002010010409943020008",
			PrintFormat: "DJ21 1000 2010 0104 0994 3020 008",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "10002010010409943020008",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	egyptConfiguration = configuration{
		CountryName:             "Egypt",
		CountryCode:             eg,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00006001880800100014553",
		IBANDefinition: definition{
			Length:      27,
			Example:     "EG1100006001880800100014553",
			PrintFormat: "EG11 0000 6001 8808 0010 0014 553",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "00006001880800100014553",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	gabonConfiguration = configuration{
		CountryName:             "Gabon",
		CountryCode:             ga,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "40002000055602673300064",
		IBANDefinition: definition{
			Length:      27,
			Example:     "GA2140002000055602673300064",
			PrintFormat: "GA21 4000 2000 0556 0267 3300 064",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "40002000055602673300064",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	guineaBissauConfiguration = configuration{
		CountryName:             "Guinea-Bissau",
		CountryCode:             gw,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "GW1430010181800637601",
		IBANDefinition: definition{
			Length:      25,
			Example:     "GW04GW1430010181800637601",
			PrintFormat: "GW04 GW14 3001 0181 8006 3760 1",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    21,
			Example:   "GW1430010181800637601",
			Structure: "aaaaaaaaaaaaaaaaaaaaa",
		},
	}

	hondurasConfiguration = configuration{
		CountryName:             "Honduras",
		CountryCode:             hn,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "27346509718683143911",
		IBANDefinition: definition{
			Length:      28,
			Example:     "HN25VOWU27346509718683143911",
			PrintFormat: "HN25 VOWU 2734 6509 7186 8314 3911",
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "VOWU27346509718683143911",
			Structure: "wwwwaaaaaaaaaaaaaaaaaaaa",
		},
	}

	iranConfiguration = configuration{
		CountryName:             "Iran",
		CountryCode:             ir,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0570029971601460641001",
		IBANDefinition: definition{
			Length:      26,
			Example:     "IR710570029971601460641001",
			PrintFormat: "IR71 0570 0299 7160 1460 6410 01",
			Structure:   "cckkbbbbaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    22,
			Example:   "0570029971601460641001",
			Structure: "bbbbaaaaaaaaaaaaaaaaaa",
		},
	}

	ivoryCoastConfiguration = configuration{
		CountryName:             "Ivory Coast",
		CountryCode:             ci,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "28456236523184811313497",
		IBANDefinition: definition{
			Length:      28,
			Example:     "CI80C28456236523184811313497",
			PrintFormat: "CI80 C284 5623 6523 1848 1131 3497",
			Structure:   "cckkbaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "C28456236523184811313497",
			Structure: "baaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	madagascarConfiguration = configuration{
		CountryName:             "Madagascar",
		CountryCode:             mg,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00005030010101914016056",
		IBANDefinition: definition{
			Length:      27,
			Example:     "MG4600005030010101914016056",
			PrintFormat: "MG46 0000 5030 0101 0191 4016 056",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    23,
			Example:   "00005030010101914016056",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	maliConfiguration = configuration{
		CountryName:             "Mali",
		CountryCode:             ml,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00890170001002120000447",
		IBANDefinition: definition{
			Length:      28,
			Example:     "ML03D00890170001002120000447",
			PrintFormat: "ML03 D008 9017 0001 0021 2000 0447",
			Structure:   "cckkbaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "D00890170001002120000447",
			Structure: "baaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	moroccoConfiguration = configuration{
		CountryName:             "Morocco",
		CountryCode:             ma,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "011519000001205000534921",
		IBANDefinition: definition{
			Length:      28,
			Example:     "MA64011519000001205000534921",
			PrintFormat: "MA64 0115 1900 0001 2050 0053 4921",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "011519000001205000534921",
			Structure: "aaaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	mozambiqueConfiguration = configuration{
		CountryName:             "Mozambique",
		CountryCode:             mz,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "000100000011834194157",
		IBANDefinition: definition{
			Length:      25,
			Example:     "MZ59000100000011834194157",
			PrintFormat: "MZ59 0001 0000 0011 8341 9415 7",
			Structure:   "cckkaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    21,
			Example:   "000100000011834194157",
			Structure: "aaaaaaaaaaaaaaaaaaaaa",
		},
	}

	nicaraguaConfiguration = configuration{
		CountryName:             "Nicaragua",
		CountryCode:             ni,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "157382578924464167824447",
		IBANDefinition: definition{
			Length:      32,
			Example:     "NI62ABCD157382578924464167824447",
			PrintFormat: "NI62 ABCD 1573 8257 8924 4641 6782 4447",
			Structure:   "cckkwwwwaaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    28,
			Example:   "ABCD157382578924464167824447",
			Structure: "wwwwaaaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	nigerConfiguration = configuration{
		CountryName:             "Niger",
		CountryCode:             ne,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0380100100130305000268",
		IBANDefinition: definition{
			Length:      28,
			Example:     "NE58NE0380100100130305000268",
			PrintFormat: "NE58 NE03 8010 0100 1303 0500 0268",
			Structure:   "cckkbbaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "NE0380100100130305000268",
			Structure: "bbaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	senegalConfiguration = configuration{
		CountryName:             "Senegal",
		CountryCode:             sn,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "00100152000048500003035",
		IBANDefinition: definition{
			Length:      28,
			Example:     "SN21S00100152000048500003035",
			PrintFormat: "SN21 S001 0015 2000 0485 0000 3035",
			Structure:   "cckkbaaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "S00100152000048500003035",
			Structure: "baaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	togoConfiguration = configuration{
		CountryName:             "Togo",
		CountryCode:             tg,
		IncludedCountryCode:     []countryCode{},
		IsSEPAAvailable:         false,
		SEPAIncludedCountryCode: []countryCode{},
		AccountNumberExample:    "0090604310346500400070",
		IBANDefinition: definition{
			Length:      28,
			Example:     "TG53TG0090604310346500400070",
			PrintFormat: "TG53 TG00 9060 4310 3465 0040 0070",
			Structure:   "cckkbbaaaaaaaaaaaaaaaaaaaaaa",
		},
		BBANDefinition: definition{
			Length:    24,
			Example:   "TG0090604310346500400070",
			Structure: "bbaaaaaaaaaaaaaaaaaaaaaa",
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
