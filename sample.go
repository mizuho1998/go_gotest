package tddbc

import (
	"errors"
	"fmt"
)

type ClosedInterval struct {
	lower int
	upper int
}

func NewClosedInterval(lower int, upper int) (*ClosedInterval, error) {
	if lower > upper {
		return nil, errors.New("lower must be smaller than upper")
	}
	ci := new(ClosedInterval)
	ci.lower = lower
	ci.upper = upper
	return ci, nil
}

func (c *ClosedInterval) String() string {
	return fmt.Sprintf("[%d, %d]", c.lower, c.upper)
}

func (c *ClosedInterval) IsIncludeDot(num int) bool {
	return c.lower <= num && num <= c.upper
}

func (c1 *ClosedInterval) IsSame(c2 ClosedInterval) bool {
	return c1.lower == c2.lower && c1.upper == c2.upper
}

func (c1 *ClosedInterval) IsInclude(c2 ClosedInterval) bool {
	return c2.lower <= c1.lower && c1.upper <= c2.upper
}
