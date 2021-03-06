package db_test

import (
	"github.com/chenbh/concourse/atc"
	"github.com/chenbh/concourse/atc/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PipelineLifecycle", func() {
	var (
		pl            db.PipelineLifecycle
		childPipeline db.Pipeline
		err           error
	)

	BeforeEach(func() {
		pl = db.NewPipelineLifecycle(dbConn, lockFactory)
	})

	JustBeforeEach(func() {
		err = pl.ArchiveAbandonedPipelines()
		Expect(err).NotTo(HaveOccurred())
	})

	Context("child pipeline is set by a job in a pipeline", func() {
		BeforeEach(func() {
			build, _ := defaultJob.CreateBuild()
			childPipeline, _, _ = build.SavePipeline("child-pipeline", defaultTeam.ID(), defaultPipelineConfig, db.ConfigVersion(0), false)
			build.Finish(db.BuildStatusSucceeded)
		})

		Context("parent pipeline is destroyed", func() {
			BeforeEach(func() {
				defaultPipeline.Destroy()
			})

			It("should archive all child pipelines", func() {
				childPipeline.Reload()
				Expect(childPipeline.Archived()).To(BeTrue())
			})
		})

		Context("parent pipeline is archived", func() {
			BeforeEach(func() {
				defaultPipeline.Archive()
			})

			It("should archive all child pipelines", func() {
				childPipeline.Reload()
				Expect(childPipeline.Archived()).To(BeTrue())
			})
		})

		Context("job is removed from the parent pipeline", func() {
			BeforeEach(func() {
				defaultPipelineConfig.Jobs = atc.JobConfigs{
					{
						Name: "a-different-job",
					},
				}
				defaultTeam.SavePipeline("default-pipeline", defaultPipelineConfig, defaultPipeline.ConfigVersion(), false)
			})

			It("archives all child pipelines set by the deleted job", func() {
				childPipeline.Reload()
				Expect(childPipeline.Archived()).To(BeTrue())
			})
		})
	})

	Context("pipeline does not have a parent job and build ID", func() {
		It("Should not archive the pipeline", func() {
			defaultPipeline.Reload()
			Expect(defaultPipeline.Archived()).To(BeFalse())
		})
	})
})
