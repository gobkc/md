package factory

const (
	DbConnect = iota
	InstallSelf
	MakeBin
	CreateGin
)

type WorkFactory struct {

}

func (w *WorkFactory) New(workType int,f func()) {
	switch workType {
	case DbConnect:
		f()
	}
}