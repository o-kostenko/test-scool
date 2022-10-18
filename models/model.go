package models

type Auth struct {
	ID     int
	ApiKey string
}

type User struct {
	ID       int
	Username string
}

type UserProfile struct {
	UserID    int
	FirstName string
	LastName  string
	Phone     string
	Address   string
	City      string
}

type UserData struct {
	UserID int
	School string
}

type Profile struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
	School    string `json:"school"`
}
