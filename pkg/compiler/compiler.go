package compiler

func Compile(file string) error {
	cc := CompilerContext{}
	newCc := cc.InitContext(file)
	return newCc
}
