package main

type IEnt struct {
	id  int
	uid int
}

type IVar struct {
	//IEnt
	id int
}

type ICon struct {
	//IEnt
	id int
}

type Model struct {
	IEnt
}

func (e IEnt) GetNewID() {

}
