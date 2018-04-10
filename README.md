# Flieger-Allisonize your code

[![Build Status](https://travis-ci.org/mikeynap/flieger.svg?branch=master)](https://travis-ci.org/mikeynap/flieger)
[![GoDoc](https://godoc.org/github.com/mikeynap/flieger?status.png)](https://godoc.org/github.com/mikeynap/flieger)

```
	var a flieger.Allistring
	var b flieger.Mapison
	a = flieger.HyphenAllistring("Ross's Mom-Flieger-Allison-Fox")
	b = flieger.HyphenMapison(a, 420.69)
	var c flieger.Floatison = b[a].(float64)
	fmt.Printf("%v: %v\n", a, c)
```
