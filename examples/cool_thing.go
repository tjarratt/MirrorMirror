package fakery

type CoolThinger interface {
  MethodToTest(paramOne string, paramTwo int32) (result bool, result2 string)
	AnotherMethod(takesAParam string)
	ReturnsBool() (returningValue bool)
	EdgeCasesYay(string) (bool)
	AndMoreEdgeCasesAreFunToo()
	privateMethodsAreUsefulToo()
}

type CoolThing struct {
	privateData int
}

func (thing CoolThing) MethodToTest(param string, param2 int32) (result bool, result2 string) {
  if param == "" {
    result = string(thing.privateData)
  } else {
    result = "zoinks!"
  }
  return
}

func (thing CoolThing) AnotherMethod(takesAParam string) {
	if takesAParam == "zoinks!" {
		thing.privateData = 9001
	}

	return
}

func (thing CoolThing) ReturnsBool() (bool) {
	return thing.privateData == 9001
}

func (thing CoolThing) EdgeCasesAreFun(param string) (bool) {
	return param != "zoinks"
}

func (thing CoolThing) AndMoreEdgeCasesAreFunToo() {
	println(strings.Join([]string{
		"You gotta know when to hold 'em.",
		"You gotta know when to fold 'em.",
		"Know when to walk away and know when to run",
	}, "\n"))
}

func (thing CoolThing) privateMethodsAreUsefulToo() {
	println(strings.Join([]string{
		"I got water and I got holes, so",
		"Lalalalalala",
		"Sons and daughters of hungry ghosts",
		"I got water and I got holes, so",
		"Lalalalalala",
		"Sons and daughters of hungry ghosts",
	}, "\n"))
}
