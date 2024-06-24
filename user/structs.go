package user

type UserService struct {
	APIKey    string
	APISecret string
}

type User struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	RealName   string     `json:"realname"`
	URL        string     `json:"url"`
	Image      []Image    `json:"image"`
	Country    string     `json:"country"`
	Age        string     `json:"age"`
	Gender     string     `json:"gender"`
	Subscriber string     `json:"subscriber"`
	Playcount  string     `json:"playcount"`
	Playlists  string     `json:"playlists"`
	Bootstrap  string     `json:"bootstrap"`
	Registered Registered `json:"registered"`
}

type Image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

type Registered struct {
	Unixtime string `json:"unixtime"`
	Text     int    `json:"#text"`
}
