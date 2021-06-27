package main

import "context"

type Context struct {
	context.Context

	valueStore map[string]interface{}
}

type Script struct {
	Before func(c *Context) error
	Action func(c *Context) error
	After  func(c *Context) error
}

func (s *Script) Run() error {
	c := &Context{valueStore: map[string]interface{}{}}
	if s.Before != nil {
		if err := s.Before(c); err != nil {
			return err
		}
	}
	if err := s.Action(c); err != nil {
		return err
	}

	if s.After != nil {
		if err := s.After(c); err != nil {
			return err
		}
	}

	return nil
}
