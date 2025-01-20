package model

type (
	Item struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Picture     string `json:"picture"`
		Price       uint   `json:"price"`
	}

	ItemFilter struct {
		Name        string `query:"name" validate:"omitempty,max=65"`
		Description string `query:"description" validate:"omitempty,max=128"`
		Paginate
	}

	Paginate struct {
		Page int64 `query:"page" validate:"required,min=1"`
		Size int64 `query:"size"  validate:"required,min=1,max=20"`
	}

	ItemResult struct {
		Items    []*Item        `json:"items"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
	}

	BuyingReq struct {
		PlayerID string
		ItemID uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity uint `json:"quantity" validate:"required,gt=0"`
	}

	SellingReq struct {
		PlayerID string
		ItemID uint64 `json:"itemID" validate:"required,gt=0"`
		Quantity uint `json:"quantity" validate:"required,gt=0"`
	}
)
