// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Rule string

const (
	RuleLimitedTag Rule = "LIMITED_TAG"
	RuleHasOwner   Rule = "HAS_OWNER"
)

var AllRule = []Rule{
	RuleLimitedTag,
	RuleHasOwner,
}

func (e Rule) IsValid() bool {
	switch e {
	case RuleLimitedTag, RuleHasOwner:
		return true
	}
	return false
}

func (e Rule) String() string {
	return string(e)
}

func (e *Rule) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Rule(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Rule", str)
	}
	return nil
}

func (e Rule) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
