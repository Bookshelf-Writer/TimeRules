package TimeRules

import (
	"bytes"
	"compress/flate"
	"github.com/mailru/easyjson"
	"golang.org/x/crypto/blake2b"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	jsonPath  = "json"  //	Папка с json-структурами
	rulesPath = "rules" //	Папка куда сохраняются валидные структуры
)

//#################################################################################################//

// jsonDecode Декодирование строки в переданую структуру
func jsonDecode(data *[]byte) (TimeConfigurationRulesObj, error) {
	obj := TimeConfigurationRulesObj{}
	err := easyjson.Unmarshal(*data, &obj)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

// jsonEncode Кодирование размеченой структуры в json-строку
func jsonEncode(obj *TimeConfigurationRulesObj) ([]byte, error) {
	jsonData, err := easyjson.Marshal(*obj)

	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

//#################################################################################################//

// isGoKeyword перевіряє, чи є рядок ключовим словом у Go
func isGoKeyword(word string) bool {
	keywords := []string{"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct", "chan", "else", "goto",
		"package", "switch", "const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var"}

	for _, keyword := range keywords {
		if word == keyword {
			return true
		}
	}
	return false
}

// createName	Получение валидного названия
func createName(input string) string {
	processed := regexp.MustCompile(`[^\w]`).ReplaceAllString(input, "_") // Видалення усіх символів, які не є літерами, цифрами або підкресленнями
	processed = regexp.MustCompile(`_+`).ReplaceAllString(processed, "_") // Видалення лишніх підкреслень і заміна послідовностей підкреслень одним підкресленням

	// Перевірка, чи починається результат літерою або підкресленням
	if len(processed) == 0 || unicode.IsDigit(rune(processed[0])) {
		processed = "_" + processed
	}

	// Перевірка на ключові слова Go
	if isGoKeyword(processed) {
		processed = "_" + processed
	}

	// Повернення обробленого рядка
	return processed
}

//.//

func stringInt(num int64) string {
	return strconv.FormatInt(num, 10)
}
func stringUint(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func stringIntSize(num int64, size uint16) (text string) {
	text = stringInt(num)
	l := len(text)
	s := int(size)

	if l < s {
		text = strings.Repeat("0", s-l) + text
	}
	return text
}

func stringFill(text string, pattern rune, size uint16) string {
	l := len(text)
	s := int(size)
	if l < s {
		text = strings.Repeat(string(pattern), s-l) + text
	}
	return text
}

func stringUintSize(num uint64, size uint16) (text string) {
	return stringFill(stringUint(num), '0', size)
}

//#################################################################################################//

// compressed Сжатие
func compressed(data *[]byte) []byte {
	var c bytes.Buffer

	writer, _ := flate.NewWriter(&c, flate.BestCompression)
	writer.Write(*data)
	writer.Close()

	return c.Bytes()
}

// decompressed Расжатие
func decompressed(data *[]byte) []byte {
	reader := flate.NewReader(bytes.NewReader(*data))

	d, _ := ioutil.ReadAll(reader)
	reader.Close()

	return d
}

//#################################################################################################//

// hashBlakeByteToByte ПОлучение хеша по данным
func hashBlakeByteToByte(data *[]byte) []byte {
	h, _ := blake2b.New(32, nil)
	h.Write(*data)
	return h.Sum(nil)
}
