package main

import (
	"fmt"
	"io"
)

// Reposter client
var Reposter Publisher

func init() {
	Reposter = &publishClient{
		repostTemplate,
	}
}

// Publisher interface
type Publisher interface {
	Publish(imageFile io.ReadCloser, username, currentCaption string) error
}

type publishClient struct {
	template string
}

func (p *publishClient) describe(username, currentCaption string) string {
	return fmt.Sprintf(p.template, username, currentCaption)
}

func (p *publishClient) Publish(imageFile io.ReadCloser, username, currentCaption string) error {
	caption := p.describe(username, currentCaption)
	return GetInstagram().Upload(imageFile, caption)
}

var repostTemplate = `Checkout the new post by @%s
---
%s
#instrubot`
