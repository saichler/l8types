package ifs

type IPlugin interface {
	Install(IVNic) error
}
