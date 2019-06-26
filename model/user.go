package model

type User struct {
	UID string `json:"uid"`
	Email string `json:"email"`
	DisplayName string `json:"displayName"`
	PhotoUrl string `json:"photoUrl"`
}

func (user *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uid": user.UID,
		"email": user.Email,
		"displayName": user.DisplayName,
		"photoUrl": user.PhotoUrl,
	}
}
