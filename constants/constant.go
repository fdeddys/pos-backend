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
	ERR_CODE_20		= "20"
	ERR_CODE_20_MSG		= "You are not allowed "

	ERR_CODE_21		= "21"
	ERR_CODE_21_MSG		= "Error send to host "
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

	ERR_CODE_12     = "12"
	ERR_CODE_12_MSG = "Failed remove data from DB"
)

const (
	ERR_CODE_30     = "30"
	ERR_CODE_30_MSG = "Not Found"
)

const (
	ERR_CODE_40     = "40"
	ERR_CODE_40_MSG = "Resto Id harus diisi"
)

const (
	STOCK_UNAVAILABLE = 0
	STOCK_UNAVAILABLE_DESC = "Tidak Tersedia"
	STOCK_AVAILABLE = 1
	STOCK_AVAILABLE_DESC = "Tersedia"
)

const (
	ORDER_STATUS_DIPESAN = 10
	ORDER_STATUS_DIPESAN_DESC = "DIPESAN"

	ORDER_STATUS_DIMASAK = 20
	ORDER_STATUS_DIMASAK_DESC = "DIMASAK"

	ORDER_STATUS_DIANTAR= 30
	ORDER_STATUS_DIANTAR_DESC = "DIANTAR"

	ORDER_STATUS_DIMEJA = 40
	ORDER_STATUS_DIMEJA_DESC = "DIMEJA"

)

const (
	NOT_YET_PAID = 10
	NOT_YET_PAID_DESC = "BELUM DIBAYAR"

	PAID = 10
	PAID_DESC = "SUDAH DIBAYAR"

)

const (
	GROUP_ACTIVE = 1
	GROUP_ACTIVE_DESC = "Active"

	GROUP_IN_ACTIVE = 0
	GROUP_IN_ACTIVE_DESC = "In Active"
)

const (
	RESTO_ACTIVE = 1
	RESTO_ACTIVE_DESC = "Active"

	RESTO_IN_ACTIVE = 0
	RESTO_IN_ACTIVE_DESC = "In Active"

	ERR_CODE_60 = "60"
	ERR_CODE_60_MSG = "Kode Resto Telah Digunakan"
)

const (
	MENU_ITEM_ACTIVE = 1
	MENU_ITEM_ACTIVE_DESC = "Active"
)

const (
	IMAGE_ACTIVE = 1
	IMAGE_ACTIVE_DESC = "Active"

	IMAGE_IN_ACTIVE = 0
	IMAGE_IN_ACTIVE_DESC = "In Active"
)

const (
	BUCKET_RESTO = "resto"
	BUCKET_MENU_ITEM = "menu-item"
)