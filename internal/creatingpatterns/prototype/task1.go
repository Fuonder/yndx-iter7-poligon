package prototype

//Напишите метод клонирования объекта указанного типа.
//
//type Config struct {
//	Version string
//	Plugins []string
//	Stat    map[string]int
//}
//
//func (cfg *Config) Clone() *Config {
//	// ...
//}

type Config struct {
	Version string
	Plugins []string
	Stat    map[string]int
}

func (cfg *Config) Clone() *Config {
	cloneObj := &Config{Version: cfg.Version}
	//cloneObj.Plugins = make([]string, len(cfg.Plugins))
	cloneObj.Stat = make(map[string]int, len(cfg.Stat))

	//for idx, item := range cfg.Plugins {
	//	cloneObj.Plugins[idx] = item
	//}

	for _, item := range cfg.Plugins {
		cloneObj.Plugins = append(cloneObj.Plugins, item)
	}

	for idx, item := range cfg.Stat {
		cloneObj.Stat[idx] = item
	}
	return cloneObj
}
