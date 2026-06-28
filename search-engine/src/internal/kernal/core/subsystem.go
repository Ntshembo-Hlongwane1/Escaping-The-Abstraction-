package core

type SubSystem interface {
	Name() string
	Init() error
	Start() error
	Stop() error
}
