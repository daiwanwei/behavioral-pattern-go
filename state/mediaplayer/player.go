package mediaplayer

import (
	"fmt"
	"strconv"
)

type Player struct {
	state        State
	playing      bool
	PlayList     []string
	CurrentTrack int
}

func NewPlayer() *Player {
	l := make([]string, 10)
	for i := 0; i < 10; i++ {
		l[i] = "track:" + strconv.Itoa(i)
	}
	p := &Player{
		playing:      true,
		PlayList:     l,
		CurrentTrack: 0,
	}
	p.state = NewReadyState(p)
	return p
}

func (p *Player) GetState() State {
	return p.state
}

//
//func (p *Player) SetPlaying(playing bool) {
//	p.playing=playing
//}
//
//func (p *Player) IsPlaying() bool {
//	return p.playing
//}

func (p *Player) StartPlayback() {
	fmt.Println("start playback:", p.PlayList[p.CurrentTrack])
}

func (p *Player) NextTrack() {
	p.CurrentTrack++
	if p.CurrentTrack >= len(p.PlayList) {
		p.CurrentTrack = 0
	}
	fmt.Println("playing:", p.PlayList[p.CurrentTrack])
}

func (p *Player) PreviousTrack() {
	p.CurrentTrack--
	if p.CurrentTrack < 0 {
		p.CurrentTrack = len(p.PlayList) - 1
	}
	fmt.Println("playing:", p.PlayList[p.CurrentTrack])
}

func (p *Player) SetCurrentTrackAfterStop() {
	p.CurrentTrack = 0
}
