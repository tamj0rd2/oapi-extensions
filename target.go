package lib

type SomethingResonse struct {
	JSON201 *struct {
		Hello string
		Inner struct {
			World string
		}
	}
}
