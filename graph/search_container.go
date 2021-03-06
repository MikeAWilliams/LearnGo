package graph

type ContainerItem interface{}

type SearchContainer interface {
	Add(item ContainerItem)
	Remove() ContainerItem
	IsEmpty() bool
}
