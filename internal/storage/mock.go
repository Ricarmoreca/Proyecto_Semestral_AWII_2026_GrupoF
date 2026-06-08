package storage

import (
	"github.com/Ricarmoreca/Proyecto_Semestral_AWII_2026_GrupoF/internal/models"
)

type AlmacenMock struct {
	GetAllCarritosFunc           func() ([]models.Carrito, error)
	GetCarritoByNumeroFunc       func(numero int) (models.Carrito, error)
	CreateCarritoFunc            func(c models.Carrito) error
	UpdateCarritoFunc            func(c models.Carrito) error
	DeleteCarritoFunc            func(numero int) error
	GetHorariosByCarritoFunc     func(numero int) ([]models.CarritoHorario, error)
	GetAllChoferesFunc           func() ([]models.Chofer, error)
	GetChoferByIDFunc            func(id int) (models.Chofer, error)
	GetChoferByLicenciaFunc      func(licencia string) (models.Chofer, error)
	CreateChoferFunc             func(c models.Chofer) (models.Chofer, error)
	UpdateChoferFunc             func(c models.Chofer) error
	DeleteChoferFunc             func(id int) error
	GetAllHorariosFunc           func() ([]models.Horario, error)
	GetHorarioByIDFunc           func(id int) (models.Horario, error)
	CreateHorarioFunc            func(h models.Horario) (models.Horario, error)
	UpdateHorarioFunc            func(h models.Horario) error
	DeleteHorarioFunc            func(id int) error
	GetCarritosByHorarioFunc     func(id int) ([]models.CarritoHorario, error)
	AsignarCarritoHorarioFunc    func(rel models.CarritoHorarioRel) error
	DesasignarCarritoHorarioFunc func(rel models.CarritoHorarioRel) error
	GetAllDespachosFunc          func() ([]models.DespachoDiario, error)
	GetDespachoByIDFunc          func(id int) (models.DespachoDiario, error)
	CreateDespachoFunc           func(d models.DespachoDiario) (models.DespachoDiario, error)
	UpdateDespachoFunc           func(d models.DespachoDiario) error
	DeleteDespachoFunc           func(id int) error
}

func (m *AlmacenMock) GetAllCarritos() ([]models.Carrito, error) {
	if m.GetAllCarritosFunc != nil {
		return m.GetAllCarritosFunc()
	}
	return []models.Carrito{}, nil
}

func (m *AlmacenMock) GetCarritoByNumero(numero int) (models.Carrito, error) {
	if m.GetCarritoByNumeroFunc != nil {
		return m.GetCarritoByNumeroFunc(numero)
	}
	return models.Carrito{}, ErrNotFound
}

func (m *AlmacenMock) CreateCarrito(c models.Carrito) error {
	if m.CreateCarritoFunc != nil {
		return m.CreateCarritoFunc(c)
	}
	return nil
}

func (m *AlmacenMock) UpdateCarrito(c models.Carrito) error {
	if m.UpdateCarritoFunc != nil {
		return m.UpdateCarritoFunc(c)
	}
	return nil
}

func (m *AlmacenMock) DeleteCarrito(numero int) error {
	if m.DeleteCarritoFunc != nil {
		return m.DeleteCarritoFunc(numero)
	}
	return nil
}

func (m *AlmacenMock) GetHorariosByCarrito(numero int) ([]models.CarritoHorario, error) {
	if m.GetHorariosByCarritoFunc != nil {
		return m.GetHorariosByCarritoFunc(numero)
	}
	return []models.CarritoHorario{}, nil
}

func (m *AlmacenMock) GetAllChoferes() ([]models.Chofer, error) {
	if m.GetAllChoferesFunc != nil {
		return m.GetAllChoferesFunc()
	}
	return []models.Chofer{}, nil
}

func (m *AlmacenMock) GetChoferByID(id int) (models.Chofer, error) {
	if m.GetChoferByIDFunc != nil {
		return m.GetChoferByIDFunc(id)
	}
	return models.Chofer{}, ErrNotFound
}

func (m *AlmacenMock) GetChoferByLicencia(licencia string) (models.Chofer, error) {
	if m.GetChoferByLicenciaFunc != nil {
		return m.GetChoferByLicenciaFunc(licencia)
	}
	return models.Chofer{}, ErrNotFound
}

func (m *AlmacenMock) CreateChofer(c models.Chofer) (models.Chofer, error) {
	if m.CreateChoferFunc != nil {
		return m.CreateChoferFunc(c)
	}
	return c, nil
}

func (m *AlmacenMock) UpdateChofer(c models.Chofer) error {
	if m.UpdateChoferFunc != nil {
		return m.UpdateChoferFunc(c)
	}
	return nil
}

func (m *AlmacenMock) DeleteChofer(id int) error {
	if m.DeleteChoferFunc != nil {
		return m.DeleteChoferFunc(id)
	}
	return nil
}

func (m *AlmacenMock) GetAllHorarios() ([]models.Horario, error) {
	if m.GetAllHorariosFunc != nil {
		return m.GetAllHorariosFunc()
	}
	return []models.Horario{}, nil
}

func (m *AlmacenMock) GetHorarioByID(id int) (models.Horario, error) {
	if m.GetHorarioByIDFunc != nil {
		return m.GetHorarioByIDFunc(id)
	}
	return models.Horario{}, ErrNotFound
}

func (m *AlmacenMock) CreateHorario(h models.Horario) (models.Horario, error) {
	if m.CreateHorarioFunc != nil {
		return m.CreateHorarioFunc(h)
	}
	return h, nil
}

func (m *AlmacenMock) UpdateHorario(h models.Horario) error {
	if m.UpdateHorarioFunc != nil {
		return m.UpdateHorarioFunc(h)
	}
	return nil
}

func (m *AlmacenMock) DeleteHorario(id int) error {
	if m.DeleteHorarioFunc != nil {
		return m.DeleteHorarioFunc(id)
	}
	return nil
}

func (m *AlmacenMock) GetCarritosByHorario(id int) ([]models.CarritoHorario, error) {
	if m.GetCarritosByHorarioFunc != nil {
		return m.GetCarritosByHorarioFunc(id)
	}
	return []models.CarritoHorario{}, nil
}

func (m *AlmacenMock) AsignarCarritoHorario(rel models.CarritoHorarioRel) error {
	if m.AsignarCarritoHorarioFunc != nil {
		return m.AsignarCarritoHorarioFunc(rel)
	}
	return nil
}

func (m *AlmacenMock) DesasignarCarritoHorario(rel models.CarritoHorarioRel) error {
	if m.DesasignarCarritoHorarioFunc != nil {
		return m.DesasignarCarritoHorarioFunc(rel)
	}
	return nil
}

func (m *AlmacenMock) GetAllDespachos() ([]models.DespachoDiario, error) {
	if m.GetAllDespachosFunc != nil {
		return m.GetAllDespachosFunc()
	}
	return []models.DespachoDiario{}, nil
}

func (m *AlmacenMock) GetDespachoByID(id int) (models.DespachoDiario, error) {
	if m.GetDespachoByIDFunc != nil {
		return m.GetDespachoByIDFunc(id)
	}
	return models.DespachoDiario{}, ErrNotFound
}

func (m *AlmacenMock) CreateDespacho(d models.DespachoDiario) (models.DespachoDiario, error) {
	if m.CreateDespachoFunc != nil {
		return m.CreateDespachoFunc(d)
	}
	return d, nil
}

func (m *AlmacenMock) UpdateDespacho(d models.DespachoDiario) error {
	if m.UpdateDespachoFunc != nil {
		return m.UpdateDespachoFunc(d)
	}
	return nil
}

func (m *AlmacenMock) DeleteDespacho(id int) error {
	if m.DeleteDespachoFunc != nil {
		return m.DeleteDespachoFunc(id)
	}
	return nil
}
