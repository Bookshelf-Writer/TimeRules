package TimeRules

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"io"
	"os"
)

const fileExJSON fileEx = "json"

//#################################################################################################//

/* Парсинг правила из json */
func JSONread(filePath string) (TimeConfigurationRulesObj, error) {
	return fileRead(&filePath, JSONreadIO)
}

/*Парсинг правила из бинарого представления json*/
func JSONreadIO(file io.Reader) (TimeConfigurationRulesObj, error) {
	obj := TimeConfigurationRulesObj{}

	//получение данных из файла
	data, err := io.ReadAll(file)
	if err != nil {
		return obj, err
	}

	//парсинг json в структуру
	obj, err = jsonDecode(&data)
	if err != nil {
		return obj, err
	}

	//Получение количества дней
	days := uint64(0)
	for _, parseObj := range obj.Month {
		days += uint64(parseObj.Days)
	}
	obj.DaysInYear = days

	return obj, nil
}

//.//

/* Запись правила в json */
func JSONwriteByte(ruleObj *TimeConfigurationRulesObj) ([]byte, error) {
	jsonData, err := easyjson.Marshal(*ruleObj)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/* Запись правила в json-файл */
func JSONwrite(ruleObj *TimeConfigurationRulesObj, dirPath string, fileName string) error {

	//Попытка создания файла
	file, err := os.Create(dirPath + "/" + fileName + "." + string(fileExJSON))
	if err != nil {
		return err
	}
	defer file.Close()

	//ПОлучение данных
	jsonData, err := JSONwriteByte(ruleObj)
	if err != nil {
		return err
	}

	file.Write(jsonData)
	return nil
}

/* Запись правила в json с отступами */
func JSONwritePrettyByte(ruleObj *TimeConfigurationRulesObj) ([]byte, error) {
	jsonData, err := json.MarshalIndent(*ruleObj, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

/* Запись правила в json-файл с отступами */
func JSONwritePretty(ruleObj *TimeConfigurationRulesObj, dirPath string, fileName string) error {

	//Попытка создания файла
	file, err := os.Create(dirPath + "/" + fileName + "." + string(fileExJSON))
	if err != nil {
		return err
	}
	defer file.Close()

	//ПОлучение данных
	jsonData, err := JSONwritePrettyByte(ruleObj)
	if err != nil {
		return err
	}

	file.Write(jsonData)
	return nil
}
