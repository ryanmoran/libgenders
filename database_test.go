package libgenders_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ryanmoran/libgenders"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDatabase(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
		path   string
	)

	it.Before(func() {
		file, err := os.CreateTemp("", "genders")
		Expect(err).NotTo(HaveOccurred())
		defer file.Close()

		path = file.Name()
	})

	it.After(func() {
		Expect(os.Remove(path)).To(Succeed())
	})

	context("NewDatabase", func() {
		it("loads the database from a path on disk", func() {
			_, err := libgenders.NewDatabase(path)
			Expect(err).NotTo(HaveOccurred())
		})

		context("failure cases", func() {
			context("when the filepath does not exist", func() {
				it("returns an error", func() {
					_, err := libgenders.NewDatabase("no-such-file")
					Expect(err).To(MatchError(ContainSubstring("no-such-file: no such file or directory")))
				})
			})
		})
	})

	context("GetNodes", func() {
		var (
			testdata = []string{
				"genders.base",
				"genders.base_comma",
				"genders.base_comments_beginning_of_line",
				"genders.base_comments_beginning_of_line_comma",
				"genders.base_comments_beginning_of_line_hostrange",
				"genders.base_comments_end_of_line",
				"genders.base_comments_end_of_line_comma",
				"genders.base_comments_end_of_line_hostrange",
				"genders.base_comments_end_of_line_with_whitespace",
				"genders.base_comments_end_of_line_with_whitespace_comma",
				"genders.base_comments_end_of_line_with_whitespace_hostrange",
				"genders.base_comments_middle_of_line",
				"genders.base_comments_middle_of_line_comma",
				"genders.base_comments_middle_of_line_hostrange",
				"genders.base_hostrange",
				"genders.base_hostrange_single",
				"genders.base_whitespace_after_attrs",
				"genders.base_whitespace_after_attrs_comma",
				"genders.base_whitespace_after_attrs_hostrange",
				"genders.base_whitespace_after_nodes",
				"genders.base_whitespace_after_nodes_comma",
				"genders.base_whitespace_after_nodes_hostrange",
				"genders.base_whitespace_before_and_after_nodes",
				"genders.base_whitespace_before_and_after_nodes_comma",
				"genders.base_whitespace_before_and_after_nodes_hostrange",
				"genders.base_whitespace_before_nodes",
				"genders.base_whitespace_before_nodes_comma",
				"genders.base_whitespace_before_nodes_hostrange",
				"genders.base_whitespace_between_nodes",
				"genders.base_whitespace_between_nodes_and_attrs",
			}
		)

		for _, filename := range testdata {
			var database libgenders.Database
			path := filepath.Join("./testdata", filename)

			it.Before(func() {
				var err error
				database, err = libgenders.NewDatabase(path)
				Expect(err).NotTo(HaveOccurred())
			})

			context(fmt.Sprintf("given %s", path), func() {
				it("returns a list of nodes", func() {
					nodes := database.GetNodes()
					Expect(nodes).To(Equal([]libgenders.Node{
						{
							Name: "node1",
							Attributes: map[string]string{
								"attr1": "",
								"attr2": "val2",
							},
						},
						{
							Name: "node2",
							Attributes: map[string]string{
								"attr1": "",
								"attr2": "val2",
							},
						},
					}))
				})
			})
		}

		context("when the file only contains node names", func() {
			var database libgenders.Database

			it.Before(func() {
				var err error
				database, err = libgenders.NewDatabase("./testdata/genders.nodes_only_many")
				Expect(err).NotTo(HaveOccurred())
			})

			it("returns a list of nodes", func() {
				nodes := database.GetNodes()
				Expect(nodes).To(Equal([]libgenders.Node{
					{Name: "node1"},
					{Name: "node2"},
					{Name: "node3"},
					{Name: "node4"},
					{Name: "node5"},
				}))
			})
		})

		context("nodes share common attributes", func() {
			var (
				testdata = []string{
					"genders.query_1",
					"genders.query_1_comma",
					"genders.query_1_hostrange",
				}
			)

			for _, filename := range testdata {
				var database libgenders.Database
				path := filepath.Join("./testdata", filename)

				it.Before(func() {
					var err error
					database, err = libgenders.NewDatabase(path)
					Expect(err).NotTo(HaveOccurred())
				})

				context(fmt.Sprintf("given %s", path), func() {
					it("returns a list of nodes", func() {
						nodes := database.GetNodes()
						Expect(nodes).To(Equal([]libgenders.Node{
							{
								Name: "node1",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr3": "",
									"attr4": "val4",
									"attr7": "",
									"attr8": "val8",
								},
							},
							{
								Name: "node2",
								Attributes: map[string]string{
									"attr1":  "",
									"attr2":  "val2",
									"attr3":  "",
									"attr4":  "val4",
									"attr9":  "",
									"attr10": "val10",
								},
							},
							{
								Name: "node3",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr3": "",
									"attr4": "val4",
									"attr7": "",
									"attr8": "val8",
								},
							},
							{
								Name: "node4",
								Attributes: map[string]string{
									"attr1":  "",
									"attr2":  "val2",
									"attr3":  "",
									"attr4":  "val4",
									"attr9":  "",
									"attr10": "val10",
								},
							},
							{
								Name: "node5",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr5": "",
									"attr6": "val6",
									"attr7": "",
									"attr8": "val8",
								},
							},
							{
								Name: "node6",
								Attributes: map[string]string{
									"attr1":  "",
									"attr2":  "val2",
									"attr5":  "",
									"attr6":  "val6",
									"attr9":  "",
									"attr10": "val10",
								},
							},
							{
								Name: "node7",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr5": "",
									"attr6": "val6",
									"attr7": "",
									"attr8": "val8",
								},
							},
							{
								Name: "node8",
								Attributes: map[string]string{
									"attr1":  "",
									"attr2":  "val2",
									"attr5":  "",
									"attr6":  "val6",
									"attr9":  "",
									"attr10": "val10",
								},
							},
						}))
					})
				})
			}
		})

		context("node attributes reference the node name", func() {
			var (
				testdata = []string{
					"genders.subst_nodename",
					"genders.subst_nodename_comma",
					"genders.subst_nodename_hostrange",
				}
			)

			for _, filename := range testdata {
				var database libgenders.Database
				path := filepath.Join("./testdata", filename)

				it.Before(func() {
					var err error
					database, err = libgenders.NewDatabase(path)
					Expect(err).NotTo(HaveOccurred())
				})

				context(fmt.Sprintf("given %s", path), func() {
					it("returns a list of nodes", func() {
						nodes := database.GetNodes()
						Expect(nodes).To(Equal([]libgenders.Node{
							{
								Name: "node1",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr3": "node1",
								},
							},
							{
								Name: "node2",
								Attributes: map[string]string{
									"attr1": "",
									"attr2": "val2",
									"attr3": "node2",
								},
							},
						}))
					})
				})
			}
		})

		context("when the node attributes contain special characters", func() {
			var database libgenders.Database

			it.Before(func() {
				var err error
				database, err = libgenders.NewDatabase("./testdata/genders.query_special_chars")
				Expect(err).NotTo(HaveOccurred())
			})

			it("handles them correctly", func() {
				nodes := database.GetNodes()
				Expect(nodes).To(Equal([]libgenders.Node{
					{
						Name: "node1",
						Attributes: map[string]string{
							"attr%percent":      "",
							"attr|pipe":         "",
							"attr&ampersand":    "",
							"attr-minus":        "",
							"attr:colon":        "",
							"attr\\backslash":   "",
							"attr/forwardslash": "",
						},
					},
					{
						Name: "node2",
						Attributes: map[string]string{
							"attr%foo%percent":      "",
							"attr|foo|pipe":         "",
							"attr&foo&ampersand":    "",
							"attr-foo-minus":        "",
							"attr:foo:colon":        "",
							"attr\\foo\\backslash":  "",
							"attr/foo/forwardslash": "",
						},
					},
					{
						Name: "node3",
						Attributes: map[string]string{
							"attr1": "attr1%percent",
							"attr2": "attr2|pipe",
							"attr3": "attr3&ampersand",
							"attr4": "attr4-minus",
							"attr5": "attr5:colon",
							"attr6": "attr6\\backslash",
							"attr7": "attr7/forwardslash",
						},
					},
					{
						Name: "node4",
						Attributes: map[string]string{
							"attr1+plus":     "",
							"attr2+foo+plus": "",
							"attr3":          "val3+plus",
							"attr4":          "val4+foo+plus",
						},
					},
				}))
			})
		})

		context("when the node attributes equal signs in their values", func() {
			var database libgenders.Database

			it.Before(func() {
				var err error
				database, err = libgenders.NewDatabase("./testdata/genders.equal_sign_in_value")
				Expect(err).NotTo(HaveOccurred())
			})

			it("handles them correctly", func() {
				nodes := database.GetNodes()
				Expect(nodes).To(Equal([]libgenders.Node{
					{
						Name: "node1",
						Attributes: map[string]string{
							"attr1": "foo=bar",
						},
					},
					{
						Name: "node2",
						Attributes: map[string]string{
							"attr1": "foo=baz",
							"attr2": "",
						},
					},
				}))
			})
		})

		context("when the node attributes contain escape characters", func() {
			it("handles them correctly", func() {
				database, err := libgenders.NewDatabase("./testdata/genders.subst_escape_char")
				Expect(err).NotTo(HaveOccurred())

				nodes := database.GetNodes()
				Expect(nodes).To(Equal([]libgenders.Node{
					{
						Name: "node1",
						Attributes: map[string]string{
							"attr1":   "",
							"attr2":   "val2",
							"escape1": "%t",
							"escape2": "%t",
							"escape3": "%n",
						},
					},
					{
						Name: "node2",
						Attributes: map[string]string{
							"attr1":   "",
							"attr2":   "val2",
							"escape1": "%t",
							"escape2": "%t",
							"escape3": "%n",
						},
					},
				}))
			})

			it("handles all cases correctly", func() {
				database, err := libgenders.NewDatabase("./testdata/genders.flag_test_raw_values")
				Expect(err).NotTo(HaveOccurred())

				nodes := database.GetNodes()
				Expect(nodes).To(Equal([]libgenders.Node{
					{
						Name: "node1",
						Attributes: map[string]string{
							"escape1": "%t",
							"escape2": "%t",
							"escape3": "%n",
							"escape4": "node1",
						},
					},
				}))
			})
		})
	})

	context("GetNodeAttr", func() {
		var database libgenders.Database

		it.Before(func() {
			var err error
			database, err = libgenders.NewDatabase("./testdata/genders.query_2_hostrange")
			Expect(err).NotTo(HaveOccurred())
		})

		it("retrieves the value of an attribute for a given node", func() {
			value, found := database.GetNodeAttr("node1", "attr2")
			Expect(found).To(BeTrue())
			Expect(value).To(Equal("valB"))
		})

		context("when the node does not exist", func() {
			it("returns false", func() {
				_, found := database.GetNodeAttr("no-such-node", "attr2")
				Expect(found).To(BeFalse())
			})
		})

		context("when the attr does not exist", func() {
			it("returns false", func() {
				_, found := database.GetNodeAttr("node1", "no-such-attr")
				Expect(found).To(BeFalse())
			})
		})
	})

	context("Query", func() {
		var database libgenders.Database

		it.Before(func() {
			var err error
			database, err = libgenders.NewDatabase("./testdata/genders.query_1_hostrange")
			Expect(err).NotTo(HaveOccurred())
		})

		context("when given simple queries", func() {
			data := map[string][]string{
				"attr1":  {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr2":  {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr3":  {"node1", "node2", "node3", "node4"},
				"attr4":  {"node1", "node2", "node3", "node4"},
				"attr5":  {"node5", "node6", "node7", "node8"},
				"attr6":  {"node5", "node6", "node7", "node8"},
				"attr7":  {"node1", "node3", "node5", "node7"},
				"attr8":  {"node1", "node3", "node5", "node7"},
				"attr9":  {"node2", "node4", "node6", "node8"},
				"attr10": {"node2", "node4", "node6", "node8"},

				"attr2=val2":   {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr4=val4":   {"node1", "node2", "node3", "node4"},
				"attr6=val6":   {"node5", "node6", "node7", "node8"},
				"attr8=val8":   {"node1", "node3", "node5", "node7"},
				"attr10=val10": {"node2", "node4", "node6", "node8"},
			}

			for query, result := range data {
				q, r := query, result

				it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
					nodes, err := database.Query(q)
					Expect(err).NotTo(HaveOccurred())

					var names []string
					for _, node := range nodes {
						names = append(names, node.Name)
					}

					Expect(names).To(Equal(r))
				})
			}
		})

		context("when the query produces an empty set", func() {
			data := map[string][]string{
				"fakeattr":                            nil,
				"attr1=fakeval":                       nil,
				"attr2=fakeval":                       nil,
				"attr3&&attr5":                        nil,
				"attr4&&attr6":                        nil,
				"attr4=val4&&attr6=val6":              nil,
				"attr7&&attr9":                        nil,
				"attr8&&attr10":                       nil,
				"attr8=val8&&attr10=val10":            nil,
				"(attr1&&attr3)--(attr1&&attr4=val4)": nil,
				"(attr2&&attr3)--(attr2&&attr4=val4)": nil,
				"(attr2=val2&&attr3)--(attr2=val2&&attr4=val4)":   nil,
				"(attr1&&attr5)--(attr1&&attr6=val6)":             nil,
				"(attr2&&attr5)--(attr2&&attr6=val6)":             nil,
				"(attr2=val2&&attr5)--(attr2=val2&&attr6=val6)":   nil,
				"(attr1&&attr7)--(attr1&&attr8=val8)":             nil,
				"(attr2&&attr7)--(attr2&&attr8=val8)":             nil,
				"(attr2=val2&&attr7)--(attr2=val2&&attr8=val8)":   nil,
				"(attr1&&attr9)--(attr1&&attr10=val10)":           nil,
				"(attr2&&attr9)--(attr2&&attr10=val10)":           nil,
				"(attr2=val2&&attr9)--(attr2=val2&&attr10=val10)": nil,
				"~attr1":      nil,
				"~attr2":      nil,
				"~attr2=val2": nil,
			}

			for query, result := range data {
				q, r := query, result

				it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
					nodes, err := database.Query(q)
					Expect(err).NotTo(HaveOccurred())

					var names []string
					for _, node := range nodes {
						names = append(names, node.Name)
					}

					Expect(names).To(Equal(r))
				})
			}
		})

		context("when the query contains binary operations", func() {
			data := map[string][]string{
				"attr1&&attr2":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr1&&attr2=val2":        {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr1||attr2":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr1||attr2=val2":        {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr3||attr5":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr4||attr6":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr4=val4||attr6=val6":   {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr7||attr9":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr8||attr10":            {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"attr8=val8||attr10=val10": {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
			}

			for query, result := range data {
				q, r := query, result

				it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
					nodes, err := database.Query(q)
					Expect(err).NotTo(HaveOccurred())

					var names []string
					for _, node := range nodes {
						names = append(names, node.Name)
					}

					Expect(names).To(Equal(r))
				})
			}
		})

		context("when the query contains complement operations", func() {
			data := map[string][]string{
				"~attr1":                      nil,
				"~attr2":                      nil,
				"~attr3":                      {"node5", "node6", "node7", "node8"},
				"~attr4":                      {"node5", "node6", "node7", "node8"},
				"~attr5":                      {"node1", "node2", "node3", "node4"},
				"~attr6":                      {"node1", "node2", "node3", "node4"},
				"~attr7":                      {"node2", "node4", "node6", "node8"},
				"~attr8":                      {"node2", "node4", "node6", "node8"},
				"~attr9":                      {"node1", "node3", "node5", "node7"},
				"~attr10":                     {"node1", "node3", "node5", "node7"},
				"~(attr3&&attr5)":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr4&&attr6)":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr4=val4&&attr6=val6)":   {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr7&&attr9)":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr8&&attr10)":            {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr8=val8&&attr10=val10)": {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
			}

			for query, result := range data {
				q, r := query, result

				it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
					nodes, err := database.Query(q)
					Expect(err).NotTo(HaveOccurred())

					var names []string
					for _, node := range nodes {
						names = append(names, node.Name)
					}

					Expect(names).To(Equal(r))
				})
			}
		})

		context("when the query contains mixed operations", func() {
			data := map[string][]string{
				"(attr1&&attr3)||(attr1&&attr5)":                       {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr1&&attr4=val4)||(attr1&&attr6=val6)":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr1&&attr7)||(attr1&&attr9)":                       {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr1&&attr8=val8)||(attr1&&attr10=val10)":           {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr2&&attr3)||(attr2&&attr5)":                       {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr2&&attr4=val4)||(attr2&&attr6=val6)":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr2&&attr7)||(attr2&&attr9)":                       {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr2&&attr8=val8)||(attr2&&attr10=val10)":           {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr2=val2&&attr3)||(attr2=val2&&attr5)":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr2=val2&&attr4=val4)||(attr2=val2&&attr6=val6)":   {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr2=val2&&attr7)||(attr2=val2&&attr9)":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"(attr2=val2&&attr8=val8)||(attr2=val2&&attr10=val10)": {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
			}

			for query, result := range data {
				q, r := query, result

				it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
					nodes, err := database.Query(q)
					Expect(err).NotTo(HaveOccurred())

					var names []string
					for _, node := range nodes {
						names = append(names, node.Name)
					}

					Expect(names).To(Equal(r))
				})
			}
		})

		context("when the query contains complex operations", func() {
			data := map[string][]string{
				"attr1&&attr3":             {"node1", "node2", "node3", "node4"},
				"attr2&&attr3":             {"node1", "node2", "node3", "node4"},
				"attr2=val2&&attr3":        {"node1", "node2", "node3", "node4"},
				"attr1&&attr4":             {"node1", "node2", "node3", "node4"},
				"attr2&&attr4":             {"node1", "node2", "node3", "node4"},
				"attr2=val2&&attr4":        {"node1", "node2", "node3", "node4"},
				"attr1&&attr4=val4":        {"node1", "node2", "node3", "node4"},
				"attr2&&attr4=val4":        {"node1", "node2", "node3", "node4"},
				"attr2=val2&&attr4=val4":   {"node1", "node2", "node3", "node4"},
				"attr1--attr5":             {"node1", "node2", "node3", "node4"},
				"attr1--attr6":             {"node1", "node2", "node3", "node4"},
				"attr1--attr6=val6":        {"node1", "node2", "node3", "node4"},
				"attr2--attr5":             {"node1", "node2", "node3", "node4"},
				"attr2--attr6":             {"node1", "node2", "node3", "node4"},
				"attr2--attr6=val6":        {"node1", "node2", "node3", "node4"},
				"attr2=val2--attr5":        {"node1", "node2", "node3", "node4"},
				"attr2=val2--attr6":        {"node1", "node2", "node3", "node4"},
				"attr2=val2--attr6=val6":   {"node1", "node2", "node3", "node4"},
				"attr1&&attr5":             {"node5", "node6", "node7", "node8"},
				"attr2&&attr5":             {"node5", "node6", "node7", "node8"},
				"attr2=val2&&attr5":        {"node5", "node6", "node7", "node8"},
				"attr1&&attr6":             {"node5", "node6", "node7", "node8"},
				"attr2&&attr6":             {"node5", "node6", "node7", "node8"},
				"attr2=val2&&attr6":        {"node5", "node6", "node7", "node8"},
				"attr1&&attr6=val6":        {"node5", "node6", "node7", "node8"},
				"attr2&&attr6=val6":        {"node5", "node6", "node7", "node8"},
				"attr2=val2&&attr6=val6":   {"node5", "node6", "node7", "node8"},
				"attr1--attr3":             {"node5", "node6", "node7", "node8"},
				"attr1--attr4":             {"node5", "node6", "node7", "node8"},
				"attr1--attr4=val4":        {"node5", "node6", "node7", "node8"},
				"attr2--attr3":             {"node5", "node6", "node7", "node8"},
				"attr2--attr4":             {"node5", "node6", "node7", "node8"},
				"attr2--attr4=val4":        {"node5", "node6", "node7", "node8"},
				"attr2=val2--attr3":        {"node5", "node6", "node7", "node8"},
				"attr2=val2--attr4":        {"node5", "node6", "node7", "node8"},
				"attr2=val2--attr4=val4":   {"node5", "node6", "node7", "node8"},
				"attr1&&attr7":             {"node1", "node3", "node5", "node7"},
				"attr2&&attr7":             {"node1", "node3", "node5", "node7"},
				"attr2=val2&&attr7":        {"node1", "node3", "node5", "node7"},
				"attr1&&attr8":             {"node1", "node3", "node5", "node7"},
				"attr2&&attr8":             {"node1", "node3", "node5", "node7"},
				"attr2=val2&&attr8":        {"node1", "node3", "node5", "node7"},
				"attr1&&attr8=val8":        {"node1", "node3", "node5", "node7"},
				"attr2&&attr8=val8":        {"node1", "node3", "node5", "node7"},
				"attr2=val2&&attr8=val8":   {"node1", "node3", "node5", "node7"},
				"attr1--attr9":             {"node1", "node3", "node5", "node7"},
				"attr1--attr10":            {"node1", "node3", "node5", "node7"},
				"attr1--attr10=val10":      {"node1", "node3", "node5", "node7"},
				"attr2--attr9":             {"node1", "node3", "node5", "node7"},
				"attr2--attr10":            {"node1", "node3", "node5", "node7"},
				"attr2--attr10=val10":      {"node1", "node3", "node5", "node7"},
				"attr2=val2--attr9":        {"node1", "node3", "node5", "node7"},
				"attr2=val2--attr10":       {"node1", "node3", "node5", "node7"},
				"attr2=val2--attr10=val10": {"node1", "node3", "node5", "node7"},
				"attr1&&attr9":             {"node2", "node4", "node6", "node8"},
				"attr2&&attr9":             {"node2", "node4", "node6", "node8"},
				"attr2=val2&&attr9":        {"node2", "node4", "node6", "node8"},
				"attr1&&attr10":            {"node2", "node4", "node6", "node8"},
				"attr2&&attr10":            {"node2", "node4", "node6", "node8"},
				"attr2=val2&&attr10":       {"node2", "node4", "node6", "node8"},
				"attr1&&attr10=val10":      {"node2", "node4", "node6", "node8"},
				"attr2&&attr10=val10":      {"node2", "node4", "node6", "node8"},
				"attr2=val2&&attr10=val10": {"node2", "node4", "node6", "node8"},
				"attr1--attr7":             {"node2", "node4", "node6", "node8"},
				"attr1--attr8":             {"node2", "node4", "node6", "node8"},
				"attr1--attr8=val8":        {"node2", "node4", "node6", "node8"},
				"attr2--attr7":             {"node2", "node4", "node6", "node8"},
				"attr2--attr8":             {"node2", "node4", "node6", "node8"},
				"attr2--attr8=val8":        {"node2", "node4", "node6", "node8"},
				"attr2=val2--attr7":        {"node2", "node4", "node6", "node8"},
				"attr2=val2--attr8":        {"node2", "node4", "node6", "node8"},
				"attr2=val2--attr8=val8":   {"node2", "node4", "node6", "node8"},
			}

			for query, result := range data {
				q, r := query, result

				it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
					nodes, err := database.Query(q)
					Expect(err).NotTo(HaveOccurred())

					var names []string
					for _, node := range nodes {
						names = append(names, node.Name)
					}

					Expect(names).To(Equal(r))
				})
			}
		})

		context("when the query results in the empty complement", func() {
			data := map[string][]string{
				"~fakeattr":            {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~attr1=fakeval":       {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~attr2=fakeval":       {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr1--attr1)":      {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr1--attr2)":      {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr1--attr2=val2)": {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr1--((attr1&&attr3)||(attr1&&attr5)))":                            {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr1--((attr1&&attr4=val4)||(attr1&&attr6=val6)))":                  {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr1--((attr1&&attr7)||(attr1&&attr9)))":                            {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr1--((attr1&&attr8=val8)||(attr1&&attr10=val10)))":                {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr2--((attr2&&attr3)||(attr2&&attr5)))":                            {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr2--((attr2&&attr4=val4)||(attr2&&attr6=val6)))":                  {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr2--((attr2&&attr7)||(attr2&&attr9)))":                            {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr2--((attr2&&attr8=val8)||(attr2&&attr10=val10)))":                {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr2=val2--((attr2=val2&&attr3)||(attr2=val2&&attr5)))":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr2=val2--((attr2=val2&&attr4=val4)||(attr2=val2&&attr6=val6)))":   {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr2=val2--((attr2=val2&&attr7)||(attr2=val2&&attr9)))":             {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(attr2=val2--((attr2=val2&&attr8=val8)||(attr2=val2&&attr10=val10)))": {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
			}

			for query, result := range data {
				q, r := query, result

				it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
					nodes, err := database.Query(q)
					Expect(err).NotTo(HaveOccurred())

					var names []string
					for _, node := range nodes {
						names = append(names, node.Name)
					}

					Expect(names).To(Equal(r))
				})
			}
		})

		context("when the query contains double negation", func() {
			data := map[string][]string{
				"~(~(attr1))":      {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(~(attr2))":      {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
				"~(~(attr2=val2))": {"node1", "node2", "node3", "node4", "node5", "node6", "node7", "node8"},
			}

			for query, result := range data {
				q, r := query, result

				it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
					nodes, err := database.Query(q)
					Expect(err).NotTo(HaveOccurred())

					var names []string
					for _, node := range nodes {
						names = append(names, node.Name)
					}

					Expect(names).To(Equal(r))
				})
			}
		})
	})
}
