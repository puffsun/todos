package app_test

import (
	//. "github.com/puffsun/todos/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("AddNewTodo", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		// {android|chrome|firefox|htmlunit|internet explorer|iPhone|iPad|opera|safari}
		page, err = agoutiDriver.NewPage(agouti.Browser("chrome"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should make below URL available", func() {
		By("visiting the URLs", func() {
			urls := []string{
				"http://localhost:8000/",
				"http://localhost:8000/active",
				"http://localhost:8000/completed",
			}

			for _, url := range urls {
				Expect(page.Navigate(url)).To(Succeed())
			}
		})
	})

	It("should manage to add new todo item", func() {
		By("visiting todo app home", func() {
			Expect(page.Navigate("http://localhost:8000")).To(Succeed())
			// By default, todo app will add a empty hash to URL
			Expect(page).To(HaveURL("http://localhost:8000/#/"))
		})

		By("adding one todo item", func() {
			Eventually(page.Find("#toggle-all")).Should(BeFound())
			Eventually(page.Find("#todo-form")).Should(BeFound())
			Eventually(page.Find("#todo-count")).Should(BeFound())
			Eventually(page.Find("#clear-completed")).Should(BeFound())

			Expect(page.Find("#new-todo").Fill("Learning Golang")).To(Succeed())
			Expect(page.Find("#todo-form").Submit()).To(Succeed())
		})
	})
})
