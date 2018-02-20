package config

import "os"

var IsDevelopment bool

// Gender
const MALE = "L"
const FEMALE = "P"

// User Verification
const USER_NOT_VERIFIED = "USER_NOT_VERIFIED"
const EMAIL_NOT_VERIFIED = "EMAIL_NOT_VERIFIED"
const USER_VERIFIED = "USER_VERIFIED"
const USER_BLOCKED = "USER_BLOCKED"
const EMAIL_VERIFICATION_TO_BE_SENT = "EMAIL_VERIFICATION_TO_BE_SENT"
const EMAIL_VERIFICATION_RESENT = "EMAIL_VERIFICATION_RESENT"
const WAITING_VERIFICATION = "WAITING_VERIFICATION"

// Channel & Post Verification
const NOT_APPROVED = "NOT_APPROVED"
const APPROVED = "APPROVED"
const REJECTED = "REJECTED"
const REPORT_SPAM = "REPORT_SPAM"
const BLOCKED = "BLOCKED"
const WAITING_APPROVAL = "WAITING_APPROVAL"

// User type
const USER = "USER"

func init() {

	isResult := os.Getenv("IS_PRODUCTION")
	if isResult == "2" {
		IsDevelopment = true
	} else {
		IsDevelopment = false
	}
}
