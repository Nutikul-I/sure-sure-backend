package handler

type PackageHandler interface {
}

type packageHandler struct {
}

func NewPackageHandler() packageHandler {
	return packageHandler{}
}
