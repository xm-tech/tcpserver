package cmd

type Command interface {
	Exec(data ...interface{}) interface{}

	OnErr(error)
}
