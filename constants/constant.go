package constants

const (
	TokenSecretKey        = "RAHASIA-KEY_123*999"
	TokenExpiredInMinutes = 8 * 60 * 60
)

// ERR code Global
const (
	ERR_CODE_00     = "00"
	ERR_CODE_00_MSG = "SUCCESS.."

	ERR_CODE_03     = "03"
	ERR_CODE_03_MSG = "Error, unmarshall body Request"

	ERR_CODE_02     = "02"
	ERR_CODE_02_MSG = "Error parameter"
)

const (
	ERR_CODE_50     = "50"
	ERR_CODE_50_MSG = "Invalid username / password"

	ERR_CODE_51     = "51"
	ERR_CODE_51_MSG = "Error connection to database"

	ERR_CODE_52     = "52"
	ERR_CODE_52_MSG = "Failed Generate token"

	ERR_CODE_53     = "53"
	ERR_CODE_53_MSG = "Failed validate token"
)

const (
	ERR_CODE_10     = "10"
	ERR_CODE_10_MSG = "Failed save data to DB"

	ERR_CODE_11     = "11"
	ERR_CODE_11_MSG = "Failed get data from DB"
)

const (
	ERR_CODE_30     = "30"
	ERR_CODE_30_MSG = "Not Found"
)
