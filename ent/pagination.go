// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"school/ent/class"
	"school/ent/school"
	"school/ent/student"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type ClassPager struct {
	Order  class.OrderOption
	Filter func(*ClassQuery) (*ClassQuery, error)
}

// ClassPaginateOption enables pagination customization.
type ClassPaginateOption func(*ClassPager)

// DefaultClassOrder is the default ordering of Class.
var DefaultClassOrder = Desc(class.FieldID)

func newClassPager(opts []ClassPaginateOption) (*ClassPager, error) {
	pager := &ClassPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultClassOrder
	}
	return pager, nil
}

func (p *ClassPager) ApplyFilter(query *ClassQuery) (*ClassQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// ClassPageList is Class PageList result.
type ClassPageList struct {
	List        []*Class     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (c *ClassQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...ClassPaginateOption,
) (*ClassPageList, error) {

	pager, err := newClassPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.ApplyFilter(c); err != nil {
		return nil, err
	}

	ret := &ClassPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := c.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		c = c.Order(pager.Order)
	} else {
		c = c.Order(DefaultClassOrder)
	}

	c = c.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type SchoolPager struct {
	Order  school.OrderOption
	Filter func(*SchoolQuery) (*SchoolQuery, error)
}

// SchoolPaginateOption enables pagination customization.
type SchoolPaginateOption func(*SchoolPager)

// DefaultSchoolOrder is the default ordering of School.
var DefaultSchoolOrder = Desc(school.FieldID)

func newSchoolPager(opts []SchoolPaginateOption) (*SchoolPager, error) {
	pager := &SchoolPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultSchoolOrder
	}
	return pager, nil
}

func (p *SchoolPager) ApplyFilter(query *SchoolQuery) (*SchoolQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// SchoolPageList is School PageList result.
type SchoolPageList struct {
	List        []*School    `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (s *SchoolQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...SchoolPaginateOption,
) (*SchoolPageList, error) {

	pager, err := newSchoolPager(opts)
	if err != nil {
		return nil, err
	}

	if s, err = pager.ApplyFilter(s); err != nil {
		return nil, err
	}

	ret := &SchoolPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := s.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		s = s.Order(pager.Order)
	} else {
		s = s.Order(DefaultSchoolOrder)
	}

	s = s.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type StudentPager struct {
	Order  student.OrderOption
	Filter func(*StudentQuery) (*StudentQuery, error)
}

// StudentPaginateOption enables pagination customization.
type StudentPaginateOption func(*StudentPager)

// DefaultStudentOrder is the default ordering of Student.
var DefaultStudentOrder = Desc(student.FieldID)

func newStudentPager(opts []StudentPaginateOption) (*StudentPager, error) {
	pager := &StudentPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultStudentOrder
	}
	return pager, nil
}

func (p *StudentPager) ApplyFilter(query *StudentQuery) (*StudentQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// StudentPageList is Student PageList result.
type StudentPageList struct {
	List        []*Student   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (s *StudentQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...StudentPaginateOption,
) (*StudentPageList, error) {

	pager, err := newStudentPager(opts)
	if err != nil {
		return nil, err
	}

	if s, err = pager.ApplyFilter(s); err != nil {
		return nil, err
	}

	ret := &StudentPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := s.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		s = s.Order(pager.Order)
	} else {
		s = s.Order(DefaultStudentOrder)
	}

	s = s.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
