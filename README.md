# EgoConf
Модуль позволяет сохранять и загружать файлы конфигураций четырех популярных форматов данных: XML, JSON, INI, YAML. Вам не нужно заботиться о конвертации данных в своем приложении, так как модуль понимает расширение файла и работает с тем форматом данных, который присутствует в файле.
### Установка
```
go get github.com/egovorukhin/egoconf
```
### Начало работы
* **Инициализация**
```golang
  import "github.com/egovorukhin/egoconf"
  ...
  
  type Config struct {
	  Host     string `xml:"Host" json:"host" yaml:"host" ini:"host"`
	  Port     int    `xml:"Port" json:"port" yaml:"port" ini:"port"`
	  Username string `xml:"Username" json:"username" yaml:"username" ini:"username"`
	  Password string `xml:"Password" json:"password" yaml:"password" ini:"password"`
	  DBName   string `xml:"DBName" json:"db_name" yaml:"dbName" ini:"dbName"`
  }  
  ...
```
Описываем структуру конфигурации, указываем в полях тэги для необходимого формата данных
* **Сохранение**
```golang
  cfg := db.Config{
		Host: "server",
		Port: 5432,
		Username: "user",
		Password: "pass",
		DBName: "name",
	}
  err = egoconf.Save("config.yml", cfg)
	if err != nil {
		log.Fatal(err)
	}
```
* **Загрузка**
```golang
  err = egoconf.Load("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
```

* **config.xml**
```
  <Config>
    <Host>server</Host>
    <Port>5432</Port>
    <Username>user</Username>
    <Password>pass</Password>
    <DBName>name</DBName>
</Config>
```
* **config.json**
```
{
    "host": "server",
    "port": 5432,
    "username": "user",
    "password": "pass",
    "db_name": "name"
}
```
* **config.ini**
```

```
