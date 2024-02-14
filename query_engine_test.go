package libgenders_test

import (
	"fmt"
	"testing"

	"github.com/ryanmoran/libgenders"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testQueryEngine(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		engine libgenders.QueryEngine
	)

	it.Before(func() {
		database, err := libgenders.NewDatabase("./testdata/genders.query_1_hostrange")
		Expect(err).NotTo(HaveOccurred())

		engine = libgenders.NewQueryEngine(database.GetNodes())
	})

	context("when given simple queries", func() {
		data := map[string][]int{
			"attr1":  {0, 1, 2, 3, 4, 5, 6, 7},
			"attr2":  {0, 1, 2, 3, 4, 5, 6, 7},
			"attr3":  {0, 1, 2, 3},
			"attr4":  {0, 1, 2, 3},
			"attr5":  {4, 5, 6, 7},
			"attr6":  {4, 5, 6, 7},
			"attr7":  {0, 2, 4, 6},
			"attr8":  {0, 2, 4, 6},
			"attr9":  {1, 3, 5, 7},
			"attr10": {1, 3, 5, 7},

			"attr2=val2":   {0, 1, 2, 3, 4, 5, 6, 7},
			"attr4=val4":   {0, 1, 2, 3},
			"attr6=val6":   {4, 5, 6, 7},
			"attr8=val8":   {0, 2, 4, 6},
			"attr10=val10": {1, 3, 5, 7},
		}

		for query, result := range data {
			q, r := query, result

			it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
				Expect(engine.Query(q)).To(Equal(r))
			})
		}
	})

	context("when the query produces an empty set", func() {
		data := map[string][]int{
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
				Expect(engine.Query(q)).To(Equal(r))
			})
		}
	})

	context("when the query contains binary operations", func() {
		data := map[string][]int{
			"attr1&&attr2":             {0, 1, 2, 3, 4, 5, 6, 7},
			"attr1&&attr2=val2":        {0, 1, 2, 3, 4, 5, 6, 7},
			"attr1||attr2":             {0, 1, 2, 3, 4, 5, 6, 7},
			"attr1||attr2=val2":        {0, 1, 2, 3, 4, 5, 6, 7},
			"attr3||attr5":             {0, 1, 2, 3, 4, 5, 6, 7},
			"attr4||attr6":             {0, 1, 2, 3, 4, 5, 6, 7},
			"attr4=val4||attr6=val6":   {0, 1, 2, 3, 4, 5, 6, 7},
			"attr7||attr9":             {0, 1, 2, 3, 4, 5, 6, 7},
			"attr8||attr10":            {0, 1, 2, 3, 4, 5, 6, 7},
			"attr8=val8||attr10=val10": {0, 1, 2, 3, 4, 5, 6, 7},
		}

		for query, result := range data {
			q, r := query, result

			it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
				Expect(engine.Query(q)).To(Equal(r), fmt.Sprintf("%#v\n", engine))
			})
		}
	})

	context("when the query contains complement operations", func() {
		data := map[string][]int{
			"~attr1":                      nil,
			"~attr2":                      nil,
			"~attr3":                      {4, 5, 6, 7},
			"~attr4":                      {4, 5, 6, 7},
			"~attr5":                      {0, 1, 2, 3},
			"~attr6":                      {0, 1, 2, 3},
			"~attr7":                      {1, 3, 5, 7},
			"~attr8":                      {1, 3, 5, 7},
			"~attr9":                      {0, 2, 4, 6},
			"~attr10":                     {0, 2, 4, 6},
			"~(attr3&&attr5)":             {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr4&&attr6)":             {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr4=val4&&attr6=val6)":   {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr7&&attr9)":             {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr8&&attr10)":            {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr8=val8&&attr10=val10)": {0, 1, 2, 3, 4, 5, 6, 7},
		}

		for query, result := range data {
			q, r := query, result

			it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
				Expect(engine.Query(q)).To(Equal(r), fmt.Sprintf("%#v\n", engine))
			})
		}
	})

	context("when the query contains mixed operations", func() {
		data := map[string][]int{
			"(attr1&&attr3)||(attr1&&attr5)":                       {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr1&&attr4=val4)||(attr1&&attr6=val6)":             {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr1&&attr7)||(attr1&&attr9)":                       {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr1&&attr8=val8)||(attr1&&attr10=val10)":           {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr2&&attr3)||(attr2&&attr5)":                       {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr2&&attr4=val4)||(attr2&&attr6=val6)":             {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr2&&attr7)||(attr2&&attr9)":                       {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr2&&attr8=val8)||(attr2&&attr10=val10)":           {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr2=val2&&attr3)||(attr2=val2&&attr5)":             {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr2=val2&&attr4=val4)||(attr2=val2&&attr6=val6)":   {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr2=val2&&attr7)||(attr2=val2&&attr9)":             {0, 1, 2, 3, 4, 5, 6, 7},
			"(attr2=val2&&attr8=val8)||(attr2=val2&&attr10=val10)": {0, 1, 2, 3, 4, 5, 6, 7},
		}

		for query, result := range data {
			q, r := query, result

			it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
				Expect(engine.Query(q)).To(Equal(r), fmt.Sprintf("%#v\n", engine))
			})
		}
	})

	context("when the query contains complex operations", func() {
		data := map[string][]int{
			"attr1&&attr3":             {0, 1, 2, 3},
			"attr2&&attr3":             {0, 1, 2, 3},
			"attr2=val2&&attr3":        {0, 1, 2, 3},
			"attr1&&attr4":             {0, 1, 2, 3},
			"attr2&&attr4":             {0, 1, 2, 3},
			"attr2=val2&&attr4":        {0, 1, 2, 3},
			"attr1&&attr4=val4":        {0, 1, 2, 3},
			"attr2&&attr4=val4":        {0, 1, 2, 3},
			"attr2=val2&&attr4=val4":   {0, 1, 2, 3},
			"attr1--attr5":             {0, 1, 2, 3},
			"attr1--attr6":             {0, 1, 2, 3},
			"attr1--attr6=val6":        {0, 1, 2, 3},
			"attr2--attr5":             {0, 1, 2, 3},
			"attr2--attr6":             {0, 1, 2, 3},
			"attr2--attr6=val6":        {0, 1, 2, 3},
			"attr2=val2--attr5":        {0, 1, 2, 3},
			"attr2=val2--attr6":        {0, 1, 2, 3},
			"attr2=val2--attr6=val6":   {0, 1, 2, 3},
			"attr1&&attr5":             {4, 5, 6, 7},
			"attr2&&attr5":             {4, 5, 6, 7},
			"attr2=val2&&attr5":        {4, 5, 6, 7},
			"attr1&&attr6":             {4, 5, 6, 7},
			"attr2&&attr6":             {4, 5, 6, 7},
			"attr2=val2&&attr6":        {4, 5, 6, 7},
			"attr1&&attr6=val6":        {4, 5, 6, 7},
			"attr2&&attr6=val6":        {4, 5, 6, 7},
			"attr2=val2&&attr6=val6":   {4, 5, 6, 7},
			"attr1--attr3":             {4, 5, 6, 7},
			"attr1--attr4":             {4, 5, 6, 7},
			"attr1--attr4=val4":        {4, 5, 6, 7},
			"attr2--attr3":             {4, 5, 6, 7},
			"attr2--attr4":             {4, 5, 6, 7},
			"attr2--attr4=val4":        {4, 5, 6, 7},
			"attr2=val2--attr3":        {4, 5, 6, 7},
			"attr2=val2--attr4":        {4, 5, 6, 7},
			"attr2=val2--attr4=val4":   {4, 5, 6, 7},
			"attr1&&attr7":             {0, 2, 4, 6},
			"attr2&&attr7":             {0, 2, 4, 6},
			"attr2=val2&&attr7":        {0, 2, 4, 6},
			"attr1&&attr8":             {0, 2, 4, 6},
			"attr2&&attr8":             {0, 2, 4, 6},
			"attr2=val2&&attr8":        {0, 2, 4, 6},
			"attr1&&attr8=val8":        {0, 2, 4, 6},
			"attr2&&attr8=val8":        {0, 2, 4, 6},
			"attr2=val2&&attr8=val8":   {0, 2, 4, 6},
			"attr1--attr9":             {0, 2, 4, 6},
			"attr1--attr10":            {0, 2, 4, 6},
			"attr1--attr10=val10":      {0, 2, 4, 6},
			"attr2--attr9":             {0, 2, 4, 6},
			"attr2--attr10":            {0, 2, 4, 6},
			"attr2--attr10=val10":      {0, 2, 4, 6},
			"attr2=val2--attr9":        {0, 2, 4, 6},
			"attr2=val2--attr10":       {0, 2, 4, 6},
			"attr2=val2--attr10=val10": {0, 2, 4, 6},
			"attr1&&attr9":             {1, 3, 5, 7},
			"attr2&&attr9":             {1, 3, 5, 7},
			"attr2=val2&&attr9":        {1, 3, 5, 7},
			"attr1&&attr10":            {1, 3, 5, 7},
			"attr2&&attr10":            {1, 3, 5, 7},
			"attr2=val2&&attr10":       {1, 3, 5, 7},
			"attr1&&attr10=val10":      {1, 3, 5, 7},
			"attr2&&attr10=val10":      {1, 3, 5, 7},
			"attr2=val2&&attr10=val10": {1, 3, 5, 7},
			"attr1--attr7":             {1, 3, 5, 7},
			"attr1--attr8":             {1, 3, 5, 7},
			"attr1--attr8=val8":        {1, 3, 5, 7},
			"attr2--attr7":             {1, 3, 5, 7},
			"attr2--attr8":             {1, 3, 5, 7},
			"attr2--attr8=val8":        {1, 3, 5, 7},
			"attr2=val2--attr7":        {1, 3, 5, 7},
			"attr2=val2--attr8":        {1, 3, 5, 7},
			"attr2=val2--attr8=val8":   {1, 3, 5, 7},
		}

		for query, result := range data {
			q, r := query, result

			it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
				Expect(engine.Query(q)).To(Equal(r), fmt.Sprintf("%#v\n", engine))
			})
		}
	})

	context("when the query results in the empty complement", func() {
		data := map[string][]int{
			"~fakeattr":            {0, 1, 2, 3, 4, 5, 6, 7},
			"~attr1=fakeval":       {0, 1, 2, 3, 4, 5, 6, 7},
			"~attr2=fakeval":       {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr1--attr1)":      {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr1--attr2)":      {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr1--attr2=val2)": {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr1--((attr1&&attr3)||(attr1&&attr5)))":                            {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr1--((attr1&&attr4=val4)||(attr1&&attr6=val6)))":                  {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr1--((attr1&&attr7)||(attr1&&attr9)))":                            {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr1--((attr1&&attr8=val8)||(attr1&&attr10=val10)))":                {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr2--((attr2&&attr3)||(attr2&&attr5)))":                            {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr2--((attr2&&attr4=val4)||(attr2&&attr6=val6)))":                  {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr2--((attr2&&attr7)||(attr2&&attr9)))":                            {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr2--((attr2&&attr8=val8)||(attr2&&attr10=val10)))":                {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr2=val2--((attr2=val2&&attr3)||(attr2=val2&&attr5)))":             {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr2=val2--((attr2=val2&&attr4=val4)||(attr2=val2&&attr6=val6)))":   {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr2=val2--((attr2=val2&&attr7)||(attr2=val2&&attr9)))":             {0, 1, 2, 3, 4, 5, 6, 7},
			"~(attr2=val2--((attr2=val2&&attr8=val8)||(attr2=val2&&attr10=val10)))": {0, 1, 2, 3, 4, 5, 6, 7},
		}

		for query, result := range data {
			q, r := query, result

			it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
				Expect(engine.Query(q)).To(Equal(r), fmt.Sprintf("%#v\n", engine))
			})
		}
	})

	context("when the query contains double negation", func() {
		data := map[string][]int{
			"~(~(attr1))":      {0, 1, 2, 3, 4, 5, 6, 7},
			"~(~(attr2))":      {0, 1, 2, 3, 4, 5, 6, 7},
			"~(~(attr2=val2))": {0, 1, 2, 3, 4, 5, 6, 7},
		}

		for query, result := range data {
			q, r := query, result

			it(fmt.Sprintf("finds the correct results for the query %q", q), func() {
				Expect(engine.Query(q)).To(Equal(r), fmt.Sprintf("%#v\n", engine))
			})
		}
	})
}
