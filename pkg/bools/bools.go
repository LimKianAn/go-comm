package bools

type Bools []bool

func (bools Bools) AllTrue() bool {
	for _, e := range bools {
		if e == false {
			return false
		}
	}
	return true
}
