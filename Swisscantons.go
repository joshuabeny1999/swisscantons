package swisscantons

import "strings"

// Canton Struct
type Canton struct {
	NameDE string
	NameOrig string
	ShortCode string
}

// ByName returns a canton structs. It tries to guess from input correct canton.
func ByName(name string) Canton {
	normalisedName := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(name)), " ", "")
	switch normalisedName {
	case "aargau", "ag":
		return Canton{NameDE: "Aargau", NameOrig: "Aargau", ShortCode: "AG"}
	case "appenzellinnerrhoden", "ai":
		return Canton{NameDE: "Appenzell Innerrhoden", NameOrig: "Appenzell Innerrhoden", ShortCode: "AI"}
	case "appenzellausserrhoden", "ar":
		return Canton{NameDE: "Appenzell Ausserrhoden", NameOrig: "Appenzell Ausserrhoden", ShortCode: "AR"}
	case "bern", "berne", "be":
		return Canton{NameDE: "Bern", NameOrig: "Bern", ShortCode: "BE"}	
	case "basellandschaft", "basel-landschaft", "bl":
		return Canton{NameDE: "Basel-Landschaft", NameOrig: "Basel-Landschaft", ShortCode: "BL"}	
	case "baselstadt", "basel", "basel-stadt", "bs":
		return Canton{NameDE: "Basel-Stadt", NameOrig: "Basel-Stadt", ShortCode: "BS"}	
	case "fribourg", "friburg", "freiburg", "fr":
		return Canton{NameDE: "Freiburg", NameOrig: "Fribourg", ShortCode: "FR"}	
	case "genf", "geneva", "genève", "ge":
		return Canton{NameDE: "Genf", NameOrig: "Genève", ShortCode: "GE"}	
	case "glarus", "gl":
		return Canton{NameDE: "Glarus", NameOrig: "Glarus", ShortCode: "GL"}	
	case "graubuenden", "graubünden", "grisons", "grischun", "grigioni","gr":
		return Canton{NameDE: "Graubünden", NameOrig: "Grisons", ShortCode: "GR"}	
	case "jura", "ju":
		return Canton{NameDE: "Jura", NameOrig: "Jura", ShortCode: "JU"}	
	case "luzern", "lucerne", "lu":
		return Canton{NameDE: "Luzern", NameOrig: "Luzern", ShortCode: "LU"}	
	case "neuchâtel", "neuchatel", "neuenburg", "ne":
		return Canton{NameDE: "Neuenburg", NameOrig: "Neuchâtel", ShortCode: "NE"}	
	case "nidwalden", "nidwald", "nw":
		return Canton{NameDE: "Nidwalden", NameOrig: "Nidwalden", ShortCode: "NW"}	
	case "obwalden", "obwald", "ow":
		return Canton{NameDE: "Obwalden", NameOrig: "Obwalden", ShortCode: "OW"}	
	case "stgallen", "st.gallen", "sankt gallen", "sg":
		return Canton{NameDE: "St. Gallen", NameOrig: "St. Gallen", ShortCode: "SG"}	
	case "schaffhausen", "sh":
		return Canton{NameDE: "Schaffhausen", NameOrig: "Schaffhausen", ShortCode: "SH"}	
	case "solothurn", "soleure", "so":
		return Canton{NameDE: "Solothurn", NameOrig: "Solothurn", ShortCode: "SO"}	
	case "schwyz", "sz":
		return Canton{NameDE: "Schwyz", NameOrig: "Schwyz", ShortCode: "SZ"}	
	case "thurgau", "tg":
		return Canton{NameDE: "Thurgau", NameOrig: "Thurgau", ShortCode: "TG"}	
	case "tessin", "ticino", "ti":
		return Canton{NameDE: "Tessin", NameOrig: "Ticino", ShortCode: "TI"}	
	case "uri", "ur":
		return Canton{NameDE: "Uri", NameOrig: "Uri", ShortCode: "UR"}	
	case "vaud", "waadt",  "vd":
		return Canton{NameDE: "Waadt", NameOrig: "Vaud", ShortCode: "VD"}	
	case "valais", "wallis", "vs":
		return Canton{NameDE: "Wallis", NameOrig: "Valais", ShortCode: "VS"}	
	case "zug", "zg":
		return Canton{NameDE: "Zug", NameOrig: "Zug", ShortCode: "ZG"}	
	case "zurich", "zürich", "zuerich", "zh":
		return Canton{NameDE: "Zürich", NameOrig: "Zürich", ShortCode: "ZH"}
	case "lichtenstein", "liechtenstein", "li":
		return Canton{NameDE: "Liechtenstein", NameOrig: "Liechtenstein", ShortCode: "LI"}
	default: 
		return Canton{NameDE: "Unbekannt", NameOrig: "Unknown", ShortCode: "UNKNOWN"}
	}
}

// GetTwoLetterCode Returns shortcode as String
func (c Canton) GetTwoLetterCode() string {
	return c.ShortCode
}

// GermanName Returns german canton name as String
func (c Canton) GermanName() string {
	return c.NameDE
}

// OrigName Returns canton name in their origin langauge as String
func (c Canton) OrigName() string {
	return c.NameOrig
}
