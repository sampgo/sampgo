package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Konst struct {
	T     string
	name  string
	value string
}

type Arg struct {
	in            bool
	out           bool
	T             string
	name          string
	default_value string
}

type Callback struct {
	badret bool
	ret    string
	name   string
	args   []Arg
}

type Native struct {
	noimpl bool
	ret    string
	name   string
	args   []Arg
}

type IDL struct {
	konsts    []Konst
	callbacks []Callback
	natives   []Native
}

func ProcessIDL(file string) []string {

	f, e := ioutil.ReadFile(file)
	if e != nil {
		panic(e)
	}

	idl := string(f)

	idl = strings.ReplaceAll(idl, "(", " ")
	idl = strings.ReplaceAll(idl, ")", " ")
	idl = strings.ReplaceAll(idl, "[", " ")
	idl = strings.ReplaceAll(idl, "]", " ")
	idl = strings.ReplaceAll(idl, ",", " ")
	idl = strings.ReplaceAll(idl, ";", " ;")

	//println(idl)

	return strings.Fields(idl)
}

func (idl *IDL) Parse(file string) {
	src := ProcessIDL(file)

	var i int = 0

	for i < len(src) {
		switch src[i] {
		case "const":
			{
				// this will work assuming the following format:
				// const T name = value ;
				var k Konst
				i++ // T
				k.T = src[i]
				i++ // name
				k.name = src[i]
				i++ // =
				i++ // value
				k.value = src[i]
				i++ // ;
				i++ // skip ;

				println("konst", k.T, k.name, "=", k.value)
				idl.konsts = append(idl.konsts, k)
				break
			}
		case "native":
			{
				// this will work assuming the following format:
				// native (noimpl)? T name T a T b T c T x (= y)? T z ;
				var n Native
				i++ // noimpl or T

				if src[i] == "noimpl" {
					n.noimpl = true
					i++ // T
				}

				n.ret = src[i]

				i++ // name
				n.name = src[i]
				i++ // arg list

				println("native", n.ret, n.name, "(noimpl:", n.noimpl, ")")
				for src[i] != ";" {
					var a Arg

					if src[i] == "in" {
						a.in = true
						i++ // out or T
					}
					if src[i] == "out" {
						a.out = true
						i++ // T
					}
					if src[i] == "bind" {
						a.out = true
						i++ // "boundString"
						i++ // T
					}

					a.T = src[i]
					i++ // name
					a.name = src[i]
					i++ // =
					if src[i] == "=" {
						i++ // default_value
						a.default_value = src[i]
						i++ // ;
					}

					if a.name == "type" {
						a.name = "type_"
					}
					if a.name == "range" {
						a.name = "range_"
					}
					if a.name == "len" {
						a.name = "len_"
					}

					println("\targ", a.T, a.name, "=", a.default_value, "out:", a.out, "in:", a.in)
					n.args = append(n.args, a)
				}

				i++ // skip ;

				idl.natives = append(idl.natives, n)
				break
			}
		case "callback":
			{
				// this will work assuming the following format:
				// callback (badret true/false)? T name T a T b T c
				var c Callback
				i++ // badret/T/nonstandard
				if src[i] == "nonstandard" {
					// The nonstandard "tag" has been used only in a_http, so it is *safe* to remove
					continue
					//i++ // badret/T
				}
				if src[i] == "badret" {
					i++ // true/false
					b, err := strconv.ParseBool(src[i])
					c.badret = b
					i++ // T
					if err != nil {
						panic(err) // clearly the .IDL is malformed, so we should panic
					}
				} else {
					c.badret = false
				}
				c.ret = src[i]
				i++ // name
				c.name = src[i]
				i++ // arg list

				println("callback", c.ret, c.name, "badret:", c.badret)

				for src[i] != ";" {
					var a Arg

					if src[i] == "in" {
						a.in = true
						i++
					}
					if src[i] == "out" {
						a.out = true
						i++
					}

					a.T = src[i]
					i++
					a.name = src[i]
					i++
					if src[i] == "=" {
						i++
						a.default_value = src[i]
						i++
					}
					if a.name == "type" {
						a.name = "type_"
					}
					if a.name == "range" {
						a.name = "range_"
					}
					println("\targ", a.T, a.name, "=", a.default_value, "out:", a.out, "in:", a.in)
					c.args = append(c.args, a)
				}
				i++ // skip ;
				idl.callbacks = append(idl.callbacks, c)
				break
			}
		default:
			{
				//println("Unknown statement: ", src[i]) // leave this here, it's usefull for debugging
				i++
			}
		}
	}

}

func (k Konst) GenGo() string {
	split := strings.Split(strings.ToLower(k.name), "_")

	var b strings.Builder

	for _, s := range split {

		rs := []rune(s)

		rs[0] = unicode.ToUpper(rs[0])

		b.WriteString(string(rs))
	}

	//return fmt.Sprintf("%s %s = %s", b.String(), k.T, k.value)

	return "\t" + b.String() + " " + k.T + " = " + k.value + "\n"
}

func (c *Callback) GenGoBadRet(b *strings.Builder) {
	b.WriteString("\tif !ok {\n\t\treturn ")
	b.WriteString(strconv.FormatBool(c.badret))
	b.WriteString("\n\t}")
}

func (a *Arg) GoType() string {
	switch a.T {
	case "float":
		return "float32"
	default:
		return a.T
	}
}

func (a *Arg) SampGoType() string {
	if (a.name == "playerid" || a.name == "killerid" || a.name == "issuerid") && a.T == "int" {
		return "Player"
	}
	return a.GoType()
}

func (a *Arg) CType() string {
	if a.T == "string" {
		return "*C.char_t"
	}
	return "C." + a.T
}

func (a *Arg) ToGo() string {
	if (a.name == "playerid" || a.name == "killerid" || a.name == "issuerid") && a.T == "int" {
		return "Player{ID: int(" + a.name + ")}"
	}
	if a.T == "string" {
		return "C.GoString(C.constToNonConst(" + a.name + "))"
	}
	return a.GoType() + "(" + a.name + ")"
}

func (a *Arg) ToC() string {
	if a.T == "string" {
		return "C.nonConstToConst(cs" + a.name + ")"
	}
	return "C." + a.T + "(" + a.name + ")"
}

func (c *Callback) GenGo() string {
	var b strings.Builder

	rs := []rune(c.name)

	rs[0] = unicode.ToLower(rs[0])
	goname := string(rs)

	rs[2] = unicode.ToLower(rs[2])

	evname := string(rs[2:])

	b.WriteString("//export ")
	b.WriteString(goname)
	b.WriteRune('\n')

	b.WriteString("func ")
	b.WriteString(goname)
	b.WriteRune('(')
	for i, a := range c.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(a.name)
		b.WriteRune(' ')
		b.WriteString(a.CType())
	}
	b.WriteString(") ")
	b.WriteString(c.ret)
	b.WriteString(" {\n")

	b.WriteString("\tevt, ok := events[\"")
	b.WriteString(evname)
	b.WriteString("\"]\n")
	c.GenGoBadRet(&b)
	b.WriteString("\n\n\tfn, ok := evt.Handler.(func(")
	for i, a := range c.args {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(a.SampGoType())
	}
	b.WriteString(") ")
	b.WriteString(c.ret)
	b.WriteString(")\n")
	c.GenGoBadRet(&b)

	b.WriteString("\n\tfn(")
	for i, a := range c.args {
		if i != 0 {
			b.WriteString(", ")
		}

		b.WriteString(a.ToGo())
	}

	b.WriteString(")\n\treturn ")
	b.WriteString(strconv.FormatBool(!c.badret))
	b.WriteString(";\n}\n")

	return b.String()
}

func (n *Native) GenGo() string {
	if n.noimpl {
		return "// Native " + n.name + " was not generated because it is marked as 'noimpl'"
	}
	var b strings.Builder

	b.WriteString("// For documentation, please visit https://open.mp/docs/scripting/functions/")
	b.WriteString(n.name)

	b.WriteString("\nfunc ")
	b.WriteString(n.name)
	b.WriteRune('(')

	getter := false

	for i, a := range n.args {
		if i != 0 {
			b.WriteString(", ")
		}

		b.WriteString(a.name)
		if a.out {
			b.WriteString(" *")
			getter = true
		} else {
			b.WriteRune(' ')
		}

		b.WriteString(a.GoType())
	}

	b.WriteString(") ")
	b.WriteString(n.ret)
	if n.ret == "float" {
		b.WriteString("32")
	}

	b.WriteString(" {\n")

	for _, a := range n.args {
		if a.out || a.T != "string" {
			continue
		}

		b.WriteString("\tcs" + a.name + " := C.CString(" + a.name + ")\n")
		b.WriteString("\tdefer C.free(unsafe.Pointer(cs" + a.name + "))\n")
	}

	if getter {
		b.WriteString("\tvar ret " + n.ret)
		if n.ret == "float" {
			b.WriteString("32")
		}
		b.WriteRune('\n')

		for _, a := range n.args {
			if a.out {
				b.WriteString("\tvar c")
				b.WriteString(a.name)
				b.WriteString(" ")

				if a.T == "string" {
					b.WriteString("*C.char")
				} else {
					b.WriteString(a.CType())
				}

				b.WriteRune('\n')
				if a.T == "string" {

					size := "size"

					for i := range n.args {
						if (n.args[i].name == (a.name + "_size")) || (n.args[i].name == (a.name + "_len")) || (strings.Contains(n.args[i].name, "len")) {
							size = n.args[i].name
							break
						}
					}

					// make sure we can store strings safely
					b.WriteString("\tc" + a.name + " = (*C.char)(C.malloc(C.uint(" + size + ")))\n")
					b.WriteString("\tdefer C.free(unsafe.Pointer(c" + a.name + "))\n")
				}
			}
		}

		b.WriteString("\tret = ")

	} else {
		b.WriteString("\treturn ")
	}

	b.WriteString(n.ret)
	if n.ret == "float" {
		b.WriteString("32")
	}

	b.WriteRune('(')

	b.WriteString("C.")
	b.WriteString(n.name)
	b.WriteRune('(')

	for i, a := range n.args {
		if i != 0 {
			b.WriteString(", ")
		}
		if a.out {
			if a.T != "string" {
				b.WriteRune('&')
			}

			b.WriteRune('c')
			b.WriteString(a.name)
		} else {
			b.WriteString(a.ToC())
		}
	}

	if !getter {
		b.WriteString("))\n}\n")
		return b.String()
	}

	b.WriteString("))\n")

	for _, a := range n.args {
		if a.out {
			b.WriteString("\t*")
			b.WriteString(a.name)
			b.WriteString(" = ")
			if a.T == "string" {
				b.WriteString("C.GoString(C.constToNonConst")
			} else {
				b.WriteString(a.GoType())
			}
			b.WriteString("(c")
			b.WriteString(a.name)
			if a.T == "string" {
				b.WriteRune(')')
			}
			b.WriteString(")\n")
		}
	}

	b.WriteString("\treturn ret\n}\n\n")

	return b.String()
}

const cgoHeader = `package sampgo

/*
#cgo CFLAGS: -I./sampgdk
#ifndef GOLANG_APP
#define GOLANG_APP
#include "main.h"
#endif
*/
import "C"


`

func (idl *IDL) GenerateConstants() string {
	var b strings.Builder
	b.WriteString("package sampgo\n\nconst(\n")
	for _, k := range idl.konsts {
		b.WriteString(k.GenGo())
	}
	b.WriteString(")")
	return b.String()
}

func (idl *IDL) GenerateNatives() string {
	var b strings.Builder

	b.WriteString(cgoHeader)
	b.WriteString("import \"unsafe\"\n\n")

	for _, n := range idl.natives {
		b.WriteString(n.GenGo())
		b.WriteRune('\n')
	}

	return b.String()

}

func (idl *IDL) GenerateCallbacks() string {
	var b strings.Builder

	b.WriteString(cgoHeader)

	for _, c := range idl.callbacks {
		b.WriteString(c.GenGo())
	}

	return b.String()

}

func main() {
	if len(os.Args) < 2 || len(os.Args[1]) == 0 {
		println("syntax:", os.Args[0], " [files.idl]")
		return
	}

	var idl IDL

	for i, k := range os.Args {
		if i == 0 {
			continue
		}
		idl.Parse(k)
		println("Parsed", k)
	}

	ioutil.WriteFile("../constants.go", []byte(idl.GenerateConstants()), 0666)
	ioutil.WriteFile("../natives.go", []byte(idl.GenerateNatives()), 0666)
	ioutil.WriteFile("../callbacks.go", []byte(idl.GenerateCallbacks()), 0666)
}
