package repository

// NOTE по пути /internal/core/port описываются контракты по которым приложение вызывает зависимые пакеты

// IDB контракт по работе с базой
type IDB interface {
	CloseDB() (err error)
}
