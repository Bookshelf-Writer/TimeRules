package TimeRules

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type TestObj struct {
	t *testing.T

	fail  bool
	timer time.Time
}

var testObj = TestObj{nil, false, time.Now()}

//#################################################################################################//

func (obj *TestObj) randomInt(from int64, to int64) int64 {
	if from > to {
		buf := to
		to = from
		from = buf
	}

	//Сумма для реально больших пределов
	sum := uint64(0)
	if from < 0 {
		sum = uint64(to) + uint64(-1*from)
	} else {
		sum = uint64(to - from)
	}

	//Урезание осетра если все так плохо
	if math.MaxInt64 < sum {
		sum = math.MaxInt64 - 4
	}

	randomNumber := from + rand.Int63n(int64(sum)+1)
	return randomNumber
}
func (obj *TestObj) randomDate(ruleObj *TimeConfigurationRulesObj) DateObj {
	year := testObj.randomInt(ruleObj.Year.Min, ruleObj.Year.Max)
	month := testObj.randomInt(0, int64(len(ruleObj.Month)))
	if month > 0 {
		month -= 1
	}
	day := testObj.randomInt(0, int64(ruleObj.Month[month].Days))

	return ruleObj.Date(year, uint16(month), uint16(day))
}
func (obj *TestObj) randomTime(ruleObj *TimeConfigurationRulesObj) TimeObj {
	hour := testObj.randomInt(0, int64(ruleObj.MaxHour))
	minute := testObj.randomInt(0, int64(ruleObj.MaxMin))

	return ruleObj.Time(uint16(hour), uint16(minute))
}
func (obj *TestObj) randomTimezone(ruleObj *TimeConfigurationRulesObj) string {
	timezones := ruleObj.TimezoneList()
	return timezones[rand.Intn(len(timezones))]
}

func (obj *TestObj) begin(t *testing.T) {
	obj.t = t

	if obj.fail {
		obj.t.SkipNow()
	}
	obj.timer = time.Now()
}
func (obj *TestObj) next(t *testing.T) {
	obj.t = t
}
func (obj *TestObj) run(name string, f func()) {
	obj.t.Run(name, func(t *testing.T) {
		testObj.next(t)
		f()
		obj.end()
	})
}
func (obj *TestObj) end() {
	if obj.t.Failed() {
		obj.fail = true
	}
}
func (obj *TestObj) time() time.Duration {
	buf := time.Since(obj.timer)
	obj.timer = time.Now()
	return buf
}
func (obj *TestObj) timePrint() {
	fmt.Printf("### TIMER %s\n", obj.time())
}

func (obj *TestObj) error(err error) {
	if err != nil {
		obj.t.Error(err)
	}
}
func (obj *TestObj) errorText(text string) {
	obj.t.Error(fmt.Errorf("%s", text))
}
func (obj *TestObj) errorSkip(err error) {
	if err != nil {
		obj.t.Error(err)
		obj.fail = true
		obj.t.SkipNow()
	}
}

func (obj *TestObj) fileInPath(path string) []string {

	//Сканируем директорию
	fileArr, err := os.ReadDir(path)
	if err != nil {
		obj.t.Error(err)
		obj.fail = true
		obj.t.Fail()
		return []string{}
	}

	//формируем список названий файла
	var retArr []string
	for _, fileObj := range fileArr {
		retArr = append(retArr, fileObj.Name())
	}

	return retArr
}
func (obj *TestObj) clearDir(path string) {
	for _, fileName := range obj.fileInPath(path) {
		os.Remove(filepath.Join(path, fileName))
	}
}

/* ################################################################################################## */

// TestValidateJSON Проверка json-файлов на корректность
func TestValidateJSON(t *testing.T) {
	testObj.begin(t)

	//Перебираем все файлы в директории jsonPath
	for _, fileName := range testObj.fileInPath(jsonPath) {
		t.Run(fileName, func(t *testing.T) {
			testObj.next(t)

			//читаем файл
			fileJsonObj, err := FileRead(jsonPath + "/" + fileName)
			testObj.errorSkip(err)

			//Проверяем структуру
			for _, textError := range fileJsonObj.CheckErrors() {
				testObj.errorText(textError)
			}

			testObj.timePrint()
			testObj.end()
		})
	}

	testObj.end()
}

// TestTimeTransform Проверка операций над временен
func TestTimeTransform(t *testing.T) {
	testObj.begin(t)

	//Перебираем все файлы в директории jsonPath
	for _, fileName := range testObj.fileInPath(jsonPath) {
		t.Run(fileName, func(t *testing.T) {
			rand.Seed(time.Now().UnixNano())
			testObj.next(t)

			//читаем файл
			fileJsonObj, err := FileRead(jsonPath + "/" + fileName)
			testObj.errorSkip(err)

			//Создание случайной даты
			timeObj := fileJsonObj.DateTime(
				testObj.randomDate(&fileJsonObj),
				testObj.randomTime(&fileJsonObj),
				testObj.randomTimezone(&fileJsonObj),
			)

			//проверка созданой датой валидатором
			err = fileJsonObj.Valid(&timeObj)
			testObj.errorSkip(err)

			//Переводим в другой часовой пояс
			newTimeObj, err := fileJsonObj.LocationSet(timeObj, testObj.randomTimezone(&fileJsonObj))
			testObj.errorSkip(err)

			fmt.Println(">>>>>>", timeObj.String(), " 	||	 ", newTimeObj.String(), "	 METHOD")

			//проверка даты методами time.* если только дата попадает в ожидаемые пределы
			if timeObj.Date.Year < math.MaxInt16 && timeObj.Date.Year > math.MinInt16 {
				if len(fileJsonObj.Month) == 12 && fileJsonObj.MaxHour == 23 && fileJsonObj.MaxMin == 59 {
					testObj.run("Go-time", func() {
						newLocalTime := newTimeObj

						//Создание метки времени в time.Date
						location, _ := time.LoadLocation(timeObj.Timezone)
						newTime := time.Date(int(newLocalTime.Date.Year), time.Month(newLocalTime.Date.Month), int(newLocalTime.Date.Day), int(newLocalTime.Time.Hour), int(newLocalTime.Time.Min), 0, 0, location)

						fmt.Println(newTime)

						//проверка условий
						if int(newLocalTime.Date.Year) != newTime.Year() {
							testObj.errorText("Break in YEAR")
						}
						if newLocalTime.Date.Month != uint16(newTime.Month()) {
							testObj.errorText("Break in MONTH")
						}
						if newLocalTime.Date.Day != uint16(newTime.Day()) {
							testObj.errorText("Break in DAY")
						}
						if newLocalTime.Time.Hour != uint16(newTime.Hour()) {
							testObj.errorText("Break in Hour")
						}
						if newLocalTime.Time.Min != uint16(newTime.Minute()) {
							testObj.errorText("Break in Minute")
						}
					})
				}
			}
			testObj.timePrint()
			testObj.end()
		})
	}
	testObj.end()
}

// TestFormats базовый тест проверки форматирования (возможно когда-то напишется нормальный тест)
func TestFormats(t *testing.T) {
	testObj.begin(t)
	rand.Seed(time.Now().UnixNano())

	//читаем файл
	fileJsonObj, err := DefRules()
	testObj.errorSkip(err)

	//Создание случайной даты
	timeObj := fileJsonObj.DateTime(
		testObj.randomDate(&fileJsonObj),
		testObj.randomTime(&fileJsonObj),
		"",
	)

	testObj.run("Year", func() {
		timeObj.Date = fileJsonObj.Date(math.MaxInt, timeObj.Date.Month, timeObj.Date.Day)

		//смещение
		numbString := stringInt(timeObj.Date.Year)
		if fileJsonObj.Format(&timeObj, "${Y}") != numbString {
			testObj.errorText("Break full-print")
		}
		if fileJsonObj.Format(&timeObj, "${Ys2}") != numbString[len(numbString)-2:] {
			testObj.errorText("Break base slice")
		}
		if len(fileJsonObj.Format(&timeObj, "${Ys5:3}")) != 3 {
			testObj.errorText("Break slice")
		}

		//деление
		numbString = stringInt(timeObj.Date.Year / 4)
		if fileJsonObj.Format(&timeObj, "${Ym4}") != numbString {
			testObj.errorText("Break division")
		}
		if fileJsonObj.Format(&timeObj, "${Ym4:2}") != numbString[len(numbString)-2:] {
			testObj.errorText("Break base slice in division")
		}
		if len(fileJsonObj.Format(&timeObj, "${Ym5:5:3}")) != 3 {
			testObj.errorText("Break slice in division")
		}

	})

	testObj.timePrint()
	testObj.end()
}
