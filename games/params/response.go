package params

type User struct {
	Id      string `json:"id"`      // 玩家ID
	Name    string `json:"name"`    // 玩家昵称
	Level   int    `json:"level"`   // 玩家等级
	Server  string `json:"server"`  // 玩家所在服务器ID
	Country string `json:"country"` // 玩家注册国家，例如CN，US等
}

type UserSearchResponseData struct {
	User User `json:"user"`
}

type Server struct {
	Id   string `json:"id"`   // 服务器ID
	Name string `json:"name"` // 服务器名称
}

type ServerListResponseData struct {
	Servers []Server `json:"servers"`
}
