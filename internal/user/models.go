package user

type User struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Username          string   `json:"username"`
	Profile_image_irl string   `json:"profile_image_irl"`
	Grade             string   `json:"grade"`
	Specialization    string   `json:"specialization"`
	Technologies      []string `json:"technologies"`
}
