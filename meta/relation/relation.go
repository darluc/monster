package relation

type Type uint8

const (
	TypeMasterSlave Type = iota
	TypeMasterReference
)

type Ratio uint8

const (
	OneToOne Ratio = iota
	OneToMany
)
