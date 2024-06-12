package TimeRules

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

const fileExTVR fileEx = "tvr" //Time Variance Rule

//#################################################################################################//

/* Запись правила в файл с расширением fileExTVR в массив байт */
func FileTVRwriteByte(ruleObj *TimeConfigurationRulesObj) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(*ruleObj); err != nil {
		return nil, err
	}

	byteBuf := buf.Bytes()
	byteBuf = compressed(&byteBuf)
	return byteBuf, nil
}

/* Запись правила в файл с расширением fileExTVR */
func FileTVRwrite(ruleObj *TimeConfigurationRulesObj, dirPath string, fileName string) error {

	//Попытка создания файла
	file, err := os.Create(dirPath + "/" + fileName + "." + string(fileExTVR))
	if err != nil {
		return err
	}
	defer file.Close()

	//Получение данных из структуры
	byteBuf, err := FileTVRwriteByte(ruleObj)
	if err != nil {
		return err
	}

	//Формирование структуры на запись
	writeObj := FF{
		hashBlakeByteToByte(&byteBuf),
		byteBuf,
	}

	//	Запись структуры в файл
	encoder := gob.NewEncoder(file)
	if err = encoder.Encode(writeObj); err != nil {
		return err
	}

	return nil
}

//.//

/* Парсинг правила из файла с расширением fileExTVR */
func FileTVRread(filePath string) (TimeConfigurationRulesObj, error) {
	return fileRead(&filePath, FileTVRreadIO)
}

/* Парсинг правила из файла с расширением fileExTVR в массив байт */
func FileTVRreadByte(filePath string) ([]byte, error) {
	return fileReadBytes(&filePath, FileTVRreadByteIO)
}

/*Парсинг правила из бинарого представления фала с расширением fileExTVR в массив байт  */
func FileTVRreadByteIO(file io.Reader) ([]byte, error) {

	//получение первичной структуры из файла
	decoder := gob.NewDecoder(file)
	bufObj := FF{}
	if err := decoder.Decode(&bufObj); err != nil {
		return nil, err
	}

	//проверка целостности
	hash := hashBlakeByteToByte(&bufObj.D)
	if !bytes.Equal(hash, bufObj.H) {
		return nil, fmt.Errorf("Break time struct\n")
	}

	//получение данных
	return decompressed(&bufObj.D), nil
}

/*Парсинг правила из бинарого представления фала с расширением fileExTVR */
func FileTVRreadIO(file io.Reader) (TimeConfigurationRulesObj, error) {
	obj := TimeConfigurationRulesObj{}

	//Получение данных из файла
	data, err := FileTVRreadByteIO(file)
	if err != nil {
		return obj, err
	}

	//получение основной структуры
	bufData := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(bufData)
	if err := decoder.Decode(&obj); err != nil {
		return obj, err
	}

	return obj, nil
}
