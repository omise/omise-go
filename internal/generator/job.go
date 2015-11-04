package main

type Job interface {
	// returns the template's name, the job struct itself is passed as the context.
	TemplateName() string
}

func Execute(j Job) error {
	// TODO:
	return nil
}
