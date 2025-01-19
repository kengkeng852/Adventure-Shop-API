package exception

import "fmt"

type PlayerItemFinding struct {
	ItemID uint64
}

func (e *PlayerItemFinding) Error() string {
	return fmt.Sprintf("finding player itemID: %d failed", e.ItemID)
}
