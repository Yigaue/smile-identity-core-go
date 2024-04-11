package jobtype

import "errors"

const (
	BIOMETRIC_KYC                  = 1
	SMART_SELFIE_AUTHENTICATION    = 2
	SMART_SELFIE_REGISTRATION      = 4
	BASIC_KYC                      = 5
	ENHANCED_KYC                   = 5
	DOCUMENT_VERIFICATION          = 6
	BUSINESS_VERIFICATION          = 7
	UPDATE_PHOTO                   = 8
	COMPARE_USER_INFO              = 9
	ENHANCED_DOCUMENT_VERIFICATION = 11
)


var (
	ErrInvalidJobtype       = errors.New("invalid jobtype")
)
