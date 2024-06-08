package types

type Task interface {
	Run(*ReqContext) error
	Deps() []string
}
