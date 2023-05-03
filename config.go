package egoconf

import (
	"os"
	"path/filepath"
)

// Возвращаем путь до лога
func getPath(filename string) (string, error) {

	// Если укахан абсолютный путь,
	// то выходим и используем его
	if filepath.IsAbs(filename) {
		return filename, nil
	}

	// Определяем путь до приложения
	app, err := os.Executable()
	if err != nil {
		return "", err
	}

	// Получаем путь директории приложения
	// и путь директории указанного файла конфигурации
	Path := filepath.Join(filepath.Dir(app), filepath.Dir(filename))

	// Проверяем на существование директории
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

// SaveExtension передаем переопределенный интерфейс для своих расширений
func SaveExtension(filename string, i interface{}, ext IExtension) error {
	Path, err := getPath(filename)
	if err != nil {
		return err
	}

	if ext == nil {
		//Получаем расширение файла
		ext, err = getFileExtension(Path)
		if err != nil {
			return err
		}
	}

	//Создаем и открываем файл
	file, err := os.Create(Path)
	defer file.Close()
	if err != nil {
		return err
	}

	//Сериализуем данные
	buffer, err := ext.Marshal(i)
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

// Save Сохраняем конфигурацию. i interface - любая структура
func Save(filename string, i interface{}) error {
	return SaveExtension(filename, i, nil)
}

// LoadExtension передаем переопределенный интерфейс для своих расширений
func LoadExtension(filename string, i interface{}, ext IExtension) error {

	//Получаем путь до файла
	p, err := getPath(filename)
	if err != nil {
		return err
	}

	if ext == nil {
		// Получаем Extension на основе расширения файла
		ext, err = getFileExtension(p)
		if err != nil {
			return err
		}
	}
	// Проверка файла
	file, err := checkFile(p)
	if err != nil {
		return err
	}

	// Десериализуем данные в струтуру
	return ext.Unmarshal(file, i)
}

// Load Указываем файл ОБЯЗАТЕЛЬНО С РАСШИРЕНИЕМ файла, по нему будем определять сериализатор
func Load(filename string, i interface{}) error {
	return LoadExtension(filename, i, nil)
}

func checkFile(path string) (file []byte, err error) {

	// Получаем инфо о существовании файла
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return
	}
	// Читаем данные из файла
	return os.ReadFile(path)
}
