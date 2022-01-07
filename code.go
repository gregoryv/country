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

// Comments must follow: ALPHA3 NUMERIC COUNTRY_NAME...
// for the methods to work.
const (
	AF Code = iota // AFG 004 Afghanistan
	AL             // ALB 008 Albania
	DZ             // DZA 012 Algeria
	AS             // ASM 016 American Samoa
	AD             // AND 020 Andorra
	AO             // AGO 024 Angola
	AI             // AIA 660 Anguilla
	AQ             // ATA 010 Antarctica
	AG             // ATG 028 Antigua and Barbuda
	AR             // ARG 032 Argentina
	AM             // ARM 051 Armenia
	AW             // ABW 533 Aruba
	AU             // AUS 036 Australia
	AT             // AUT 040 Austria
	AZ             // AZE 031 Azerbaijan
	BS             // BHS 044 Bahamas (the)
	BH             // BHR 048 Bahrain
	BD             // BGD 050 Bangladesh
	BB             // BRB 052 Barbados
	BY             // BLR 112 Belarus
	BE             // BEL 056 Belgium
	BZ             // BLZ 084 Belize
	BJ             // BEN 204 Benin
	BM             // BMU 060 Bermuda
	BT             // BTN 064 Bhutan
	BO             // BOL 068 Bolivia (Plurinational State of)
	BQ             // BES 535 Bonaire, Sint Eustatius and Saba
	BA             // BIH 070 Bosnia and Herzegovina
	BW             // BWA 072 Botswana
	BV             // BVT 074 Bouvet Island
	BR             // BRA 076 Brazil
	IO             // IOT 086 British Indian Ocean Territory (the)
	BN             // BRN 096 Brunei Darussalam
	BG             // BGR 100 Bulgaria
	BF             // BFA 854 Burkina Faso
	BI             // BDI 108 Burundi
	CV             // CPV 132 Cabo Verde
	KH             // KHM 116 Cambodia
	CM             // CMR 120 Cameroon
	CA             // CAN 124 Canada
	KY             // CYM 136 Cayman Islands (the)
	CF             // CAF 140 Central African Republic (the)
	TD             // TCD 148 Chad
	CL             // CHL 152 Chile
	CN             // CHN 156 China
	CX             // CXR 162 Christmas Island
	CC             // CCK 166 Cocos (Keeling) Islands (the)
	CO             // COL 170 Colombia
	KM             // COM 174 Comoros (the)
	CD             // COD 180 Congo (the Democratic Republic of the)
	CG             // COG 178 Congo (the)
	CK             // COK 184 Cook Islands (the)
	CR             // CRI 188 Costa Rica
	HR             // HRV 191 Croatia
	CU             // CUB 192 Cuba
	CW             // CUW 531 Curaçao
	CY             // CYP 196 Cyprus
	CZ             // CZE 203 Czechia
	CI             // CIV 384 Côte d'Ivoire
	DK             // DNK 208 Denmark
	DJ             // DJI 262 Djibouti
	DM             // DMA 212 Dominica
	DO             // DOM 214 Dominican Republic (the)
	EC             // ECU 218 Ecuador
	EG             // EGY 818 Egypt
	SV             // SLV 222 El Salvador
	GQ             // GNQ 226 Equatorial Guinea
	ER             // ERI 232 Eritrea
	EE             // EST 233 Estonia
	SZ             // SWZ 748 Eswatini
	ET             // ETH 231 Ethiopia
	FK             // FLK 238 Falkland Islands (the) [Malvinas]
	FO             // FRO 234 Faroe Islands (the)
	FJ             // FJI 242 Fiji
	FI             // FIN 246 Finland
	FR             // FRA 250 France
	GF             // GUF 254 French Guiana
	PF             // PYF 258 French Polynesia
	TF             // ATF 260 French Southern Territories (the)
	GA             // GAB 266 Gabon
	GM             // GMB 270 Gambia (the)
	GE             // GEO 268 Georgia
	DE             // DEU 276 Germany
	GH             // GHA 288 Ghana
	GI             // GIB 292 Gibraltar
	GR             // GRC 300 Greece
	GL             // GRL 304 Greenland
	GD             // GRD 308 Grenada
	GP             // GLP 312 Guadeloupe
	GU             // GUM 316 Guam
	GT             // GTM 320 Guatemala
	GG             // GGY 831 Guernsey
	GN             // GIN 324 Guinea
	GW             // GNB 624 Guinea-Bissau
	GY             // GUY 328 Guyana
	HT             // HTI 332 Haiti
	HM             // HMD 334 Heard Island and McDonald Islands
	VA             // VAT 336 Holy See (the)
	HN             // HND 340 Honduras
	HK             // HKG 344 Hong Kong
	HU             // HUN 348 Hungary
	IS             // ISL 352 Iceland
	IN             // IND 356 India
	ID             // IDN 360 Indonesia
	IR             // IRN 364 Iran (Islamic Republic of)
	IQ             // IRQ 368 Iraq
	IE             // IRL 372 Ireland
	IM             // IMN 833 Isle of Man
	IL             // ISR 376 Israel
	IT             // ITA 380 Italy
	JM             // JAM 388 Jamaica
	JP             // JPN 392 Japan
	JE             // JEY 832 Jersey
	JO             // JOR 400 Jordan
	KZ             // KAZ 398 Kazakhstan
	KE             // KEN 404 Kenya
	KI             // KIR 296 Kiribati
	KP             // PRK 408 Korea (the Democratic People's Republic of)
	KR             // KOR 410 Korea (the Republic of)
	KW             // KWT 414 Kuwait
	KG             // KGZ 417 Kyrgyzstan
	LA             // LAO 418 Lao People's Democratic Republic (the)
	LV             // LVA 428 Latvia
	LB             // LBN 422 Lebanon
	LS             // LSO 426 Lesotho
	LR             // LBR 430 Liberia
	LY             // LBY 434 Libya
	LI             // LIE 438 Liechtenstein
	LT             // LTU 440 Lithuania
	LU             // LUX 442 Luxembourg
	MO             // MAC 446 Macao
	MG             // MDG 450 Madagascar
	MW             // MWI 454 Malawi
	MY             // MYS 458 Malaysia
	MV             // MDV 462 Maldives
	ML             // MLI 466 Mali
	MT             // MLT 470 Malta
	MH             // MHL 584 Marshall Islands (the)
	MQ             // MTQ 474 Martinique
	MR             // MRT 478 Mauritania
	MU             // MUS 480 Mauritius
	YT             // MYT 175 Mayotte
	MX             // MEX 484 Mexico
	FM             // FSM 583 Micronesia (Federated States of)
	MD             // MDA 498 Moldova (the Republic of)
	MC             // MCO 492 Monaco
	MN             // MNG 496 Mongolia
	ME             // MNE 499 Montenegro
	MS             // MSR 500 Montserrat
	MA             // MAR 504 Morocco
	MZ             // MOZ 508 Mozambique
	MM             // MMR 104 Myanmar
	NA             // NAM 516 Namibia
	NR             // NRU 520 Nauru
	NP             // NPL 524 Nepal
	NL             // NLD 528 Netherlands (the)
	NC             // NCL 540 New Caledonia
	NZ             // NZL 554 New Zealand
	NI             // NIC 558 Nicaragua
	NE             // NER 562 Niger (the)
	NG             // NGA 566 Nigeria
	NU             // NIU 570 Niue
	NF             // NFK 574 Norfolk Island
	MP             // MNP 580 Northern Mariana Islands (the)
	NO             // NOR 578 Norway
	OM             // OMN 512 Oman
	PK             // PAK 586 Pakistan
	PW             // PLW 585 Palau
	PS             // PSE 275 Palestine, State of
	PA             // PAN 591 Panama
	PG             // PNG 598 Papua New Guinea
	PY             // PRY 600 Paraguay
	PE             // PER 604 Peru
	PH             // PHL 608 Philippines (the)
	PN             // PCN 612 Pitcairn
	PL             // POL 616 Poland
	PT             // PRT 620 Portugal
	PR             // PRI 630 Puerto Rico
	QA             // QAT 634 Qatar
	MK             // MKD 807 Republic of North Macedonia
	RO             // ROU 642 Romania
	RU             // RUS 643 Russian Federation (the)
	RW             // RWA 646 Rwanda
	RE             // REU 638 Réunion
	BL             // BLM 652 Saint Barthélemy
	SH             // SHN 654 Saint Helena, Ascension and Tristan da Cunha
	KN             // KNA 659 Saint Kitts and Nevis
	LC             // LCA 662 Saint Lucia
	MF             // MAF 663 Saint Martin (French part)
	PM             // SPM 666 Saint Pierre and Miquelon
	VC             // VCT 670 Saint Vincent and the Grenadines
	WS             // WSM 882 Samoa
	SM             // SMR 674 San Marino
	ST             // STP 678 Sao Tome and Principe
	SA             // SAU 682 Saudi Arabia
	SN             // SEN 686 Senegal
	RS             // SRB 688 Serbia
	SC             // SYC 690 Seychelles
	SL             // SLE 694 Sierra Leone
	SG             // SGP 702 Singapore
	SX             // SXM 534 Sint Maarten (Dutch part)
	SK             // SVK 703 Slovakia
	SI             // SVN 705 Slovenia
	SB             // SLB 090 Solomon Islands
	SO             // SOM 706 Somalia
	ZA             // ZAF 710 South Africa
	GS             // SGS 239 South Georgia and the South Sandwich Islands
	SS             // SSD 728 South Sudan
	ES             // ESP 724 Spain
	LK             // LKA 144 Sri Lanka
	SD             // SDN 729 Sudan (the)
	SR             // SUR 740 Suriname
	SJ             // SJM 744 Svalbard and Jan Mayen
	SE             // SWE 752 Sweden
	CH             // CHE 756 Switzerland
	SY             // SYR 760 Syrian Arab Republic
	TW             // TWN 158 Taiwan (Province of China)
	TJ             // TJK 762 Tajikistan
	TZ             // TZA 834 Tanzania, United Republic of
	TH             // THA 764 Thailand
	TL             // TLS 626 Timor-Leste
	TG             // TGO 768 Togo
	TK             // TKL 772 Tokelau
	TO             // TON 776 Tonga
	TT             // TTO 780 Trinidad and Tobago
	TN             // TUN 788 Tunisia
	TR             // TUR 792 Turkey
	TM             // TKM 795 Turkmenistan
	TC             // TCA 796 Turks and Caicos Islands (the)
	TV             // TUV 798 Tuvalu
	UG             // UGA 800 Uganda
	UA             // UKR 804 Ukraine
	AE             // ARE 784 United Arab Emirates (the)
	GB             // GBR 826 United Kingdom of Great Britain and Northern Ireland (the)
	UM             // UMI 581 United States Minor Outlying Islands (the)
	US             // USA 840 United States of America (the)
	UY             // URY 858 Uruguay
	UZ             // UZB 860 Uzbekistan
	VU             // VUT 548 Vanuatu
	VE             // VEN 862 Venezuela (Bolivarian Republic of)
	VN             // VNM 704 Viet Nam
	VG             // VGB 092 Virgin Islands (British)
	VI             // VIR 850 Virgin Islands (U.S.)
	WF             // WLF 876 Wallis and Futuna
	EH             // ESH 732 Western Sahara
	YE             // YEM 887 Yemen
	ZM             // ZMB 894 Zambia
	ZW             // ZWE 716 Zimbabwe
	AX             // ALA 248 Åland Islands
)

func (c Code) Alpha3() string { return c.String()[:3] }
func (c Code) Numeric() int {
	v, err := strconv.Atoi(c.String()[4:7])
	if err != nil {
		log.Fatal(err)
	}
	return v
}
func (c Code) Country() string { return c.String()[8:] }