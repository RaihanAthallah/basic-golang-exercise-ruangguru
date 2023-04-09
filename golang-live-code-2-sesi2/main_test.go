package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Student Management", func() {
		var students map[string][]interface{}

		BeforeEach(func() {
			students = make(map[string][]interface{})
		})

		Describe("RemoveStudent", func() {
			It("should remove a student from the map", func() {
				students["John Doe"] = []interface{}{"123 Main St", "555-1234", 80}
				main.RemoveStudent(&students)("John Doe")
				Expect(len(students)).To(Equal(0))
				Expect(students["John Doe"]).To(BeNil())
			})
		})

		Describe("UpdateScore", func() {
			Context("when the student exists", func() {
				It("should update the score of the student", func() {
					students["John Doe"] = []interface{}{"123 Main St", "555-1234", 80}
					main.UpdateScore(&students)("John Doe", 90)
					Expect(students["John Doe"][2]).To(Equal(90))
				})
			})
		})
	})
})
