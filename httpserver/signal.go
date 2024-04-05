package httpserver

type ServerSignal string

func (s ServerSignal) Signal() {

}

func (s ServerSignal) String() string {
	return string(s)
}

const SIGRUNERR = ServerSignal("httpserver listen and serve err")
