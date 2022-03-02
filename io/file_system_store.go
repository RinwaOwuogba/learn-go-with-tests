package main

import (
	"encoding/json"
	"os"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(league, Player{name, 1})
	}

	f.database.Encode(league)
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{database}),
		league:   league,
	}

}
