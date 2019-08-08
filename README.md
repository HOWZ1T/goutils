# Simple Utility Code
Simple utility code that I use in my go projects.

# Utils by package
- log
    - `func GetLevel() Level`
    - `func GetFMT() string`
    - `func GetIoStreams() []io.Writer`
    - `func SetLevel(level Level)`
    - `func SetIOStreams(streams []io.Writer)`
    - `func SetFMT(fmt string)`
    - `func AddStream(stream io.Writer)`
    - `func RemoveStream(stream io.Writer) error`
    - `func Delete()`
    - `func Fatal(format string, a ...interface{}) (n int, err error)`
    - `func Error(format string, a ...interface{}) (n int, err error)`
    - `func Warning(format string, a ...interface{}) (n int, err error)`
    - `func Info(format string, a ...interface{}) (n int, err error)`
    - `func Debug(format string, a ...interface{}) (n int, err error)`

- reflect
    - `func IsFunction(v interface{}) bool`
    - `func GetFunctionName(v interface{}) string`