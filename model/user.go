package model

import "fmt"

type User struct {
	ID    int64 `pg:"pk"`
	CTFID string
}

func (u User) String() string {
	return fmt.Sprintf("User{ID=%d CTFID=%s}", u.ID, u.CTFID)
}
