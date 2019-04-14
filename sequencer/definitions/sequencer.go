package definitions

import (
	"fmt"
	"io/ioutil"

	"github.com/bspaans/bleep/channels"
	. "github.com/bspaans/bleep/sequencer/sequences"
	"github.com/bspaans/bleep/util"
	"gopkg.in/yaml.v2"
)

type SequencerDef struct {
	BPM         float64              `yaml:"bpm"`
	Granularity int                  `yaml:"granularity"`
	Channels    channels.ChannelsDef `yaml:",inline"`
	Sequences   []SequenceDef        `yaml:"sequences"`
}

func (s *SequencerDef) GetSequences() ([]Sequence, error) {
	sequences := []Sequence{}
	for i, se := range s.Sequences {
		sequence, err := se.GetSequence(s.Granularity)
		if err != nil {
			return nil, util.WrapError(fmt.Sprintf("sequence [%d]", i), err)
		}
		sequences = append(sequences, sequence)
	}
	return sequences, nil
}

func NewSequencerDefFromFile(file string) (*SequencerDef, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	result := SequencerDef{}
	if err := yaml.Unmarshal(contents, &result); err != nil {
		return nil, err
	}
	if len(result.Sequences) == 0 {
		return nil, fmt.Errorf("No sequences in sequencer def %s", file)
	}
	return &result, nil
}
