package form

import (
	"errors"
	"net/url"
	"time"
)

// ErrNoToken indicates that provided form data has no security token.
var ErrNoToken = errors.New("No token provided")

// DefaultFormHandler is a form handler initialized with an in-memory
// cache and a 24 hour expiration on form data.
var DefaultFormHandler = NewFormHandler(NewCache(), time.Hour*24)

// FormHandler manages the cache-and-load lifecycle of forms.
//
// FormHandler enforces security constraints on forms, and will modify
// forms in place.
type FormHandler struct {
	cache Cache
	// The duration a form will be kept before it expires.
	Expiration time.Duration
}

// NewFormHandler creates a new FormHandler.
func NewFormHandler(c Cache, expiration time.Duration) *FormHandler {
	return &FormHandler{
		cache:      c,
		Expiration: expiration,
	}
}

// Prepare modifies the form for caching and security, then inserts it into cache.
//
// This will add a security field to the end of the form's Fields list. The
// generated ID will be returned. And the form will be placed into the cache.
//
// This form can later be retrieved using the returned ID.
func (f *FormHandler) Prepare(form *Form) (string, error) {
	sf := SecurityField()
	form.Fields = append(form.Fields, sf)
	f.cache.Set(sf.Value, form, time.Now().Add(f.Expiration))

	return sf.Value, nil
}

// Retrieve uses a request's key/value pairs to populate a cached form.
//
// It then decodes the submission data into the relevant cached form,
// and returns the form. The Value fields on each form element will be
// set according to the data. If the form data was not set for a particular
// field, that field will be left alone (which means that if it had a default
// value, that will remain in effect).
//
// Finally, Retrieve will remove the form from the cache, since a form
// cannot be re-used.
//
// The implementing function must pass in the appropriate set of values.
// The "net/http" library makes Get, Post, Put, and Patch variables all
// available as *url.Values.
func (f *FormHandler) Retrieve(data *url.Values) (*Form, error) {
	id := data.Get(SecureTokenName)
	if id == "" {
		return nil, ErrNoToken
	}

	fm, err := f.Get(id)
	if err != nil {
		return nil, err
	}

	if err := Reconcile(fm, data); err != nil {
		// Form might still be useful in this case.
		return fm, err
	}

	f.Remove(id)
	return fm, nil
}

func (f *FormHandler) Get(id string) (*Form, error) {
	return f.cache.Get(id)
}

func (f *FormHandler) Remove(id string) error {
	return f.cache.Remove(id)
}

// Reconcile modifies a form in place, merging the data into the form's Value fields.
//
// Validation is not handled by the reconciler.
//
// Normally, reconciliation will happen via the FormHandler's Retrieve method.
func Reconcile(fm *Form, data *url.Values) error {
	return reconcileFields(fm.Fields, data, fm)
}

func reconcileFields(fields []Field, data *url.Values, fm *Form) error {
	for _, field := range fields {
		// Because of the limitations on the type switch, we have to
		// enumerate each type on its own line so that f is set correctly.
		switch f := field.(type) {
		case *Div:
			reconcileFields(f.Fields, data, fm)
		case *FieldSet:
			reconcileFields(f.Fields, data, fm)
		case *Radio:
			if val := data.Get(f.Name); val == f.Value {
				//fmt.Printf("name=%q value=%q is checked\n", f.Name, f.Value)
				f.Checked = true
			}
		case *Checkbox:
			d := *data
			vals, ok := d[f.Name]
			if !ok {
				continue
			}

			for _, val := range vals {
				if f.Value == val {
					//fmt.Printf("name=%q value=%q is checked\n", f.Name, f.Value)
					f.Checked = true
					break
				}
			}

		case *Keygen:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *TextArea:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Text:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Password:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Tel:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *URL:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Email:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Date:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Time:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Number:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Range:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Color:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Hidden:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Button:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *ButtonInput:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Submit:
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Image:
			// TODO: User Agent should also pass .x and .y fields in
			// the data.Values, where those indicate where on the image
			// it was clicked.
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}
		case *Input:
			// Unlikely but possible case.
			if val := data.Get(f.Name); val != "" {
				f.Value = val
			}

		default:
			// All other field types don't have values.
			//fmt.Printf("Type: %T\n", f)
			continue
		}
	}
	return nil
}

func allFieldsNamed(name string, fm *Form) []Field {
	return recursiveFieldsNamed(name, fm.Fields)
}

func recursiveFieldsNamed(name string, ff []Field) []Field {
	found := []Field{}
	for _, item := range ff {

		if div, ok := item.(Div); ok {
			more := recursiveFieldsNamed(name, div.Fields)
			found = append(found, more...)
		} else if fg, ok := item.(FieldSet); ok {
			more := recursiveFieldsNamed(name, fg.Fields)
			found = append(found, more...)
		} else if ele, ok := item.(Input); ok {
			if ele.Name == name {
				found = append(found, ele)
			}
		}
	}
	return found
}
