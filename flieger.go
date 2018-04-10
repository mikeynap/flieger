package flieger

type Allisint = int
type Floatison = float64
type Allistring = string
type Interson = interface{}
type Mapison = map[Allistring]Interson

func HyphenAllistring(s string) Allistring {
	return Allistring(s)
}

func HyphenAllisint(i int) Allisint {
	return Allisint(i)
}

func HyphenFloatison(f float64) Floatison {
	return Floatison(f)
}

func HypenInterson(i interface{}) Interson {
	return Interson(i)
}

func HyphenMapison(s string, i interface{}) Mapison {
	return Mapison{s: i}
}
