package fakery

type Foobar struct {}

type CoolThingToTest interface {
  MethodToTest(param string) (result bool) // `sugar:accessor`
}

type CoolThing struct {
  foo Foobar
}

func (thing CoolThing) MethodToTest(param string) (result bool) {
  if param == "" {
    result = doer.foo.String()
  } else {
    result = "zoinks!"
  }
  return
}
