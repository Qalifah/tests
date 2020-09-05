package poker

import (
	"encoding/json"
	"os"
	"fmt"
	"sort"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	League League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := initialisePlayerDBFile(file)
    if err != nil {
        return nil, fmt.Errorf("problem initialising player db file, %v", err)
    }

    league, err := NewLeague(file)

    if err != nil {
        return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
    }

    return &FileSystemPlayerStore{json.NewEncoder(&Tape{file}),league}, nil
}

func(f *FileSystemPlayerStore) GetPlayerScore(name string) int{

	player := f.League.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func(f *FileSystemPlayerStore) RecordWin(name string) {

	player := f.League.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.League = append(f.League, Player{name, 1})
	}

	f.Database.Encode(f.League)
}

func (f *FileSystemPlayerStore) GetLeague() League {
    sort.Slice(f.League, func(i, j int) bool {
        return f.League[i].Wins > f.League[j].Wins
    })
    return f.League
}

func initialisePlayerDBFile(file *os.File) error {
    file.Seek(0, 0)

    info, err := file.Stat()

    if err != nil {
        return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
    }

    if info.Size()==0 {
        file.Write([]byte("[]"))
        file.Seek(0, 0)
    }

    return nil
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