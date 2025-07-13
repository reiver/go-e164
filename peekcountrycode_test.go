package e164

import (
	"testing"
)

func TestPeekCountryCode(t *testing.T) {

	tests := []struct{
		String string
		ExpectedLength int
		ExpectedFound bool
	}{
		{
			String: "",
			ExpectedLength: 0,
			ExpectedFound: false,
		},



		{
			String: "apple",
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String: "BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String: "Cherry",
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String: "dAtE",
			ExpectedLength: 0,
			ExpectedFound: false,
		},



		{
			String: "once",
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String: "twice",
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String: "thrice",
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String: "fource",
			ExpectedLength: 0,
			ExpectedFound: false,
		},



		{
			String: "1",
			ExpectedLength: 1,
			ExpectedFound: true,
		},



		{
			String: "0BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},


		// 1 – United States, including United States territories
		// 1 – Canada
		// 1 – Caribbean nations, Dutch and British Overseas Territories
		{
			String:             "1BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},
		{
			String: "16045551234",
			ExpectedLength: 1,
			ExpectedFound: true,

		},

		// 1 (340) – United States Virgin Islands
		{
			String:             "1BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (670) – Northern Mariana Islands
		{
			String:             "1670BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (671) – Guam
		{
			String:             "1671BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (684) – American Samoa
		{
			String:             "1684BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (787, 939) – Puerto Rico
		{
			String:             "1787BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},
		{
			String:             "1939BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (242) – Bahamas
		{
			String:             "1242BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (246) – Barbados
		{
			String:             "1246BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (264) – Anguilla
		{
			String:             "1264BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (268) – Antigua and Barbuda
		{
			String:             "1268BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (284) – British Virgin Islands
		{
			String:             "1284BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (345) – Cayman Islands
		{
			String:             "1345BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (441) – Bermuda
		{
			String:             "1441BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (473) – Grenada
		{
			String:             "1473BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (649) – Turks and Caicos Islands
		{
			String:             "1649BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (658, 876) – Jamaica
		{
			String:             "1658BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},
		{
			String:             "1876BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (664) – Montserrat
		{
			String:             "1664BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (721) – Sint Maarten
		{
			String:             "1721BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (758) – Saint Lucia
		{
			String:             "1758BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (767) – Dominica
		{
			String:             "1767BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (784) – Saint Vincent and the Grenadines
		{
			String:             "1784BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (809, 829, 849) – Dominican Republic
		{
			String:             "1809BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},
		{
			String:             "1829BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},
		{
			String:             "1849BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (868) – Trinidad and Tobago
		{
			String:             "1868BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		// 1 (869) – Saint Kitts and Nevis
		{
			String:             "1869BANANA",
			ExpectedLength: len("1"),
			ExpectedFound: true,
		},

		{
			String:             "2BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 20 – Egypt
		{
			String:             "20BANANA",
			ExpectedLength: len("20"),
			ExpectedFound: true,
		},

		{
			String:             "21BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 210 – unassigned
		{
			String:             "210BANANA",
			ExpectedLength: len("210"),
			ExpectedFound: true,
		},

		// 211 – South Sudan
		{
			String:             "211BANANA",
			ExpectedLength: len("211"),
			ExpectedFound: true,
		},

		// 212 – Morocco (including Western Sahara)
		{
			String:             "212BANANA",
			ExpectedLength: len("212"),
			ExpectedFound: true,
		},

		// 213 – Algeria
		{
			String:             "213BANANA",
			ExpectedLength: len("213"),
			ExpectedFound: true,
		},

		// 214 – unassigned
		{
			String:             "214BANANA",
			ExpectedLength: len("214"),
			ExpectedFound: true,
		},

		// 215 – unassigned
		{
			String:             "215BANANA",
			ExpectedLength: len("215"),
			ExpectedFound: true,
		},

		// 216 – Tunisia
		{
			String:             "216BANANA",
			ExpectedLength: len("216"),
			ExpectedFound: true,
		},

		// 217 – unassigned
		{
			String:             "217BANANA",
			ExpectedLength: len("217"),
			ExpectedFound: true,
		},

		// 218 – Libya
		{
			String:             "218BANANA",
			ExpectedLength: len("218"),
			ExpectedFound: true,
		},

		// 219 – unassigned
		{
			String:             "219BANANA",
			ExpectedLength: len("219"),
			ExpectedFound: true,
		},

		{
			String:             "22BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 220 – Gambia
		{
			String:             "220BANANA",
			ExpectedLength: len("220"),
			ExpectedFound: true,
		},

		// 221 – Senegal
		{
			String:             "221BANANA",
			ExpectedLength: len("221"),
			ExpectedFound: true,
		},

		// 222 – Mauritania
		{
			String:             "222BANANA",
			ExpectedLength: len("222"),
			ExpectedFound: true,
		},

		// 223 – Mali
		{
			String:             "223BANANA",
			ExpectedLength: len("223"),
			ExpectedFound: true,
		},

		// 224 – Guinea
		{
			String:             "224BANANA",
			ExpectedLength: len("224"),
			ExpectedFound: true,
		},

		// 225 – Ivory Coast
		{
			String:             "225BANANA",
			ExpectedLength: len("225"),
			ExpectedFound: true,
		},

		// 226 – Burkina Faso
		{
			String:             "226BANANA",
			ExpectedLength: len("226"),
			ExpectedFound: true,
		},

		// 227 – Niger
		{
			String:             "227BANANA",
			ExpectedLength: len("227"),
			ExpectedFound: true,
		},

		// 228 – Togo
		{
			String:             "228BANANA",
			ExpectedLength: len("228"),
			ExpectedFound: true,
		},

		// 229 – Benin
		{
			String:             "229BANANA",
			ExpectedLength: len("229"),
			ExpectedFound: true,
		},

		{
			String:             "23BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 230 – Mauritius
		{
			String:             "230BANANA",
			ExpectedLength: len("230"),
			ExpectedFound: true,
		},

		// 231 – Liberia
		{
			String:             "231BANANA",
			ExpectedLength: len("231"),
			ExpectedFound: true,
		},

		// 232 – Sierra Leone
		{
			String:             "232BANANA",
			ExpectedLength: len("232"),
			ExpectedFound: true,
		},

		// 233 – Ghana
		{
			String:             "233BANANA",
			ExpectedLength: len("233"),
			ExpectedFound: true,
		},

		// 234 – Nigeria
		{
			String:             "234BANANA",
			ExpectedLength: len("234"),
			ExpectedFound: true,
		},

		// 235 – Chad
		{
			String:             "235BANANA",
			ExpectedLength: len("235"),
			ExpectedFound: true,
		},

		// 236 – Central African Republic
		{
			String:             "236BANANA",
			ExpectedLength: len("236"),
			ExpectedFound: true,
		},

		// 237 – Cameroon
		{
			String:             "237BANANA",
			ExpectedLength: len("237"),
			ExpectedFound: true,
		},

		// 238 – Cape Verde
		{
			String:             "238BANANA",
			ExpectedLength: len("238"),
			ExpectedFound: true,
		},

		// 239 – São Tomé and Príncipe
		{
			String:             "239BANANA",
			ExpectedLength: len("239"),
			ExpectedFound: true,
		},

		{
			String:             "24BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 240 – Equatorial Guinea
		{
			String:             "240BANANA",
			ExpectedLength: len("240"),
			ExpectedFound: true,
		},

		// 241 – Gabon
		{
			String:             "241BANANA",
			ExpectedLength: len("241"),
			ExpectedFound: true,
		},

		// 242 – Republic of the Congo
		{
			String:             "242BANANA",
			ExpectedLength: len("242"),
			ExpectedFound: true,
		},

		// 243 – Democratic Republic of the Congo
		{
			String:             "243BANANA",
			ExpectedLength: len("243"),
			ExpectedFound: true,
		},

		// 244 – Angola
		{
			String:             "244BANANA",
			ExpectedLength: len("244"),
			ExpectedFound: true,
		},

		// 245 – Guinea-Bissau
		{
			String:             "245BANANA",
			ExpectedLength: len("245"),
			ExpectedFound: true,
		},

		// 246 – British Indian Ocean Territory
		{
			String:             "246BANANA",
			ExpectedLength: len("246"),
			ExpectedFound: true,
		},

		// 247 – Ascension Island
		{
			String:             "247BANANA",
			ExpectedLength: len("247"),
			ExpectedFound: true,
		},

		// 248 – Seychelles
		{
			String:             "248BANANA",
			ExpectedLength: len("248"),
			ExpectedFound: true,
		},

		// 249 – Sudan
		{
			String:             "249BANANA",
			ExpectedLength: len("249"),
			ExpectedFound: true,
		},

		{
			String:             "25BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 250 – Rwanda
		{
			String:             "250BANANA",
			ExpectedLength: len("250"),
			ExpectedFound: true,
		},

		// 251 – Ethiopia
		{
			String:             "251BANANA",
			ExpectedLength: len("251"),
			ExpectedFound: true,
		},

		// 252 – Somalia (including Somaliland)
		{
			String:             "252BANANA",
			ExpectedLength: len("252"),
			ExpectedFound: true,
		},

		// 253 – Djibouti
		{
			String:             "253BANANA",
			ExpectedLength: len("253"),
			ExpectedFound: true,
		},

		// 254 – Kenya
		{
			String:             "254BANANA",
			ExpectedLength: len("254"),
			ExpectedFound: true,
		},

		// 255 – Tanzania
		{
			String:             "255BANANA",
			ExpectedLength: len("255"),
			ExpectedFound: true,
		},

		// 255 (24) – Zanzibar, in place of never-implemented 259
		{
			String:             "25524BANANA",
			ExpectedLength: len("255"),
			ExpectedFound: true,
		},

		// 256 – Uganda
		{
			String:             "256BANANA",
			ExpectedLength: len("256"),
			ExpectedFound: true,
		},

		// 257 – Burundi
		{
			String:             "257BANANA",
			ExpectedLength: len("257"),
			ExpectedFound: true,
		},

		// 258 – Mozambique
		{
			String:             "258BANANA",
			ExpectedLength: len("258"),
			ExpectedFound: true,
		},

		// 259 – unassigned (was intended for People's Republic of Zanzibar but never implemented – see 255 Tanzania)
		{
			String:             "259BANANA",
			ExpectedLength: len("259"),
			ExpectedFound: true,
		},

		{
			String:             "26BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 260 – Zambia
		{
			String:             "260BANANA",
			ExpectedLength: len("260"),
			ExpectedFound: true,
		},

		// 261 – Madagascar
		{
			String:             "261BANANA",
			ExpectedLength: len("261"),
			ExpectedFound: true,
		},

		// 262 – Réunion
		{
			String:             "262BANANA",
			ExpectedLength: len("262"),
			ExpectedFound: true,
		},

		// 262 (269,639) – Mayotte (formerly at 269 Comoros)
		{
			String:             "262269BANANA",
			ExpectedLength: len("262"),
			ExpectedFound: true,
		},
		{
			String:             "262639BANANA",
			ExpectedLength: len("262"),
			ExpectedFound: true,
		},

		// 263 – Zimbabwe
		{
			String:             "263BANANA",
			ExpectedLength: len("263"),
			ExpectedFound: true,
		},

		// 264 – Namibia (formerly 27 (6x) as South West Africa)
		{
			String:             "264BANANA",
			ExpectedLength: len("264"),
			ExpectedFound: true,
		},

		// 265 – Malawi
		{
			String:             "265BANANA",
			ExpectedLength: len("265"),
			ExpectedFound: true,
		},

		// 266 – Lesotho
		{
			String:             "266BANANA",
			ExpectedLength: len("266"),
			ExpectedFound: true,
		},

		// 267 – Botswana
		{
			String:             "267BANANA",
			ExpectedLength: len("267"),
			ExpectedFound: true,
		},

		// 268 – Eswatini
		{
			String:             "268BANANA",
			ExpectedLength: len("268"),
			ExpectedFound: true,
		},

		// 269 – Comoros (formerly assigned to Mayotte, now at 262)
		{
			String:             "269BANANA",
			ExpectedLength: len("269"),
			ExpectedFound: true,
		},

		// 27 – South Africa
		{
			String:             "27BANANA",
			ExpectedLength: len("27"),
			ExpectedFound: true,
		},

		{
			String:             "28BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 28x – unassigned (reserved for country code expansion)[1]
		{
			String:             "280BANANA",
			ExpectedLength: len("280"),
			ExpectedFound: true,
		},
		{
			String:             "281BANANA",
			ExpectedLength: len("281"),
			ExpectedFound: true,
		},
		{
			String:             "282BANANA",
			ExpectedLength: len("282"),
			ExpectedFound: true,
		},
		{
			String:             "283BANANA",
			ExpectedLength: len("283"),
			ExpectedFound: true,
		},
		{
			String:             "284BANANA",
			ExpectedLength: len("284"),
			ExpectedFound: true,
		},
		{
			String:             "285BANANA",
			ExpectedLength: len("285"),
			ExpectedFound: true,
		},
		{
			String:             "286BANANA",
			ExpectedLength: len("286"),
			ExpectedFound: true,
		},
		{
			String:             "287BANANA",
			ExpectedLength: len("287"),
			ExpectedFound: true,
		},
		{
			String:             "288BANANA",
			ExpectedLength: len("288"),
			ExpectedFound: true,
		},
		{
			String:             "289BANANA",
			ExpectedLength: len("289"),
			ExpectedFound: true,
		},

		{
			String:             "29BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 290 – Saint Helena
		{
			String:             "290BANANA",
			ExpectedLength: len("290"),
			ExpectedFound: true,
		},

		// 290 (8) – Tristan da Cunha
		{
			String:             "2908BANANA",
			ExpectedLength: len("290"),
			ExpectedFound: true,
		},

		// 291 – Eritrea
		{
			String:             "291BANANA",
			ExpectedLength: len("291"),
			ExpectedFound: true,
		},

		// 292 – unassigned
		{
			String:             "292BANANA",
			ExpectedLength: len("292"),
			ExpectedFound: true,
		},

		// 293 – unassigned
		{
			String:             "293BANANA",
			ExpectedLength: len("293"),
			ExpectedFound: true,
		},

		// 294 – unassigned
		{
			String:             "294BANANA",
			ExpectedLength: len("294"),
			ExpectedFound: true,
		},

		// 295 – unassigned (formerly assigned to San Marino, now at 378)
		{
			String:             "295BANANA",
			ExpectedLength: len("295"),
			ExpectedFound: true,
		},

		// 296 – unassigned
		{
			String:             "296BANANA",
			ExpectedLength: len("296"),
			ExpectedFound: true,
		},

		// 297 – Aruba
		{
			String:             "297BANANA",
			ExpectedLength: len("297"),
			ExpectedFound: true,
		},

		// 298 – Faroe Islands
		{
			String:             "298BANANA",
			ExpectedLength: len("298"),
			ExpectedFound: true,
		},

		// 299 – Greenland
		{
			String:             "299BANANA",
			ExpectedLength: len("299"),
			ExpectedFound: true,
		},

		{
			String:             "3BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 30 – Greece
		{
			String:             "30BANANA",
			ExpectedLength: len("30"),
			ExpectedFound: true,
		},

		// 31 – Netherlands
		{
			String:             "31BANANA",
			ExpectedLength: len("31"),
			ExpectedFound: true,
		},

		// 32 – Belgium
		{
			String:             "32BANANA",
			ExpectedLength: len("32"),
			ExpectedFound: true,
		},

		// 33 – France
		{
			String:             "33BANANA",
			ExpectedLength: len("33"),
			ExpectedFound: true,
		},

		// 34 – Spain
		{
			String:             "34BANANA",
			ExpectedLength: len("34"),
			ExpectedFound: true,
		},

		{
			String:             "35BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 350 – Gibraltar
		{
			String:             "350BANANA",
			ExpectedLength: len("350"),
			ExpectedFound: true,
		},

		// 351 – Portugal
		{
			String:             "351BANANA",
			ExpectedLength: len("351"),
			ExpectedFound: true,
		},

		// 351 (291) – Madeira (landlines only)
		{
			String:             "351291BANANA",
			ExpectedLength: len("351"),
			ExpectedFound: true,
		},

		// 351 (292) – Azores (landlines only, Horta, Azores area)
		{
			String:             "351292BANANA",
			ExpectedLength: len("351"),
			ExpectedFound: true,
		},

		// 351 (295) – Azores (landlines only, Angra do Heroísmo area)
		{
			String:             "351292BANANA",
			ExpectedLength: len("351"),
			ExpectedFound: true,
		},

		// 351 (296) – Azores (landlines only, Ponta Delgada and São Miguel Island area)
		{
			String:             "351292BANANA",
			ExpectedLength: len("351"),
			ExpectedFound: true,
		},

		// 352 – Luxembourg
		{
			String:             "352BANANA",
			ExpectedLength: len("352"),
			ExpectedFound: true,
		},

		// 353 – Ireland
		{
			String:             "353BANANA",
			ExpectedLength: len("353"),
			ExpectedFound: true,
		},

		// 354 – Iceland
		{
			String:             "354BANANA",
			ExpectedLength: len("354"),
			ExpectedFound: true,
		},

		// 355 – Albania
		{
			String:             "355BANANA",
			ExpectedLength: len("355"),
			ExpectedFound: true,
		},

		// 356 – Malta
		{
			String:             "356BANANA",
			ExpectedLength: len("356"),
			ExpectedFound: true,
		},

		// 357 – Cyprus (including Akrotiri and Dhekelia)
		{
			String:             "357BANANA",
			ExpectedLength: len("357"),
			ExpectedFound: true,
		},

		// 358 – Finland
		{
			String:             "358BANANA",
			ExpectedLength: len("358"),
			ExpectedFound: true,
		},

		// 358 (18) – Åland
		{
			String:             "35818BANANA",
			ExpectedLength: len("358"),
			ExpectedFound: true,
		},

		// 359 – Bulgaria
		{
			String:             "359BANANA",
			ExpectedLength: len("359"),
			ExpectedFound: true,
		},

		// 36 – Hungary (formerly assigned to Turkey, now at 90)
		{
			String:             "36BANANA",
			ExpectedLength: len("36"),
			ExpectedFound: true,
		},

		{
			String:             "37BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 370 – Lithuania
		{
			String:             "370BANANA",
			ExpectedLength: len("370"),
			ExpectedFound: true,
		},

		// 371 – Latvia
		{
			String:             "371BANANA",
			ExpectedLength: len("371"),
			ExpectedFound: true,
		},

		// 372 – Estonia
		{
			String:             "372BANANA",
			ExpectedLength: len("372"),
			ExpectedFound: true,
		},

		// 373 – Moldova
		{
			String:             "373BANANA",
			ExpectedLength: len("373"),
			ExpectedFound: true,
		},

		// 374 – Armenia
		{
			String:             "374BANANA",
			ExpectedLength: len("374"),
			ExpectedFound: true,
		},

		// 375 – Belarus
		{
			String:             "375BANANA",
			ExpectedLength: len("375"),
			ExpectedFound: true,
		},

		// 376 – Andorra (formerly 33 628)
		{
			String:             "376BANANA",
			ExpectedLength: len("376"),
			ExpectedFound: true,
		},

		// 377 – Monaco (formerly 33 93)
		{
			String:             "377BANANA",
			ExpectedLength: len("377"),
			ExpectedFound: true,
		},

		// 378 – San Marino (interchangeably with 39 0549; earlier was allocated 295 but never used)
		{
			String:             "378BANANA",
			ExpectedLength: len("378"),
			ExpectedFound: true,
		},

		// 379 – Vatican City (assigned but uses 39 06698).
		{
			String:             "379BANANA",
			ExpectedLength: len("379"),
			ExpectedFound: true,
		},

		{
			String:             "38BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 380 – Ukraine
		{
			String:             "380BANANA",
			ExpectedLength: len("380"),
			ExpectedFound: true,
		},

		// 381 – Serbia
		{
			String:             "381BANANA",
			ExpectedLength: len("381"),
			ExpectedFound: true,
		},

		// 382 – Montenegro
		{
			String:             "382BANANA",
			ExpectedLength: len("382"),
			ExpectedFound: true,
		},

		// 383 – Kosovo
		{
			String:             "383BANANA",
			ExpectedLength: len("383"),
			ExpectedFound: true,
		},

		// 384 – unassigned
		{
			String:             "384BANANA",
			ExpectedLength: len("384"),
			ExpectedFound: true,
		},

		// 385 – Croatia
		{
			String:             "385BANANA",
			ExpectedLength: len("385"),
			ExpectedFound: true,
		},

		// 386 – Slovenia
		{
			String:             "386BANANA",
			ExpectedLength: len("386"),
			ExpectedFound: true,
		},

		// 387 – Bosnia and Herzegovina
		{
			String:             "387BANANA",
			ExpectedLength: len("387"),
			ExpectedFound: true,
		},

		// 388 – unassigned (formerly assigned to the European Telephony Numbering Space)[1][2]
		{
			String:             "388BANANA",
			ExpectedLength: len("388"),
			ExpectedFound: true,
		},

		// 389 – North Macedonia
		{
			String:             "389BANANA",
			ExpectedLength: len("389"),
			ExpectedFound: true,
		},

		// 39 – Italy
		{
			String:             "39BANANA",
			ExpectedLength: len("39"),
			ExpectedFound: true,
		},

		// 39 (0549) – San Marino (interchangeably with 378)
		{
			String:             "390549BANANA",
			ExpectedLength: len("39"),
			ExpectedFound: true,
		},

		// 39 (06 698) – Vatican City (assigned 379 but not in use)
		{
			String:             "3906698BANANA",
			ExpectedLength: len("39"),
			ExpectedFound: true,
		},

		{
			String:             "4BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 40 – Romania
		{
			String:             "40BANANA",
			ExpectedLength: len("40"),
			ExpectedFound: true,
		},

		// 41 –  Switzerland
		{
			String:             "41BANANA",
			ExpectedLength: len("41"),
			ExpectedFound: true,
		},

		// 41 (91) – Campione d'Italia, an Italian enclave. 91 is the prefix for the Swiss canton Ticino in which the enclave resides. Its phone system is fully integrated into the Swiss system.
		{
			String:             "4191BANANA",
			ExpectedLength: len("41"),
			ExpectedFound: true,
		},

		// 42 – formerly assigned to Czechoslovakia, later to its breakup successors (CZ, SK) until 1997
		{
			String:             "42BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 420 – Czech Republic
		{
			String:             "420BANANA",
			ExpectedLength: len("420"),
			ExpectedFound: true,
		},

		// 421 – Slovakia
		{
			String:             "421BANANA",
			ExpectedLength: len("421"),
			ExpectedFound: true,
		},

		// 422 – unassigned
		{
			String:             "422BANANA",
			ExpectedLength: len("422"),
			ExpectedFound: true,
		},

		// 423 – Liechtenstein (formerly at 41 (75))
		{
			String:             "423BANANA",
			ExpectedLength: len("423"),
			ExpectedFound: true,
		},

		// 424 – unassigned
		{
			String:             "424BANANA",
			ExpectedLength: len("424"),
			ExpectedFound: true,
		},

		// 425 – unassigned
		{
			String:             "425BANANA",
			ExpectedLength: len("425"),
			ExpectedFound: true,
		},

		// 426 – unassigned
		{
			String:             "426BANANA",
			ExpectedLength: len("426"),
			ExpectedFound: true,
		},

		// 427 – unassigned
		{
			String:             "427BANANA",
			ExpectedLength: len("427"),
			ExpectedFound: true,
		},

		// 428 – unassigned
		{
			String:             "428BANANA",
			ExpectedLength: len("428"),
			ExpectedFound: true,
		},

		// 429 – unassigned
		{
			String:             "429BANANA",
			ExpectedLength: len("429"),
			ExpectedFound: true,
		},

		// 43 – Austria
		{
			String:             "43BANANA",
			ExpectedLength: len("43"),
			ExpectedFound: true,
		},

		// 44 – United Kingdom
		{
			String:             "44BANANA",
			ExpectedLength: len("44"),
			ExpectedFound: true,
		},

		// 44 (1481) – Guernsey
		{
			String:             "441481BANANA",
			ExpectedLength: len("44"),
			ExpectedFound: true,
		},

		// 44 (1534) – Jersey
		{
			String:             "441534BANANA",
			ExpectedLength: len("44"),
			ExpectedFound: true,
		},

		// 44 (1624) – Isle of Man
		{
			String:             "441624BANANA",
			ExpectedLength: len("44"),
			ExpectedFound: true,
		},

		// 45 – Denmark
		{
			String:             "45BANANA",
			ExpectedLength: len("45"),
			ExpectedFound: true,
		},

		// 46 – Sweden
		{
			String:             "46BANANA",
			ExpectedLength: len("46"),
			ExpectedFound: true,
		},

		// 47 – Norway
		{
			String:             "47BANANA",
			ExpectedLength: len("47"),
			ExpectedFound: true,
		},

		// 47 (79) – Svalbard
		{
			String:             "4779BANANA",
			ExpectedLength: len("47"),
			ExpectedFound: true,
		},

		// 48 – Poland
		{
			String:             "48BANANA",
			ExpectedLength: len("48"),
			ExpectedFound: true,
		},

		// 49 – Germany
		{
			String:             "49BANANA",
			ExpectedLength: len("49"),
			ExpectedFound: true,
		},

		{
			String:             "5BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		{
			String:             "50BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 500 – Falkland Islands
		{
			String:             "500BANANA",
			ExpectedLength: len("500"),
			ExpectedFound: true,
		},

		// 500 – South Georgia and the South Sandwich Islands
		{
			String:             "500BANANA",
			ExpectedLength: len("500"),
			ExpectedFound: true,
		},

		// 501 – Belize
		{
			String:             "501BANANA",
			ExpectedLength: len("501"),
			ExpectedFound: true,
		},

		// 502 – Guatemala
		{
			String:             "502BANANA",
			ExpectedLength: len("502"),
			ExpectedFound: true,
		},

		// 503 – El Salvador
		{
			String:             "503BANANA",
			ExpectedLength: len("503"),
			ExpectedFound: true,
		},

		// 504 – Honduras
		{
			String:             "504BANANA",
			ExpectedLength: len("504"),
			ExpectedFound: true,
		},

		// 505 – Nicaragua
		{
			String:             "505BANANA",
			ExpectedLength: len("505"),
			ExpectedFound: true,
		},

		// 506 – Costa Rica
		{
			String:             "506BANANA",
			ExpectedLength: len("506"),
			ExpectedFound: true,
		},

		// 507 – Panama
		{
			String:             "507BANANA",
			ExpectedLength: len("507"),
			ExpectedFound: true,
		},

		// 508 – Saint-Pierre and Miquelon
		{
			String:             "508BANANA",
			ExpectedLength: len("508"),
			ExpectedFound: true,
		},

		// 509 – Haiti
		{
			String:             "509BANANA",
			ExpectedLength: len("509"),
			ExpectedFound: true,
		},

		// 51 – Peru
		{
			String:             "51BANANA",
			ExpectedLength: len("51"),
			ExpectedFound: true,
		},

		// 52 – Mexico
		{
			String:             "52BANANA",
			ExpectedLength: len("52"),
			ExpectedFound: true,
		},

		// 53 – Cuba
		{
			String:             "53BANANA",
			ExpectedLength: len("53"),
			ExpectedFound: true,
		},

		// 54 – Argentina
		{
			String:             "54BANANA",
			ExpectedLength: len("54"),
			ExpectedFound: true,
		},

		// 55 – Brazil
		{
			String:             "55BANANA",
			ExpectedLength: len("55"),
			ExpectedFound: true,
		},

		// 56 – Chile
		{
			String:             "56BANANA",
			ExpectedLength: len("56"),
			ExpectedFound: true,
		},

		// 57 – Colombia
		{
			String:             "57BANANA",
			ExpectedLength: len("57"),
			ExpectedFound: true,
		},

		// 58 – Venezuela
		{
			String:             "58BANANA",
			ExpectedLength: len("58"),
			ExpectedFound: true,
		},

		{
			String: "59BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 590 – Guadeloupe (including Saint Barthélemy, Saint Martin)
		{
			String:             "590BANANA",
			ExpectedLength: len("590"),
			ExpectedFound: true,
		},

		// 591 – Bolivia
		{
			String:             "591BANANA",
			ExpectedLength: len("591"),
			ExpectedFound: true,
		},

		// 592 – Guyana
		{
			String:             "592BANANA",
			ExpectedLength: len("592"),
			ExpectedFound: true,
		},

		// 593 – Ecuador
		{
			String:             "593BANANA",
			ExpectedLength: len("593"),
			ExpectedFound: true,
		},

		// 594 – French Guiana
		{
			String:             "594BANANA",
			ExpectedLength: len("594"),
			ExpectedFound: true,
		},

		// 595 – Paraguay
		{
			String:             "595BANANA",
			ExpectedLength: len("595"),
			ExpectedFound: true,
		},

		// 596 – Martinique (formerly assigned to Peru, now 51)
		{
			String:             "596BANANA",
			ExpectedLength: len("596"),
			ExpectedFound: true,
		},

		// 597 – Suriname
		{
			String:             "597BANANA",
			ExpectedLength: len("597"),
			ExpectedFound: true,
		},

		// 598 – Uruguay
		{
			String:             "598BANANA",
			ExpectedLength: len("598"),
			ExpectedFound: true,
		},

		// 599 – Former Netherlands Antilles, now grouped as follows:
		{
			String:             "599BANANA",
			ExpectedLength: len("599"),
			ExpectedFound: true,
		},

		// 599 3 – Sint Eustatius
		{
			String:             "5993BANANA",
			ExpectedLength: len("599"),
			ExpectedFound: true,
		},

		// 599 4 – Saba
		{
			String:             "5994BANANA",
			ExpectedLength: len("599"),
			ExpectedFound: true,
		},

		// 599 5 – unassigned (formerly assigned to Sint Maarten, now included in NANP as 1 (721))
		{
			String:             "5995BANANA",
			ExpectedLength: len("599"),
			ExpectedFound: true,
		},

		// 599 7 – Bonaire
		{
			String:             "5997BANANA",
			ExpectedLength: len("599"),
			ExpectedFound: true,
		},

		// 599 8 – unassigned (formerly assigned to Aruba, now at 297)
		{
			String:             "5998BANANA",
			ExpectedLength: len("599"),
			ExpectedFound: true,
		},

		// 599 9 – Curaçao
		{
			String:             "5999BANANA",
			ExpectedLength: len("599"),
			ExpectedFound: true,
		},

		{
			String:             "6BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 60 – Malaysia
		{
			String:             "60BANANA",
			ExpectedLength: len("60"),
			ExpectedFound: true,
		},

		// 61 – Australia (see also 672 below)
		{
			String:             "61BANANA",
			ExpectedLength: len("61"),
			ExpectedFound: true,
		},

		// 61 (8 9162) – Cocos Islands
		{
			String:             "6189162BANANA",
			ExpectedLength: len("61"),
			ExpectedFound: true,
		},

		// 61 (8 9164) – Christmas Island
		{
			String:             "6189164BANANA",
			ExpectedLength: len("61"),
			ExpectedFound: true,
		},

		// 62 – Indonesia
		{
			String:             "62BANANA",
			ExpectedLength: len("62"),
			ExpectedFound: true,
		},

		// 63 – Philippines
		{
			String:             "63BANANA",
			ExpectedLength: len("63"),
			ExpectedFound: true,
		},

		// 64 – New Zealand
		// 64 – Pitcairn Islands
		{
			String:             "64BANANA",
			ExpectedLength: len("64"),
			ExpectedFound: true,
		},

		// 65 – Singapore
		{
			String:             "65BANANA",
			ExpectedLength: len("65"),
			ExpectedFound: true,
		},

		// 66 – Thailand
		{
			String:             "66BANANA",
			ExpectedLength: len("66"),
			ExpectedFound: true,
		},

		{
			String:             "67BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 670 – East Timor (formerly 62/39 during the Indonesian occupation; formerly assigned to Northern Mariana Islands, now part of NANP as 1 (670))
		{
			String:             "670BANANA",
			ExpectedLength: len("670"),
			ExpectedFound: true,
		},

		// 671 – unassigned (formerly assigned to Guam, now part of NANP as 1 (671))
		{
			String:             "671BANANA",
			ExpectedLength: len("671"),
			ExpectedFound: true,
		},

		// 672 – Australian External Territories (see also 61 Australia above); formerly assigned to Portuguese Timor (see 670)
		{
			String:             "672BANANA",
			ExpectedLength: len("672"),
			ExpectedFound: true,
		},

		// 672 (1) – Australia Australian Antarctic Territory
		{
			String:             "6721BANANA",
			ExpectedLength: len("672"),
			ExpectedFound: true,
		},

		// 672 (3) – Norfolk Island
		{
			String:             "6723BANANA",
			ExpectedLength: len("672"),
			ExpectedFound: true,
		},

		// 673 – Brunei
		{
			String:             "673BANANA",
			ExpectedLength: len("673"),
			ExpectedFound: true,
		},

		// 674 – Nauru
		{
			String:             "674BANANA",
			ExpectedLength: len("674"),
			ExpectedFound: true,
		},

		// 675 – Papua New Guinea
		{
			String:             "675BANANA",
			ExpectedLength: len("675"),
			ExpectedFound: true,
		},

		// 676 – Tonga
		{
			String:             "676BANANA",
			ExpectedLength: len("676"),
			ExpectedFound: true,
		},

		// 677 – Solomon Islands
		{
			String:             "677BANANA",
			ExpectedLength: len("677"),
			ExpectedFound: true,
		},

		// 678 – Vanuatu
		{
			String:             "678BANANA",
			ExpectedLength: len("678"),
			ExpectedFound: true,
		},

		// 679 – Fiji
		{
			String:             "679BANANA",
			ExpectedLength: len("679"),
			ExpectedFound: true,
		},

		{
			String:             "68BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 680 – Palau
		// 681 – Wallis and Futuna
		// 682 – Cook Islands
		// 683 – Niue
		// 684 – unassigned (formerly assigned to American Samoa, now part of NANP as 1 (684))
		// 685 – Samoa
		// 686 – Kiribati
		// 687 – New Caledonia
		// 688 – Tuvalu
		// 689 – French Polynesia

		{
			String:             "69BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 690 – Tokelau
		// 691 – Federated States of Micronesia
		// 692 – Marshall Islands
		// 693 – unassigned
		// 694 – unassigned
		// 695 – unassigned
		// 696 – unassigned
		// 697 – unassigned
		// 698 – unassigned
		// 699 – unassigned

		{
			String:             "7BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},

		// 7 (1–5, 8–9) – Russia
		{
			String:             "71BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "72BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "73BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "74BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "75BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "78BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "79BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},

		// 7 (840, 940) – Abkhazia (formerly 995 (44))
		{
			String:             "7840BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "7940BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},

		// 7 (850, 929) – South Ossetia (formerly 995 (34))
		{
			String:             "7850BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "7929BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},

		// 7 (0, 6–7) – Kazakhstan (reserved 997 but abandoned in November 2023)
		{
			String:             "70BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "76BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},
		{
			String:             "77BANANA",
			ExpectedLength: len("7"),
			ExpectedFound: true,
		},

		{
			String:             "8BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		{
			String:             "80BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 800 – Universal International Freephone Service
		{
			String:             "800BANANA",
			ExpectedLength: len("800"),
			ExpectedFound: true,
		},

		// 801 – unassigned
		{
			String:             "801BANANA",
			ExpectedLength: len("801"),
			ExpectedFound: true,
		},

		// 802 – unassigned
		{
			String:             "802BANANA",
			ExpectedLength: len("802"),
			ExpectedFound: true,
		},

		// 803 – unassigned
		{
			String:             "803BANANA",
			ExpectedLength: len("803"),
			ExpectedFound: true,
		},

		// 804 – unassigned
		{
			String:             "804BANANA",
			ExpectedLength: len("804"),
			ExpectedFound: true,
		},

		// 805 – unassigned
		{
			String:             "805BANANA",
			ExpectedLength: len("805"),
			ExpectedFound: true,
		},

		// 806 – unassigned
		{
			String:             "806BANANA",
			ExpectedLength: len("806"),
			ExpectedFound: true,
		},

		// 807 – unassigned
		{
			String:             "807BANANA",
			ExpectedLength: len("807"),
			ExpectedFound: true,
		},

		// 808 – Universal International Shared Cost Numbers
		{
			String:             "808BANANA",
			ExpectedLength: len("808"),
			ExpectedFound: true,
		},

		// 809 – unassigned
		{
			String:             "809BANANA",
			ExpectedLength: len("809"),
			ExpectedFound: true,
		},

		// 81 – Japan
		{
			String:             "81BANANA",
			ExpectedLength: len("81"),
			ExpectedFound: true,
		},

		// 82 – South Korea
		{
			String:             "82BANANA",
			ExpectedLength: len("82"),
			ExpectedFound: true,
		},

		// 83x – unassigned (reserved for country code expansion)[1]
		{
			String:             "83BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String:             "830BANANA",
			ExpectedLength: len("830"),
			ExpectedFound: true,
		},
		{
			String:             "831BANANA",
			ExpectedLength: len("831"),
			ExpectedFound: true,
		},
		{
			String:             "832BANANA",
			ExpectedLength: len("832"),
			ExpectedFound: true,
		},
		{
			String:             "833BANANA",
			ExpectedLength: len("833"),
			ExpectedFound: true,
		},
		{
			String:             "834BANANA",
			ExpectedLength: len("834"),
			ExpectedFound: true,
		},
		{
			String:             "835BANANA",
			ExpectedLength: len("835"),
			ExpectedFound: true,
		},
		{
			String:             "836BANANA",
			ExpectedLength: len("836"),
			ExpectedFound: true,
		},
		{
			String:             "837BANANA",
			ExpectedLength: len("837"),
			ExpectedFound: true,
		},
		{
			String:             "838BANANA",
			ExpectedLength: len("838"),
			ExpectedFound: true,
		},
		{
			String:             "839BANANA",
			ExpectedLength: len("839"),
			ExpectedFound: true,
		},

		// 84 – Vietnam
		// 850 – North Korea
		// 851 – unassigned
		// 852 – Hong Kong
		// 853 – Macau
		// 854 – unassigned
		// 855 – Cambodia
		// 856 – Laos
		// 857 – unassigned (formerly assigned to ANAC satellite service)
		// 858 – unassigned (formerly assigned to ANAC satellite service)
		// 859 – unassigned
		// 86 – China
		// 870 – Global Mobile Satellite System (Inmarsat)
		// 871 – unassigned (formerly assigned to Inmarsat Atlantic East, discontinued in 2008)
		// 872 – unassigned (formerly assigned to Inmarsat Pacific, discontinued in 2008)
		// 873 – unassigned (formerly assigned to Inmarsat Indian, discontinued in 2008)
		// 874 – unassigned (formerly assigned to Inmarsat Atlantic West, discontinued in 2008)
		// 875 – unassigned (reserved for future maritime mobile service)
		// 876 – unassigned (reserved for future maritime mobile service)
		// 877 – unassigned (reserved for future maritime mobile service)
		// 878 – unassigned (formerly used for Universal Personal Telecommunications Service, discontinued in 2022)
		// 879 – unassigned (reserved for national non-commercial purposes)

		{
			String:             "88BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 880 – Bangladesh
		{
			String:             "880BANANA",
			ExpectedLength: len("880"),
			ExpectedFound: true,
		},

		// 881 – Global Mobile Satellite System
		{
			String:             "881BANANA",
			ExpectedLength: len("881"),
			ExpectedFound: true,
		},

		// 882 – International Networks
		{
			String:             "882BANANA",
			ExpectedLength: len("882"),
			ExpectedFound: true,
		},

		// 883 – International Networks
		{
			String:             "883BANANA",
			ExpectedLength: len("883"),
			ExpectedFound: true,
		},

		// 884 – unassigned
		{
			String:             "884BANANA",
			ExpectedLength: len("884"),
			ExpectedFound: true,
		},

		// 885 – unassigned
		{
			String:             "885BANANA",
			ExpectedLength: len("885"),
			ExpectedFound: true,
		},


		// 886 – Taiwan
		{
			String:             "886BANANA",
			ExpectedLength: len("886"),
			ExpectedFound: true,
		},

		// 887 – unassigned
		{
			String:             "887BANANA",
			ExpectedLength: len("887"),
			ExpectedFound: true,
		},

		// 888 – unassigned (formerly assigned to OCHA for Telecommunications for Disaster Relief service)
		{
			String:             "888BANANA",
			ExpectedLength: len("888"),
			ExpectedFound: true,
		},

		// 889 – unassigned
		{
			String:             "889BANANA",
			ExpectedLength: len("889"),
			ExpectedFound: true,
		},

		// 89x – unassigned (reserved for country code expansion)
		{
			String:             "89BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},
		{
			String:             "890BANANA",
			ExpectedLength: len("890"),
			ExpectedFound: true,
		},
		{
			String:             "891BANANA",
			ExpectedLength: len("891"),
			ExpectedFound: true,
		},
		{
			String:             "892BANANA",
			ExpectedLength: len("892"),
			ExpectedFound: true,
		},
		{
			String:             "893BANANA",
			ExpectedLength: len("893"),
			ExpectedFound: true,
		},
		{
			String:             "894BANANA",
			ExpectedLength: len("894"),
			ExpectedFound: true,
		},
		{
			String:             "895BANANA",
			ExpectedLength: len("895"),
			ExpectedFound: true,
		},
		{
			String:             "896BANANA",
			ExpectedLength: len("896"),
			ExpectedFound: true,
		},
		{
			String:             "897BANANA",
			ExpectedLength: len("897"),
			ExpectedFound: true,
		},
		{
			String:             "898BANANA",
			ExpectedLength: len("898"),
			ExpectedFound: true,
		},
		{
			String:             "899BANANA",
			ExpectedLength: len("899"),
			ExpectedFound: true,
		},


		{
			String:             "9BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 90 – Turkey
		{
			String:             "90BANANA",
			ExpectedLength: len("90"),
			ExpectedFound: true,
		},

		// 90 (392) – Northern Cyprus
		{
			String:             "90392BANANA",
			ExpectedLength: len("90"),
			ExpectedFound: true,
		},
		// 91 – India
		{
			String:             "91BANANA",
			ExpectedLength: len("91"),
			ExpectedFound: true,
		},

		// 91 (191) – Jammu
		{
			String:             "91191BANANA",
			ExpectedLength: len("91"),
			ExpectedFound: true,
		},

		// 91 (194) – Kashmir
		{
			String:             "91194BANANA",
			ExpectedLength: len("91"),
			ExpectedFound: true,
		},

		// 92 – Pakistan
		{
			String:             "92BANANA",
			ExpectedLength: len("92"),
			ExpectedFound: true,
		},

		// 92 (581) – Gilgit Baltistan
		{
			String:             "92581BANANA",
			ExpectedLength: len("92"),
			ExpectedFound: true,
		},

		// 92 (582) – Azad Kashmir
		{
			String:             "92582BANANA",
			ExpectedLength: len("92"),
			ExpectedFound: true,
		},

		// 93 – Afghanistan
		{
			String:             "93BANANA",
			ExpectedLength: len("93"),
			ExpectedFound: true,
		},

		// 94 – Sri Lanka
		{
			String:             "94BANANA",
			ExpectedLength: len("94"),
			ExpectedFound: true,
		},

		// 95 – Myanmar
		{
			String:             "95BANANA",
			ExpectedLength: len("95"),
			ExpectedFound: true,
		},

		{
			String:             "96BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 960 – Maldives
		{
			String:             "960BANANA",
			ExpectedLength: len("960"),
			ExpectedFound: true,
		},

		// 961 – Lebanon
		{
			String:             "961BANANA",
			ExpectedLength: len("961"),
			ExpectedFound: true,
		},

		// 962 – Jordan
		{
			String:             "962BANANA",
			ExpectedLength: len("962"),
			ExpectedFound: true,
		},

		// 963 – Syria
		{
			String:             "963BANANA",
			ExpectedLength: len("963"),
			ExpectedFound: true,
		},

		// 964 – Iraq
		{
			String:             "964BANANA",
			ExpectedLength: len("964"),
			ExpectedFound: true,
		},

		// 965 – Kuwait
		{
			String:             "965BANANA",
			ExpectedLength: len("965"),
			ExpectedFound: true,
		},

		// 966 – Saudi Arabia
		{
			String:             "966BANANA",
			ExpectedLength: len("966"),
			ExpectedFound: true,
		},

		// 967 – Yemen
		{
			String:             "967BANANA",
			ExpectedLength: len("967"),
			ExpectedFound: true,
		},

		// 968 – Oman
		{
			String:             "968BANANA",
			ExpectedLength: len("968"),
			ExpectedFound: true,
		},

		// 969 – unassigned (formerly assigned to South Yemen until its unification with North Yemen, now part of 967 Yemen)
		{
			String:             "969BANANA",
			ExpectedLength: len("969"),
			ExpectedFound: true,
		},

		{
			String:             "97BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 970 – Palestine (interchangeably with 972)
		{
			String:             "970BANANA",
			ExpectedLength: len("970"),
			ExpectedFound: true,
		},

		// 971 – United Arab Emirates
		{
			String:             "971BANANA",
			ExpectedLength: len("971"),
			ExpectedFound: true,
		},

		// 972 – Israel (also Palestine, interchangeably with 970)
		{
			String:             "972BANANA",
			ExpectedLength: len("972"),
			ExpectedFound: true,
		},

		// 973 – Bahrain
		{
			String:             "973BANANA",
			ExpectedLength: len("973"),
			ExpectedFound: true,
		},

		// 974 – Qatar
		{
			String:             "974BANANA",
			ExpectedLength: len("974"),
			ExpectedFound: true,
		},

		// 975 – Bhutan
		{
			String:             "975BANANA",
			ExpectedLength: len("975"),
			ExpectedFound: true,
		},

		// 976 – Mongolia
		{
			String:             "976BANANA",
			ExpectedLength: len("976"),
			ExpectedFound: true,
		},

		// 977 – Nepal
		{
			String:             "977BANANA",
			ExpectedLength: len("977"),
			ExpectedFound: true,
		},

		// 978 – unassigned (formerly assigned to Dubai, now part of 971 United Arab Emirates)
		{
			String:             "978BANANA",
			ExpectedLength: len("978"),
			ExpectedFound: true,
		},

		// 979 – Universal International Premium Rate Service (UIPRS); (formerly assigned to Abu Dhabi, now part of 971 United Arab Emirates)
		{
			String:             "979BANANA",
			ExpectedLength: len("979"),
			ExpectedFound: true,
		},

		// 98 – Iran
		{
			String:             "98BANANA",
			ExpectedLength: len("98"),
			ExpectedFound: true,
		},

		{
			String:             "99BANANA",
			ExpectedLength: 0,
			ExpectedFound: false,
		},

		// 990 – unassigned
		{
			String:             "990BANANA",
			ExpectedLength: len("990"),
			ExpectedFound: true,
		},

		// 991 – unassigned (formerly used for International Telecommunications Public Correspondence Service)
		{
			String:             "991BANANA",
			ExpectedLength: len("991"),
			ExpectedFound: true,
		},

		// 992 – Tajikistan
		{
			String:             "992BANANA",
			ExpectedLength: len("992"),
			ExpectedFound: true,
		},

		// 993 – Turkmenistan
		{
			String:             "993BANANA",
			ExpectedLength: len("993"),
			ExpectedFound: true,
		},

		// 994 – Azerbaijan
		{
			String:             "994BANANA",
			ExpectedLength: len("994"),
			ExpectedFound: true,
		},

		// 995 – Georgia
		{
			String:             "995BANANA",
			ExpectedLength: len("995"),
			ExpectedFound: true,
		},

		// 995 (34) – formerly South Ossetia (now 7 (850, 929))
		{
			String:             "99534BANANA",
			ExpectedLength: len("995"),
			ExpectedFound: true,
		},

		// 995 (44) – formerly Abkhazia[4][5] (now 7 (840, 940))
		{
			String:             "99544BANANA",
			ExpectedLength: len("995"),
			ExpectedFound: true,
		},

		// 996 – Kyrgyzstan
		{
			String:             "996BANANA",
			ExpectedLength: len("996"),
			ExpectedFound: true,
		},

		// 997 – Kazakhstan (reserved but abandoned in November 2023; uses 7 (6xx, 7xx))
		{
			String:             "997BANANA",
			ExpectedLength: len("997"),
			ExpectedFound: true,
		},

		// 998 – Uzbekistan
		{
			String:             "998BANANA",
			ExpectedLength: len("998"),
			ExpectedFound: true,
		},

		// 999 – unassigned (reserved for future global service)
		{
			String:             "999BANANA",
			ExpectedLength: len("999"),
			ExpectedFound: true,
		},
	}

	for _, areaCode := range areaCodesCA {
		{
			test := struct{
				String string
				ExpectedLength int
				ExpectedFound bool
			}{
				String: "1" + areaCode,
				ExpectedLength: 1,
				ExpectedFound: true,
			}

			tests = append(tests, test)
		}

		{
			test := struct{
				String string
				ExpectedLength int
				ExpectedFound bool
			}{
				String: "1" + areaCode + "5551234",
				ExpectedLength: 1,
				ExpectedFound: true,
			}

			tests = append(tests, test)
		}
	}

	for _, areaCode := range areaCodesUS {
		{
			test := struct{
				String string
				ExpectedLength int
				ExpectedFound bool
			}{
				String: "1" + areaCode,
				ExpectedLength: 1,
				ExpectedFound: true,
			}

			tests = append(tests, test)
		}

		{
			test := struct{
				String string
				ExpectedLength int
				ExpectedFound bool
			}{
				String: "1" + areaCode + "5551234",
				ExpectedLength: 1,
				ExpectedFound: true,
			}

			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		actualLength, actualFound := peekCountryCode(test.String)

		{
			actual := actualFound
			expected := test.ExpectedFound

			if expected != actual {
				t.Errorf("For test #%d, the actual 'found' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			actual := actualLength
			expected := test.ExpectedLength

			if expected != actual {
				t.Errorf("For test #%d, the actual 'length' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
