package snail_matrix

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ()

var _ = Describe("Snail matrix solution specs", func() {
	var ()

	BeforeEach(func() {
	})

	Describe("Matrix wrapper function specs", func() {
		Context("when the input matrix is empty", func() {
			It("returns a matrix with just the number", func() {
				empty_matrix := [][]int{}

				newMatrix := matrix_wrapper_with(empty_matrix, 0)

				Expect(newMatrix[0][0]).To(Equal(0))
			})
		})
		Context("when the input matrix is just a number", func() {
			It("returns a matrix with the number wraped matrix", func() {
				matrixWithZero := [][]int{[]int{0}}

				newMatrix := matrix_wrapper_with(matrixWithZero, 1)

				Expect(newMatrix[0][0]).To(Equal(1))
				Expect(newMatrix[1][1]).To(Equal(0))
				Expect(newMatrix[2][2]).To(Equal(1))
			})
		})
		Context("when the input matrix length is 3x3", func() {
			It("returns a matrix with length 5x5", func() {
				inputMatrix := [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}

				newMatrix := matrix_wrapper_with(inputMatrix, 2)

				Expect(len(newMatrix)).To(Equal(5))
			})
			It("returns a matrix with wrapped with input number", func() {
				inputMatrix := [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}

				newMatrix := matrix_wrapper_with(inputMatrix, 2)

				Expect(newMatrix[0][0]).To(Equal(2))
				Expect(newMatrix[0][4]).To(Equal(2))
				Expect(newMatrix[4][4]).To(Equal(2))
			})
		})
	})

	Describe("Snail matrix function specs", func() {
		Context("when the input number is 0", func() {
			It("returns a matrix with just the number", func() {
				snailMatrix := snail_matrix_for(0)

				Expect(snailMatrix[0][0]).To(Equal(0))
			})
		})
		Context("when the input number is 1", func() {
			It("returns a matrix with 0 sourounded with 1", func() {
				snailMatrix := snail_matrix_for(1)

				Expect(snailMatrix[0][0]).To(Equal(1))
				Expect(snailMatrix[1][1]).To(Equal(0))
				Expect(snailMatrix[2][2]).To(Equal(1))
			})
		})
		Context("when the input number is 2", func() {
			It("returns a snail matrix sourounded with 2", func() {
				snailMatrix := snail_matrix_for(2)

				Expect(snailMatrix[0][0]).To(Equal(2))
				Expect(snailMatrix[2][2]).To(Equal(0))
				Expect(snailMatrix[4][4]).To(Equal(2))
			})
		})
		Context("when the input number is 3", func() {
			It("returns a snail matrix sourounded with 3", func() {
				snailMatrix := snail_matrix_for(3)

				Expect(snailMatrix[0][0]).To(Equal(3))
				Expect(snailMatrix[3][3]).To(Equal(0))
				Expect(snailMatrix[6][6]).To(Equal(3))
			})
		})
	})

})
