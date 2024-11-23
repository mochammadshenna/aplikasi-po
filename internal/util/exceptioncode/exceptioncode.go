package exceptioncode

import "errors"

var (
	ErrEmptyResult     = errors.New("empty result")
	ErrTokenInvalid    = errors.New("token invalid")
	ErrTokenExpired    = errors.New("token expired")
	ErrInvalidRequest  = errors.New("invalid request")
	ErrFailedReadEnum  = errors.New("failed to read enum")
	ErrUnableToLock    = errors.New("unable to lock")
	ErrInvalidMimeType = errors.New("invalid mime types")

	// spesific postgre error
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrUniqueViolation     = errors.New("unique violation")

	// employer wallet
	ErrNotEnoughBalance = errors.New("wallet balance is not sufficient")
)

const (
	CodeDataNotFound                      = "DATA_NOT_FOUND"
	CodeAccountLocked                     = "ACCOUNT_LOCKED"
	CodeOtpFailed                         = "OTP_FAILED"
	CodeReferralExist                     = "REFERRAL_EXIST"
	CodeOtpInvalid                        = "OTP_INVALID"
	CodeRequestTooFast                    = "REQUEST_TOO_FAST"
	CodeInvalidCredential                 = "INVALID_CREDENTIAL"
	CodeDelimiterInvalid                  = "DELIMITER_INVALID"
	CodeJobApplicantExist                 = "JOB_APPLICANT_EXIST"
	CodeQuizWrongData                     = "QUIZ_WRONG_DATA"
	CodeInvalidRequest                    = "INVALID_REQUEST"
	CodeInvalidWANumber                   = "INVALID_WA_NUMBER"
	CodeEmailExist                        = "EMAIL_EXIST"
	CodePhoneNumberExist                  = "PHONE_NUMBER_EXIST"
	CodeInvalidValidation                 = "INVALID_VALIDATION"
	CodeDataPersonalRequired              = "DATA_PERSONAL_REQUIRED"
	CodeDataKartuCVEducationRequired      = "DATA_KARTU_CV_EDUCATION_REQUIRED"
	CodeDataCVPintarEducationRequired     = "DATA_CV_PINTAR_EDUCATION_REQUIRED"
	CodeMinimumAgeNotPass                 = "MINIMUM_AGE_NOT_PASS"
	CodeMaximumAgeNotPass                 = "MAXIMUM_AGE_NOT_PASS"
	CodeInvalidAge                        = "INVALID_AGE"
	CodeMaximumTotalWorkExperienceNotPass = "MAX_TOTAL_WORK_EXP_NOT_PASS"
	CodeWorkExperienceAboveAge            = "WORK_EXPERIENCE_ABOVE_AGE"
	CodeInvalidTotalWorkExp               = "INVALID_TOTAL_WORK_EXP"
	CodeFirebaseTokenUpdateDone           = "FIREBASE_TOKEN_UPDATE_DONE"
	CodeForbidden                         = "FORBIDDEN"
	CodeDataLocked                        = "DATA_LOCKED"
	CodeJobClosed                         = "JOB_CLOSED"
	CodeBadRequest                        = "BAD_REQUEST"
	CodeDateExpired                       = "DATE_EXPIRED"
	CodeInvalidFileExtension              = "INVALID_FILE_EXTENSION"
	CodeInvalidFileSize                   = "INVALID_FILE_SIZE"
	CodeMissingRequiredData               = "MISSING_REQUIRED_DATA"
	CodeAccountBlocked                    = "ACCOUNT_BLOCKED"
	CodeAccountDeleted                    = "ACCOUNT_DELETED"
	CodeEmailNotVerified                  = "EMAIL_NOT_VERIFIED"
	CodeInvalidData                       = "INVALID_DATA"
	CodeConflict                          = "CONFLICT"
	CodeQuotaLimitReached                 = "QUOTA_LIMIT_REACHED"
	CodeRefundFailed                      = "REFUND_FAILED"
	CodeChecksumMismatch                  = "CHECKSUM_MISMATCH"
	CodeServiceUnavailable                = "SERVICE_UNAVAILABLE"
	CodeRequestFailed                     = "REQUEST_FAILED"
	CodeDataAlreadyExist                  = "DATA_ALREADY_EXIST"
	CodeInternalServerError               = "INTERNAL_SERVER_ERROR"

	// cari cuan
	CodeMissionExpired             = "MISSION_EXPIRED"
	CodeQuotaFull                  = "QUOTA_FULL"
	CodeCompleteOldMission         = "COMPLETE_OLD_MISSION"
	CodeMaxSubmissionReach         = "MAX_SUBMISSION_REACH"
	CodeBalanceNotEnough           = "BALANCE_NOT_ENOUGH"
	CodeNikExist                   = "NIK_EXIST"
	CodeMissionApplicationExpired  = "MISSION_APPLICATION_EXPIRED"
	CodeMaxKycAttemptReach         = "MAX_KYC_ATTEMPT_REACH"
	CodeCashoutAccountInvalid      = "CASHOUT_ACCOUNT_INVALID"
	CodeMissionGroupAlreadyApplied = "MISSION_GROUP_ALREADY_APPLIED"

	// in-app chat
	CodeChatUserNotFound = "CHAT_USER_NOT_FOUND"

	// rekrut
	CodeEmployerProfileInvalid = "EMPLOYER_PROFILE_INVALID"

	// Webinar
	CodeWrongDateFormat       = "WRONG_DATE_FORMAT"
	CodePhonNumberWrongFormat = "WRONG_PHONE_NUMBER_FORMAT"

	// Upskilling
	CodeFileNotProcessedYet  = "FILE_NOT_PROCESSED_YET"
	CodeFileTypeNotSupported = "FILE_TYPE_NOT_SUPPORTED"
	CodeSubscriptionInactive = "SUBSCRIPTION_INACTIVE"

	// SSO Login
	CodeTokenExpired = "TOKEN_EXPIRED"
	CodeTokenInvalid = "TOKEN_INVALID"

	// Quiz
	CodeInvalidQuizFormat = "INVALID_QUIZ_FORMAT"
)

type (
	errorType struct {
		ErrorMessage string
	}
	ErrorNotFound            errorType
	ErrorOTPFailed           errorType
	ErrorForeignKeyViolation errorType
)
