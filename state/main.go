package main

import "behavioral-pattern-go/state/mediaplayer"

func main() {
	p := mediaplayer.NewPlayer()
	p.GetState().OnPlay()
	p.GetState().OnNext()
	p.GetState().OnPrevious()
	p.GetState().OnLock()
	p.GetState().OnPlay()
	p.GetState().OnPlay()
	//p.GetState().OnPlay()
	//p.GetState().OnPlay()

}
