package ffxiv

// defining listings for use case
type Listings struct {
	Listings []*Listing
}
type Listing struct {
	DataCenter  string
	Duty        string `selector:".left .duty"`
	Description string
	Creator     string `selector:".right .creator .text"`
	World       string `selector:".right .world .text"`
	Expires     string `selector:".right .expires .text"`
	Updated     string `selector:".right .updated .text"`
	Party       []*Slot
}

// defining what makes a slot
type Job int
type role int

type Slot struct {
	Roles  Roles
	Job    Job
	Filled bool
}
type Roles struct {
	Roles []role
}

func NewSlot() *Slot {
	return &Slot{
		Roles: Roles{Roles: []role{}},
	}
}
func (ls *Listings) Add(l *Listing) {
	for _, existingListing := range ls.Listings {
		if existingListing.Creator == l.Creator {
			return
		}
	}
	ls.Listings = append(ls.Listings, l)
}

func (l *Listings) GetListings(ls *Listings, duty string) *Listings {
	listings := &Listings{Listings: []*Listing{}}
	for _, l := range ls.Listings {
		if l.Duty == DutyHandler(duty) {
			if l.DataCenter == "Aether" {
				listings.Listings = append(listings.Listings, l)
			}
		}
	}
	return listings
}

func (l *Listings) GetUltimateListings(ls *Listings) *Listings {
	listings := &Listings{Listings: []*Listing{}}
	for _, l := range ls.Listings {
		if l.DataCenter == "Aether" {
			if l.Duty == DutyHandler("ucob") {
				listings.Listings = append(listings.Listings, l)
			}
			if l.Duty == DutyHandler("uwu") {
				listings.Listings = append(listings.Listings, l)
			}
			if l.Duty == DutyHandler("tea") {
				listings.Listings = append(listings.Listings, l)
			}
			if l.Duty == DutyHandler("dsr") {
				listings.Listings = append(listings.Listings, l)
			}
			if l.Duty == DutyHandler("top") {
				listings.Listings = append(listings.Listings, l)
			}
		}
	}
	return listings
}

func DutyHandler(duty string) string {
	switch duty {
	case "ucob":
		return "The Unending Coil of Bahamut (Ultimate)"
	case "uwu":
		return "The Weapon's Refrain (Ultimate)"
	case "tea":
		return "The Epic of Alexander (Ultimate)"
	case "dsr":
		return "Dragonsong's Reprise (Ultimate)"
	case "top":
		return "The Omega Protocol (Ultimate)"
	}
	return ""
}

func GetJob(abbr string) Job {
	switch abbr {
	case "GNB":
		return GNB
	case "PLD":
		return PLD
	case "DRK":
		return DRK
	case "WAR":
		return WAR
	case "SCH":
		return SCH
	case "SGE":
		return SGE
	case "AST":
		return WHM
	case "SAM":
		return SAM
	case "DRG":
		return DRG
	case "NIN":
		return NIN
	case "MNK":
		return MNK
	case "RPR":
		return RPR
	case "BRD":
		return BRD
	case "MCH":
		return MCH
	case "DNC":
		return DNC
	case "BLM":
		return BLM
	case "SMN":
		return SMN
	case "RDM":
		return RDM
	case "BLU":
		return BLU
	}
	return Unknown
}

const (
	Empty role = iota
	Tank
	Healer
	Dps
)
const (
	Unknown Job = iota
	GNB
	PLD
	GLD
	DRK
	WAR
	SCH
	SGE
	AST
	WHM
	SAM
	DRG
	NIN
	MNK
	RPR
	BRD
	MCH
	DNC
	BLM
	SMN
	RDM
	BLU
)
