package storage

type Storage interface {
	PutDB() (int, error)
	GetDB() (string, error)
}

type DataPut struct {
	Url1 string
}

type DataGet struct {
	IdUrlRedirect int
}

//Хранение значений
var bd = map[int]string{}
var index int

func (d DataPut) PutDB() (out int, err error) {
	index = len(bd)
	bd[index] = d.Url1
	return index, err
}

func (dDG DataGet) GetDB() (url2Redirect string, err error) {
	return bd[dDG.IdUrlRedirect], err
}
