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
type Slot struct {
	Roles  Roles
	Job    Job
	Filled bool
}
type Roles struct {
	Roles []role
}
type Job int
type role int

func NewSlot() *Slot {
	return &Slot{
		Roles: Roles{Roles: []role{}},
	}
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

func (ls *Listings) Add(l *Listing) {
	for _, existingListing := range ls.Listings {
		if existingListing.Creator == l.Creator {
			return
		}
	}
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
