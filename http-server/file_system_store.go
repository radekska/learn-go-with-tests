package poker

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   League
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

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&Tape{file}),
		league:   league,
	}, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.league.Sort()
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
		f.league = append(f.league, Player{
			Name: name,
			Wins: 1,
		})
	}

	f.Database.Encode(f.league)
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v ", err)
	}

	return store, closeFunc, nil
}
