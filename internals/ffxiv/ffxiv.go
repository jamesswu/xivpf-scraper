package ffxiv

// defining listings for use case
type Listings struct {
	Listings []*Listing
}
type Listing struct {
	data_center string
	duty        string `selector:".left .duty"`
	description string
	creator     string `selector:".right .creator .text"`
	world       string `selector:".right .world .text"`
	expires     string `selector:".right .expires .text"`
	updated     string `selector:".right .updated .text"`
	party       []*Slot
}

// defining what makes a slot
type Slot struct {
	roles  Roles
	job    Job
	filled bool
}
type Roles struct {
	Roles []role
}

type Job int
type role int

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
