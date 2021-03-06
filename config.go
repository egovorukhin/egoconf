package egoconf

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//Возвращаем путь до лога
func getPath(filename string) (string, error) {

	//Если укахан абсолютный путь,
	//то выходим и используем его
	if filepath.IsAbs(filename) {
		return filename, nil
	}

	//Определяем путь до приложения
	app, err := os.Executable()
	if err != nil {
		return "", err
	}

	//Получаем путь директории приложения
	//и путь директории указанного файла конфигурации
	Path := filepath.Join(filepath.Dir(app), filepath.Dir(filename))

	//Проверяем на существование директории
	if _, err := os.Stat(Path); os.IsNotExist(err) {

		//Если нет директорий, то создаём их
		err = os.MkdirAll(Path, 0777)
		if err != nil {
			return "", err
		}
	}

	if filename == "" {
		filename = "config"
	}

	return filepath.Join(Path, filepath.Base(filename)), nil
}

//Сохраняем конфигурацию
//i interface - любая структура
func Save(filename string, i interface{}) error {

	Path, err := getPath(filename)
	if err != nil {
		return err
	}

	//Получаем расширение файла
	ext := getFileExtension(Path)

	//Создаем и открываем файл
	file, err := os.Create(Path)
	defer file.Close()
	if err != nil {
		return err
	}

	//Сериализуем данные
	buffer, err := ext.marshal(i)
	if err != nil {
		return err
	}

	//Пишем в файл
	_, err = file.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}

//Либо указываем расширение файла (приоритет),
//либо указываем Extension
func Load(filename string, i interface{}) error {

	//Получаем путь до файла
	Path, err := getPath(filename)
	if err != nil {
		return err
	}

	//Получаем Extension на основе расширения файла
	ext := getFileExtension(Path)

	//Если нет файла то создаём и сохраняем с сериализованными данными
	_, err = os.Stat(Path)
	if os.IsNotExist(err) {
		return err
	}

	//Читаем данные из файла
	file, err := ioutil.ReadFile(Path)
	if err != nil {
		return err
	}

	//Десериализуем данные в струтуру
	err = ext.unmarshal(file, i)
	if err != nil {
		return err
	}

	return nil
}
