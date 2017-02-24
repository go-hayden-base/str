package str

type IEnumerableString interface {
	Filter(reg string) IEnumerableString
	FilterFunc(f func(line string, stop *bool) bool) IEnumerableString
	Enumerate(f func(line string, err error))
}
