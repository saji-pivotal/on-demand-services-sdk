package serviceadapter_test

import (
	"encoding/json"

	"github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Domain", func() {
	Context("RequestParameters", func() {
		Context("when arbitraryParams are present", func() {
			It("can extract arbitraryParams", func() {
				params := serviceadapter.RequestParameters{"parameters": map[string]interface{}{"foo": "bar"}}
				Expect(params.ArbitraryParams()).To(Equal(map[string]interface{}{"foo": "bar"}))
			})
		})

		Context("when arbitraryParams are absent", func() {
			It("arbitrary params are empty", func() {
				params := serviceadapter.RequestParameters{"plan_id": "baz"}
				Expect(params.ArbitraryParams()).To(Equal(map[string]interface{}{}))
			})
		})
	})
	Context("DashboardUrl", func() {
		It("serializes dashboard_url", func() {
			dashboardUrl := serviceadapter.DashboardUrl{DashboardUrl: "https://someurl.com"}
			Expect(json.Marshal(dashboardUrl)).To(MatchJSON(`{ "dashboard_url": "https://someurl.com"}`))
		})
	})
	Context("plan", func() {
		Describe("(de)serialising", func() {
			planJson := []byte(`{
				"instance_groups": [
					{
						"name": "example-server",
						"vm_type": "small",
						"persistent_disk_type": "ten",
						"networks": [
							"example-network"
						],
						"azs": [
							"example-az"
						],
						"instances": 1,
						"lifecycle": "errand"
					}
				],
				"properties": {
					"example": "property"
				 }
			}`)

			expectedPlan := serviceadapter.Plan{
				InstanceGroups: []serviceadapter.InstanceGroup{{
					Name:           "example-server",
					VMType:         "small",
					PersistentDisk: "ten",
					Networks:       []string{"example-network"},
					AZs:            []string{"example-az"},
					Instances:      1,
					Lifecycle:      "errand",
				}},
				Properties: serviceadapter.Properties{"example": "property"},
			}

			It("deserialises plan object containing all optional fields from json", func() {
				var plan serviceadapter.Plan
				Expect(json.Unmarshal(planJson, &plan)).To(Succeed())
				Expect(plan).To(Equal(expectedPlan))
			})

			It("serialises plan object containing all optional fields to json", func() {
				Expect(json.Marshal(expectedPlan)).To(MatchJSON(planJson))
			})

			It("serialises plan object containing only mandatory fields to json", func() {
				expectedPlan := serviceadapter.Plan{
					InstanceGroups: []serviceadapter.InstanceGroup{{
						Name:      "example-server",
						VMType:    "small",
						Networks:  []string{"example-network"},
						AZs:       []string{"az1"},
						Instances: 1,
					}},
					Properties: serviceadapter.Properties{},
				}

				planJson := []byte(`{
				"instance_groups": [
					{
						"name": "example-server",
						"vm_type": "small",
						"networks": [
							"example-network"
						],
						"instances": 1,
						"azs": ["az1"]
					}
				],
				"properties": {}
			}`)
				Expect(json.Marshal(expectedPlan)).To(MatchJSON(planJson))
			})

			It("serializes required fields for instance_groups", func() {
				plan := serviceadapter.Plan{
					InstanceGroups: []serviceadapter.InstanceGroup{{}},
					Properties:     serviceadapter.Properties{},
				}

				planJson, err := json.Marshal(plan)
				Expect(err).ToNot(HaveOccurred())

				Expect(planJson).To(ContainSubstring("name"))
				Expect(planJson).To(ContainSubstring("vm_type"))
				Expect(planJson).To(ContainSubstring("instances"))
				Expect(planJson).To(ContainSubstring("networks"))
				Expect(planJson).To(ContainSubstring("azs"))
			})
		})

		Describe("validation", func() {
			var plan serviceadapter.Plan

			BeforeEach(func() {
				plan = serviceadapter.Plan{
					InstanceGroups: []serviceadapter.InstanceGroup{{
						Name:      "example-server",
						VMType:    "small",
						Networks:  []string{"example-network"},
						Instances: 1,
						AZs:       []string{"az1"},
					}},
					Properties: serviceadapter.Properties{},
				}
			})

			Context("when nothing is missing", func() {
				It("returns no error", func() {
					Expect(plan.Validate()).ToNot(HaveOccurred())
				})
			})

			Context("when instance groups are missing", func() {
				BeforeEach(func() {
					plan.InstanceGroups = nil
				})

				It("returns an error", func() {
					Expect(plan.Validate()).To(HaveOccurred())
				})
			})

			Context("when vm type is missing", func() {
				BeforeEach(func() {
					plan.InstanceGroups[0].VMType = ""
				})

				It("returns an error", func() {
					Expect(plan.Validate()).To(HaveOccurred())
				})
			})

			Context("when networks is missing", func() {
				BeforeEach(func() {
					plan.InstanceGroups[0].Networks = nil
				})

				It("returns an error", func() {
					Expect(plan.Validate()).To(HaveOccurred())
				})
			})

			Context("when azs are missing", func() {
				BeforeEach(func() {
					plan.InstanceGroups[0].AZs = nil
				})

				It("returns an error", func() {
					Expect(plan.Validate()).To(HaveOccurred())
				})
			})

			Context("when azs are empty", func() {
				BeforeEach(func() {
					plan.InstanceGroups[0].AZs = []string{}
				})

				It("returns an error", func() {
					Expect(plan.Validate()).To(HaveOccurred())
				})
			})

			Context("when instances is 0", func() {
				BeforeEach(func() {
					plan.InstanceGroups[0].Instances = 0
				})

				It("returns an error", func() {
					Expect(plan.Validate()).To(HaveOccurred())
				})
			})

			Context("when instance group name is missing", func() {
				BeforeEach(func() {
					plan.InstanceGroups[0].Name = ""
				})

				It("returns an error", func() {
					Expect(plan.Validate()).To(HaveOccurred())
				})
			})
		})
	})
})
