// Package isp shows how Interface Segregation Principle can be applied in Go.
//
// https://en.wikipedia.org/wiki/Interface_segregation_principle
package isp

// Notifier implements all notification types
type Notifier interface {
	SendSMS()
	SendPush()
	SendEmail()
	// other methods...
}

// ----------------------------------------------------------------

// User is any client which can signup and receive otp
type User struct {
	notifier Notifier
}

// SignUp uses SendEmail() of Notifier
func (u *User) SignUp() {
	u.notifier.SendEmail()
}

// SendOTP uses SendSMS() of Notifier
func (u *User) SendOTP() {
	u.notifier.SendSMS()
}

// We have SendPush() and potentially other methods unused.
// For test purposes Notifier is difficult to stub.

// ----------------------------------------------------------------

// EmailSender shrinks interface to just what is needed
type EmailSender interface {
	SendEmail()
}

// SMSSender shrinks interface to just what is needed
type SMSSender interface {
	SendSMS()
}

type User2 struct {
	userEmailNotifier EmailSender
	userSMSNotifier   SMSSender
}

// SignUp uses SendEmail() of Notifier
func (u *User2) SignUp() {
	u.userEmailNotifier.SendEmail()
}

// SendOTP uses SendSMS() of Notifier
func (u *User2) SendOTP() {
	u.userSMSNotifier.SendSMS()
}

// Only used methods present in interfaces.
// Small interfaces is much easier to stub.
