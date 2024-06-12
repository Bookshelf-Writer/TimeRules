package TimeRules

import (
	"fmt"
	"math/big"
)

//#################################################################################################//

/* Проверка времени на соответствие стандартам */
func (obj *TimeConfigurationRulesObj) Valid(tvr *DateTimeObj) error {
	err := func(name string) error {
		return fmt.Errorf("Invalid: %s\n", name)
	}

	if tvr.Type != obj.Name {
		return err("Type")
	}
	if len(tvr.Timezone) > 0 {
		_, status := obj.Timezones[tvr.Timezone]
		if !status {
			return err("Timezone")
		}
	}

	if tvr.Date.Year < obj.Year.Min || tvr.Date.Year > obj.Year.Max {
		return err("Year")
	}
	if tvr.Date.Month > uint16(len(obj.Month)) {
		return err("Month")
	}
	if tvr.Date.Day == 0 || tvr.Date.Day > obj.Month[tvr.Date.Month-1].Days {
		return err("Day")
	}

	if tvr.Time.Hour > obj.MaxHour {
		return err("Hour")
	}
	if tvr.Time.Min > obj.MaxMin {
		return err("Hour")
	}

	return nil
}

//#################################################################################################//

/*	Создание только даты (автоформатирование если зашли за пределы) */
func (obj *TimeConfigurationRulesObj) Date(year int64, month uint16, day uint16) DateObj {
	if year < obj.Year.Min {
		year = obj.Year.Min
	}
	if year > obj.Year.Max {
		year = obj.Year.Max
	}
	if month > uint16(len(obj.Month)) {
		month = uint16(len(obj.Month))
	}
	if month == 0 {
		month = 1
	}
	if month > uint16(len(obj.Month)) {
		month = uint16(len(obj.Month))
	}
	if day == 0 {
		day = 1
	}
	if day > obj.Month[month-1].Days {
		day = obj.Month[month-1].Days
	}

	return DateObj{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

/* Создание только времени (автоформатирование если зашли за пределы) */
func (obj *TimeConfigurationRulesObj) Time(hour uint16, minute uint16) TimeObj {
	if hour > obj.MaxHour {
		hour = obj.MaxHour
	}
	if minute > obj.MaxMin {
		minute = obj.MaxMin
	}

	return TimeObj{
		Hour: hour,
		Min:  minute,
	}
}

/*	Создание точки времени (никаких проверок)	*/
func (obj *TimeConfigurationRulesObj) DateTime(dateObj DateObj, timeObj TimeObj, timezone string) DateTimeObj {
	retObj := DateTimeObj{
		Type:     obj.Name,
		Timezone: timezone,

		Date: dateObj,
		Time: timeObj,
	}

	return retObj
}

/* Создание точки времени напрямую */
func (obj *TimeConfigurationRulesObj) DateFull(year int64, month uint16, day uint16, timezone string) DateTimeObj {
	return obj.DateTime(obj.Date(year, month, day), obj.Time(0, 0), timezone)
}

//#################################################################################################//

/*	Получение списка всех доступных таймзон	*/
func (obj *TimeConfigurationRulesObj) TimezoneList() []string {
	var timezones []string
	for k := range obj.Timezones {
		timezones = append(timezones, k)
	}
	return timezones
}

//#################################################################################################//

/* Перевод даты в число минут */
func (obj *TimeConfigurationRulesObj) DateTimeToMinutes(dateObj DateObj, timeObj TimeObj) *big.Int {
	totalMinutes := big.NewInt(int64(timeObj.Min))

	//Расчитываемые константы
	minutesPerDay := big.NewInt(0).Mul(big.NewInt(int64(obj.MaxHour+1)), big.NewInt(int64(obj.MaxMin+1)))
	minutesPerYear := big.NewInt(0).Mul(minutesPerDay, big.NewInt(int64(obj.DaysInYear)))

	// Подсчет минут до начала указанного месяца в году
	for i := uint16(0); i < dateObj.Month-1; i++ {
		minutesPerMonth := big.NewInt(0).Mul(minutesPerDay, big.NewInt(int64(obj.Month[i].Days)))
		totalMinutes.Add(totalMinutes, minutesPerMonth)
	}

	//Подсчет минут от начала до указаного года
	minutesPerYear.Mul(minutesPerYear, big.NewInt(dateObj.Year))
	totalMinutes.Add(totalMinutes, minutesPerYear)

	return totalMinutes
}

// reduceFromBigInt Получение смещения в int64 с вычитаем произведения этого смещения
func reduceFromBigInt(totalMinutes, componentSize *big.Int) (int64, *big.Int) {
	count := big.NewInt(0)
	count.DivMod(totalMinutes, componentSize, totalMinutes)
	return count.Int64(), totalMinutes
}

/* Перевод числа минут в дату */
func (obj *TimeConfigurationRulesObj) MinutesToDate(totalMinutes *big.Int) (dateObj DateObj, timeObj TimeObj) {
	minutesPerDay := big.NewInt(0).Mul(big.NewInt(int64(obj.MaxHour+1)), big.NewInt(int64(obj.MaxMin+1)))
	minutesPerYear := big.NewInt(0).Mul(minutesPerDay, big.NewInt(int64(obj.DaysInYear)))
	minutesPerHour := big.NewInt(int64(obj.MaxMin + 1))
	bufInt64 := int64(0)
	isNegative := false

	//Проверка на отрицательность
	if totalMinutes.Sign() < 0 {
		isNegative = true
		totalMinutes.Mul(totalMinutes, big.NewInt(-1))
	}

	// Вычисление года
	dateObj.Year, totalMinutes = reduceFromBigInt(totalMinutes, minutesPerYear)
	if isNegative {
		dateObj.Year *= -1
	}

	// Вычисление месяца
	for i, month := range obj.Month {
		minutesPerMonth := big.NewInt(0).Mul(minutesPerDay, big.NewInt(int64(month.Days)))
		if totalMinutes.Cmp(minutesPerMonth) < 0 {
			dateObj.Month = uint16(i + 1) //Установка месяца по точке отхода
			break
		}
		totalMinutes.Sub(totalMinutes, minutesPerMonth)
	}

	// Вычисление дня
	bufInt64, totalMinutes = reduceFromBigInt(totalMinutes, minutesPerDay)
	bufInt64++ // Так как дни начинаются с 1
	dateObj.Day = uint16(bufInt64)

	// Вычисление часа
	bufInt64, totalMinutes = reduceFromBigInt(totalMinutes, minutesPerHour)
	timeObj.Hour = uint16(bufInt64)

	// Оставшиеся минуты
	timeObj.Min = uint16(totalMinutes.Int64())

	return dateObj, timeObj
}

//..//

/* Изменение таймзоны со смешением временных диапазонов */
func (obj *TimeConfigurationRulesObj) LocationSet(DateTime DateTimeObj, timezone string) (dateTime DateTimeObj, err error) {

	oldOffset := int64(0)
	if len(DateTime.Timezone) > 0 {
		valid := false
		oldOffset, valid = obj.Timezones[DateTime.Timezone]
		if !valid {
			return dateTime, fmt.Errorf("Invalid Timezone (OLD)[%s]", timezone)
		}
	}

	newOffset := int64(0)
	if len(timezone) > 0 {
		valid := false
		newOffset, valid = obj.Timezones[timezone]
		if !valid {
			return dateTime, fmt.Errorf("Invalid Timezone (NEW)[ %s]", timezone)
		}
	}

	//Получение глобальной точки времени и смешение меж часовыми поясами
	bigTime := obj.DateTimeToMinutes(DateTime.Date, DateTime.Time)
	bigTime.Sub(bigTime, big.NewInt(oldOffset))
	bigTime.Add(bigTime, big.NewInt(newOffset))

	//Формирование обьекта даты
	objDate, objTime := obj.MinutesToDate(bigTime)
	dateTime = obj.DateTime(objDate, objTime, timezone)

	return dateTime, nil
}
