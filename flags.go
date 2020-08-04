package translate

// flags maps unicode emoji flag to a Google Translate supported language.
var flags = map[string]string{
  "\U0001f1e6\U0001f1e8": "en",    // Ascension Island
  "\U0001f1e6\U0001f1e9": "ca",    // Andorra
  "\U0001f1e6\U0001f1ea": "ar",    // United Arab Emirates
  "\U0001f1e6\U0001f1eb": "ps",    // Afghanistan
  "\U0001f1e6\U0001f1ec": "en",    // Antigua & Barbuda
  "\U0001f1e6\U0001f1ee": "en",    // Anguilla
  "\U0001f1e6\U0001f1f1": "sq",    // Albania
  "\U0001f1e6\U0001f1f2": "hy",    // Armenia
  "\U0001f1e6\U0001f1f4": "pt",    // Angola
  "\U0001f1e6\U0001f1f6": "en",    // Antarctica
  "\U0001f1e6\U0001f1f7": "es",    // Argentina
  "\U0001f1e6\U0001f1f8": "sm",    // American Samoa
  "\U0001f1e6\U0001f1f9": "de",    // Austria
  "\U0001f1e6\U0001f1fa": "en",    // Australia
  "\U0001f1e6\U0001f1fc": "nl",    // Aruba
  "\U0001f1e6\U0001f1fd": "sv",    // Åland Islands
  "\U0001f1e6\U0001f1ff": "az",    // Azerbaijan
  "\U0001f1e7\U0001f1e6": "bs",    // Bosnia & Herzegovina
  "\U0001f1e7\U0001f1e7": "en",    // Barbados
  "\U0001f1e7\U0001f1e9": "bn",    // Bangladesh
  "\U0001f1e7\U0001f1ea": "nl",    // Belgium
  "\U0001f1e7\U0001f1eb": "fr",    // Burkina Faso
  "\U0001f1e7\U0001f1ec": "bg",    // Bulgaria
  "\U0001f1e7\U0001f1ed": "ar",    // Bahrain
  "\U0001f1e7\U0001f1ee": "fr",    // Burundi
  "\U0001f1e7\U0001f1ef": "fr",    // Benin
  "\U0001f1e7\U0001f1f1": "fr",    // St. Barthélemy
  "\U0001f1e7\U0001f1f2": "en",    // Bermuda
  "\U0001f1e7\U0001f1f3": "ms",    // Brunei
  "\U0001f1e7\U0001f1f4": "es",    // Bolivia
  "\U0001f1e7\U0001f1f6": "nl",    // Caribbean Netherlands
  "\U0001f1e7\U0001f1f7": "pt",    // Brazil
  "\U0001f1e7\U0001f1f8": "en",    // Bahamas
  "\U0001f1e7\U0001f1f9": "ne",    // Bhutan
  "\U0001f1e7\U0001f1fb": "no",    // Bouvet Island
  "\U0001f1e7\U0001f1fc": "en",    // Botswana
  "\U0001f1e7\U0001f1fe": "be",    // Belarus
  "\U0001f1e7\U0001f1ff": "en",    // Belize
  "\U0001f1e8\U0001f1e6": "en",    // Canada
  "\U0001f1e8\U0001f1e8": "en",    // Cocos (Keeling) Islands
  "\U0001f1e8\U0001f1e9": "fr",    // Congo - Kinshasa
  "\U0001f1e8\U0001f1eb": "fr",    // Central African Republic
  "\U0001f1e8\U0001f1ec": "fr",    // Congo - Brazzaville
  "\U0001f1e8\U0001f1ed": "de",    // Switzerland
  "\U0001f1e8\U0001f1ee": "fr",    // Côte d’Ivoire
  "\U0001f1e8\U0001f1f0": "en",    // Cook Islands
  "\U0001f1e8\U0001f1f1": "es",    // Chile
  "\U0001f1e8\U0001f1f2": "fr",    // Cameroon
  "\U0001f1e8\U0001f1f3": "zh",    // China
  "\U0001f1e8\U0001f1f4": "es",    // Colombia
  "\U0001f1e8\U0001f1f5": "fr",    // Clipperton Island
  "\U0001f1e8\U0001f1f7": "es",    // Costa Rica
  "\U0001f1e8\U0001f1fa": "es",    // Cuba
  "\U0001f1e8\U0001f1fb": "pt",    // Cape Verde
  "\U0001f1e8\U0001f1fc": "nl",    // Curaçao
  "\U0001f1e8\U0001f1fd": "en",    // Christmas Island
  "\U0001f1e8\U0001f1fe": "el",    // Cyprus
  "\U0001f1e8\U0001f1ff": "cs",    // Czechia
  "\U0001f1e9\U0001f1ea": "de",    // Germany
  "\U0001f1e9\U0001f1ec": "en",    // Diego Garcia
  "\U0001f1e9\U0001f1ef": "fr",    // Djibouti
  "\U0001f1e9\U0001f1f0": "da",    // Denmark
  "\U0001f1e9\U0001f1f2": "en",    // Dominica
  "\U0001f1e9\U0001f1f4": "es",    // Dominican Republic
  "\U0001f1e9\U0001f1ff": "ar",    // Algeria
  "\U0001f1ea\U0001f1e6": "es",    // Ceuta & Melilla
  "\U0001f1ea\U0001f1e8": "es",    // Ecuador
  "\U0001f1ea\U0001f1ea": "et",    // Estonia
  "\U0001f1ea\U0001f1ec": "ar",    // Egypt
  "\U0001f1ea\U0001f1ed": "ar",    // Western Sahara
  "\U0001f1ea\U0001f1f7": "ar",    // Eritrea
  "\U0001f1ea\U0001f1f8": "es",    // Spain
  "\U0001f1ea\U0001f1f9": "am",    // Ethiopia
  "\U0001f1ea\U0001f1fa": "en",    // European Union
  "\U0001f1eb\U0001f1ee": "fi",    // Finland
  "\U0001f1eb\U0001f1ef": "en",    // Fiji
  "\U0001f1eb\U0001f1f0": "en",    // Falkland Islands
  "\U0001f1eb\U0001f1f2": "en",    // Micronesia
  "\U0001f1eb\U0001f1f4": "nl",    // Faroe Islands
  "\U0001f1eb\U0001f1f7": "fr",    // France
  "\U0001f1ec\U0001f1e6": "fr",    // Gabon
  "\U0001f1ec\U0001f1e7": "en",    // United Kingdom
  "\U0001f1ec\U0001f1e9": "en",    // Grenada
  "\U0001f1ec\U0001f1ea": "ka",    // Georgia
  "\U0001f1ec\U0001f1eb": "fr",    // French Guiana
  "\U0001f1ec\U0001f1ec": "en",    // Guernsey
  "\U0001f1ec\U0001f1ed": "en",    // Ghana
  "\U0001f1ec\U0001f1ee": "en",    // Gibraltar
  "\U0001f1ec\U0001f1f1": "da",    // Greenland
  "\U0001f1ec\U0001f1f2": "en",    // Gambia
  "\U0001f1ec\U0001f1f3": "fr",    // Guinea
  "\U0001f1ec\U0001f1f5": "fr",    // Guadeloupe
  "\U0001f1ec\U0001f1f6": "es",    // Equatorial Guinea
  "\U0001f1ec\U0001f1f7": "el",    // Greece
  "\U0001f1ec\U0001f1f8": "en",    // South Georgia & South Sandwich Islands
  "\U0001f1ec\U0001f1f9": "es",    // Guatemala
  "\U0001f1ec\U0001f1fa": "es",    // Guam
  "\U0001f1ec\U0001f1fc": "pt",    // Guinea-Bissau
  "\U0001f1ec\U0001f1fe": "en",    // Guyana
  "\U0001f1ed\U0001f1f0": "zh",    // Hong Kong SAR China
  "\U0001f1ed\U0001f1f2": "en",    // Heard & McDonald Islands
  "\U0001f1ed\U0001f1f3": "es",    // Honduras
  "\U0001f1ed\U0001f1f7": "hr",    // Croatia
  "\U0001f1ed\U0001f1f9": "ht",    // Haiti
  "\U0001f1ed\U0001f1fa": "hu",    // Hungary
  "\U0001f1ee\U0001f1e8": "es",    // Canary Islands
  "\U0001f1ee\U0001f1e9": "id",    // Indonesia
  "\U0001f1ee\U0001f1ea": "en",    // Ireland
  "\U0001f1ee\U0001f1f1": "he",    // Israel
  "\U0001f1ee\U0001f1f2": "en",    // Isle of Man
  "\U0001f1ee\U0001f1f3": "hi",    // India
  "\U0001f1ee\U0001f1f4": "en",    // British Indian Ocean Territory
  "\U0001f1ee\U0001f1f6": "ku",    // Iraq
  "\U0001f1ee\U0001f1f7": "fa",    // Iran
  "\U0001f1ee\U0001f1f8": "is",    // Iceland
  "\U0001f1ee\U0001f1f9": "it",    // Italy
  "\U0001f1ef\U0001f1ea": "en",    // Jersey
  "\U0001f1ef\U0001f1f2": "en",    // Jamaica
  "\U0001f1ef\U0001f1f4": "ar",    // Jordan
  "\U0001f1ef\U0001f1f5": "ja",    // Japan
  "\U0001f1f0\U0001f1ea": "sw",    // Kenya
  "\U0001f1f0\U0001f1ec": "ky",    // Kyrgyzstan
  "\U0001f1f0\U0001f1ed": "km",    // Cambodia
  "\U0001f1f0\U0001f1ee": "en",    // Kiribati
  "\U0001f1f0\U0001f1f2": "fr",    // Comoros
  "\U0001f1f0\U0001f1f3": "en",    // St. Kitts & Nevis
  "\U0001f1f0\U0001f1f5": "ko",    // North Korea
  "\U0001f1f0\U0001f1f7": "ko",    // South Korea
  "\U0001f1f0\U0001f1fc": "ar",    // Kuwait
  "\U0001f1f0\U0001f1fe": "en",    // Cayman Islands
  "\U0001f1f0\U0001f1ff": "kk",    // Kazakhstan
  "\U0001f1f1\U0001f1e6": "lo",    // Laos
  "\U0001f1f1\U0001f1e7": "ar",    // Lebanon
  "\U0001f1f1\U0001f1e8": "en",    // St. Lucia
  "\U0001f1f1\U0001f1ee": "de",    // Liechtenstein
  "\U0001f1f1\U0001f1f0": "si",    // Sri Lanka
  "\U0001f1f1\U0001f1f7": "en",    // Liberia
  "\U0001f1f1\U0001f1f8": "st",    // Lesotho
  "\U0001f1f1\U0001f1f9": "lt",    // Lithuania
  "\U0001f1f1\U0001f1fa": "lb",    // Luxembourg
  "\U0001f1f1\U0001f1fb": "lv",    // Latvia
  "\U0001f1f1\U0001f1fe": "ar",    // Libya
  "\U0001f1f2\U0001f1e6": "ar",    // Morocco
  "\U0001f1f2\U0001f1e8": "fr",    // Monaco
  "\U0001f1f2\U0001f1e9": "ro",    // Moldova
  "\U0001f1f2\U0001f1ea": "bs",    // Montenegro
  "\U0001f1f2\U0001f1eb": "fr",    // St. Martin
  "\U0001f1f2\U0001f1ec": "mg",    // Madagascar
  "\U0001f1f2\U0001f1ed": "en",    // Marshall Islands
  "\U0001f1f2\U0001f1f0": "mk",    // North Macedonia
  "\U0001f1f2\U0001f1f1": "fr",    // Mali
  "\U0001f1f2\U0001f1f2": "my",    // Myanmar (Burma)
  "\U0001f1f2\U0001f1f3": "mn",    // Mongolia
  "\U0001f1f2\U0001f1f4": "zh",    // Macao SAR China
  "\U0001f1f2\U0001f1f5": "en",    // Northern Mariana Islands
  "\U0001f1f2\U0001f1f6": "fr",    // Martinique
  "\U0001f1f2\U0001f1f7": "ar",    // Mauritania
  "\U0001f1f2\U0001f1f8": "en",    // Montserrat
  "\U0001f1f2\U0001f1f9": "mt",    // Malta
  "\U0001f1f2\U0001f1fa": "en",    // Mauritius
  "\U0001f1f2\U0001f1fb": "en",    // Maldives
  "\U0001f1f2\U0001f1fc": "en",    // Malawi
  "\U0001f1f2\U0001f1fd": "es",    // Mexico
  "\U0001f1f2\U0001f1fe": "ms",    // Malaysia
  "\U0001f1f2\U0001f1ff": "pt",    // Mozambique
  "\U0001f1f3\U0001f1e6": "es",    // Namibia
  "\U0001f1f3\U0001f1e8": "fr",    // New Caledonia
  "\U0001f1f3\U0001f1ea": "fr",    // Niger
  "\U0001f1f3\U0001f1eb": "es",    // Norfolk Island
  "\U0001f1f3\U0001f1ec": "en",    // Nigeria
  "\U0001f1f3\U0001f1ee": "es",    // Nicaragua
  "\U0001f1f3\U0001f1f1": "nl",    // Netherlands
  "\U0001f1f3\U0001f1f4": "no",    // Norway
  "\U0001f1f3\U0001f1f5": "ne",    // Nepal
  "\U0001f1f3\U0001f1f7": "en",    // Nauru
  "\U0001f1f3\U0001f1fa": "en",    // Niue
  "\U0001f1f3\U0001f1ff": "en",    // New Zealand
  "\U0001f1f4\U0001f1f2": "ar",    // Oman
  "\U0001f1f5\U0001f1e6": "es",    // Panama
  "\U0001f1f5\U0001f1ea": "es",    // Peru
  "\U0001f1f5\U0001f1eb": "fr",    // French Polynesia
  "\U0001f1f5\U0001f1ec": "en",    // Papua New Guinea
  "\U0001f1f5\U0001f1ed": "tl",    // Philippines
  "\U0001f1f5\U0001f1f0": "ur",    // Pakistan
  "\U0001f1f5\U0001f1f1": "pl",    // Poland
  "\U0001f1f5\U0001f1f2": "fr",    // St. Pierre & Miquelon
  "\U0001f1f5\U0001f1f3": "en",    // Pitcairn Islands
  "\U0001f1f5\U0001f1f7": "es",    // Puerto Rico
  "\U0001f1f5\U0001f1f8": "ar",    // Palestinian Territories
  "\U0001f1f5\U0001f1f9": "pt",    // Portugal
  "\U0001f1f5\U0001f1fc": "es",    // Palau
  "\U0001f1f5\U0001f1fe": "es",    // Paraguay
  "\U0001f1f6\U0001f1e6": "ar",    // Qatar
  "\U0001f1f7\U0001f1ea": "fr",    // Réunion
  "\U0001f1f7\U0001f1f4": "ro",    // Romania
  "\U0001f1f7\U0001f1f8": "sr",    // Serbia
  "\U0001f1f7\U0001f1fa": "ru",    // Russia
  "\U0001f1f7\U0001f1fc": "sw",    // Rwanda
  "\U0001f1f8\U0001f1e6": "ar",    // Saudi Arabia
  "\U0001f1f8\U0001f1e7": "en",    // Solomon Islands
  "\U0001f1f8\U0001f1e8": "en",    // Seychelles
  "\U0001f1f8\U0001f1e9": "ar",    // Sudan
  "\U0001f1f8\U0001f1ea": "sv",    // Sweden
  "\U0001f1f8\U0001f1ec": "ms",    // Singapore
  "\U0001f1f8\U0001f1ed": "en",    // St. Helena
  "\U0001f1f8\U0001f1ee": "sl",    // Slovenia
  "\U0001f1f8\U0001f1ef": "no",    // Svalbard & Jan Mayen
  "\U0001f1f8\U0001f1f0": "sk",    // Slovakia
  "\U0001f1f8\U0001f1f1": "en",    // Sierra Leone
  "\U0001f1f8\U0001f1f2": "it",    // San Marino
  "\U0001f1f8\U0001f1f3": "fr",    // Senegal
  "\U0001f1f8\U0001f1f4": "so",    // Somalia
  "\U0001f1f8\U0001f1f7": "nl",    // Suriname
  "\U0001f1f8\U0001f1f8": "en",    // South Sudan
  "\U0001f1f8\U0001f1f9": "pt",    // São Tomé & Príncipe
  "\U0001f1f8\U0001f1fb": "es",    // El Salvador
  "\U0001f1f8\U0001f1fd": "nl",    // Sint Maarten
  "\U0001f1f8\U0001f1fe": "ar",    // Syria
  "\U0001f1f8\U0001f1ff": "en",    // Eswatini
  "\U0001f1f9\U0001f1e6": "en",    // Tristan da Cunha
  "\U0001f1f9\U0001f1e8": "en",    // Turks & Caicos Islands
  "\U0001f1f9\U0001f1e9": "fr",    // Chad
  "\U0001f1f9\U0001f1eb": "fr",    // French Southern Territories
  "\U0001f1f9\U0001f1ec": "fr",    // Togo
  "\U0001f1f9\U0001f1ed": "th",    // Thailand
  "\U0001f1f9\U0001f1ef": "tg",    // Tajikistan
  "\U0001f1f9\U0001f1f0": "en",    // Tokelau
  "\U0001f1f9\U0001f1f1": "pt",    // Timor-Leste
  "\U0001f1f9\U0001f1f2": "tk",    // Turkmenistan
  "\U0001f1f9\U0001f1f3": "ar",    // Tunisia
  "\U0001f1f9\U0001f1f4": "en",    // Tonga
  "\U0001f1f9\U0001f1f7": "tr",    // Turkey
  "\U0001f1f9\U0001f1f9": "en",    // Trinidad & Tobago
  "\U0001f1f9\U0001f1fb": "en",    // Tuvalu
  "\U0001f1f9\U0001f1fc": "zh-TW", // Taiwan
  "\U0001f1f9\U0001f1ff": "sw",    // Tanzania
  "\U0001f1fa\U0001f1e6": "uk",    // Ukraine
  "\U0001f1fa\U0001f1ec": "sw",    // Uganda
  "\U0001f1fa\U0001f1f2": "en",    // U.S. Outlying Islands
  "\U0001f1fa\U0001f1f3": "en",    // United Nations
  "\U0001f1fa\U0001f1f8": "en",    // United States
  "\U0001f1fa\U0001f1fe": "es",    // Uruguay
  "\U0001f1fa\U0001f1ff": "uz",    // Uzbekistan
  "\U0001f1fb\U0001f1e6": "it",    // Vatican City
  "\U0001f1fb\U0001f1e8": "en",    // St. Vincent & Grenadines
  "\U0001f1fb\U0001f1ea": "es",    // Venezuela
  "\U0001f1fb\U0001f1ec": "en",    // British Virgin Islands
  "\U0001f1fb\U0001f1ee": "en",    // U.S. Virgin Islands
  "\U0001f1fb\U0001f1f3": "vi",    // Vietnam
  "\U0001f1fb\U0001f1fa": "fr",    // Vanuatu
  "\U0001f1fc\U0001f1eb": "fr",    // Wallis & Futuna
  "\U0001f1fc\U0001f1f8": "sm",    // Samoa
  "\U0001f1fd\U0001f1f0": "sq",    // Kosovo
  "\U0001f1fe\U0001f1ea": "ar",    // Yemen
  "\U0001f1fe\U0001f1f9": "fr",    // Mayotte
  "\U0001f1ff\U0001f1e6": "af",    // South Africa
  "\U0001f1ff\U0001f1f2": "en",    // Zambia
  "\U0001f1ff\U0001f1fc": "sn",    // Zimbabwe
  "\U0001f3f4\U000e0067\U000e0062\U000e0065\U000e006e\U000e0067\U000e007f": "en", // England
  "\U0001f3f4\U000e0067\U000e0062\U000e0073\U000e0063\U000e0074\U000e007f": "en", // Scotland
  "\U0001f3f4\U000e0067\U000e0062\U000e0077\U000e006c\U000e0073\U000e007f": "en", // Wales
}
