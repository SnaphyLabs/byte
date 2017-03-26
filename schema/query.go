package schema

import (
	"regexp"
	"encoding/json"
	"errors"
	"fmt"
)

// Query defines an expression against a schema to perform a match on schema's data.
type Query []Expression

// Expression is a query or query component that can be matched against a payoad.
type Expression interface {
	Match(payload map[string]interface{}) bool
}

// Value represents any kind of value to use in query
type Value interface{}

//Name of user defined collections.
type COLLECTION struct {
	Value string
}


//------------------------------------------------------COMPARISION QUERY OPERATORS------------------------------------------------------------------------

// Equal matches all values that are equal to a specified value.
type Equal struct {
	Field string
	Value Value
	//Value float64
}

// NotEqual matches all values that are not equal to a specified value.
type NotEqual struct {
	Field string
	Value Value
	//Value float64
}


// GreaterThan matches values that are greater than a specified value.
type GreaterThan struct {
	Field string
	Value Value
}

// GreaterOrEqual matches values that are greater than or equal to a specified value.
type GreaterOrEqual struct {
	Field string
	Value Value
	//Value float64
}

// LowerThan matches values that are less than a specified value.
type LowerThan struct {
	Field string
	//Value float64
	Value Value
}

// LowerOrEqual matches values that are less than or equal to a specified value.
type LowerOrEqual struct {
	Field string
	Value Value
	//Value float64
}



// In natches any of the values specified in an array.
type In struct {
	Field  string
	Values []Value
}

// NotIn matches none of the values specified in an array.
type NotIn struct {
	Field  string
	Values []Value

}



//-------------------------------------------------------------LOGICAL QUERY OPERATOR----------------------------------------------------------------


// And joins query clauses with a logical AND, returns all documents
// that match the conditions of both clauses.
type And []Query

// Or joins query clauses with a logical OR, returns all documents that
// match the conditions of either clause.
type Or []Query



type Not struct{
	Field string
	Value Expression
}

type Nor []Query


//--------------------------------------------------------------ELEMENT QUERY OPERATOR-----------------------------------------------------------------

// Exist matches all values which are present, even if nil
type Exist struct {
	Field string
	Value bool
}


//--------------------------------------------------------------EVALUATION QUERY OPERATOR--------------------------------------------------------------
//Mod operation..
type Mod struct {
	Field string
	Divisor float64
	Remainder float64
}


//Regex matches values that match to a specified regular expression.
type Regex struct {
	Field string
	Value *regexp.Regexp
}

//Supports text based search..
type Text struct {
	Search string
	Language string
	CaseSensitive bool
	DiacriticSensitive bool
}

//NewQuery returns a new query with the provided key/value validated against validator
func NewQuery(q map[string]interface{}) (Query, error) {
	return validateQuery(q, "")
}

// ParseQuery parses and validate a query as string
//pattern
// {
// 	"field":{
//		"$regex": /* */, etc
// 	}
// }
func ParseQuery(query string) (Query, error) {
	var j interface{}
	if err := json.Unmarshal([]byte(query), &j); err != nil {
		return nil, errors.New("must be valid JSON")
	}
	q, ok := j.(map[string]interface{})
	if !ok {
		return nil, errors.New("must be a JSON object")
	}
	return validateQuery(q, "")
}


// validateQuery recursively validates and cast a query
func validateQuery(q map[string]interface{}, parentKey string) (Query, error) {
	queries := Query{}
	//Also check for collection at first level and throw error if not present..
	if parentKey == ""{
		//Now check if the query has collection key defined or not.
		if q["$collection"] != nil{
			var (
				c string
				ok bool
			)
			if c, ok = q["$collection"].(string); !ok{
				return nil, errors.New("$collection must be of string type")
			}

			queries = append(queries, COLLECTION{Value: c})
			//Now remove the $collection key from the query..
			delete(q, "$collection")
		}else{
			return nil, errors.New("$collection type not found at parent level.")
		}
	}


	for key, exp := range q {
		switch key {
		case "$text":
			if parentKey != ""{
				return nil, errors.New("$text can't be at second level.")
			}
			if subexp, ok := exp.(map[string]interface{}); ok {
				//Now check for search type..
				if subexp["$search"] == nil{
					return nil, errors.New("$search not found during processing $text query")
				}
				var searchValue string
				if searchValue, ok = subexp["$search"].(string); ok {
					t := Text{
						Search: searchValue,
					}

					if subexp["$language"] != nil{
						var lang string
						if lang, ok = subexp["$language"].(string); ok{
							t.Language = lang
						}else{
							return nil, errors.New("$language must be of string type")
						}
					}

					if subexp["$caseSensitive"] != nil{
						var cs bool
						if cs, ok = subexp["$caseSensitive"].(bool); ok{
							t.CaseSensitive = cs
						}else{
							return nil, errors.New("$caseSensitive must be of boolean type")
						}
					}

					if subexp["$diacriticSensitive"] != nil{
						var ds bool
						if ds, ok = subexp["$diacriticSensitive"].(bool); ok{
							t.DiacriticSensitive = ds
						}else{
							return nil, errors.New("$diacriticSensitive must be of boolean type")
						}
					}

					queries = append(queries, t)

				}else{
					return nil, errors.New("$search query value must be of string type")
				}

			}else{
				return nil, errors.New("$text invalid format found.")
			}

		case "$regex":
			if parentKey == "" {
				return nil, errors.New("$regex can't be at first level")
			}
			if regex, ok := exp.(string); ok {
				v, err := regexp.Compile(regex)
				if err != nil {
					return nil, fmt.Errorf("$regex: invalid regex: %v", err)
				}
				queries = append(queries, Regex{Field: parentKey, Value: v})
			}
		case "$mod":
			if parentKey == ""{
				return nil, errors.New("$mod can't be at first level")
			}

			if v, ok := exp.([2]float64); ok {
				queries = append(queries, Mod{
					Field:parentKey,
					Divisor: v[0],
					Remainder: v[1],
				})
			}else{
				return nil, errors.New("Invalid type found for $mod operation. Divisor and Remainder required")
			}

		//Element query operations cases..
		case "$exists":
			if parentKey == "" {
				return nil, errors.New("$exists can't be at first level")
			}
			positive, ok := exp.(bool)
			if !ok {
				return nil, errors.New("$exists can only get Boolean as value")
			}
			queries = append(queries, Exist{Field: parentKey, Value: positive})

		//Logical Query Operations..
		case "$or", "$and":
			op := key
			var subQueries []interface{}
			var ok bool
			if subQueries, ok = exp.([]interface{}); !ok {
				return nil, fmt.Errorf("value for %s must be an array of dicts", op)
			}
			if len(subQueries) < 2 {
				return nil, fmt.Errorf("%s must contain at least two elements", op)
			}
			//TODO: editing as wrong use here.. Cast map to Query object
			subQList := []Query{}
			for _, subQuery := range subQueries {
				sq, ok := subQuery.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("value for %s must be an array of dicts", op)
				}
				query, err := validateQuery(sq, "")
				if err != nil {
					return nil, err
				}
				//TODO: Error chances..here it must be an array of queries an not and expression..
				subQList = append(subQList, query)
			}
			switch op {
			case "$or":
				queries = append(queries, Or(subQList))
			case "$and":
				queries = append(queries, And(subQList))
			}
		case "$not":
			//TODO: Test chances of error..
			//Example db.inventory.find( { price: { $not: { $gt: 1.99 } } } )
			if parentKey == ""{
				return nil, errors.New("$not can't be a first level")
			}
			if subexp, ok := exp.(map[string]interface{}); ok {
				subqueries, err := validateQuery(subexp, "$not")
				if err != nil{
					return nil, err
				}
				if len(subqueries) != 0{
					queries = append(queries, Not{
						Field: key,
						//Add first element of the queries only..
						Value: subqueries[0],
					})
				}

			}else{
				return nil, errors.New("$not value must be a dict type only")
			}

		case "$nor":
			//TODO: Test chances of error..
			//Example db.inventory.find( { price: { $not: { $gt: 1.99 } } } )
			if parentKey != ""{
				return nil, errors.New("$nor can't be a second level")
			}
			var subQueries []interface{}
			var ok bool

			if subQueries, ok = exp.([]interface{}); !ok {
				return nil, errors.New("$nor value must be an array of dict type only")
			}

			subQList := []Query{}
			for _, subQuery := range subQueries {
				sq, ok := subQuery.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("value for $nor must be an array of dicts")
				}
				query, err := validateQuery(sq, "")
				if err != nil {
					return nil, err
				}
				//TODO: Error chances..here it must be an array of queries an not and expression..
				subQList = append(subQList, query)
			}

			queries = append(queries, subQList)

		case "$ne":
			op := key
			if parentKey == "" {
				return nil, fmt.Errorf("%s can't be at first level", op)
			}
			//TODO: Validate the field present in the value..
			/*if field := validator.GetField(parentKey); field != nil {
				if field.Validator != nil {
					if _, err := field.Validator.Validate(exp); err != nil {
						return nil, fmt.Errorf("invalid query expression for field `%s': %s", parentKey, err)
					}
				}
			}*/
			queries = append(queries, NotEqual{Field: parentKey, Value: exp})
		case "$gt", "$gte", "$lt", "$lte":
			op := key
			if parentKey == "" {
				return nil, fmt.Errorf("%s can't be at first level", op)
			}
			//TODO: Validate the fields and type of data..as date or number else throw error..
			/*n, ok := isNumber(exp)
			if !ok {
				return nil, fmt.Errorf("%s: value for %s must be a number", parentKey, op)
			}
			if field := validator.GetField(parentKey); field != nil {
				if field.Validator != nil {
					switch field.Validator.(type) {
					case *Integer, *Float, Integer, Float:
						if _, err := field.Validator.Validate(exp); err != nil {
							return nil, fmt.Errorf("invalid query expression for field `%s': %s", parentKey, err)
						}
					default:
						return nil, fmt.Errorf("%s: cannot apply %s operation on a non numerical field", parentKey, op)
					}
				}
			}*/
			switch op {
				case "$gt":
					queries = append(queries, GreaterThan{Field: parentKey, Value: exp})
				case "$gte":
					//queries = append(queries, GreaterOrEqual{Field: parentKey, Value: n})
					queries = append(queries, GreaterOrEqual{Field: parentKey, Value: exp})
				case "$lt":
					//queries = append(queries, LowerThan{Field: parentKey, Value: n})
					queries = append(queries, LowerThan{Field: parentKey, Value: exp})
				case "$lte":
					//queries = append(queries, LowerOrEqual{Field: parentKey, Value: n})
					queries = append(queries, LowerOrEqual{Field: parentKey, Value: exp})
			}
		case "$in", "$nin":
			op := key
			if parentKey == "" {
				return nil, fmt.Errorf("%s can't be at first level", op)
			}
			if _, ok := exp.(map[string]interface{}); ok {
				return nil, fmt.Errorf("%s: value for %s can't be a dict", parentKey, op)
			}
			values := []Value{}
			//TODO: Validation removed add data...here..
			//if field := validator.GetField(parentKey); field != nil {
				vals, ok := exp.([]interface{})
				if !ok {
					vals = []interface{}{exp}
				}
			//TODO: Validation removed here..
				/*if field.Validator != nil {
					for _, v := range vals {
						if _, err := field.Validator.Validate(v); err != nil {
							return nil, fmt.Errorf("invalid query expression (%s) for field `%s': %s", v, parentKey, err)
						}
					}
				}*/

				for _, v := range vals {
					values = append(values, v)
				}
			//}
			switch op {
				case "$in":
					queries = append(queries, In{Field: parentKey, Values: values})
				case "$nin":
					queries = append(queries, NotIn{Field: parentKey, Values: values})
			}

		default:
			//TODO:" Validation removed..
			// Field query
			/*field := validator.GetField(key)
			if field == nil {
				return nil, fmt.Errorf("unknown query field: %s", key)
			}
			if !field.Filterable {
				return nil, fmt.Errorf("field is not filterable: %s", key)
			}*/
			if parentKey != "" {
				return nil, fmt.Errorf("%s: invalid expression", parentKey)
			}
			if subQuery, ok := exp.(map[string]interface{}); ok {
				sq, err := validateQuery(subQuery, key)
				if err != nil {
					return nil, err
				}
				queries = append(queries, sq...)
			} else {
				//TODO: Validation removed..
				/*// Exact match
				if field.Validator != nil {
					if _, err := field.Validator.Validate(exp); err != nil {
						return nil, fmt.Errorf("invalid query expression for field `%s': %s", key, err)
					}
				}*/
				queries = append(queries, Equal{Field: key, Value: exp})
			}
		}
	}
	return queries, nil
}



/*
// Match implements Expression interface
func (e Query) Match(payload map[string]interface{}) bool {
	// Run each sub queries like a root query, stop/pass on first match
	for _, subQuery := range e {
		if !subQuery.Match(payload) {
			return false
		}
	}
	return true
}

// Match implements Expression interface
func (e And) Match(payload map[string]interface{}) bool {
	// Run each sub queries like a root query, stop/pass on first match
	for _, subQuery := range e {
		if !subQuery.Match(payload) {
			return false
		}
	}
	return true
}

// Match implements Expression interface
func (e Or) Match(payload map[string]interface{}) bool {
	// Run each sub queries like a root query, stop/pass on first match
	for _, subQuery := range e {
		if subQuery.Match(payload) {
			return true
		}
	}
	return false
}

// Match implements Expression interface
func (e In) Match(payload map[string]interface{}) bool {
	value := getField(payload, e.Field)
	for _, v := range e.Values {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Match implements Expression interface
func (e NotIn) Match(payload map[string]interface{}) bool {
	value := getField(payload, e.Field)
	for _, v := range e.Values {
		if reflect.DeepEqual(v, value) {
			return false
		}
	}
	return true
}

// Match implements Expression interface
func (e Equal) Match(payload map[string]interface{}) bool {
	return reflect.DeepEqual(getField(payload, e.Field), e.Value)
}

// Match implements Expression interface
func (e NotEqual) Match(payload map[string]interface{}) bool {
	return !reflect.DeepEqual(getField(payload, e.Field), e.Value)
}

// Match implements Expression interface
func (e Exist) Match(payload map[string]interface{}) bool {
	_, found := getFieldExist(payload, e.Field)
	return found
}

// Match implements Expression interface
func (e NotExist) Match(payload map[string]interface{}) bool {
	_, found := getFieldExist(payload, e.Field)
	return !found
}

// Match implements Expression interface
func (e GreaterThan) Match(payload map[string]interface{}) bool {
	n, ok := isNumber(getField(payload, e.Field))
	return ok && (n > e.Value)
}

// Match implements Expression interface
func (e GreaterOrEqual) Match(payload map[string]interface{}) bool {
	n, ok := isNumber(getField(payload, e.Field))
	return ok && (n >= e.Value)
}

// Match implements Expression interface
func (e LowerThan) Match(payload map[string]interface{}) bool {
	n, ok := isNumber(getField(payload, e.Field))
	return ok && (n < e.Value)
}

// Match implements Expression interface
func (e LowerOrEqual) Match(payload map[string]interface{}) bool {
	n, ok := isNumber(getField(payload, e.Field))
	return ok && (n <= e.Value)
}

// Match implements Expression interface
func (e Regex) Match(payload map[string]interface{}) bool {
	return e.Value.MatchString(payload[e.Field].(string))
}
*/

