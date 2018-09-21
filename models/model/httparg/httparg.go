package httparg

type LoginArg struct {
	Account string `json:"account"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	PicCode string `json:"pcode"`
	PicId string `json:"pid"`
	PhoneCode string `json:"pcode"`
}

type LoginReply struct {
	Token string `json:"token, omitempty"`
}


type LogoutArg struct {
	Token string `json:"token"`
}

type ArticleHotReply struct {
	Title string
	ID string
	
}

