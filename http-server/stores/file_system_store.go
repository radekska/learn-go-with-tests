package stores

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"learn-go-with-tests/http-server/player"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   player.League
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	if err := initialisePlayerDBFile(file); err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := player.NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&Tape{file}),
		league:   league,
	}, nil
}

func (f *FileSystemPlayerStore) GetLeague() player.League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	p := f.league.FindPlayer(name)
	if p != nil {
		return p.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	p := f.league.FindPlayer(name)

	// When you range over a slice you are returned the current index of the loop
	// (in our case i) and a copy of the element at that index.
	// Changing the Wins value of a copy won't have any effect on the league slice that we iterate on.
	// For that reason, we need to get the reference to the actual value by
	// doing league[i] and then changing that value instead.
	if p != nil {
		p.Wins++
	} else {
		f.league = append(f.league, player.Player{
			Name: name,
			Wins: 1,
		})
	}

	f.Database.Encode(f.league)
}
