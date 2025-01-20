package exception

import "fmt"

type ItemQuantityNotEnough struct {
	ItemID uint64
}

func (e *ItemQuantityNotEnough) Error() string {
	return fmt.Sprintf("itemID: %d is not enough", e.ItemID)
}
