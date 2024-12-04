package mapper

import (
	"golipors/internal/user/domain"
	"golipors/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

func ToDomainUser(m *types.User) *domain.User {
	if m == nil {
		return nil
	}

	return &domain.User{
		ID:            domain.UserID(m.ID),
		NationalID:    m.NationalID,
		Email:         m.Email,
		Password:      m.Password,
		FirstName:     m.FirstName,
		LastName:      m.LastName,
		Birthday:      m.Birthday,
		City:          m.City,
		WalletBalance: m.WalletBalance,
		VoteBalance:   m.VoteBalance,
	}
}

func ToModelUser(d *domain.User) *types.User {
	if d == nil {
		return nil
	}

	return &types.User{
		Model: gorm.Model{
			ID:        uint(d.ID),
			CreatedAt: d.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(d.DeletedAt)),
		},
		NationalID:    d.NationalID,
		Email:         d.Email,
		Password:      d.Password,
		FirstName:     d.FirstName,
		LastName:      d.LastName,
		Birthday:      d.Birthday,
		City:          d.City,
		WalletBalance: d.WalletBalance,
		VoteBalance:   d.VoteBalance,
	}
}
