package model

type Todo struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func GetAll() ([]*Todo, error) {
	var tls []*Todo = []*Todo{
		{Title: "aaa", Detail: "iiii"},
		{Title: "フロント実装", Detail: "フロントを実装する"},
	}
	return tls, nil
}
