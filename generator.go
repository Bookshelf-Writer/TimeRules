package TimeRules

import (
	"embed"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strings"
)

// Структура сохраняемого файла
type FF struct {
	H []byte //	hash
	D []byte //	data
}

func init() {
	gob.Register(FF{})
}

type fileEx string

//.//

//go:embed rules/Earth0.tvr
var defFileRule embed.FS //	Встроеная переменная с правилом по умолчанию

//#################################################################################################//

/* Получение структуры времени 'по умолчанию' (встроена в библиотеку) */
func DefRules() (TimeConfigurationRulesObj, error) {

	//получение имени файла
	fileArr, err := defFileRule.ReadDir(rulesPath)
	if err != nil {
		return TimeConfigurationRulesObj{}, err
	}

	//открываем на чтение встраивание
	fileObj, err := defFileRule.Open(rulesPath + "/" + fileArr[0].Name())
	if err != nil {
		return TimeConfigurationRulesObj{}, err
	}

	//получаем расширение и отправляем на распаковку
	ex := strings.Split(fileArr[0].Name(), ".")
	obj, err := FileReadIO(fileObj, fileEx(ex[len(ex)-1]))
	fileObj.Close()

	return obj, err
}

// fileRead Чтение файла на структуру
func fileRead(filePath *string, retFunction func(file io.Reader) (TimeConfigurationRulesObj, error)) (TimeConfigurationRulesObj, error) {
	file, err := os.Open(*filePath)
	if err != nil {
		return TimeConfigurationRulesObj{}, err
	}

	obj, err := retFunction(file)
	file.Close()

	return obj, err
}

// fileReadBytes Чтение файла на байты
func fileReadBytes(filePath *string, retFunction func(file io.Reader) ([]byte, error)) ([]byte, error) {
	file, err := os.Open(*filePath)
	if err != nil {
		return nil, err
	}

	data, err := retFunction(file)
	file.Close()

	return data, err
}

//#################################################################################################//

/* Парсинг правила из файла */
func FileRead(filePath string) (TimeConfigurationRulesObj, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return TimeConfigurationRulesObj{}, err
	}

	parts := strings.Split(filePath, ".")
	obj, err := FileReadIO(file, fileEx(parts[len(parts)-1]))
	file.Close()

	return obj, err
}

/*Парсинг правила из бинарого представления фала */
func FileReadIO(file io.Reader, ex fileEx) (TimeConfigurationRulesObj, error) {
	switch ex {

	case fileExTVR:
		return FileTVRreadIO(file)

	case fileExJSON:
		return JSONreadIO(file)

	}

	return TimeConfigurationRulesObj{}, fmt.Errorf("Unknown file extension *.%s\n", ex)
}

//.//

/* Запись правила в файл с указаным расширением */
func FileWrite(ruleObj *TimeConfigurationRulesObj, dirPath string, ex fileEx) error {
	switch ex {

	case fileExTVR:
		return FileTVRwrite(ruleObj, dirPath, ruleObj.Name)

	case fileExJSON:
		return JSONwrite(ruleObj, dirPath, ruleObj.Name)

	}

	return fmt.Errorf("Unknown file extension *.%s\n", ex)
}
