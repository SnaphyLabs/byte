package resource


import (
	"github.com/SnaphyLabs/SnaphyByte/schema"
)

// Lookup holds filter and sort used to select items in a resource collection
type Lookup struct {
	// The client supplied filter. Filter is a MongoDB inspired query with a more limited
	// set of capabilities. See https://github.com/rs/rest-layer#filtering
	// for more info.
	filter schema.Query
	// The client supplied soft. Sort is a list of resource fields or sub-fields separated
	// by comas (,). To invert the sort, a minus (-) can be prefixed.
	// See https://github.com/rs/rest-layer#sorting for more info.
	sort []string
	// The client supplied selector. Selector is a way for the client to reformat the
	// resource representation at runtime by defining which fields should be included
	// in the document. The REST Layer selector language allows field aliasing, field
	// transformation with parameters and sub-item/collection embedding.
	//Remove fields here..
	selector []string
}

//Get list of fields..
func (l *Lookup) Field() []string  {
	return l.selector
}

//Add fields selector to selector
func (l *Lookup)AddField(selector []string) {
	l.selector = selector
}




// Sort is a list of resource fields or sub-fields separated
// by comas (,). To invert the sort, a minus (-) can be prefixed.
//
// See https://github.com/rs/rest-layer#sorting for more info.
func (l *Lookup) Sort() []string {
	return l.sort
}



// Filter is a MongoDB inspired query with a more limited set of capabilities.
//
// See https://github.com/rs/rest-layer#filtering for more info.
func (l *Lookup) Filter() *schema.Query {
	return &l.filter
}



// SetSorts set the sort fields with a pre-parsed list of fields to sort on.
// This method doesn't validate sort fields.
func (l *Lookup) SetSorts(sorts []string) {
	l.sort = sorts
}


/*
// SetSort parses and validate a sort parameter and set it as lookup's Sort
func (l *Lookup) SetSort(sort string, v schema.Validator) error {
	sorts := []string{}
	for _, f := range strings.Split(sort, ",") {
		f = strings.Trim(f, " ")
		if f == "" {
			return errors.New("empty soft field")
		}
		// If the field start with - (to indicate descended sort), shift it before
		// validator lookup
		i := 0
		if f[0] == '-' {
			i = 1
		}
		// Make sure the field exists
		field := v.GetField(f[i:])
		if field == nil {
			return fmt.Errorf("invalid sort field: %s", f[i:])
		}
		if !field.Sortable {
			return fmt.Errorf("field is not sortable: %s", f[i:])
		}
		sorts = append(sorts, f)
	}
	l.sort = sorts
	return nil
}*/
/*

// AddFilter parses and validate a filter parameter and add it to lookup's filter
//
// The filter query is validated against the provided validator to ensure all queried
// fields exists and are of the right type.
func (l *Lookup) AddFilter(filter string, v schema.Validator) error {
	f, err := schema.ParseQuery(filter, v)
	if err != nil {
		return err
	}
	l.AddQuery(f)
	return nil
}
*/



// AddQuery add an existing schema.Query to the lookup's filters
func (l *Lookup) AddQuery(query schema.Query) {
	if l.filter == nil {
		l.filter = query
		return
	}
	for _, exp := range query {
		l.filter = append(l.filter, exp)
	}
}



/*
// SetSelector parses a selector expression, validates it and assign it to the current Lookup.
func (l *Lookup) SetSelector(s string, v schema.Validator) error {
	pos := 0
	selector, err := parseSelectorExpression([]byte(s), &pos, len(s), false)
	if err != nil {
		return err
	}
	if err = validateSelector(selector, v); err != nil {
		return err
	}
	l.selector = selector
	return nil
}

// ReferenceResolver is a function resolving a reference to another field
type ReferenceResolver func(path string) (*Resource, error)

// ApplySelector applies fields filtering / rename to the payload fields
func (l *Lookup) ApplySelector(ctx context.Context, v schema.Validator, p map[string]interface{}, resolver ReferenceResolver) (map[string]interface{}, error) {
	payload, err := applySelector(ctx, l.selector, v, p, resolver)
	if err == nil {
		// The resulting payload may contain some asyncSelector, we must execute them
		// concurrently until there's no more
		err = resolveAsyncSelectors(ctx, payload)
	}
	return payload, err
}

*/
