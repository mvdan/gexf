// Copyright (c) 2015, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package gexf

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

func doTest(t *testing.T, enc string, doc *Doc) {
	got := New()
	if err := xml.Unmarshal([]byte(enc), got); err != nil {
		t.Errorf("Could not unmarshal document: %v", err)
	}
	if !reflect.DeepEqual(got, doc) {
		t.Errorf("Documents are different.\nExp: %#v\nGot: %#v", doc, got)
	}
	b, err := xml.MarshalIndent(doc, "", "\t")
	if err != nil {
		t.Errorf("Could not marshal document: %v", err)
	}
	wantEnc := string(b)
	if wantEnc != enc {
		t.Errorf("Encoded documents are different.\nExp:\n%s\nGot:\n%s", enc, wantEnc)
	}
}

func TestHelloWorld(t *testing.T) {
	enc := `<gexf xmlns="http://www.gexf.net/1.2draft" version="1.2">
	<meta lastmodifieddate="2009-03-20">
		<creator>Gephi.org</creator>
		<description>A hello world! file</description>
	</meta>
	<graph>
		<nodes>
			<node id="0" label="Hello"></node>
			<node id="1" label="World"></node>
		</nodes>
		<edges>
			<edge id="0" source="0" target="1"></edge>
		</edges>
	</graph>
</gexf>`
	doc := New()
	doc.Meta = Meta{
		LastModified: Date{time.Date(2009, time.March, 20, 0, 0, 0, 0, time.UTC)},
		Creator:      "Gephi.org",
		Description:  "A hello world! file",
	}
	doc.Graph = Graph{
		Mode:        Static,
		DefEdgeType: Directed,
		Nodes: &Nodes{
			Node: []Node{
				{
					ID:    "0",
					Label: "Hello",
				},
				{
					ID:    "1",
					Label: "World",
				},
			},
		},
		Edges: &Edges{
			Edge: []Edge{
				{
					ID:     "0",
					Source: "0",
					Target: "1",
				},
			},
		},
	}
	doTest(t, enc, doc)
}

func TestAttributes(t *testing.T) {
	enc := `<gexf xmlns="http://www.gexf.net/1.2draft" version="1.2">
	<meta lastmodifieddate="2009-03-20">
		<creator>Gephi.org</creator>
		<description>A Web network</description>
	</meta>
	<graph>
		<attributes class="node">
			<attribute id="0" title="url" type="string"></attribute>
			<attribute id="1" title="indegree" type="float"></attribute>
			<attribute id="2" title="frog" type="boolean">
				<default>true</default>
			</attribute>
		</attributes>
		<nodes>
			<node id="0" label="Gephi">
				<attvalues>
					<attvalue for="0" value="http://gephi.org"></attvalue>
					<attvalue for="2" value="false"></attvalue>
				</attvalues>
			</node>
			<node id="1" label="Webatlas">
				<attvalues>
					<attvalue for="1" value="2"></attvalue>
					<attvalue for="2" value="true"></attvalue>
				</attvalues>
			</node>
		</nodes>
	</graph>
</gexf>`
	doc := New()
	doc.Meta = Meta{
		LastModified: Date{time.Date(2009, time.March, 20, 0, 0, 0, 0, time.UTC)},
		Creator:      "Gephi.org",
		Description:  "A Web network",
	}
	doc.Graph = Graph{
		DefEdgeType: Directed,
		Attrs: &Attributes{
			Class: ClassNode,
			Attribute: []Attribute{
				{
					ID:    "0",
					Title: "url",
					Type:  "string",
				},
				{
					ID:    "1",
					Title: "indegree",
					Type:  "float",
				},
				{
					ID:      "2",
					Title:   "frog",
					Type:    "boolean",
					Default: "true",
				},
			},
		},
		Nodes: &Nodes{
			Node: []Node{
				{
					ID:    "0",
					Label: "Gephi",
					AttValues: &AttValues{
						AttValue: []AttValue{
							{
								For:   "0",
								Value: "http://gephi.org",
							},
							{
								For:   "2",
								Value: "false",
							},
						},
					},
				},
				{
					ID:    "1",
					Label: "Webatlas",
					AttValues: &AttValues{
						AttValue: []AttValue{
							{
								For:   "1",
								Value: "2",
							},
							{
								For:   "2",
								Value: "true",
							},
						},
					},
				},
			},
		},
	}
	doTest(t, enc, doc)
}

func TestViz(t *testing.T) {
	enc := `<gexf xmlns="http://www.gexf.net/1.2draft" version="1.2">
	<meta lastmodifieddate="2009-03-20">
		<creator>Gephi.org</creator>
		<description>A hello world! file</description>
	</meta>
	<graph>
		<nodes>
			<node id="0" label="Hello">
				<size xmlns="http://www.gexf.net/1.2draft/viz" value="20.5"></size>
			</node>
		</nodes>
	</graph>
</gexf>`
	doc := New()
	doc.Meta = Meta{
		LastModified: Date{time.Date(2009, time.March, 20, 0, 0, 0, 0, time.UTC)},
		Creator:      "Gephi.org",
		Description:  "A hello world! file",
	}
	doc.Graph = Graph{
		Mode:        Static,
		DefEdgeType: Directed,
		Nodes: &Nodes{
			Node: []Node{
				{
					ID:    "0",
					Label: "Hello",
					Size:  &Size{Value: 20.5},
				},
			},
		},
	}
	doTest(t, enc, doc)
}
