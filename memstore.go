package graphql_go_pets_example

var petStore = map[string]Pet{}

var userStore = map[string]User{
	"frank": {
		ID:   "frank",
		Name: "Frank",
	},
}

var tagStore = map[string]Tag{
	"1": {
		ID:    "1",
		Title: "test",
	},
}
