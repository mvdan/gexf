// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// Package gexf implements the GEXF file format
package gexf

import (
	"encoding/xml"
	"time"
)

const (
	Version = "1.2"
	Space   = "http://www.gexf.net/1.2draft"
	Local   = "gexf"

	dateFormat = "2006-01-02"
)

type Doc struct {
	XMLName xml.Name
	Version string `xml:"version,attr"`
	Meta    Meta   `xml:"meta"`

	Graph Graph `xml:"graph"`
}

func New() *Doc {
	return &Doc{
		XMLName: xml.Name{
			Space: Space,
			Local: Local,
		},
		Version: Version,
	}
}

type Date struct {
	time.Time
}

func (d *Date) MarshalText() ([]byte, error) {
	return []byte(d.UTC().Format(dateFormat)), nil
}

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(dateFormat, string(text))
	if err != nil {
		return err
	}
	*d = Date{t}
	return nil
}

type Meta struct {
	LastMod  Date   `xml:"lastmodifieddate,attr"`
	Creator  string `xml:"creator,omitempty"`
	Keywords string `xml:"keywords,omitempty"`
	Desc     string `xml:"description,omitempty"`
}

type Graph struct {
	Mode    GraphMode `xml:"mode,attr,omitempty"`
	IDType  IDType    `xml:"idtype,attr,omitempty"`
	DefEdge EdgeType  `xml:"defaultedgetype,attr,omitempty"`

	Attrs *Attributes `xml:"attributes,omitempty"`
	Nodes *[]Node     `xml:"nodes>node,omitempty"`
	Edges *[]Edge     `xml:"edges>edge,omitempty"`
}

type EdgeType int

const (
	Directed EdgeType = iota
	Undirected
	Mutual
)

func (t EdgeType) String() string {
	switch t {
	case Directed:
		return "directed"
	case Undirected:
		return "undirected"
	case Mutual:
		return "mutual"
	}
	return ""
}

func (t EdgeType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

type IDType int

const (
	String IDType = iota
	Integer
)

func (t IDType) String() string {
	switch t {
	case String:
		return "string"
	case Integer:
		return "integer"
	}
	return ""
}

func (t IDType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

type GraphMode int

const (
	Static GraphMode = iota
	Dynamic
)

func (t GraphMode) String() string {
	switch t {
	case Static:
		return "static"
	case Dynamic:
		return "dynamic"
	}
	return ""
}

func (t GraphMode) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

type ClassType int

const (
	ClassNode ClassType = iota
	ClassEdge
)

func (t ClassType) String() string {
	switch t {
	case ClassNode:
		return "node"
	case ClassEdge:
		return "edge"
	}
	return ""
}

func (t ClassType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

type Attributes struct {
	Class ClassType `xml:"class,attr"`
	Attrs []Attr    `xml:"attribute"`
}

type Attr struct {
	ID      string `xml:"id,attr"`
	Title   string `xml:"title,attr"`
	Type    string `xml:"type,attr"`
	Default string `xml:"default,omitempty"`
}

type Node struct {
	ID    string     `xml:"id,attr"`
	Label string     `xml:"label,attr,omitempty"`
	Attrs *[]AttrVal `xml:"attvalues>attvalue,omiempty"`
	Size  *Size      `xml:"http://www.gexf.net/1.2draft/viz size,omitempty"`
	Pos   *Pos       `xml:"http://www.gexf.net/1.2draft/viz position,omitempty"`
	Color *Color     `xml:"http://www.gexf.net/1.2draft/viz color,omitempty"`
}

type Size struct {
	Value float64 `xml:"value,attr"`
}

type Pos struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
	Z float64 `xml:"z,attr"`
}

type Color struct {
	R uint8 `xml:"r,attr"`
	G uint8 `xml:"g,attr"`
	B uint8 `xml:"b,attr"`
}

type AttrVal struct {
	For   string `xml:"for,attr"`
	Value string `xml:"value,attr"`
}

type Edge struct {
	ID     string     `xml:"id,attr"`
	Label  string     `xml:"label,attr,omitempty"`
	Type   EdgeType   `xml:"type,attr,omitempty"`
	Source string     `xml:"source,attr"`
	Target string     `xml:"target,attr"`
	Weight float64    `xml:"weight,attr,omitempty"`
	Attrs  *[]AttrVal `xml:"attvalues>attvalue,omiempty"`
}
