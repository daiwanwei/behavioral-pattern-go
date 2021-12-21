package mediaplayer

import "fmt"

type State interface {
	OnLock()
	OnPlay()
	OnNext()
	OnPrevious()
}

type LockedState struct {
	player *Player
}

func NewLockedState(player *Player) State {
	return LockedState{player: player}
}

func (state LockedState) OnLock() {
	if state.player.playing {
		state.player.state = NewLockedState(state.player)
		fmt.Println("stop playing")
		return
	}
	fmt.Println("locked ....")
}
func (state LockedState) OnPlay() {
	state.player.state = NewReadyState(state.player)
	fmt.Println("ready")
}

func (state LockedState) OnNext() {
	fmt.Println("locked ....")
}

func (state LockedState) OnPrevious() {
	fmt.Println("locked ....")
}

type ReadyState struct {
	player *Player
}

func NewReadyState(player *Player) State {
	return ReadyState{player: player}
}

func (state ReadyState) OnLock() {
	state.player.state = NewLockedState(state.player)
	fmt.Println("Locked ...")
}
func (state ReadyState) OnPlay() {
	state.player.StartPlayback()
	state.player.state = NewPlayingState(state.player)
}

func (state ReadyState) OnNext() {
	fmt.Println("locked ....")
}

func (state ReadyState) OnPrevious() {
	fmt.Println("locked ....")
}

type PlayingState struct {
	player *Player
}

func NewPlayingState(player *Player) State {
	return PlayingState{player: player}
}

func (state PlayingState) OnLock() {
	state.player.state = NewLockedState(state.player)
	state.player.SetCurrentTrackAfterStop()
	fmt.Println("stop playing")
}

func (state PlayingState) OnPlay() {
	state.player.state = NewReadyState(state.player)
	fmt.Println("Paused....")
}

func (state PlayingState) OnNext() {
	state.player.NextTrack()
}

func (state PlayingState) OnPrevious() {
	state.player.PreviousTrack()
}
