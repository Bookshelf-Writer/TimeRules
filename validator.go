package TimeRules

import (
	"math"
	"regexp"
)

//#################################################################################################//

type validateObj struct {
	errors []string
}

type validateValueObj struct {
	valid *validateObj
	name  string
	value interface{}
}

//..//

func (obj *validateObj) add(text string) {
	obj.errors = append(obj.errors, text)
}

func (obj *validateObj) valid(value interface{}, name string) *validateValueObj {
	validObj := validateValueObj{
		obj,
		name,
		value,
	}
	return &validObj
}
func (obj *validateValueObj) max(limit int64) *validateValueObj {
	funcError := func() {
		obj.valid.add("Variable should be: `" + obj.name + "` < " + stringInt(limit))
	}

	switch v := obj.value.(type) {
	case string:
		if len(v) > int(limit) {
			funcError()
		}
	case []any:
		if len(v) > int(limit) {
			funcError()
		}
	case []MonthObj:
		if len(v) > int(limit) {
			funcError()
		}
	case map[string]int64:
		if len(v) > int(limit) {
			funcError()
		}
	case int:
		if v > int(limit) {
			funcError()
		}
	case int32:
		if v > int32(limit) {
			funcError()
		}
	case int64:
		if v > limit {
			funcError()
		}
	case uint16:
		if v > uint16(limit) {
			funcError()
		}
	case uint32:
		if v > uint32(limit) {
			funcError()
		}
	case uint64:
		if v > uint64(limit) {
			funcError()
		}

	default:
		obj.valid.add("Value `" + obj.name + "`: Type not supported")
	}

	return obj
}
func (obj *validateValueObj) min(limit int64) *validateValueObj {
	funcError := func() {
		obj.valid.add("Variable should be: `" + obj.name + "` > " + stringInt(limit))
	}

	switch v := obj.value.(type) {
	case string:
		if len(v) < int(limit) {
			funcError()
		}
	case []any:
		if len(v) < int(limit) {
			funcError()
		}
	case []MonthObj:
		if len(v) < int(limit) {
			funcError()
		}
	case map[string]int64:
		if len(v) < int(limit) {
			funcError()
		}
	case int:
		if v < int(limit) {
			funcError()
		}
	case int16:
		if v < int16(limit) {
			funcError()
		}
	case int32:
		if v < int32(limit) {
			funcError()
		}
	case int64:
		if v < int64(limit) {
			funcError()
		}
	case uint16:
		if v < uint16(limit) {
			funcError()
		}
	case uint32:
		if v < uint32(limit) {
			funcError()
		}
	case uint64:
		if v < uint64(limit) {
			funcError()
		}

	default:
		obj.valid.add("Value `" + obj.name + "`: Type not supported")
	}

	return obj
}
func (obj *validateValueObj) text(regExp string) *validateValueObj {
	funcValid := func(text string) {
		re := regexp.MustCompile(regExp)
		if !re.MatchString(text) {
			obj.valid.add("Value {`" + obj.name + "`}['" + text + "'] not valid, Allowed only: `" + regExp + "`")
		}
	}

	switch v := obj.value.(type) {
	case string:
		funcValid(v)
	default:
		obj.valid.add("Value `" + obj.name + "`: only string-type")
	}

	return obj
}

//#################################################################################################//

/* Проверяет структуру на синтаксические ошибки */
func (obj *TimeConfigurationRulesObj) CheckErrors() (errors []string) {
	bufObj := validateObj{}

	bufObj.valid(obj.Name, "name").min(3).max(32).text(`^[a-zA-Z0-9 \-.:;_+@]*$`)
	bufObj.valid(obj.Description, "description").max(1000).text(`^[^'"\\]+$`)

	bufObj.valid(obj.INF.Ver, "inf.ver").min(5).max(12).text(`^[0-9\.]*$`)
	bufObj.valid(obj.INF.Creator, "inf.creator").max(32).text(`^[^'"\\%&,/]+$`)

	bufObj.valid(obj.Year.Min, "year.min").min(math.MinInt64).max(math.MaxInt64)
	bufObj.valid(obj.Year.Max, "year.max").min(math.MinInt64).max(math.MaxInt64)
	bufObj.valid(obj.MaxHour, "MaxHour").min(2).max(math.MaxUint16)
	bufObj.valid(obj.MaxMin, "MaxMin").min(2).max(math.MaxUint16)

	bufObj.valid(obj.Month, "month").min(1)
	bufObj.valid(obj.Timezones, "timezones").min(1)

	//Перебираем месяцы если до этого ошибок не было
	if len(bufObj.errors) == 0 {
		for _, month := range obj.Month {
			bufObj.valid(month.FullName, "month.fullName").min(1).max(32).text(`^[^'"\\%&,/]+$`)
			bufObj.valid(month.ShortName, "month.fullName").min(1).max(4).text(`^[a-zA-Z0-9]*$`)
			bufObj.valid(month.Days, "year.days").min(1).max(math.MaxUint16)
		}
	}

	return bufObj.errors
}
