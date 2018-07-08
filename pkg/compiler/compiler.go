package compiler

func Compile(file string) error {
	cc := CompilerContext{}
	newCc := cc.InitContext(file)
	if newCc.Error() != "" {
		return newCc
	}
	return nil
}
