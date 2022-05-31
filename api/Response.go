package api

type Response[T ImgData | TokenData] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ImgData struct {
	Links ImgLink `json:"links"`
}

type TokenData struct {
	Token string `json:"token"`
}

type ImgLink struct {
	Url string `json:"url"`
}
