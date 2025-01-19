package exception

import "fmt"

type PlayerItemListing struct {
	PlayerID string
}

func (e *PlayerItemListing) Error() string {
	return fmt.Sprintf("listing items playerID: %s failed", e.PlayerID)
}
