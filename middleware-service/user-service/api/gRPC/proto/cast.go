package proto

import (
	"user-service/core/meta"
	"user-service/core/model"
)

func (m *User) ToNative() *model.User {
	return &model.User{
		UniqueID:     m.UniqueID,
		Username:     m.Username,
		UsernameN:    m.UsernameN,
		Email:        m.Email,
		EmailN:       m.EmailN,
		PasswordHash: m.PasswordHash,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		PhoneNumber:  m.PhoneNumber,
		Avatar:       m.Avatar,
		Location:     m.Location,
		PaymentInfo:  m.PaymentInfo.ToNative(),
		Liked:        m.Liked,
		Confirmed:    m.Confirmed,
		Role:         model.UserRole(m.Role),
		RegisterDate: m.RegisterDate.AsTime(),
		Provider:     model.UserProvider(m.Provider),
		// ConnectedProviders: m.ConnectedProviders,
	}
}

func (m *User) FromNative(n *model.User) *User {
	m.UniqueID = n.UniqueID
	m.Username = n.Username
	m.UsernameN = n.UsernameN
	m.Email = n.Email
	m.EmailN = n.EmailN
	m.PasswordHash = n.PasswordHash
	m.FirstName = n.FirstName
	m.LastName = n.LastName
	m.PhoneNumber = n.PhoneNumber
	m.Avatar = n.Avatar
	m.Location = n.Location
	// m.PaymentInfo = n.PaymentInfo.ToNative()
	m.Liked = n.Liked
	m.Confirmed = n.Confirmed
	// model.UserRole(m.Role) = model.UserRole(m.Role)
	// m.RegisterDate.AsTime() = n.RegisterDate.AsTime()
	// model.UserProvider(m.Provider) = model.UserProvider(m.Provider)
	return m
}

func NativeToUsers(native []*model.User) []*User {
	users := make([]*User, 0)
	for _, user := range native {
		users = append(users, (&User{}).FromNative(user))
	}
	return users
}

func UsersToNative(in []*User) []*model.User {
	users := make([]*model.User, 0)
	for _, user := range in {
		users = append(users, user.ToNative())
	}
	return users
}

func (m *AddressInfo) ToNative() model.AddressInfo {
	return model.AddressInfo{
		Country: m.Country,
		City: m.City,
		Address: m.Address,
		Address2: m.Address2,
		Region: m.Region,
		PostalCode: m.PostalCode,
	}
}

func (m *PaymentInfo) ToNative() model.PaymentInfo {
	return model.PaymentInfo{
		CardNumber: m.CardNumber,
		Expires: m.Expires,
		CVV: m.CVV,
		BillingInfo: m.BillingInfo.ToNative(),
	}
}

func (m *RequestParams) ToNative() *meta.RequestParams {
	n := &meta.RequestParams{}
	n.SetLimit(int(m.Limit))
	n.SetOffset(int(m.Offset))
	n.SetSortBy(m.SortBy)
	n.SetSortDirection(m.SortDirection)
	return n
}