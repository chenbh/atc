package db_test

import (
	"github.com/chenbh/concourse/atc/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PipelineRef", func() {
	var (
		pr       db.PipelineRef
	)

	BeforeEach(func(){
		pr = db.NewPipelineRef(defaultPipeline.ID(), defaultPipeline.Name(), dbConn, lockFactory)
	})

	It("id should be correct", func() {
		Expect(pr.PipelineID()).To(Equal(defaultPipeline.ID()))
	})

	It("name should be correct", func() {
		Expect(pr.PipelineName()).To(Equal(defaultPipeline.Name()))
	})

	It("pipeline should be correct", func() {
		p, found, err := pr.Pipeline()
		Expect(err).ToNot(HaveOccurred())
		Expect(found).To(BeTrue())
		Expect(p).To(Equal(defaultPipeline))
	})
})
