package queue

import "sort"

type Polynomial indeterminate

type indeterminate struct {
	coefficient int
	power       uint
	next        *indeterminate
}

func NewPolynomial(vs []indeterminate) *Polynomial {
	for i := 0; i < len(vs)-1; i++ {
		vs[i].next = &vs[i+1]
	}
	v := Polynomial{}
	if len(vs) > 0 {
		v = Polynomial(vs[0])
	}
	return &v
}

func (p *Polynomial) Add(another *Polynomial) *Polynomial {
	m := map[uint]int{}

	ind := (*indeterminate)(p)
	for ind != nil {
		m[p.power] = p.coefficient
		ind = ind.next
	}

	ind = (*indeterminate)(another)
	for ind != nil {
		if _, ok := m[ind.power]; ok {
			m[ind.power] += ind.coefficient
		}
		ind = ind.next
	}

	vs := make([]indeterminate, 0)
	for power, coefficient := range m {
		vs = append(
			vs, indeterminate{
				coefficient: coefficient,
				power:       power,
			},
		)
	}

	sort.Slice(
		vs, func(i, j int) bool {
			return vs[i].power > vs[j].power
		},
	)

	return NewPolynomial(vs)
}

func (p *Polynomial) MultipliedBy(another *Polynomial) *Polynomial {
	if p == nil || another == nil {
		return nil
	}
	m := map[uint]int{}
	pi := (*indeterminate)(p)
	for pi != nil {
		ai := (*indeterminate)(another)
		for ai != nil {
			coefficient := pi.coefficient * ai.coefficient
			power := pi.power + ai.power
			if _, ok := m[power]; ok {
				m[power] += coefficient
			} else {
				m[power] = coefficient
			}
			ai = ai.next
		}
		pi = pi.next
	}
	vs := make([]indeterminate, 0)
	for power, coefficient := range m {
		vs = append(
			vs, indeterminate{
				coefficient: coefficient,
				power:       power,
			},
		)
	}
	sort.Slice(
		vs, func(i, j int) bool {
			return vs[i].power > vs[j].power
		},
	)
	return NewPolynomial(vs)
}
