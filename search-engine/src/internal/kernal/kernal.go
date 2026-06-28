package kernal

import (
	"fmt"

	"searchengine.com/src/internal/kernal/core"
)

type Kernal struct {
	subsystems []core.SubSystem
}

func NewKernal() *Kernal {
	return &Kernal{
		subsystems: make([]core.SubSystem, 0),
	}
}

func (k *Kernal) Register(s core.SubSystem) {
	fmt.Printf("\n %s] Registered", s.Name())
	k.subsystems = append(k.subsystems, s)
}

func (k *Kernal) InitAll() error {
	fmt.Printf("\n PHASE: INIT")

	for _, s := range k.subsystems {
		if err := s.Init(); err != nil {
			return fmt.Errorf("Init failed for %s: %w", s.Name(), err)
		}
	}

	return nil
}

func (k *Kernal) StartAll() error {
	fmt.Printf("\n PHASE: START")

	for _, s := range k.subsystems {
		if err := s.Start(); err != nil {
			return fmt.Errorf("Start failed for %s: %w", s.Name(), err)
		}
	}

	return nil
}

func (k *Kernal) Shutdown() {
	fmt.Printf("\n PHASE: Shutdown")

	for i := len(k.subsystems) - 1; i >= 0; i-- {
		s := k.subsystems[i]
		if err := s.Stop(); err != nil {
			fmt.Printf("\n WARNING %s stop error: %v", s.Name(), err)
		}
	}
}
