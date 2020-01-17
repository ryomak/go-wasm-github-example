package common


type Repository struct{
  Name string `json:"name"`
  User string `json:"user"`
  Star uint`json:"star"`
  URL string  `json:"url"`
  Description string `json:"description"`
}

type Repositories []Repository
