package util

// NewProgress return a pointer of progress struct
func NewProgress() *Progress {
	return &Progress{percent: -1}
}

// progress indicates a progress of some action
type Progress struct {
	percent int8
}

func (p *Progress) SetPercent(percent int8) {
	p.percent = percent
}

func (p *Progress) Percent() int8 {
	return p.percent
}

func (p *Progress) Done() {
	p.percent = 100
}

func (p *Progress) IsDone() bool {
	return p.percent == 100
}
