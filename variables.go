package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randomChar() rune {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(26)
	return rune('a' + i)
}

func randomString(len int) string {
	result := []rune{}
	for i := 0; i < len; i++ {
		c := randomChar()
		result = append(result, c)
	}
	return string(result)
}

func randomType() string {
	types := []string{"int", "float", "bool", "string", "error"}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(types))
	result := types[i]
	rand.Seed(time.Now().UnixNano())
	switch result {
	case "int":
		i := rand.Intn(500) % 6
		switch i {
		case 0:
			return "int"
		case 1:
			return "int32"
		case 2:
			return "int64"
		case 3:
			return "uint"
		case 4:
			return "uint32"
		default:
			return "uint64"
		}
	case "float":
		i := rand.Intn(500) % 2
		switch i {
		case 0:
			return "float32"
		default:
			return "float64"
		}
	}
	return types[i]
}

func randomValue(vtype string) string {
	rand.Seed(time.Now().UnixNano())
	switch vtype {
	case "int", "int32", "int64":
		i := rand.Intn(1000) - 500
		return strconv.Itoa(i)
	case "uint32", "uint", "uint64":
		i := rand.Intn(1000)
		return strconv.Itoa(i)
	case "float32", "float64":
		i := float64(rand.Intn(500)) / 20
		return fmt.Sprintf("%.2f", i)
	case "bool":
		i := bool(rand.Intn(500)%2 == 0)
		return fmt.Sprintf("%t", i)
	case "string":
		return fmt.Sprintf("\"%s\"", randomString(5))
	case "error":
		return fmt.Sprintf("errors.New(\"%s\")", randomString(5))
	default:
		return ""
	}
}

type Var struct {
	Name  string
	Type  string
	Value string
	Str   string
}

func randomVariable(name string) Var {
	result := Var{Name: name}
	result.Type = randomType()
	result.Value = randomValue(result.Type)
	result.Str = fmt.Sprintf("переменную %s с типом %s значением %s;", result.Name, result.Type, result.Value)
	return result
}

func main() {
	for i := 0; i < 26; i++ {
		var name1, name2, name3 string
		name1 = string(randomChar())
		for true {
			name2 = string(randomChar())
			if name1 != name2 {
				break
			}
		}
		for true {
			name3 = string(randomChar())
			if name3 != name1 && name3 != name2 {
				break
			}
		}
		var1 := randomVariable(string(name1)).Str
		var2 := randomVariable(string(name2)).Str
		var3 := randomVariable(string(name3)).Str
		fmt.Printf("%d. Инициализируйте %s %s %s и выведите их на экран\n", i, var1, var2, var3)
	}

}
