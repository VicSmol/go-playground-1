package setx

type SetX interface {
	NewSet()
	Add()
	Remove()
	Contains()
	Size()
	ToSlice()
	IsEmpty()
}
