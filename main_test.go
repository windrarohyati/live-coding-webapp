package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Go EduHub 3", Ordered, func() {
	var apiServer *gin.Engine
	var studentRepo repo.StudentRepository
	var courseRepo repo.CourseRepository

	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	studentRepo = repo.NewStudentRepo(conn)
	courseRepo = repo.NewCourseRepo(conn)

	BeforeEach(func() {
		gin.SetMode(gin.ReleaseMode) //release
		err = conn.AutoMigrate(&model.Student{}, &model.Course{})
		Expect(err).ShouldNot(HaveOccurred())
		err = db.Reset(conn, "students")
		err = db.Reset(conn, "courses")
		Expect(err).ShouldNot(HaveOccurred())

		apiServer = gin.New()
		apiServer = main.RunServer(conn, apiServer)
	})

	Describe("Repository", func() {
		Describe("Student repository", func() {
			When("deleting student data in students table in the database", func() {
				It("should delete the existing student data in students table in the database", func() {
					student := model.Student{
						Name:     "John",
						Email:    "Jl. Raya Cilandak",
						Phone:    "081345435355",
						CourseID: 1,
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					err = studentRepo.Delete(1)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).Should(HaveOccurred())
					Expect(result).To(BeNil())

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Describe("Course repository", func() {
			When("deleting course data in courses table in the database", func() {
				It("should delete the existing course data in courses table in the database", func() {
					course := model.Course{
						Name:       "Data Structures and Algorithms",
						Schedule:   "Monday, Wednesday, and Friday, 9am - 11am",
						Grade:      3.2,
						Attendance: 80,
					}
					err := courseRepo.Store(&course)
					Expect(err).ShouldNot(HaveOccurred())

					err = courseRepo.Delete(1)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := courseRepo.FetchByID(1)
					Expect(err).Should(HaveOccurred())
					Expect(result).To(BeNil())

					err = db.Reset(conn, "courses")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})

	Describe("API", func() {
		Describe("/student/delete/:id", func() {
			When("deleting existing student", func() {
				It("should return status code 200", func() {
					student := model.Student{
						Name:     "Test Student",
						Email:    "test@student.com",
						Phone:    "123456789",
						CourseID: 1,
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/student/delete/1", nil)
					apiServer.ServeHTTP(w, r)

					var response model.SuccessResponse
					json.Unmarshal(w.Body.Bytes(), &response)

					Expect(w.Code).To(Equal(http.StatusOK))
					Expect(response.Message).To(Equal("student delete success"))
				})
			})

			When("deleting non-existent student", func() {
				It("should return status code 404", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/student/delete/999", nil)
					apiServer.ServeHTTP(w, r)

					var response model.ErrorResponse
					json.Unmarshal(w.Body.Bytes(), &response)

					Expect(w.Code).To(Equal(http.StatusNotFound))
					Expect(response.Error).NotTo(BeNil())
				})
			})

			When("deleting student with invalid ID", func() {
				It("should return status code 400", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/student/delete/invalid", nil)
					apiServer.ServeHTTP(w, r)

					var response model.ErrorResponse
					json.Unmarshal(w.Body.Bytes(), &response)

					Expect(w.Code).To(Equal(http.StatusBadRequest))
					Expect(response.Error).NotTo(BeNil())
				})
			})
		})

		Describe("/course/delete/:id", func() {
			When("deleting existing course", func() {
				It("should return status code 200", func() {
					course := model.Course{
						Name:       "Data Structures and Algorithms",
						Schedule:   "Monday, Wednesday, and Friday, 9am - 11am",
						Grade:      3.2,
						Attendance: 80,
					}
					err := courseRepo.Store(&course)
					Expect(err).ShouldNot(HaveOccurred())

					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/course/delete/1", nil)
					apiServer.ServeHTTP(w, r)

					var response model.SuccessResponse
					json.Unmarshal(w.Body.Bytes(), &response)

					Expect(w.Code).To(Equal(http.StatusOK))
					Expect(response.Message).To(Equal("course delete success"))
				})
			})

			When("deleting non-existent course", func() {
				It("should return status code 404", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/course/delete/999", nil)
					apiServer.ServeHTTP(w, r)

					var response model.ErrorResponse
					json.Unmarshal(w.Body.Bytes(), &response)

					Expect(w.Code).To(Equal(http.StatusNotFound))
					Expect(response.Error).NotTo(BeNil())
				})
			})

			When("deleting course with invalid ID", func() {
				It("should return status code 400", func() {
					w := httptest.NewRecorder()
					r := httptest.NewRequest("DELETE", "/course/delete/invalid", nil)
					apiServer.ServeHTTP(w, r)

					var response model.ErrorResponse
					json.Unmarshal(w.Body.Bytes(), &response)

					Expect(w.Code).To(Equal(http.StatusBadRequest))
					Expect(response.Error).NotTo(BeNil())
				})
			})
		})
	})
})
