package TimeRules

import (
	"os"
	"path/filepath"
	"testing"
)

// Компиляция правил из json
func TestBuild(t *testing.T) {
	testObj.begin(t)

	//Получение файлов с разметкой правил в json
	fileArr, err := os.ReadDir(jsonPath)
	if err == nil {

		//Удаление всех ранее сгенерированных файлов
		delFiles, err := os.ReadDir(rulesPath)
		if err == nil {
			for _, file := range delFiles {
				filePath := filepath.Join(rulesPath, file.Name())
				os.Remove(filePath)
			}
		}

		//Перебор правил
		for _, fileJsonObj := range fileArr {
			filePath := fileJsonObj.Name() //Получение название файла
			t.Run(filePath, func(t *testing.T) {
				testObj.next(t)

				//загрузка данных из файла
				ruleObj, err := FileRead(jsonPath + "/" + filePath)
				if err != nil {
					t.Error(err)
					return
				}

				//Выводим информацию о правиле
				t.Log(ruleObj.INF.Ver, ruleObj.Name, "\n", ruleObj.Description, ruleObj.INF.Creator)

				//Проверяем структуру
				for _, textError := range ruleObj.CheckErrors() {
					testObj.errorText(textError)
				}

				//сохранение TVR-файла
				if !t.Failed() {
					fileName := createName(ruleObj.Name)
					err = FileTVRwrite(&ruleObj, rulesPath, fileName)
					if err != nil {
						t.Error(err)
					}
				}

				testObj.timePrint()
				testObj.end()
			})
		}
	}

	testObj.end()
}
