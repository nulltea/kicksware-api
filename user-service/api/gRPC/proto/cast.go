package proto

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/timoth-y/kicksware-api/user-service/core/meta"
	"github.com/timoth-y/kicksware-api/user-service/core/model"
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
	m.PaymentInfo = PaymentInfo{}.FromNative(n.PaymentInfo)
	m.Liked = n.Liked
	m.Confirmed = n.Confirmed
	// model.UserRole(m.Role) = model.UserRole(m.Role)
	m.RegisterDate = timestamppb.New(n.RegisterDate)
	// model.UserProvider(m.Provider) = model.UserProvider(m.Provider)
	return m
}

func NativeToUsers(native []*model.User) []*User {
	users := make([]*User, 0)
	for _, user := range native {
		if user == nil {
			continue;
		}
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

func (m AddressInfo) FromNative(n model.AddressInfo) *AddressInfo {
	m.Country      = n.Country
	m.City         = n.City
	m.Address      = n.Address
	m.Address2     = n.Address2
	m.Region       = n.Region
	m.PostalCode   = n.PostalCode
	return &m
}

func (m *PaymentInfo) ToNative() model.PaymentInfo {
	return model.PaymentInfo{
		CardNumber: m.CardNumber,
		Expires: model.YearMonth{
			Year: m.Expires.Year,
			Month: m.Expires.Month,
		},
		CVV: m.CVV,
		BillingInfo: m.BillingInfo.ToNative(),
	}
}

func (m PaymentInfo) FromNative(n model.PaymentInfo) *PaymentInfo {
	m.CardNumber   = n.CardNumber
	m.Expires      = &YearMonth{
		Year: n.Expires.Year,
		Month: n.Expires.Month,
	}
	m.CVV          = n.CVV
	m.BillingInfo  = AddressInfo{}.FromNative(n.BillingInfo)
	return &m
}

func (m *RequestParams) ToNative() *meta.RequestParams {
	n := &meta.RequestParams{}
	n.SetLimit(int(m.Limit))
	n.SetOffset(int(m.Offset))
	if m.SortBy != nil {
		n.SetSortBy(m.SortBy.Value)
	}
	if n.SortDirection != nil {
		n.SetSortDirection(m.SortDirection.Value)
	}
	return n
}

func (m RequestParams) FromNative(n *meta.RequestParams) *RequestParams {
	m.Limit = int32(n.Limit())
	m.Offset = int32(n.Offset())
	m.SortBy = &wrappers.StringValue{Value: n.SortBy()}
	m.SortDirection = &wrappers.StringValue{Value: n.SortDirection()}
	return &m
}

func (m *AuthToken) ToNative() *meta.AuthToken {
	return &meta.AuthToken{
		Token: m.Token,
		Success: m.Success,
		Expires: m.Expires.AsTime(),
	}
}

func (m AuthToken) FromNative(n *meta.AuthToken) *AuthToken {
	m.Token = n.Token
	m.Success = n.Success
	m.Expires = timestamppb.New(n.Expires)
	return &m
}