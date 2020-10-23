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
	normalisedName = strings.ReplaceAll(strings.Trim(strings.toLower(name)), " ", "")
	switch normalisedName {
	case "aargau", "ag":
		return Canton{NameDE: "Aargau", NameOrig: "Aargau", Shortcode: "AG"}
	case "appenzellinnerrhoden", "ai":
		return Canton{NameDE: "Appenzell Innerrhoden", NameOrig: "Appenzell Innerrhoden", Shortcode: "AI"}
	case "appenzellausserrhoden", "ar":
		return Canton{NameDE: "Appenzell Ausserrhoden", NameOrig: "Appenzell Ausserrhoden", Shortcode: "AR"}
	case "bern", "berne", "be":
		return Canton{NameDE: "Bern", NameOrig: "Bern", Shortcode: "BE"}	
	case "basellandschaft", "basel-landschaft", "bl":
		return Canton{NameDE: "Basel-Landschaft", NameOrig: "Basel-Landschaft", Shortcode: "BL"}	
	case "baselstadt", "basel", "basel-stadt", "bs":
		return Canton{NameDE: "Basel-Stadt", NameOrig: "Basel-Stadt", Shortcode: "BS"}	
	case "fribourg", "friburg", "freiburg", "fr":
		return Canton{NameDE: "Freiburg", NameOrig: "Fribourg", Shortcode: "FR"}	
	case "genf", "geneva", "genève", "ge":
		return Canton{NameDE: "Genf", NameOrig: "Genève", Shortcode: "GE"}	
	case "glarus", "gl":
		return Canton{NameDE: "Glarus", NameOrig: "Glarus", Shortcode: "GL"}	
	case "graubuenden", "graubünden", "grisons", "grischun", "grigioni","gr":
		return Canton{NameDE: "Graubünden", NameOrig: "Grisons", Shortcode: "GR"}	
	case "jura", "ju":
		return Canton{NameDE: "Jura", NameOrig: "Jura", Shortcode: "JU"}	
	case "luzern", "lucerne", "ai":
		return Canton{NameDE: "Luzern", NameOrig: "Luzern", Shortcode: "LU"}	
	case "neuchâtel", "neuchatel", "neuenburg", "ai":
		return Canton{NameDE: "Neuenburg", NameOrig: "Neuchâtel", Shortcode: "NE"}	
	case "nidwalden", "nidwald", "nw":
		return Canton{NameDE: "Nidwalden", NameOrig: "Nidwalden", Shortcode: "NW"}	
	case "obwalden", "obwald", "ow":
		return Canton{NameDE: "Obwalden", NameOrig: "Obwalden", Shortcode: "OW"}	
	case "stgallen", "st.gallen", "sg":
		return Canton{NameDE: "St. Gallen", NameOrig: "St. Gallen", Shortcode: "SG"}	
	case "schaffhausen", "sh":
		return Canton{NameDE: "Schaffhausen", NameOrig: "Schaffhausen", Shortcode: "SH"}	
	case "solothurn", "soleure", "so":
		return Canton{NameDE: "Solothurn", NameOrig: "Solothurn", Shortcode: "SO"}	
	case "schwyz", "sz":
		return Canton{NameDE: "Schwyz", NameOrig: "Schwyz", Shortcode: "SZ"}	
	case "thurgau", "tg":
		return Canton{NameDE: "Thurgau", NameOrig: "Thurgau", Shortcode: "TG"}	
	case "tessin", "ticino", "ti":
		return Canton{NameDE: "Tessin", NameOrig: "Ticino", Shortcode: "TI"}	
	case "uri", "ur":
		return Canton{NameDE: "Uri", NameOrig: "Uri", Shortcode: "UR"}	
	case "vaud", "waadt",  "vd":
		return Canton{NameDE: "Waadt", NameOrig: "Vaud", Shortcode: "VD"}	
	case "valais", "wallis", "vs":
		return Canton{NameDE: "Wallis", NameOrig: "Valais", Shortcode: "VS"}	
	case "zug", "zg":
		return Canton{NameDE: "Zug", NameOrig: "Zug", Shortcode: "ZG"}	
	case "zurich", "zürich", "zuerich", "zh":
		return Canton{NameDE: "Zürich", NameOrig: "Zürich", Shortcode: "ZH"}
	default: 
		return Canton{NameDE: "Unbekannt", NameOrig: "Unknown", Shortcode: "UNKNOWN"}
	}
}

// ShortCode Returns shortcode as String
func (c Canton) ShortCode() string {
	return c.ShortCode
}

// NameDE... Returns german canton name as String
func (c Canton) NameDE() string {
	return c.NameDE
}

// NameOrig Returns canton name in their origin langauge as String
func (c Canton) NameOrig() string {
	return c.NameOrig
}