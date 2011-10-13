package main

func init() {
	addTestCases(netdialTests, nil)
}

var netdialTests = []testCase{
	{
		Name: "netdial.0",
		Fn:   netdial,
		In: `package main

import "net"

func f() {
	c, err := net.Dial(net, "", addr)
	c, err = net.Dial(net, "", addr)
}
`,
		Out: `package main

import "net"

func f() {
	c, err := net.Dial(net, addr)
	c, err = net.Dial(net, addr)
}
`,
	},

	{
		Name: "netlookup.0",
		Fn:   netlookup,
		In: `package main

import "net"

func f() {
	foo, bar, _ := net.LookupHost(host)
	foo, bar, _ = net.LookupHost(host)
}
`,
		Out: `package main

import "net"

func f() {
	foo, bar := net.LookupHost(host)
	foo, bar = net.LookupHost(host)
}
`,
	},
}
