package constants

// variables
const (
	CookieUser = "user"
)

// fiber contexts
const (
	ContextUid = "userId"
)

// params
const (
	ParamUid = "userId"
)

// Success messages
// ...

// Fail messages
// ...
const (
	Unauthenticated    = "unauthenticated to access resource"
	Unauthorized       = "unauthorized to access resource"
	InvalidCredentials = "invalid credenticals"
	UserNotExist       = "user does not exists"
	AppNotExists       = "app does not exists"
	ErrInValidFields   = "invalid fields"
	InvalidAppName     = "invalid appname"
)

// Error messages
const (
	ErrGetUser         = "error while get user"
	ErrLoginUser       = "error while login user"
	ErrInsertUser      = "error while creating user, please try after sometime"
	ErrHealthCheckDb   = "error while checking health of database"
	ErrUnauthenticated = "error verifing user identity"
	ErrGetApps         = "error while getting apps name"
	ErrGetReviews      = "error while getting reviews"
	ErrGetCategories   = "error while getting categories"
)

// Events
const (
	EventUserRegistered = "event:userRegistered"
)

// CSV file names
const (
	AppsInfoFile       = "googleplaystore.csv"
	AppsInfoTestFile   = "googleplaystore_dummy_apps.csv"
	AppsReviewInfoFile = "googleplaystore_user_reviews.csv"
)

// Page
const (
	QueryPage    = "page"
	QueryPerPage = "per_page"
)

// AppName
const (
	QueryAppName = "app"
)
