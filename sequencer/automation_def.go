package sequencer

import "fmt"

type AutomationDef struct {
	BackAndForth *[]int    `yaml:"back_and_forth"`
	Cycle        *[]int    `yaml:"cycle"`
	Range        *RangeDef `yaml:"range"`
	Sweep        *RangeDef `yaml:"sweep"`
	FadeIn       *RangeDef `yaml:"fade_in"`
}

func (a *AutomationDef) GetAutomation() (IntAutomation, error) {
	if a.BackAndForth != nil {
		return IntBackAndForthAutomation(*a.BackAndForth), nil
	} else if a.Cycle != nil {
		return IntCycleAutomation(*a.Cycle), nil
	} else if a.Range != nil {
		return IntRangeAutomation(a.Range.From, a.Range.To, a.Range.Step), nil
	} else if a.Sweep != nil {
		return IntSweepAutomation(a.Sweep.From, a.Sweep.To, a.Sweep.ChangeEvery), nil
	} else if a.FadeIn != nil {
		return IntFadeInAutomation(a.FadeIn.From, a.FadeIn.To, a.FadeIn.ChangeEvery), nil
	}
	return nil, fmt.Errorf("Missing automation")
}

type FloatAutomationDef struct {
	BackAndForth *[]float64 `yaml:"back_and_forth"`
}

func (a *FloatAutomationDef) GetAutomation() (FloatAutomation, error) {
	if a.BackAndForth != nil {
		return FloatBackAndForthAutomation(*a.BackAndForth), nil
	}
	return nil, fmt.Errorf("Missing automation")
}

type RangeDef struct {
	From        int
	To          int
	Step        int
	ChangeEvery int `yaml:"change_every"`
}

type CycleChordsDef struct {
	Count  int     `yaml:"count"`
	Chords [][]int `yaml:"chords"`
}

func (c *CycleChordsDef) GetAutomation(seq *Sequencer) (IntArrayAutomation, error) {
	return ChordCycleArrayAutomation(c.Count, c.Chords), nil
}

type ArrayAutomationDef struct {
	CycleChords *CycleChordsDef `yaml:"cycle_chords"`
}

func (a *ArrayAutomationDef) GetAutomation(seq *Sequencer) (IntArrayAutomation, error) {
	if a.CycleChords != nil {
		return a.CycleChords.GetAutomation(seq)
	}
	return nil, fmt.Errorf("Missing array automation")
}