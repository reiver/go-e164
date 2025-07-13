package e164

// ISOCountryCode is given an E.164 country-code (ex: "1", "82", "98"), and a national-destination-code (ex: "604")
// and returns the corresponding ISO 3166-1 alpha-2 country-code (ex: "CA", "KR", "IR").
//
// For example:
//
//	isoCountryCode := e164.ISOCountryCode("1", "604")
//	// isoCountryCode == "CA"
func ISOCountryCode(countryCode string, nationalDestinationCode string) string {

	// A switch in Go is faster than map.
	//
	// THIS SWITCH WILL NEED TO BE UPDATED WHENEVER THE ITU-T MAKES AN UPDATE.
	switch countryCode {
	case "1":
		switch nationalDestinationCode {
		case	"907",                                                                                     // Alaska
			"205","251","256","334",                                                                   // Alabama
			"479","501","870",                                                                         // Arkansas
			"480","520","602","623","928",                                                             // Arizona
			"209","213","310","323","408","415","510","530","559","562","619","626","650","661","707","714","760","805","818","831","858","909","916","925","949","951", // California
			"303","719","970",                                                                         // Colorado
			"203","860",                                                                               // Connecticut
			"202",                                                                                     // District of Columbia
			"302",                                                                                     // Delaware
			"239","305","321","352","386","407","561","727","772","813","850","863","904","941","954", // Florida
			"229","404","478","706","770","912",                                                       // Georgia
			"808",                                                                                     // Hawaii
			"319","515","563","641","712",                                                             // Iowa
			"208",                                                                                     // Idaho
			"217","309","312","618","630","708","773","815","847",                                     // Illinois
			"219","260","317","574","765","812",                                                       // Indiana
			"316","620","785","913",                                                                   // Kansas
			"270","502","606","859",                                                                   // Kentucky
			"225","318","337","504","985",                                                             // Louisiana
			"413","508","617","781","978",                                                             // Massachusetts
			"301","410",                                                                               // Maryland
			"207",                                                                                     // Maine
			"231","248","269","313","517","586","616","734","810","906","989",                         // Michigan
			"218","320","507","612","651","763","952",                                                 // Minnesota
			"314","417","573","636","660","816",                                                       // Missouri
			"228","601","662",                                                                         // Mississippi
			"406",                                                                                     // Montana
			"252","336","704","828","910","919",                                                       // North Carolina
			"701",                                                                                     // North Dakota
			"308","402",                                                                               // Nebraska
			"603",                                                                                     // New Hampshire
			"201","609","732","856","908","973",                                                       // New Jersey
			"505","575",                                                                               // New Mexico
			"702","775",                                                                               // Nevada
			"212","315","516","518","585","607","631","716","718","845","914",                         // New York
			"216","330","419","440","513","614","740","937",                                           // Ohio
			"405","580","918",                                                                         // Oklahoma
			"503","541",                                                                               // Oregon
			"215","412","570","610","717","724","814",                                                 // Pennsylvania
			"401",                                                                                     // Rhode Island
			"803","843","864",                                                                         // South Carolina
			"605",                                                                                     // South Dakota
			"423","615","731","865","901","931",                                                       // Tennessee
			"210","214","254","281","325","361","409","432","512","713","806","817","830","903","915","936","940","956","972","979", // Texas
			"435","801",                                                                               // Utah
			"276","434","540","703","757","804",                                                       // Virginia
			"802",                                                                                     // Vermont
			"206","253","360","425","509",                                                             // Washington
			"262","414","608","715","920",                                                             // Wisconsin
			"304",                                                                                     // West Virginia
			"307":                                                                                     // Wyoming
			return "US"
		case	"204", // Manitoba
			"226", // London, ON
			"236", // Vancouver, BC
			"249", // Sudbury, ON
			"250", // Kelowna, BC
			"263", // Montreal, QC
			"289", // Hamilton, ON
			"306", // Saskatchewan
			"343", // Ottawa, ON
			"354", // Granby, QC
			"365", // Hamilton, ON
			"367", // Quebec, QC
			"368", // Calgary, AB
			"382", // London, ON
			"403", // Calgary, AB
			"416", // Toronto, ON
			"418", // Quebec, QC
			"428", // New Brunswick
			"431", // Manitoba
			"437", // Toronto, ON
			"438", // Montreal, QC
			"450", // Granby, QC
			"468", // Sherbrooke, QC
			"474", // Saskatchewan
			"506", // New Brunswick
			"514", // Montreal, QC
			"519", // London, ON
			"548", // London, ON
			"579", // Granby, QC
			"581", // Quebec, QC
			"584", // Manitoba
			"587", // Calgary, AB
			"604", // Vancouver, BC
			"613", // Ottawa, ON
			"639", // Saskatchewan
			"647", // Toronto, ON
			"672", // Vancouver, BC
			"683", // Sudbury, ON
			"705", // Sudbury, ON
			"709", // Newfoundland/Labrador
			"742", // Hamilton, ON
			"753", // Ottawa, ON
			"778", // Vancouver, BC
			"780", // Edmonton, AB
			"782", // Nova Scotia/PE Island
			"807", // Kenora, ON
			"819", // Sherbrooke, QC
			"825", // Calgary, AB
			"867", // Northern Canada
			"873", // Sherbrooke, QC
			"879", // Newfoundland/Labrador
			"902", // Nova Scotia/PE Island
			"905": // Hamilton, ON
			return "CA"
		// 1 (242) – Bahamas
		case "242":
			return "BS"
		// 1 (246) – Barbados
		case "246":
			return "BB"
		// 1 (264) – Anguilla
		case "264":
			return "AI"
		// 1 (268) – Antigua and Barbuda
		case "268":
			return "AG"
		// 1 (284) – British Virgin Islands
		case "284":
			return "VG"
		// 1 (340) – United States Virgin Islands
		case "340":
			return "VI"
		// 1 (345) – Cayman Islands
		case "345":
			return "KY"
		// 1 (441) – Bermuda
		case "441":
			return "BM"
		// 1 (473) – Grenada
		case "473":
			return "GD"
		// 1 (649) – Turks and Caicos Islands
		case "649":
			return "TC"
		// 1 (658, 876) – Jamaica
		case "658", "876":
			return "JM"
		// 1 (664) – Montserrat
		case "664":
			return "MS"
		// 1 (670) – Northern Mariana Islands
		case "670":
			return "MP"
		// 1 (671) – Guam
		case "671":
			return "GU"
		// 1 (684) – American Samoa
		case "684":
			return "AS"
		// 1 (721) – Sint Maarten
		case "721":
			return "SX"
		// 1 (758) – Saint Lucia
		case "758":
			return "LC"
		// 1 (767) – Dominica
		case "767":
			return "DM"
		// 1 (784) – Saint Vincent and the Grenadines
		case "784":
			return "VC"
		// 1 (787, 939) – Puerto Rico
		case "787","939":
			return "PR"
		// 1 (809, 829, 849) – Dominican Republic
		case "809", "829", "849":
			return "DO"
		// 1 (868) – Trinidad and Tobago
		case "868":
			return "TT"
		// 1 (869) – Saint Kitts and Nevis
		case "869":
			return "KN"
		default:
			return ""
		}

	// 20 – Egypt
	case "20":
		return "EG"

	// 211 – South Sudan
	case "211":
		return "SS"

	// 212 – Morocco (including Western Sahara)
	case "212":
		return "MA"

	// 213 – Algeria
	case "213":
		return "DZ"

	// 216 – Tunisia
	case "216":
		return "TN"

	// 218 – Libya
	case "218":
		return "LY"

	// 220 – Gambia
	case "220":
		return "GM"

	// 221 – Senegal
	case "221":
		return "SN"

	// 222 – Mauritania
	case "222":
		return "MR"

	// 223 – Mali
	case "223":
		return "ML"

	// 224 – Guinea
	case "224":
		return "GN"

	// 225 – Ivory Coast
	case "225":
		return "CI"

	// 226 – Burkina Faso
	case "226":
		return "BF"

	// 227 – Niger
	case "227":
		return "NE"

	// 228 – Togo
	case "228":
		return "TG"

	// 229 – Benin
	case "229":
		return "BJ"

	// 230 – Mauritius
	case "230":
		return "MU"

	// 231 – Liberia
	case "231":
		return "LR"

	// 232 – Sierra Leone
	case "232":
		return "SL"

	// 233 – Ghana
	case "233":
		return "GH"

	// 234 – Nigeria
	case "234":
		return "NG"

	// 235 – Chad
	case "235":
		return "TD"

	// 236 – Central African Republic
	case "236":
		return "CF"

	// 237 – Cameroon
	case "237":
		return "CM"

	// 238 – Cape Verde
	case "238":
		return "CV"

	// 239 – São Tomé and Príncipe
	case "239":
		return "ST"

	// 240 – Equatorial Guinea
	case "240":
		return "GQ"

	// 241 – Gabon
	case "241":
		return "GA"

	// 242 – Republic of the Congo
	case "242":
		return "CG"

	// 243 – Democratic Republic of the Congo
	case "243":
		return "CD"

	// 244 – Angola
	case "244":
		return "AO"

	// 245 – Guinea-Bissau
	case "245":
		return "GW"

	// 246 – British Indian Ocean Territory
	case "246":
		return "IO"

	// 247 – Ascension Island
	case "247":
		return "AC"

	// 248 – Seychelles
	case "248":
		return "SC"

	// 249 – Sudan
	case "249":
		return "SD"

	// 250 – Rwanda
	case "250":
		return "RW"

	// 251 – Ethiopia
	case "251":
		return "ET"

	// 252 – Somalia (including Somaliland)
	case "252":
		return "SO"

	// 253 – Djibouti
	case "253":
		return "DJ"

	// 254 – Kenya
	case "254":
		return "KE"

	// 255 – Tanzania
	// 255 (24) – Zanzibar, in place of never-implemented 259
	case "255":
		return "TZ"

	// 256 – Uganda
	case "256":
		return "UG"

	// 257 – Burundi
	case "257":
		return "BI"

	// 258 – Mozambique
	case "258":
		return "MZ"

	// 260 – Zambia
	case "260":
		return "ZM"

	// 261 – Madagascar
	case "261":
		return "MG"

	// 262 – Réunion
	// 262 (269,639) – Mayotte (formerly at 269 Comoros)
	case "262":
		switch nationalDestinationCode {
		case "269","639":
			return "YT"
		default:
			return "RE"
		}

	// 263 – Zimbabwe
	case "263":
		return "ZW"

	// 264 – Namibia (formerly 27 (6x) as South West Africa)
	case "264":
		return "NA"

	// 265 – Malawi
	case "265":
		return "MW"

	// 266 – Lesotho
	case "266":
		return "LS"

	// 267 – Botswana
	case "267":
		return "BW"

	// 268 – Eswatini
	case "268":
		return "SZ"

	// 269 – Comoros (formerly assigned to Mayotte, now at 262)
	case "269":
		return "KM"

	// 27 – South Africa
	case "27":
		return "ZA"

	// 290 – Saint Helena
        // 290 (8) – Tristan da Cunha
	case "290":
		return "SH"

	// 291 – Eritrea
	case "291":
		return "ER"

	// 297 – Aruba
	case "297":
		return "AW"

	// 298 – Faroe Islands
	case "298":
		return "FO"

	// 299 – Greenland
	case "299":
		return "GL"

	// 30 – Greece
	case "30":
		return "GR"

	// 31 – Netherlands
	case "31":
		return "NL"

	// 32 – Belgium
	case "32":
		return "BE"

	// 33 – France
	case "33":
		return "FR"

	// 34 – Spain
	case "34":
		return "ES"

	// 350 – Gibraltar
	case "350":
		return "GI"

	// 351 – Portugal
	// 351 (291) – Madeira (landlines only)
	// 351 (292) – Azores (landlines only, Horta, Azores area)
	// 351 (295) – Azores (landlines only, Angra do Heroísmo area)
	// 351 (296) – Azores (landlines only, Ponta Delgada and São Miguel Island area)
	case "351":
		return "PT"

	// 352 – Luxembourg
	case "352":
		return "LU"

	// 353 – Ireland
	case "353":
		return "IE"

	// 354 – Iceland
	case "354":
		return "IS"

	// 355 – Albania
	case "355":
		return "AL"

	// 356 – Malta
	case "356":
		return "MT"

	// 357 – Cyprus (including Akrotiri and Dhekelia)
	case "357":
		return "CY"

	// 358 – Finland
	// 358 (18) – Åland
	case "358":
		switch nationalDestinationCode {
		case "18":
			return "AX"
		default:
			return "FI"
		}

	// 359 – Bulgaria
	case "359":
		return "BG"

	// 36 – Hungary (formerly assigned to Turkey, now at 90)
	case "36":
		return "HU"

	// 370 – Lithuania
	case "370":
		return "LT"

	// 371 – Latvia
	case "371":
		return "LV"

	// 372 – Estonia
	case "372":
		return "EE"

	// 373 – Moldova
	case "373":
		return "MD"

	// 374 – Armenia
	case "374":
		return "AM"

	// 375 – Belarus
	case "375":
		return "BY"

	// 376 – Andorra (formerly 33 628)
	case "376":
		return "AD"

	// 377 – Monaco (formerly 33 93)
	case "377":
		return "MC"

	// 378 – San Marino (interchangeably with 39 0549; earlier was allocated 295 but never used)
	case "378":
		return "SM"

	// 379 – Vatican City (assigned but uses 39 06698).
	case "379":
		return "VA"

	// 380 – Ukraine
	case "380":
		return "UA"

	// 381 – Serbia
	case "381":
		return "RS"

	// 382 – Montenegro
	case "382":
		return "ME"

	// 383 – Kosovo
	case "383":
		return "XK"

	// 385 – Croatia
	case "385":
		return "HR"

	// 386 – Slovenia
	case "386":
		return "SI"

	// 387 – Bosnia and Herzegovina
	case "387":
		return "BA"

	// 389 – North Macedonia
	case "389":
		return "MK"

	// 39 – Italy
	// 39 (0549) – San Marino (interchangeably with 378)
	// 39 (06 698) – Vatican City (assigned 379 but not in use)
	case "39":
		switch nationalDestinationCode {
		case "0549":
			return "SM"
		case "06698":
			return "VA"
		default:
			return "IT"
		}

	// 40 – Romania
	case "40":
		return "RO"

	// 41 –  Switzerland
	// 41 (91) – Campione d'Italia, an Italian enclave. 91 is the prefix for the Swiss canton Ticino in which the enclave resides. Its phone system is fully integrated into the Swiss system.
	case "41":
		switch nationalDestinationCode {
		case "91":
			return "IT"
		default:
			return "CH"
		}

	// 420 – Czech Republic
	case "420":
		return "CZ"

	// 421 – Slovakia
	case "421":
		return "SK"

	// 423 – Liechtenstein (formerly at 41 (75))
	case "423":
		return "LI"

	// 43 – Austria
	case "43":
		return "AT"

	// 44 – United Kingdom
	// 44 (1481) – Guernsey
	// 44 (1534) – Jersey
	// 44 (1624) – Isle of Man
	case "44":
		switch nationalDestinationCode {
		case "1481":
			return "GG"
		case "1534":
			return "JE"
		case "1624":
			return "IM"
		default:
			return "GB"
		}

	// 45 – Denmark
	case "45":
		return "DK"

	// 46 – Sweden
	case "46":
		return "SE"

	// 47 – Norway
	// 47 (79) – Svalbard
	case "47":
		switch nationalDestinationCode {
		case "79":
			return "SJ"
		default:
			return "NO"
		}

	// 48 – Poland
	case "48":
		return "PL"

	// 49 – Germany
	case "49":
		return "DE"

	// 500 – Falkland Islands
	// 500 – South Georgia and the South Sandwich Islands
	case "500":
		return "FK"

	// 501 – Belize
	case "501":
		return "BZ"

	// 502 – Guatemala
	case "502":
		return "GT"

	// 503 – El Salvador
	case "503":
		return "SV"

	// 504 – Honduras
	case "504":
		return "HN"

	// 505 – Nicaragua
	case "505":
		return "NI"

	// 506 – Costa Rica
	case "506":
		return "CR"

	// 507 – Panama
	case "507":
		return "PA"

	// 508 – Saint-Pierre and Miquelon
	case "508":
		return "PM"

	// 509 – Haiti
	case "509":
		return "HT"

	// 51 – Peru
	case "51":
		return "PE"

	// 52 – Mexico
	case "52":
		return "MX"

	// 53 – Cuba
	case "53":
		return "CU"

	// 54 – Argentina
	case "54":
		return "AR"

	// 55 – Brazil
	case "55":
		return "BR"

	// 56 – Chile
	case "56":
		return "CL"

	// 57 – Colombia
	case "57":
		return "CO"

	// 58 – Venezuela
	case "58":
		return "VE"

	// 590 – Guadeloupe (including Saint Barthélemy, Saint Martin)
	case "590":
		return "GP"

	// 591 – Bolivia
	case "591":
		return "BO"

	// 592 – Guyana
	case "592":
		return "GY"

	// 593 – Ecuador
	case "593":
		return "EC"

	// 594 – French Guiana
	case "594":
		return "GF"

	// 595 – Paraguay
	case "595":
		return "PY"

	// 596 – Martinique (formerly assigned to Peru, now 51)
	case "596":
		return "MQ"

	// 597 – Suriname
	case "597":
		return "SR"

	// 598 – Uruguay
	case "598":
		return "UY"

	// 599 – Former Netherlands Antilles, now grouped as follows:
	// 599 3 – Sint Eustatius
	// 599 4 – Saba
	// 599 7 – Bonaire
	// 599 9 – Curaçao
	case "599":
		switch nationalDestinationCode {
		case "3":
			return "BQ"
		case "4":
			return "BQ"
		case "7":
			return "BQ"
		case "9":
			return "CW"
		default:
			return "AN"
		}

	// 60 – Malaysia
	case "60":
		return "MY"

	// 61 – Australia (see also 672 below)
	// 61 (8 9162) – Cocos Islands
	// 61 (8 9164) – Christmas Island
	case "61":
		return ""

	// 62 – Indonesia
	case "62":
		return "ID"

	// 63 – Philippines
	case "63":
		return "PH"

	// 64 – New Zealand
	// 64 – Pitcairn Islands
	case "64":
		return "NZ"

	// 65 – Singapore
	case "65":
		return "SG"

	// 66 – Thailand
	case "66":
		return "TH"

	// 670 – East Timor (formerly 62/39 during the Indonesian occupation; formerly assigned to Northern Mariana Islands, now part of NANP as 1 (670))
	case "670":
		return "TL"

	// 672 – Australian External Territories (see also 61 Australia above); formerly assigned to Portuguese Timor (see 670)
	// 672 (1) – Australia Australian Antarctic Territory
	// 672 (3) – Norfolk Island
	case "672":
		switch nationalDestinationCode {
		case "1":
			return "AU"
		case "3":
			return "NF"
		default:
			return "AU"
		}

	// 673 – Brunei
	case "673":
		return "BN"

	// 674 – Nauru
	case "674":
		return "NR"

	// 675 – Papua New Guinea
	case "675":
		return "PG"

	// 676 – Tonga
	case "676":
		return "TO"

	// 677 – Solomon Islands
	case "677":
		return "SB"

	// 678 – Vanuatu
	case "678":
		return "VU"

	// 679 – Fiji
	case "679":
		return "FJ"

	// 680 – Palau
	case "680":
		return "PW"

	// 681 – Wallis and Futuna
	case "681":
		return "WF"

	// 682 – Cook Islands
	case "682":
		return "CK"

	// 683 – Niue
	case "683":
		return "NU"

	// 685 – Samoa
	case "685":
		return "WS"

	// 686 – Kiribati
	case "686":
		return "KI"

	// 687 – New Caledonia
	case "687":
		return "NC"

	// 688 – Tuvalu
	case "688":
		return "TV"

	// 689 – French Polynesia
	case "689":
		return "PF"

	// 690 – Tokelau
	case "690":
		return "TK"

	// 691 – Federated States of Micronesia
	case "691":
		return "FM"

	// 692 – Marshall Islands
	case "692":
		return "MH"

	// 7 (1–5, 8–9) – Russia
	// 7 (840, 940) – Abkhazia (formerly 995 (44))
	// 7 (850, 929) – South Ossetia (formerly 995 (34))
	// 7 (0, 6–7) – Kazakhstan (reserved 997 but abandoned in November 2023)
	case "7":
		switch nationalDestinationCode {
		case "1","2","3","4","5","8","9":
			return "RU"
		case "840","940":
			return "GE"
		case "850","929":
			return "RU"
		case "0","6","7":
			return "KZ"
		default:
			return ""
		}

	// 81 – Japan
	case "81":
		return "JP"

	// 82 – South Korea
	case "82":
		return "KR"

	// 84 – Vietnam
	case "84":
		return "VN"

	// 850 – North Korea
	case "850":
		return "KP"

	// 852 – Hong Kong
	case "852":
		return "HK"

	// 853 – Macau
	case "853":
		return "MO"

	// 855 – Cambodia
	case "855":
		return "KH"

	// 856 – Laos
	case "856":
		return "LA"

	// 86 – China
	case "86":
		return "CN"

	// 880 – Bangladesh
	case "880":
		return "BD"

	// 886 – Taiwan
	case "886":
		return "TW"

	// 90 – Turkey
	case "90":
		return "TR"

	// 91 – India
	case "91":
		return "IN"

	// 92 – Pakistan
	case "92":
		return "PK"

	// 93 – Afghanistan
	case "93":
		return "AF"

	// 94 – Sri Lanka
	case "94":
		return "LK"

	// 95 – Myanmar
	case "95":
		return "MM"

	// 960 – Maldives
	case "960":
		return "MV"

	// 961 – Lebanon
	case "961":
		return "LB"

	// 962 – Jordan
	case "962":
		return "JO"

	// 963 – Syria
	case "963":
		return "SY"

	// 964 – Iraq
	case "964":
		return "IQ"

	// 965 – Kuwait
	case "965":
		return "KW"

	// 966 – Saudi Arabia
	case "966":
		return "SA"

	// 967 – Yemen
	case "967":
		return "YE"

	// 968 – Oman
	case "968":
		return "OM"

	// 970 – Palestine (interchangeably with 972)
	case "970":
		return "PS"

	// 971 – United Arab Emirates
	case "971":
		return "AE"

	// 972 – Israel (also Palestine, interchangeably with 970)
	case "972":
		return "IL"

	// 973 – Bahrain
	case "973":
		return "BH"

	// 974 – Qatar
	case "974":
		return "QA"

	// 975 – Bhutan
	case "975":
		return "BT"

	// 976 – Mongolia
	case "976":
		return "MN"

	// 977 – Nepal
	case "977":
		return "NP"

	// 98 – Iran
	case "98":
		return "IR"

	// 992 – Tajikistan
	case "992":
		return "TJ"

	// 993 – Turkmenistan
	case "993":
		return "TM"

	// 994 – Azerbaijan
	case "994":
		return "AZ"

	// 995 – Georgia
	case "995":
		return "GE"

	// 996 – Kyrgyzstan
	case "996":
		return "KG"

	// 997 – Kazakhstan (reserved but abandoned in November 2023; uses 7 (6xx, 7xx))
	case "997":
		return "KZ"

	// 998 – Uzbekistan
	case "998":
		return "UZ"

	default:
		return ""
	}
}
