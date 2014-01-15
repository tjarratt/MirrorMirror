package fakery

type Foobar struct {}

type CoolThingToTest interface {
  MethodToTest(param string, param2 int32) (result bool, result2 string)
	IGuessThisIsOkayToo(takesAParam string)
	AndThisIsSomethingAsWell() (returningValue bool)
	EdgeCasesYay(string) (bool)
	AndMoreEdgeCasesAreFunToo()
}

type CoolThing struct {
  foo Foobar
}

func (thing CoolThing) MethodToTest(param string, param2 int32) (result bool, result2 string) {
  if param == "" {
    result = doer.foo.String()
  } else {
    result = "zoinks!"
  }
  return
}
