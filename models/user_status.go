package models

type UserStatus string

const (
	UserStatusActive   = "active"   // default user status, user can login
	UserStatusBlocked  = "blocked"  // since security reason user can't login temporarily
	UserStatusInactive = "inactive" // user can't login
	UserStatusDelete   = "delete"   // user delete, this record is just for trace
)
