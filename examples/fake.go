package fakery

type Foobar struct {}

type CoolThingToTest interface {
  MethodToTest(paramOne string, paramTwo int32) (result bool, result2 string)
	IGuessThisIsOkayToo(takesAParam string)
	AndThisIsSomethingAsWell() (returningValue bool)
	EdgeCasesYay(string) (bool)
	AndMoreEdgeCasesAreFunToo()
	butMaybeNotThisPrivateThing()
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


/* fakery ahead
type FakeCoolThingToTest struct {
	Returns struct {
		MethodToTestResult bool
		MethodToTestResult2 string
	}

	Received struct {
		MethodToTestParam string
		MethodToTestParam2 int32
	}
}
*/
