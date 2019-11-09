package category

type TbContent struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Title      string `json:"title"`
	SubTitle   string `json:"sub_title"`
	TitleDesc  string `json:"title_desc"`
	Url        string `json:"url"`
	Pic        string `json:"pic"`
	Pic2       string `json:"pic2"`
	Content    string `json:"content"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}