package gostarter

type ConfigBase struct {
	App AppConf `yaml:"app"`
}

type AppConf struct {
	Name string `yaml:"name"`
	Port int `yaml:"port"`
}