// Package country provides codes and names of countries.
// Source of information was
//
//   https://www.iban.com/country-codes
package country

import (
	"log"
	"strconv"
)

type Code int

//go:generate stringer -linecomment -type Code

const (
	AF Code = iota // AF AFG 004 Afghanistan
	AL             // AL ALB 008 Albania
	DZ             // DZ DZA 012 Algeria
	AS             // AS ASM 016 American Samoa
	AD             // AD AND 020 Andorra
	AO             // AO AGO 024 Angola
	AI             // AI AIA 660 Anguilla
	AQ             // AQ ATA 010 Antarctica
	AG             // AG ATG 028 Antigua and Barbuda
	AR             // AR ARG 032 Argentina
	AM             // AM ARM 051 Armenia
	AW             // AW ABW 533 Aruba
	AU             // AU AUS 036 Australia
	AT             // AT AUT 040 Austria
	AZ             // AZ AZE 031 Azerbaijan
	BS             // BS BHS 044 Bahamas (the)
	BH             // BH BHR 048 Bahrain
	BD             // BD BGD 050 Bangladesh
	BB             // BB BRB 052 Barbados
	BY             // BY BLR 112 Belarus
	BE             // BE BEL 056 Belgium
	BZ             // BZ BLZ 084 Belize
	BJ             // BJ BEN 204 Benin
	BM             // BM BMU 060 Bermuda
	BT             // BT BTN 064 Bhutan
	BO             // BO BOL 068 Bolivia (Plurinational State of)
	BQ             // BQ BES 535 Bonaire, Sint Eustatius and Saba
	BA             // BA BIH 070 Bosnia and Herzegovina
	BW             // BW BWA 072 Botswana
	BV             // BV BVT 074 Bouvet Island
	BR             // BR BRA 076 Brazil
	IO             // IO IOT 086 British Indian Ocean Territory (the)
	BN             // BN BRN 096 Brunei Darussalam
	BG             // BG BGR 100 Bulgaria
	BF             // BF BFA 854 Burkina Faso
	BI             // BI BDI 108 Burundi
	CV             // CV CPV 132 Cabo Verde
	KH             // KH KHM 116 Cambodia
	CM             // CM CMR 120 Cameroon
	CA             // CA CAN 124 Canada
	KY             // KY CYM 136 Cayman Islands (the)
	CF             // CF CAF 140 Central African Republic (the)
	TD             // TD TCD 148 Chad
	CL             // CL CHL 152 Chile
	CN             // CN CHN 156 China
	CX             // CX CXR 162 Christmas Island
	CC             // CC CCK 166 Cocos (Keeling) Islands (the)
	CO             // CO COL 170 Colombia
	KM             // KM COM 174 Comoros (the)
	CD             // CD COD 180 Congo (the Democratic Republic of the)
	CG             // CG COG 178 Congo (the)
	CK             // CK COK 184 Cook Islands (the)
	CR             // CR CRI 188 Costa Rica
	HR             // HR HRV 191 Croatia
	CU             // CU CUB 192 Cuba
	CW             // CW CUW 531 Curaçao
	CY             // CY CYP 196 Cyprus
	CZ             // CZ CZE 203 Czechia
	CI             // CI CIV 384 Côte d'Ivoire
	DK             // DK DNK 208 Denmark
	DJ             // DJ DJI 262 Djibouti
	DM             // DM DMA 212 Dominica
	DO             // DO DOM 214 Dominican Republic (the)
	EC             // EC ECU 218 Ecuador
	EG             // EG EGY 818 Egypt
	SV             // SV SLV 222 El Salvador
	GQ             // GQ GNQ 226 Equatorial Guinea
	ER             // ER ERI 232 Eritrea
	EE             // EE EST 233 Estonia
	SZ             // SZ SWZ 748 Eswatini
	ET             // ET ETH 231 Ethiopia
	FK             // FK FLK 238 Falkland Islands (the) [Malvinas]
	FO             // FO FRO 234 Faroe Islands (the)
	FJ             // FJ FJI 242 Fiji
	FI             // FI FIN 246 Finland
	FR             // FR FRA 250 France
	GF             // GF GUF 254 French Guiana
	PF             // PF PYF 258 French Polynesia
	TF             // TF ATF 260 French Southern Territories (the)
	GA             // GA GAB 266 Gabon
	GM             // GM GMB 270 Gambia (the)
	GE             // GE GEO 268 Georgia
	DE             // DE DEU 276 Germany
	GH             // GH GHA 288 Ghana
	GI             // GI GIB 292 Gibraltar
	GR             // GR GRC 300 Greece
	GL             // GL GRL 304 Greenland
	GD             // GD GRD 308 Grenada
	GP             // GP GLP 312 Guadeloupe
	GU             // GU GUM 316 Guam
	GT             // GT GTM 320 Guatemala
	GG             // GG GGY 831 Guernsey
	GN             // GN GIN 324 Guinea
	GW             // GW GNB 624 Guinea-Bissau
	GY             // GY GUY 328 Guyana
	HT             // HT HTI 332 Haiti
	HM             // HM HMD 334 Heard Island and McDonald Islands
	VA             // VA VAT 336 Holy See (the)
	HN             // HN HND 340 Honduras
	HK             // HK HKG 344 Hong Kong
	HU             // HU HUN 348 Hungary
	IS             // IS ISL 352 Iceland
	IN             // IN IND 356 India
	ID             // ID IDN 360 Indonesia
	IR             // IR IRN 364 Iran (Islamic Republic of)
	IQ             // IQ IRQ 368 Iraq
	IE             // IE IRL 372 Ireland
	IM             // IM IMN 833 Isle of Man
	IL             // IL ISR 376 Israel
	IT             // IT ITA 380 Italy
	JM             // JM JAM 388 Jamaica
	JP             // JP JPN 392 Japan
	JE             // JE JEY 832 Jersey
	JO             // JO JOR 400 Jordan
	KZ             // KZ KAZ 398 Kazakhstan
	KE             // KE KEN 404 Kenya
	KI             // KI KIR 296 Kiribati
	KP             // KP PRK 408 Korea (the Democratic People's Republic of)
	KR             // KR KOR 410 Korea (the Republic of)
	KW             // KW KWT 414 Kuwait
	KG             // KG KGZ 417 Kyrgyzstan
	LA             // LA LAO 418 Lao People's Democratic Republic (the)
	LV             // LV LVA 428 Latvia
	LB             // LB LBN 422 Lebanon
	LS             // LS LSO 426 Lesotho
	LR             // LR LBR 430 Liberia
	LY             // LY LBY 434 Libya
	LI             // LI LIE 438 Liechtenstein
	LT             // LT LTU 440 Lithuania
	LU             // LU LUX 442 Luxembourg
	MO             // MO MAC 446 Macao
	MG             // MG MDG 450 Madagascar
	MW             // MW MWI 454 Malawi
	MY             // MY MYS 458 Malaysia
	MV             // MV MDV 462 Maldives
	ML             // ML MLI 466 Mali
	MT             // MT MLT 470 Malta
	MH             // MH MHL 584 Marshall Islands (the)
	MQ             // MQ MTQ 474 Martinique
	MR             // MR MRT 478 Mauritania
	MU             // MU MUS 480 Mauritius
	YT             // YT MYT 175 Mayotte
	MX             // MX MEX 484 Mexico
	FM             // FM FSM 583 Micronesia (Federated States of)
	MD             // MD MDA 498 Moldova (the Republic of)
	MC             // MC MCO 492 Monaco
	MN             // MN MNG 496 Mongolia
	ME             // ME MNE 499 Montenegro
	MS             // MS MSR 500 Montserrat
	MA             // MA MAR 504 Morocco
	MZ             // MZ MOZ 508 Mozambique
	MM             // MM MMR 104 Myanmar
	NA             // NA NAM 516 Namibia
	NR             // NR NRU 520 Nauru
	NP             // NP NPL 524 Nepal
	NL             // NL NLD 528 Netherlands (the)
	NC             // NC NCL 540 New Caledonia
	NZ             // NZ NZL 554 New Zealand
	NI             // NI NIC 558 Nicaragua
	NE             // NE NER 562 Niger (the)
	NG             // NG NGA 566 Nigeria
	NU             // NU NIU 570 Niue
	NF             // NF NFK 574 Norfolk Island
	MP             // MP MNP 580 Northern Mariana Islands (the)
	NO             // NO NOR 578 Norway
	OM             // OM OMN 512 Oman
	PK             // PK PAK 586 Pakistan
	PW             // PW PLW 585 Palau
	PS             // PS PSE 275 Palestine, State of
	PA             // PA PAN 591 Panama
	PG             // PG PNG 598 Papua New Guinea
	PY             // PY PRY 600 Paraguay
	PE             // PE PER 604 Peru
	PH             // PH PHL 608 Philippines (the)
	PN             // PN PCN 612 Pitcairn
	PL             // PL POL 616 Poland
	PT             // PT PRT 620 Portugal
	PR             // PR PRI 630 Puerto Rico
	QA             // QA QAT 634 Qatar
	MK             // MK MKD 807 Republic of North Macedonia
	RO             // RO ROU 642 Romania
	RU             // RU RUS 643 Russian Federation (the)
	RW             // RW RWA 646 Rwanda
	RE             // RE REU 638 Réunion
	BL             // BL BLM 652 Saint Barthélemy
	SH             // SH SHN 654 Saint Helena, Ascension and Tristan da Cunha
	KN             // KN KNA 659 Saint Kitts and Nevis
	LC             // LC LCA 662 Saint Lucia
	MF             // MF MAF 663 Saint Martin (French part)
	PM             // PM SPM 666 Saint Pierre and Miquelon
	VC             // VC VCT 670 Saint Vincent and the Grenadines
	WS             // WS WSM 882 Samoa
	SM             // SM SMR 674 San Marino
	ST             // ST STP 678 Sao Tome and Principe
	SA             // SA SAU 682 Saudi Arabia
	SN             // SN SEN 686 Senegal
	RS             // RS SRB 688 Serbia
	SC             // SC SYC 690 Seychelles
	SL             // SL SLE 694 Sierra Leone
	SG             // SG SGP 702 Singapore
	SX             // SX SXM 534 Sint Maarten (Dutch part)
	SK             // SK SVK 703 Slovakia
	SI             // SI SVN 705 Slovenia
	SB             // SB SLB 090 Solomon Islands
	SO             // SO SOM 706 Somalia
	ZA             // ZA ZAF 710 South Africa
	GS             // GS SGS 239 South Georgia and the South Sandwich Islands
	SS             // SS SSD 728 South Sudan
	ES             // ES ESP 724 Spain
	LK             // LK LKA 144 Sri Lanka
	SD             // SD SDN 729 Sudan (the)
	SR             // SR SUR 740 Suriname
	SJ             // SJ SJM 744 Svalbard and Jan Mayen
	SE             // SE SWE 752 Sweden
	CH             // CH CHE 756 Switzerland
	SY             // SY SYR 760 Syrian Arab Republic
	TW             // TW TWN 158 Taiwan (Province of China)
	TJ             // TJ TJK 762 Tajikistan
	TZ             // TZ TZA 834 Tanzania, United Republic of
	TH             // TH THA 764 Thailand
	TL             // TL TLS 626 Timor-Leste
	TG             // TG TGO 768 Togo
	TK             // TK TKL 772 Tokelau
	TO             // TO TON 776 Tonga
	TT             // TT TTO 780 Trinidad and Tobago
	TN             // TN TUN 788 Tunisia
	TR             // TR TUR 792 Turkey
	TM             // TM TKM 795 Turkmenistan
	TC             // TC TCA 796 Turks and Caicos Islands (the)
	TV             // TV TUV 798 Tuvalu
	UG             // UG UGA 800 Uganda
	UA             // UA UKR 804 Ukraine
	AE             // AE ARE 784 United Arab Emirates (the)
	GB             // GB GBR 826 United Kingdom of Great Britain and Northern Ireland (the)
	UM             // UM UMI 581 United States Minor Outlying Islands (the)
	US             // US USA 840 United States of America (the)
	UY             // UY URY 858 Uruguay
	UZ             // UZ UZB 860 Uzbekistan
	VU             // VU VUT 548 Vanuatu
	VE             // VE VEN 862 Venezuela (Bolivarian Republic of)
	VN             // VN VNM 704 Viet Nam
	VG             // VG VGB 092 Virgin Islands (British)
	VI             // VI VIR 850 Virgin Islands (U.S.)
	WF             // WF WLF 876 Wallis and Futuna
	EH             // EH ESH 732 Western Sahara
	YE             // YE YEM 887 Yemen
	ZM             // ZM ZMB 894 Zambia
	ZW             // ZW ZWE 716 Zimbabwe
	AX             // AX ALA 248 Åland Islands

	last // here so we can handle bad codes, ie. Code(9999)
)

// Alpha2 returns the two letter string of a country
func (c Code) Alpha2() string {
	if !c.isValid() {
		return "undefined"
	}
	return c.String()[:2]
}

// Alpha3 returns the three letter string of a country
func (c Code) Alpha3() string {
	if !c.isValid() {
		return "undefined"
	}
	return c.String()[3:6]
}

// Numeric returns the numeric code of a country
func (c Code) Numeric() int {
	if !c.isValid() {
		return -1
	}
	v, err := strconv.Atoi(c.String()[7:10])
	if err != nil {
		log.Fatal(err, c.String())
	}
	return v
}

func (c Code) isValid() bool {
	return c >= 0 && c < last
}

// Country returns the name of a country
func (c Code) Country() string {
	if !c.isValid() {
		return "undefined"
	}
	return c.String()[11:]
}
