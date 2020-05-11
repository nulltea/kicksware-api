package model

type UserRole string

var (
	Regular UserRole = ""
	Admin   UserRole = "adm"
	Guest   UserRole = "gst"
)
