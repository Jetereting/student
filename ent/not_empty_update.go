// Code generated by ent, DO NOT EDIT.

package ent

// set field if value is not empty. e.g. string does not equal to ""
func (s *StudentUpdate) SetNotEmptySort(value uint32) *StudentUpdate {
	if value != 0 {
		return s.SetSort(value)
	}
	return s
}

// set field if value is not empty. e.g. string does not equal to ""
func (s *StudentUpdateOne) SetNotEmptySort(value uint32) *StudentUpdateOne {
	if value != 0 {
		return s.SetSort(value)
	}
	return s
}

// set field if value is not empty. e.g. string does not equal to ""
func (s *StudentUpdate) SetNotEmptyStatus(value uint8) *StudentUpdate {
	if value != 0 {
		return s.SetStatus(value)
	}
	return s
}

// set field if value is not empty. e.g. string does not equal to ""
func (s *StudentUpdateOne) SetNotEmptyStatus(value uint8) *StudentUpdateOne {
	if value != 0 {
		return s.SetStatus(value)
	}
	return s
}
