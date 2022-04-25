package handler

import "drinkpipe-ui/store"

var units = map[string]*Unit{}

type Handler struct {
	Repository store.EntityStore
}
