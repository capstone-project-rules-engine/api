package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

type Rule struct {
	Id         int
	Conditions map[string]interface{}
	Action     interface{}
}

type Condition struct {
	Label     string
	Attribute string
	Operator  string
}

type Body struct {
	Name string
	Type string
}

type Action struct {
	Label     string
	Attribute string
	Type      string
}

type Description struct {
	Condition string
	Action    string
}

// di satuin
type DB struct {
	Name       string
	Endpoint   string
	Bodies     []Body
	Conditions []Condition
	Action
	Rules []Rule
	Description
}

func GenerateDummy() *DB {
	return &DB{
		Name:     "RuleSet 1",
		Endpoint: "ruleset1",
		Bodies: []Body{
			{
				Name: "harga",
				Type: "number",
			},
			{
				Name: "jualan",
				Type: "number",
			},
		},
		Conditions: []Condition{
			{
				Label:     "$hargalebih",
				Attribute: "harga",
				Operator:  ">",
			},
			{
				Label:     "$jualanlebih",
				Attribute: "jualan",
				Operator:  ">",
			},
		},
		Action: Action{
			Label:     "$diskon",
			Attribute: "diskon",
			Type:      "number",
		},
		Rules: []Rule{
			{
				Id: 1,
				Conditions: map[string]interface{}{
					"$hargalebih":  10,
					"$jualanlebih": 5,
				},
				Action: 30,
			},
			{
				Id: 2,
				Conditions: map[string]interface{}{
					"$jualanlebih": 10,
				},
				Action: 20,
			},
			{
				Id:         3,
				Conditions: map[string]interface{}{},
				Action:     5,
			},
		},
	}
}

func ValidateRule(operator string, initValue, inputValue interface{}) (bool, error) {
	expressionString := fmt.Sprintf("%v %s %v", inputValue, operator, initValue)
	expression, err := govaluate.NewEvaluableExpression(expressionString)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	parameters := make(map[string]interface{})
	result, err := expression.Evaluate(parameters)
	if err != nil {
		return false, err
	}

	return result.(bool), nil
}

func CheckBoolean(slice []bool) bool {
	for _, value := range slice {
		if !value {
			return false
		}
	}
	return true
}

func Exec(ruleSet string, rulesSelected DB, inputUSer map[string]interface{}) (interface{}, error) {
	var boolComplete []bool
	for _, rule := range rulesSelected.Rules {
		boolComplete = boolComplete[:0] // emptying the slice
		for _, condition := range rulesSelected.Conditions {
			// cek empty map
			if len(rule.Conditions) == 0 {
				return rule.Action, nil
			}
			// cek key ada apa kgk
			if _, exists := rule.Conditions[condition.Label]; !exists {
				continue
			}
			// cek key di inout user ada apa kgk
			if inputUSer[condition.Attribute] == nil {
				boolComplete = append(boolComplete, false)
				continue
			}
			// validasi rule
			result, err := ValidateRule(condition.Operator, rule.Conditions[condition.Label], inputUSer[condition.Attribute])
			if err != nil {
				return nil, err
			}
			boolComplete = append(boolComplete, result)
		}
		if CheckBoolean(boolComplete) {
			return rule.Action, nil
		}
	}

	return nil, nil
}

func main() {
	input := map[string]interface{}{
		"nama":  "suki",
		"harga": 15,
		"jualan": 20,
	}

	rules := GenerateDummy()

	resultAction, err := Exec("Rule Set 1", *rules, input)
	if err != nil {
		fmt.Println("error during exec: ", err)
	}
	fmt.Println("result: ", resultAction)
}
