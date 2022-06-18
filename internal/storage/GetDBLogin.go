package storage

type GetDBLoginer interface {
	GetDBLogins() map[int]string
}

type GetDBLoginT struct {
	Login string
}

func (p GetDBLoginT) GetDBLogins() (mapURL map[int]string) {

	mapURL = useBD[p.Login]
	//	fmt.Println(map1)
	return mapURL
}
