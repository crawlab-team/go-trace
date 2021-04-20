package trace

import "github.com/ztrue/tracerr"

func PrintError(err error) {
	err = tracerr.Wrap(err)
	tracerr.Print(err)
}

func TraceError(err error) error {
	err = tracerr.Wrap(err)
	tracerr.Print(err)
	return err
}

func Error(err error) error {
	return tracerr.Wrap(err)
}
