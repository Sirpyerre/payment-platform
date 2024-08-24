package logger

type Logger interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warningf(moduleName, functionName, format string, args ...any)
	Error(moduleName, functionName string, err error)
	Errorf(moduleName, functionName, format string, args ...any)
	FatalIfError(moduleName, functionName string, errs ...error)
	Fatal(moduleName, functionName string, err error)
	Fatalf(moduleName, functionName, format string, args ...any)
}
