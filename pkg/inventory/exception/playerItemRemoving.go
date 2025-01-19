package exception

import "fmt"

type PlayerItemRemoving struct {
	ItemID uint64
}

func (e *PlayerItemRemoving) Error() string {
	return fmt.Sprintf("removing player itemsID: %d failed", e.ItemID)
}
