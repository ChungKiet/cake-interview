package constant

import "errors"

var (
	ErrUserAccountRequired             = errors.New("ACCOUNT_IS_REQUIRED")
	ErrUserInfoRequired                = errors.New("USER_INFO_REQUIRED") // must input email, phone, or user_name
	ErrUserAccountInValid              = errors.New("ACCOUNT_INVALID")    // account must be: email, phone or user_name
	ErrUserPasswordInvalid             = errors.New("PASSWORD_INVALID")   // password.length must be >= 5
	ErrUserFullNameRequired            = errors.New("USER_REQUIRED")
	ErrUserBirthdayInvalid             = errors.New("USER_BIRTHDAY_INVALID")
	ErrUserPasswordIncorrect           = errors.New("USER_PASSWORD_INCORRECT")
	ErrVoucherNameRequired             = errors.New("VOUCHER_NAME_REQUIRED")
	ErrVoucherTypeRequired             = errors.New("VOUCHER_TYPE_REQUIRED")
	ErrVoucherTypeInvalid              = errors.New("VOUCHER_TYPE_INVALID")
	ErrVoucherValueInvalid             = errors.New("VOUCHER_VALUE_INVALID")
	ErrVoucherCampaignNameInvalid      = errors.New("VOUCHER_NAME_INVALID")
	ErrVoucherCampaignStartDayInvalid  = errors.New("VOUCHER_START_DAY_INVALID")
	ErrVoucherCampaignEndDayInvalid    = errors.New("VOUCHER_END_DAY_INVALID")
	ErrVoucherCampaignLiveTimeInvalid  = errors.New("VOUCHER_LIVE_TIME_INVALID")
	ErrVoucherCampaignTotalSlotInvalid = errors.New("VOUCHER_TOTAL_SLOT_INVALID")
	ErrCampaignMustHavVouchers         = errors.New("CAMPAIGN_MUST_HAVE_VOUCHER")
	ErrVoucherNotFound                 = errors.New("VOUCHER_NOT_FOUND")
)
