package db

import (
	"reflect"
	"regexp"
	"strings"
)

// 处理表名
func (_b *builder) tableNames() {
	_b.table = reflect.TypeOf(_b.model).Elem().Name()
	_b.table = pluralizeTableName(_b.table)
}

var tableNameRules = map[string]string{
	"(s)tatus$":                         "${1}${2}tatuses",
	"(quiz)$":                           "${1}zes",
	"^(ox)$":                            "${1}${2}en",
	"([m|l])ouse$":                      "${1}ice",
	"(matr|vert|ind)(ix|ex)$":           "${1}ices",
	"(x|ch|ss|sh)$":                     "${1}es",
	"([^aeiouy]|qu)y$":                  "${1}ies",
	"(hive|gulf)$":                      "${1}s",
	"(?:([^f])fe|([lr])f)$":             "${1}${2}ves",
	"sis$":                              "ses",
	"([ti])um$":                         "${1}a",
	"(c)riterion$":                      "${1}riteria",
	"(p)erson$":                         "${1}eople",
	"(m)an$":                            "${1}en",
	"(c)hild$":                          "${1}hildren",
	"(f)oot$":                           "${1}eet",
	"(buffal|her|potat|tomat|volcan)o$": "${1}${2}oes",
	"(alumn|bacill|cact|foc|fung|nucle|radi|stimul|syllab|termin|vir)us$": "${1}i",
	"us$":                           "uses",
	"(alias)$":                      "${1}es",
	"(analys|ax|cris|test|thes)is$": "${1}es",
	"s$":                            "s",
	"^$":                            "",
	"$":                             "s",
}
var tableNameRuleSort = []string{
	"(s)tatus$",
	"(quiz)$",
	"^(ox)$",
	"([m|l])ouse$",
	"(matr|vert|ind)(ix|ex)$",
	"(x|ch|ss|sh)$",
	"([^aeiouy]|qu)y$",
	"(hive|gulf)$",
	"(?:([^f])fe|([lr])f)$",
	"sis$",
	"([ti])um$",
	"(c)riterion$",
	"(p)erson$",
	"(m)an$",
	"(c)hild$",
	"(f)oot$",
	"(buffal|her|potat|tomat|volcan)o$",
	"(alumn|bacill|cact|foc|fung|nucle|radi|stimul|syllab|termin|vir)us$",
	"us$",
	"(alias)$",
	"(analys|ax|cris|test|thes)is$",
	"s$",
	"^$",
	"$",
}

func pluralizeTableName(name string) string {

	for _, rule := range tableNameRuleSort {
		value := tableNameRules[rule]
		if ok, _ := regexp.MatchString(rule, name); ok {
			r, _ := regexp.Compile(rule)
			name = r.ReplaceAllString(name, value)
			break
		}

	}

	return snakeCase(name)

}

func snakeCase(name string) string {
	r, _ := regexp.Compile("(.)([A-Z])")
	return strings.ToLower(r.ReplaceAllString(name, "${1}_${2}"))
}
