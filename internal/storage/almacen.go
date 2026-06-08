package storage

type Almacen interface {
	CarritoRepository
	ChoferRepository
	HorarioRepository
	CarritoHorarioRepository
	DespachoDiarioRepository
}
