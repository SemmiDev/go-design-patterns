package main

import (
	"fmt"
	"time"
)

type Playlist interface {
	Play()
	Stop()
}

type CodingPlaylist struct {}
type RhythmPlaylist struct {}

func (p *CodingPlaylist) Play()  {
	fmt.Println("play coding playlist")
}

func (p *CodingPlaylist) Stop()  {
	fmt.Println("stop coding playlist")
}

func (p *RhythmPlaylist) Play()  {
	fmt.Println("play rhytm playlist")
}

func (p *RhythmPlaylist) Stop()  {
	fmt.Println("stop rhytm playlist")
}

type AreDoing string

const (
	Ngoding AreDoing = "ngoding"
	Santuy AreDoing = "Santuy"
)

type PlaylistPlaying interface {
	Playing(time time.Time) Playlist
	Type() AreDoing
}

type Sam struct {}
type Dev struct {}

func (s *Sam) Playing(time time.Time) Playlist {
	if time.Weekday().String() != "Sunday" {
		return &CodingPlaylist{}
	}
	return &RhythmPlaylist{}
}

func (s Sam) Type() AreDoing {
	return Ngoding
}

func (s *Dev) Playing(time time.Time) Playlist {
	if time.Weekday().String() != "Sunday" {
		return &RhythmPlaylist{}
	}
	return &CodingPlaylist{}
}

func (s Dev) Type() AreDoing {
	return Santuy
}

func main() {
	var playlistPlaying PlaylistPlaying
	var playlistPlaying2 PlaylistPlaying

	playlistPlaying = &Sam{}
	playlistPlaying2 = &Dev{}

	playlist := playlistPlaying.Playing(time.Now())
	playlist.Play()
	fmt.Println(playlistPlaying.Type())
	playlist.Stop()

	playlist2 := playlistPlaying2.Playing(time.Now())
	playlist2.Play()
	fmt.Println(playlistPlaying2.Type())
	playlist2.Stop()
}