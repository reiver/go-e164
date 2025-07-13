package e164

var (
	// The values in this were taken from:
	//	https://en.wikipedia.org/wiki/List_of_telephone_country_codes
	//
	// The ITU-T document for this returned an error.
	//
	// THIS LIST WILL NEED TO BE UPDATED WHENEVER THE ITU-T MAKES AN UPDATE.
	countryCodes = []string{
		"1",

		"20",  // Egypt
		"210", // unassigned
		"211", // South Sudan
		"212", // Morocco (including Western Sahara)
		"213", // Algeria
		"214", // unassigned
		"215", // unassigned
		"216", // Tunisia
		"217", // unassigned
		"218", // Libya
		"219", // unassigned
		"220", // Gambia
		"221", // Senegal
		"222", // Mauritania
		"223", // Mali
		"224", // Guinea
		"225", // Ivory Coast
		"226", // Burkina Faso
		"227", // Niger
		"228", // Togo
		"229", // Benin
		"230", // Mauritius
		"231", // Liberia
		"232", // Sierra Leone
		"233", // Ghana
		"234", // Nigeria
		"235", // Chad
		"236", // Central African Republic
		"237", // Cameroon
		"238", // Cape Verde
		"239", // São Tomé and Príncipe
		"240", // Equatorial Guinea
		"241", // Gabon
		"242", // Republic of the Congo
		"243", // Democratic Republic of the Congo
		"244", // Angola
		"245", // Guinea-Bissau
		"246", // British Indian Ocean Territory
		"247", // Ascension Island
		"248", // Seychelles
		"249", // Sudan
		"250", // Rwanda
		"251", // Ethiopia
		"252", // Somalia (including Somaliland)
		"253", // Djibouti
		"254", // Kenya
		"255",
		"256", // Uganda
		"257", // Burundi
		"258", // Mozambique
		"259", // unassigned (was intended for People's Republic of Zanzibar but never implemented – see 255 Tanzania)
		"260", // Zambia
		"261", // Madagascar
		"262",
		"263", // Zimbabwe
		"264", // Namibia (formerly 27 (6x) as South West Africa)
		"265", // Malawi
		"266", // Lesotho
		"267", // Botswana
		"268", // Eswatini
		"269", // Comoros (formerly assigned to Mayotte, now at 262)
		"27",  //  South Africa
		"280", // 28x – unassigned (reserved for country code expansion)[1]
		"281", // 28x – unassigned (reserved for country code expansion)[1]
		"282", // 28x – unassigned (reserved for country code expansion)[1]
		"283", // 28x – unassigned (reserved for country code expansion)[1]
		"284", // 28x – unassigned (reserved for country code expansion)[1]
		"285", // 28x – unassigned (reserved for country code expansion)[1]
		"286", // 28x – unassigned (reserved for country code expansion)[1]
		"287", // 28x – unassigned (reserved for country code expansion)[1]
		"288", // 28x – unassigned (reserved for country code expansion)[1]
		"289", // 28x – unassigned (reserved for country code expansion)[1]
		"290",
		"291", // Eritrea
		"292", // unassigned
		"293", // unassigned
		"294", // unassigned
		"295", // unassigned (formerly assigned to San Marino, now at 378)
		"296", // unassigned
		"297", // Aruba
		"298", // Faroe Islands
		"299", // Greenland

		"30",  // Greece
		"31",  // Netherlands
		"32",  // Belgium
		"33",  // France
		"34",  // Spain
		"350", // Gibraltar
		"351",
		"352", // Luxembourg
		"353", // Ireland
		"354", // Iceland
		"355", // Albania
		"356", // Malta
		"357", // Cyprus (including Akrotiri and Dhekelia)
		"358",
		"359", // Bulgaria
		"36",  // Hungary (formerly assigned to Turkey, now at 90)
		// 37 – formerly assigned to East Germany until its reunification with West Germany, now part of 49 Germany
		"370", // Lithuania
		"371", // Latvia
		"372", // Estonia
		"373", // Moldova
		"374", // Armenia
		"375", // Belarus
		"376", // Andorra (formerly 33 628)
		"377", // Monaco (formerly 33 93)
		"378", // San Marino (interchangeably with 39 0549; earlier was allocated 295 but never used)
		"379", // Vatican City (assigned but uses 39 06698).
		// 38 – formerly assigned to Yugoslavia until its break-up in 1991
		"380", // Ukraine
		"381", // Serbia
		"382", // Montenegro
		"383", // Kosovo
		"384", // unassigned
		"385", // Croatia
		"386", // Slovenia
		"387", // Bosnia and Herzegovina
		"388", // unassigned (formerly assigned to the European Telephony Numbering Space)[1][2]
		"389", // North Macedonia
		"39",
		"40", // Romania
		"41",
		// 42 – formerly assigned to Czechoslovakia, later to its breakup successors (CZ, SK) until 1997
		"420", // Czech Republic
		"421", // Slovakia
		"422", // unassigned
		"423", // Liechtenstein (formerly at 41 (75))
		"424", // unassigned
		"425", // unassigned
		"426", // unassigned
		"427", // unassigned
		"428", // unassigned
		"429", // unassigned
		"43",  // Austria
		"44",
		"45",  // Denmark
		"46",  // Sweden
		"47",
		"48",  // Poland
		"49",  // Germany

		"500",
		"501", // Belize
		"502", // Guatemala
		"503", // El Salvador
		"504", // Honduras
		"505", // Nicaragua
		"506", // Costa Rica
		"507", // Panama
		"508", // Saint-Pierre and Miquelon
		"509", // Haiti
		"51",  // Peru
		"52",  // Mexico
		"53",  // Cuba
		"54",  // Argentina
		"55",  // Brazil
		"56",  // Chile
		"57",  // Colombia
		"58",  // Venezuela
		"590", // Guadeloupe (including Saint Barthélemy, Saint Martin)
		"591", // Bolivia
		"592", // Guyana
		"593", // Ecuador
		"594", // French Guiana
		"595", // Paraguay
		"596", // Martinique (formerly assigned to Peru, now 51)
		"597", // Suriname
		"598", // Uruguay
		"599",

		"60", // Malaysia
		"61",
		"62", // Indonesia
		"63", // Philippines
		"64",
		"65", // Singapore
		"66", // Thailand
		"670", // East Timor (formerly 62/39 during the Indonesian occupation; formerly assigned to Northern Mariana Islands, now part of NANP as 1 (670))
		"671", // unassigned (formerly assigned to Guam, now part of NANP as 1 (671))
		"672",
		"673", // Brunei
		"674", // Nauru
		"675", // Papua New Guinea
		"676", // Tonga
		"677", // Solomon Islands
		"678", // Vanuatu
		"679", // Fiji
		"680", // Palau
		"681", // Wallis and Futuna
		"682", // Cook Islands
		"683", // Niue
		"684", // unassigned (formerly assigned to American Samoa, now part of NANP as 1 (684))
		"685", // Samoa
		"686", // Kiribati
		"687", // New Caledonia
		"688", // Tuvalu
		"689", // French Polynesia
		"690", // Tokelau
		"691", // Federated States of Micronesia
		"692", // Marshall Islands
		"693", // unassigned
		"694", // unassigned
		"695", // unassigned
		"696", // unassigned
		"697", // unassigned
		"698", // unassigned
		"699", // unassigned

		"7",

		"800", // Universal International Freephone Service
		"801", // unassigned
		"802", // unassigned
		"803", // unassigned
		"804", // unassigned
		"805", // unassigned
		"806", // unassigned
		"807", // unassigned
		"808", // Universal International Shared Cost Numbers
		"809", // unassigned
		"81",  // Japan
		"82",  // South Korea
		"830", // 83x – unassigned (reserved for country code expansion)
		"831", // 83x – unassigned (reserved for country code expansion)
		"832", // 83x – unassigned (reserved for country code expansion)
		"833", // 83x – unassigned (reserved for country code expansion)
		"834", // 83x – unassigned (reserved for country code expansion)
		"835", // 83x – unassigned (reserved for country code expansion)
		"836", // 83x – unassigned (reserved for country code expansion)
		"837", // 83x – unassigned (reserved for country code expansion)
		"838", // 83x – unassigned (reserved for country code expansion)
		"839", // 83x – unassigned (reserved for country code expansion)
		"84",  // Vietnam
		"850", // North Korea
		"851", // unassigned
		"852", // Hong Kong
		"853", // Macau
		"854", // unassigned
		"855", // Cambodia
		"856", // Laos
		"857", // unassigned (formerly assigned to ANAC satellite service)
		"858", // unassigned (formerly assigned to ANAC satellite service)
		"859", // unassigned
		"86",  // China
		"870", // Global Mobile Satellite System (Inmarsat)
		"871", // unassigned (formerly assigned to Inmarsat Atlantic East, discontinued in 2008)
		"872", // unassigned (formerly assigned to Inmarsat Pacific, discontinued in 2008)
		"873", // unassigned (formerly assigned to Inmarsat Indian, discontinued in 2008)
		"874", // unassigned (formerly assigned to Inmarsat Atlantic West, discontinued in 2008)
		"875", // unassigned (reserved for future maritime mobile service)
		"876", // unassigned (reserved for future maritime mobile service)
		"877", // unassigned (reserved for future maritime mobile service)
		"878", // unassigned (formerly used for Universal Personal Telecommunications Service, discontinued in 2022)
		"879", // unassigned (reserved for national non-commercial purposes)
		"880", // Bangladesh
		"881", // Global Mobile Satellite System
		"882", // International Networks
		"883", // International Networks
		"884", // unassigned
		"885", // unassigned
		"886", // Taiwan
		"887", // unassigned
		"888", // unassigned (formerly assigned to OCHA for Telecommunications for Disaster Relief service)
		"889", // unassigned
		"890",  // 89x – unassigned (reserved for country code expansion)
		"891",  // 89x – unassigned (reserved for country code expansion)
		"892",  // 89x – unassigned (reserved for country code expansion)
		"893",  // 89x – unassigned (reserved for country code expansion)
		"894",  // 89x – unassigned (reserved for country code expansion)
		"895",  // 89x – unassigned (reserved for country code expansion)
		"896",  // 89x – unassigned (reserved for country code expansion)
		"897",  // 89x – unassigned (reserved for country code expansion)
		"898",  // 89x – unassigned (reserved for country code expansion)
		"899",  // 89x – unassigned (reserved for country code expansion)

		"90",
		"91",
		"92",
		"93",  // Afghanistan
		"94",  // Sri Lanka
		"95",  // Myanmar
		"960", // Maldives
		"961", // Lebanon
		"962", // Jordan
		"963", // Syria
		"964", // Iraq
		"965", // Kuwait
		"966", // Saudi Arabia
		"967", // Yemen
		"968", // Oman
		"969", // unassigned (formerly assigned to South Yemen until its unification with North Yemen, now part of 967 Yemen)
		"970",
		"971", // United Arab Emirates
		"972",
		"973", // Bahrain
		"974", // Qatar
		"975", // Bhutan
		"976", // Mongolia
		"977", // Nepal
		"978", // unassigned (formerly assigned to Dubai, now part of 971 United Arab Emirates)
		"979", // Universal International Premium Rate Service (UIPRS); (formerly assigned to Abu Dhabi, now part of 971 United Arab Emirates)
		"98",  // Iran
		"990", // unassigned
		"991", // unassigned (formerly used for International Telecommunications Public Correspondence Service)
		"992", // Tajikistan
		"993", // Turkmenistan
		"994", // Azerbaijan
		"995",
		"996", // Kyrgyzstan
		"997",
		"998", // Uzbekistan
		"999", // unassigned (reserved for future global service)
	}
)
