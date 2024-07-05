package qbittorrent

type Search interface {
	Start()
	Stop()
	Status()
	Results()
	Delete()
	Plugins()
	InstallPlugins()
	UninstallPlugins()
	EnableSearchPlugins()
	UpdateSearchPlugins()
}

func (c *client) Start() {
	//TODO implement me
	panic("implement me")
}

func (c *client) Stop() {
	//TODO implement me
	panic("implement me")
}

func (c *client) Status() {
	//TODO implement me
	panic("implement me")
}

func (c *client) Results() {
	//TODO implement me
	panic("implement me")
}

func (c *client) Delete() {
	//TODO implement me
	panic("implement me")
}

func (c *client) Plugins() {
	//TODO implement me
	panic("implement me")
}

func (c *client) InstallPlugins() {
	//TODO implement me
	panic("implement me")
}

func (c *client) UninstallPlugins() {
	//TODO implement me
	panic("implement me")
}

func (c *client) EnableSearchPlugins() {
	//TODO implement me
	panic("implement me")
}

func (c *client) UpdateSearchPlugins() {
	//TODO implement me
	panic("implement me")
}
