package spacelift

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	. "github.com/spacelift-io/terraform-provider-spacelift/spacelift/internal/testhelpers"
)

func TestVCSIntegration(t *testing.T) {
	randomID := func() string { return acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum) }

	for _, resourceName := range []string{"spacelift_stack", "spacelift_module"} {
		t.Run(resourceName, func(t *testing.T) {
			t.Run("setting up ID", func(t *testing.T) {
				randomID := func() string { return acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum) }

				testCases := []struct {
					name           string
					repository     string
					branch         string
					space          string
					provider       string
					dataSource     string
					attributeValue string
					attribute      string
				}{
					// Azure Dev Ops
					{
						name:       "azure-devops-with-non-specified-integration-" + randomID(),
						repository: "spacelift-ci",
						branch:     "main",
						space:      "root",
						provider: `azure_devops {
							project = "spacelift-ci"
						}`,
						attribute:      "azure_devops.0.id",
						attributeValue: testConfig.SourceCode.AzureDevOps.Default.ID,
					},
					{
						name:       "azure-devops-with-an-empty-integration-id-" + randomID(),
						repository: "spacelift-ci",
						branch:     "main",
						space:      "root",
						provider: `azure_devops {
							project = "spacelift-ci"
							id = ""
						}`,
						attribute:      "azure_devops.0.id",
						attributeValue: testConfig.SourceCode.AzureDevOps.Default.ID,
					},
					{
						name:       "azure-devops-with-default-integration-" + randomID(),
						repository: "spacelift-ci",
						branch:     "main",
						space:      "root",
						provider: `azure_devops {
							project = "spacelift-ci"
							id = "` + testConfig.SourceCode.AzureDevOps.Default.ID + `"
						}`,
						attribute:      "azure_devops.0.id",
						attributeValue: testConfig.SourceCode.AzureDevOps.Default.ID,
					},
					{
						name:       "azure-devops-with-space-level-integration-" + randomID(),
						repository: "spacelift-ci",
						branch:     "main",
						space:      testConfig.SourceCode.AzureDevOps.SpaceLevel.Space,
						provider: `azure_devops {
							project = "spacelift-ci"
							id = "` + testConfig.SourceCode.AzureDevOps.SpaceLevel.ID + `"
						}`,
						attribute:      "azure_devops.0.id",
						attributeValue: testConfig.SourceCode.AzureDevOps.SpaceLevel.ID,
					},
					{
						name:       "azure-devops-with-default-integration-from-data-source-" + randomID(),
						repository: "spacelift-ci",
						branch:     "main",
						space:      "root",
						provider: `azure_devops {
							project = "spacelift-ci"
							id = data.spacelift_azure_devops_integration.test.id
						}`,
						dataSource:     `data "spacelift_azure_devops_integration" "test" {}`,
						attribute:      "azure_devops.0.id",
						attributeValue: testConfig.SourceCode.AzureDevOps.Default.ID,
					},
					{
						name:       "azure-devops-with-space-level-integration-from-data-source-" + randomID(),
						repository: "spacelift-ci",
						branch:     "main",
						space:      testConfig.SourceCode.AzureDevOps.SpaceLevel.Space,
						provider: `azure_devops {
							project = "spacelift-ci"
							id = data.spacelift_azure_devops_integration.test.id
						}`,
						dataSource: `data "spacelift_azure_devops_integration" "test" {
							id = "` + testConfig.SourceCode.AzureDevOps.SpaceLevel.ID + `"
						}`,
						attribute:      "azure_devops.0.id",
						attributeValue: testConfig.SourceCode.AzureDevOps.SpaceLevel.ID,
					},
					// Bitbucket Cloud
					{
						name:       "bitbucket-cloud-with-non-specified-integration-" + randomID(),
						repository: "empty",
						branch:     "master",
						space:      "root",
						provider: `bitbucket_cloud {
							namespace = "thespacelift"
						}`,
						attribute:      "bitbucket_cloud.0.id",
						attributeValue: testConfig.SourceCode.BitbucketCloud.Default.ID,
					},
					{
						name:       "bitbucket-cloud-with-an-empty-integration-id-" + randomID(),
						repository: "empty",
						branch:     "master",
						space:      "root",
						provider: `bitbucket_cloud {
							namespace = "thespacelift"
							id = ""
						}`,
						attribute:      "bitbucket_cloud.0.id",
						attributeValue: testConfig.SourceCode.BitbucketCloud.Default.ID,
					},
					{
						name:       "bitbucket-cloud-with-default-integration-" + randomID(),
						repository: "empty",
						branch:     "master",
						space:      "root",
						provider: `bitbucket_cloud {
							namespace = "thespacelift"
							id = "` + testConfig.SourceCode.BitbucketCloud.Default.ID + `"
						}`,
						attribute:      "bitbucket_cloud.0.id",
						attributeValue: testConfig.SourceCode.BitbucketCloud.Default.ID,
					},
					{
						name:       "bitbucket-cloud-with-space-level-integration-" + randomID(),
						repository: "empty",
						branch:     "master",
						space:      testConfig.SourceCode.BitbucketCloud.SpaceLevel.Space,
						provider: `bitbucket_cloud {
							namespace = "thespacelift"
							id = "` + testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID + `"
						}`,
						attribute:      "bitbucket_cloud.0.id",
						attributeValue: testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID,
					},
					{
						name:       "bitbucket-cloud-with-default-integration-from-data-source-" + randomID(),
						repository: "empty",
						branch:     "master",
						space:      "root",
						provider: `bitbucket_cloud {
							namespace = "thespacelift"
							id = data.spacelift_bitbucket_cloud_integration.test.id
						}`,
						dataSource:     `data "spacelift_bitbucket_cloud_integration" "test" {}`,
						attribute:      "bitbucket_cloud.0.id",
						attributeValue: testConfig.SourceCode.BitbucketCloud.Default.ID,
					},
					{
						name:       "bitbucket-cloud-with-space-level-integration-from-data-source-" + randomID(),
						repository: "empty",
						branch:     "master",
						space:      testConfig.SourceCode.BitbucketCloud.SpaceLevel.Space,
						provider: `bitbucket_cloud {
							namespace = "thespacelift"
							id = data.spacelift_bitbucket_cloud_integration.test.id
						}`,
						dataSource: `data "spacelift_bitbucket_cloud_integration" "test" {
							id = "` + testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID + `"
						}`,
						attribute:      "bitbucket_cloud.0.id",
						attributeValue: testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID,
					},
					// Bitbucket Datacenter
					{
						name:       "bitbucket-datacenter-with-non-specified-integration-" + randomID(),
						repository: "tfprovider-test",
						branch:     "master",
						space:      "root",
						provider: `bitbucket_datacenter {
							namespace = "E2E"
						}`,
						attribute:      "bitbucket_datacenter.0.id",
						attributeValue: testConfig.SourceCode.BitbucketDatacenter.Default.ID,
					},
					{
						name:       "bitbucket-datacenter-with-an-empty-integration-id-" + randomID(),
						repository: "tfprovider-test",
						branch:     "master",
						space:      "root",
						provider: `bitbucket_datacenter {
							namespace = "E2E"
							id = ""
						}`,
						attribute:      "bitbucket_datacenter.0.id",
						attributeValue: testConfig.SourceCode.BitbucketDatacenter.Default.ID,
					},
					{
						name:       "bitbucket-datacenter-with-default-integration-" + randomID(),
						repository: "tfprovider-test",
						branch:     "master",
						space:      "root",
						provider: `bitbucket_datacenter {
							namespace = "E2E"
							id = "` + testConfig.SourceCode.BitbucketDatacenter.Default.ID + `"
						}`,
						attribute:      "bitbucket_datacenter.0.id",
						attributeValue: testConfig.SourceCode.BitbucketDatacenter.Default.ID,
					},
					{
						name:       "bitbucket-datacenter-with-space-level-integration-" + randomID(),
						repository: "tfprovider-test",
						branch:     "master",
						space:      testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.Space,
						provider: `bitbucket_datacenter {
							namespace = "E2E"
							id = "` + testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID + `"
						}`,
						attribute:      "bitbucket_datacenter.0.id",
						attributeValue: testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID,
					},
					{
						name:       "bitbucket-datacenter-with-default-integration-from-data-source-" + randomID(),
						repository: "tfprovider-test",
						branch:     "master",
						space:      "root",
						provider: `bitbucket_datacenter {
							namespace = "E2E"
							id = data.spacelift_bitbucket_datacenter_integration.test.id
						}`,
						dataSource:     `data "spacelift_bitbucket_datacenter_integration" "test" {}`,
						attribute:      "bitbucket_datacenter.0.id",
						attributeValue: testConfig.SourceCode.BitbucketDatacenter.Default.ID,
					},
					{
						name:       "bitbucket-datacenter-with-space-level-integration-from-data-source-" + randomID(),
						repository: "tfprovider-test",
						branch:     "master",
						space:      testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.Space,
						provider: `bitbucket_datacenter {
							namespace = "E2E"
							id = data.spacelift_bitbucket_datacenter_integration.test.id
						}`,
						dataSource: `data "spacelift_bitbucket_datacenter_integration" "test" {
							id = "` + testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID + `"
						}`,
						attribute:      "bitbucket_datacenter.0.id",
						attributeValue: testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID,
					},
					// GitHub Enterprise
					{
						name:       "github-with-non-specified-integration-" + randomID(),
						repository: "empty",
						branch:     "main",
						space:      "root",
						provider: `github_enterprise {
							namespace = "spacelift-ci-org"
						}`,
						attribute:      "github_enterprise.0.id",
						attributeValue: testConfig.SourceCode.GithubEnterprise.Default.ID,
					},
					{
						name:       "github-with-an-empty-integration-id-" + randomID(),
						repository: "empty",
						branch:     "main",
						space:      "root",
						provider: `github_enterprise {
							namespace = "spacelift-ci-org"
							id = ""
						}`,
						attribute:      "github_enterprise.0.id",
						attributeValue: testConfig.SourceCode.GithubEnterprise.Default.ID,
					},
					{
						name:       "github-with-default-integration-" + randomID(),
						repository: "empty",
						branch:     "main",
						space:      "root",
						provider: `github_enterprise {
							namespace = "spacelift-ci-org"
							id = "` + testConfig.SourceCode.GithubEnterprise.Default.ID + `"
						}`,
						attribute:      "github_enterprise.0.id",
						attributeValue: testConfig.SourceCode.GithubEnterprise.Default.ID,
					},
					{
						name:       "github-with-space-level-integration-" + randomID(),
						repository: "empty",
						branch:     "main",
						space:      testConfig.SourceCode.GithubEnterprise.SpaceLevel.Space,
						provider: `github_enterprise {
							namespace = "spacelift-ci-org"
							id = "` + testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID + `"
						}`,
						attribute:      "github_enterprise.0.id",
						attributeValue: testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID,
					},
					{
						name:       "github-with-default-integration-from-data-source-" + randomID(),
						repository: "empty",
						branch:     "main",
						space:      "root",
						provider: `github_enterprise {
							namespace = "spacelift-ci-org"
							id = data.spacelift_github_enterprise_integration.test.id
						}`,
						dataSource:     `data "spacelift_github_enterprise_integration" "test" {}`,
						attribute:      "github_enterprise.0.id",
						attributeValue: testConfig.SourceCode.GithubEnterprise.Default.ID,
					},
					{
						name:       "github-with-space-level-integration-from-data-source-" + randomID(),
						repository: "empty",
						branch:     "main",
						space:      testConfig.SourceCode.GithubEnterprise.SpaceLevel.Space,
						provider: `github_enterprise {
							namespace = "spacelift-ci-org"
							id = data.spacelift_github_enterprise_integration.test.id
						}`,
						dataSource: `data "spacelift_github_enterprise_integration" "test" {
							id = "` + testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID + `"
						}`,
						attribute:      "github_enterprise.0.id",
						attributeValue: testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID,
					},
					// GitLab
					{
						name:       "gitlab-with-non-specified-integration-" + randomID(),
						repository: "multimodule",
						branch:     "main",
						space:      "root",
						provider: `gitlab {
							namespace = "spacelift-ci"
						}`,
						attribute:      "gitlab.0.id",
						attributeValue: testConfig.SourceCode.Gitlab.Default.ID,
					},
					{
						name:       "gitlab-with-an-empty-integration-id-" + randomID(),
						repository: "multimodule",
						branch:     "main",
						space:      "root",
						provider: `gitlab {
							namespace = "spacelift-ci"
							id = ""
						}`,
						attribute:      "gitlab.0.id",
						attributeValue: testConfig.SourceCode.Gitlab.Default.ID,
					},
					{
						name:       "gitlab-with-default-integration-" + randomID(),
						repository: "multimodule",
						branch:     "main",
						space:      "root",
						provider: `gitlab {
							namespace = "spacelift-ci"
							id = "` + testConfig.SourceCode.Gitlab.Default.ID + `"
						}`,
						attribute:      "gitlab.0.id",
						attributeValue: testConfig.SourceCode.Gitlab.Default.ID,
					},
					{
						name:       "gitlab-with-space-level-integration-" + randomID(),
						repository: "multimodule",
						branch:     "main",
						space:      testConfig.SourceCode.Gitlab.SpaceLevel.Space,
						provider: `gitlab {
							namespace = "spacelift-ci"
							id = "` + testConfig.SourceCode.Gitlab.SpaceLevel.ID + `"
						}`,
						attribute:      "gitlab.0.id",
						attributeValue: testConfig.SourceCode.Gitlab.SpaceLevel.ID,
					},
					{
						name:       "gitlab-with-default-integration-from-data-source-" + randomID(),
						repository: "multimodule",
						branch:     "main",
						space:      "root",
						provider: `gitlab {
							namespace = "spacelift-ci"
							id = data.spacelift_gitlab_integration.test.id
						}`,
						dataSource:     `data "spacelift_gitlab_integration" "test" {}`,
						attribute:      "gitlab.0.id",
						attributeValue: testConfig.SourceCode.Gitlab.Default.ID,
					},
					{
						name:       "gitlab-with-space-level-integration-from-data-source-" + randomID(),
						repository: "multimodule",
						branch:     "main",
						space:      testConfig.SourceCode.Gitlab.SpaceLevel.Space,
						provider: `gitlab {
							namespace = "spacelift-ci"
							id = data.spacelift_gitlab_integration.test.id
						}`,
						dataSource: `data "spacelift_gitlab_integration" "test" {
							id = "` + testConfig.SourceCode.Gitlab.SpaceLevel.ID + `"
						}`,
						attribute:      "gitlab.0.id",
						attributeValue: testConfig.SourceCode.Gitlab.SpaceLevel.ID,
					},
				}

				for _, tc := range testCases {
					tc := tc
					t.Run(tc.name, func(t *testing.T) {
						config := fmt.Sprintf(`
							%s

							resource "`+resourceName+`" "test" {
								name                            = "%s"
								repository                      = "%s"
								branch                          = "%s"
								space_id                        = "%s"
								administrative                  = false
								%s
							}`, tc.dataSource, tc.name, tc.repository, tc.branch, tc.space, tc.provider)

						var tfstateSerial int64
						testSteps(t, []resource.TestStep{
							{
								Config: config,
								Check: func(tfstate *terraform.State) error {
									tfstateSerial = tfstate.Serial
									return Resource(resourceName+".test", Attribute(tc.attribute, Equals(tc.attributeValue)))(tfstate)
								},
							},
							{
								Config: config,
								Check: func(tfstate *terraform.State) error {
									// We need to check the serials to make sure nothing changed
									if serial := tfstate.Serial; serial != tfstateSerial {
										return fmt.Errorf("serials do not match: %d != %d", serial, tfstateSerial)
									}
									return nil
								},
							},
						})
					})
				}
			})

			t.Run("change ID", func(t *testing.T) {
				type testCaseStepAttribute struct {
					key   string
					value string
				}

				type testCaseStep struct {
					dataSource string
					provider   string
					attribute  testCaseStepAttribute
				}

				testCases := []struct {
					name       string
					repository string
					branch     string
					space      string
					steps      []testCaseStep
				}{
					// Azure Dev Ops
					{
						name:       "azure-devops-with-changing-vcs-id-" + randomID(),
						repository: "spacelift-ci",
						branch:     "main",
						space:      testConfig.SourceCode.AzureDevOps.SpaceLevel.Space,
						steps: []testCaseStep{
							{
								dataSource: "",
								provider: `azure_devops {
									project = "spacelift-ci"
								}`,
								attribute: testCaseStepAttribute{
									key:   "azure_devops.0.id",
									value: testConfig.SourceCode.AzureDevOps.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `azure_devops {
									project = "spacelift-ci"
									id = "` + testConfig.SourceCode.AzureDevOps.SpaceLevel.ID + `"
								}`,
								attribute: testCaseStepAttribute{
									key:   "azure_devops.0.id",
									value: testConfig.SourceCode.AzureDevOps.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_azure_devops_integration" "test" {
									id = "` + testConfig.SourceCode.AzureDevOps.SpaceLevel.ID + `"
								}`,
								provider: `azure_devops {
									project = "spacelift-ci"
									id = data.spacelift_azure_devops_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "azure_devops.0.id",
									value: testConfig.SourceCode.AzureDevOps.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_azure_devops_integration" "test" {}`,
								provider: `azure_devops {
									project = "spacelift-ci"
									id = data.spacelift_azure_devops_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "azure_devops.0.id",
									value: testConfig.SourceCode.AzureDevOps.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `azure_devops {
									project = "spacelift-ci"
								}`,
								attribute: testCaseStepAttribute{
									key:   "azure_devops.0.id",
									value: testConfig.SourceCode.AzureDevOps.Default.ID,
								},
							},
						},
					},
					// Bitbucket Cloud
					{
						name:       "bitbucket-cloud-with-changing-vcs-id-" + randomID(),
						repository: "empty",
						branch:     "master",
						space:      testConfig.SourceCode.BitbucketCloud.SpaceLevel.Space,
						steps: []testCaseStep{
							{
								dataSource: "",
								provider: `bitbucket_cloud {
									namespace = "thespacelift"
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_cloud.0.id",
									value: testConfig.SourceCode.BitbucketCloud.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `bitbucket_cloud {
									namespace = "thespacelift"
									id = "` + testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID + `"
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_cloud.0.id",
									value: testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_bitbucket_cloud_integration" "test" {
									id = "` + testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID + `"
								}`,
								provider: `bitbucket_cloud {
									namespace = "thespacelift"
									id = data.spacelift_bitbucket_cloud_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_cloud.0.id",
									value: testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_bitbucket_cloud_integration" "test" {}`,
								provider: `bitbucket_cloud {
									namespace = "thespacelift"
									id = data.spacelift_bitbucket_cloud_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_cloud.0.id",
									value: testConfig.SourceCode.BitbucketCloud.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `bitbucket_cloud {
									namespace = "thespacelift"
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_cloud.0.id",
									value: testConfig.SourceCode.BitbucketCloud.Default.ID,
								},
							},
						},
					},
					// Bitbucket Datacenter
					{
						name:       "bitbucket-datacenter-with-changing-vcs-id-" + randomID(),
						repository: "tfprovider-test",
						branch:     "master",
						space:      testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.Space,
						steps: []testCaseStep{
							{
								dataSource: "",
								provider: `bitbucket_datacenter {
									namespace = "E2E"
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_datacenter.0.id",
									value: testConfig.SourceCode.BitbucketDatacenter.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `bitbucket_datacenter {
									namespace = "E2E"
									id = "` + testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID + `"
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_datacenter.0.id",
									value: testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_bitbucket_datacenter_integration" "test" {
									id = "` + testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID + `"
								}`,
								provider: `bitbucket_datacenter {
									namespace = "E2E"
									id = data.spacelift_bitbucket_datacenter_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_datacenter.0.id",
									value: testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_bitbucket_datacenter_integration" "test" {}`,
								provider: `bitbucket_datacenter {
									namespace = "E2E"
									id = data.spacelift_bitbucket_datacenter_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_datacenter.0.id",
									value: testConfig.SourceCode.BitbucketDatacenter.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `bitbucket_datacenter {
									namespace = "E2E"
								}`,
								attribute: testCaseStepAttribute{
									key:   "bitbucket_datacenter.0.id",
									value: testConfig.SourceCode.BitbucketDatacenter.Default.ID,
								},
							},
						},
					},
					// GitHub Enterprise
					{
						name:       "github-enterprise-with-changing-vcs-id-" + randomID(),
						repository: "empty",
						branch:     "main",
						space:      testConfig.SourceCode.GithubEnterprise.SpaceLevel.Space,
						steps: []testCaseStep{
							{
								dataSource: "",
								provider: `github_enterprise {
									namespace = "spacelift-ci-org"
								}`,
								attribute: testCaseStepAttribute{
									key:   "github_enterprise.0.id",
									value: testConfig.SourceCode.GithubEnterprise.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `github_enterprise {
									namespace = "spacelift-ci-org"
									id = "` + testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID + `"
								}`,
								attribute: testCaseStepAttribute{
									key:   "github_enterprise.0.id",
									value: testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_github_enterprise_integration" "test" {
									id = "` + testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID + `"
								}`,
								provider: `github_enterprise {
									namespace = "spacelift-ci-org"
									id = data.spacelift_github_enterprise_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "github_enterprise.0.id",
									value: testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_github_enterprise_integration" "test" {}`,
								provider: `github_enterprise {
									namespace = "spacelift-ci-org"
									id = data.spacelift_github_enterprise_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "github_enterprise.0.id",
									value: testConfig.SourceCode.GithubEnterprise.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `github_enterprise {
									namespace = "spacelift-ci-org"
								}`,
								attribute: testCaseStepAttribute{
									key:   "github_enterprise.0.id",
									value: testConfig.SourceCode.GithubEnterprise.Default.ID,
								},
							},
						},
					},
					// GitLab
					{
						name:       "gitlab-with-changing-vcs-id-" + randomID(),
						repository: "multimodule",
						branch:     "main",
						space:      testConfig.SourceCode.Gitlab.SpaceLevel.Space,
						steps: []testCaseStep{
							{
								dataSource: "",
								provider: `gitlab {
									namespace = "spacelift-ci"
								}`,
								attribute: testCaseStepAttribute{
									key:   "gitlab.0.id",
									value: testConfig.SourceCode.Gitlab.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `gitlab {
									namespace = "spacelift-ci"
									id = "` + testConfig.SourceCode.Gitlab.SpaceLevel.ID + `"
								}`,
								attribute: testCaseStepAttribute{
									key:   "gitlab.0.id",
									value: testConfig.SourceCode.Gitlab.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_gitlab_integration" "test" {
									id = "` + testConfig.SourceCode.Gitlab.SpaceLevel.ID + `"
								}`,
								provider: `gitlab {
									namespace = "spacelift-ci"
									id = data.spacelift_gitlab_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "gitlab.0.id",
									value: testConfig.SourceCode.Gitlab.SpaceLevel.ID,
								},
							},
							{
								dataSource: `data "spacelift_gitlab_integration" "test" {}`,
								provider: `gitlab {
									namespace = "spacelift-ci"
									id = data.spacelift_gitlab_integration.test.id
								}`,
								attribute: testCaseStepAttribute{
									key:   "gitlab.0.id",
									value: testConfig.SourceCode.Gitlab.Default.ID,
								},
							},
							{
								dataSource: "",
								provider: `gitlab {
									namespace = "spacelift-ci"
								}`,
								attribute: testCaseStepAttribute{
									key:   "gitlab.0.id",
									value: testConfig.SourceCode.Gitlab.Default.ID,
								},
							},
						},
					},
				}

				for _, tc := range testCases {
					tc := tc
					t.Run(tc.name, func(t *testing.T) {
						var steps []resource.TestStep
						for i := range tc.steps {
							step := tc.steps[i]
							steps = append(steps, resource.TestStep{
								Config: fmt.Sprintf(`
									%s

									resource "`+resourceName+`" "test" {
										name                            = "%s"
										repository                      = "%s"
										branch                          = "%s"
										space_id                        = "%s"
										administrative                  = false
										%s
									}`, step.dataSource, tc.name, tc.repository, tc.branch, tc.space, step.provider),
								Check: Resource(resourceName+".test", Attribute(step.attribute.key, Equals(step.attribute.value))),
							})
						}

						testSteps(t, steps)
					})
				}
			})

			t.Run("mix providers", func(t *testing.T) {
				testSteps(t, []resource.TestStep{
					{
						Config: `
							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "spacelift-ci"
								branch                          = "main"
								space_id                        = "` + testConfig.SourceCode.AzureDevOps.SpaceLevel.Space + `"
								administrative                  = false
								azure_devops {
									project = "spacelift-ci"
									id = "` + testConfig.SourceCode.AzureDevOps.SpaceLevel.ID + `"
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("azure_devops.0.id", Equals(testConfig.SourceCode.AzureDevOps.SpaceLevel.ID))),
					},
					{
						Config: `
							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "empty"
								branch                          = "master"
								space_id                        = "` + testConfig.SourceCode.BitbucketCloud.SpaceLevel.Space + `"
								administrative                  = false
								bitbucket_cloud {
									namespace = "thespacelift"
									id = "` + testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID + `"
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("bitbucket_cloud.0.id", Equals(testConfig.SourceCode.BitbucketCloud.SpaceLevel.ID))),
					},
					{
						Config: `
							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "tfprovider-test"
								branch                          = "master"
								space_id                        = "` + testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.Space + `"
								administrative                  = false
								bitbucket_datacenter {
									namespace = "E2E"
									id = "` + testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID + `"
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("bitbucket_datacenter.0.id", Equals(testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.ID))),
					},
					{
						Config: `
							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "empty"
								branch                          = "main"
								space_id                        = "` + testConfig.SourceCode.GithubEnterprise.SpaceLevel.Space + `"
								administrative                  = false
								github_enterprise {
									namespace = "spacelift-ci-org"
									id = "` + testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID + `"
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("github_enterprise.0.id", Equals(testConfig.SourceCode.GithubEnterprise.SpaceLevel.ID))),
					},
					{
						Config: `
							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "multimodule"
								branch                          = "main"
								space_id                        = "` + testConfig.SourceCode.Gitlab.SpaceLevel.Space + `"
								administrative                  = false
								gitlab {
									namespace = "spacelift-ci"
									id = "` + testConfig.SourceCode.Gitlab.SpaceLevel.ID + `"
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("gitlab.0.id", Equals(testConfig.SourceCode.Gitlab.SpaceLevel.ID))),
					},
					{
						Config: `
							data "spacelift_azure_devops_integration" "test" {}
							
							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "spacelift-ci"
								branch                          = "main"
								space_id                        = "` + testConfig.SourceCode.AzureDevOps.SpaceLevel.Space + `"
								administrative                  = false
								azure_devops {
									project = "spacelift-ci"
									id = data.spacelift_azure_devops_integration.test.id
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("azure_devops.0.id", Equals(testConfig.SourceCode.AzureDevOps.Default.ID))),
					},
					{
						Config: `
							data "spacelift_bitbucket_cloud_integration" "test"		 {}

							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "empty"
								branch                          = "master"
								space_id                        = "` + testConfig.SourceCode.BitbucketCloud.SpaceLevel.Space + `"
								administrative                  = false
								bitbucket_cloud {
									namespace = "thespacelift"
									id = data.spacelift_bitbucket_cloud_integration.test.id
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("bitbucket_cloud.0.id", Equals(testConfig.SourceCode.BitbucketCloud.Default.ID))),
					},
					{
						Config: `
							data "spacelift_bitbucket_datacenter_integration" "test" {}
							
							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "tfprovider-test"
								branch                          = "master"
								space_id                        = "` + testConfig.SourceCode.BitbucketDatacenter.SpaceLevel.Space + `"
								administrative                  = false
								bitbucket_datacenter {
									namespace = "E2E"
									id = data.spacelift_bitbucket_datacenter_integration.test.id
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("bitbucket_datacenter.0.id", Equals(testConfig.SourceCode.BitbucketDatacenter.Default.ID))),
					},
					{
						Config: `
							data "spacelift_github_enterprise_integration" "test" {}

							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "empty"
								branch                          = "main"
								space_id                        = "` + testConfig.SourceCode.GithubEnterprise.SpaceLevel.Space + `"
								administrative                  = false
								github_enterprise {
									namespace = "spacelift-ci-org"
									id = data.spacelift_github_enterprise_integration.test.id
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("github_enterprise.0.id", Equals(testConfig.SourceCode.GithubEnterprise.Default.ID))),
					},
					{
						Config: `
							data "spacelift_gitlab_integration" "test" {}

							resource "` + resourceName + `" "test" {
								name                            = "mix-different-providers-` + randomID() + `"
								repository                      = "multimodule"
								branch                          = "main"
								space_id                        = "` + testConfig.SourceCode.Gitlab.SpaceLevel.Space + `"
								administrative                  = false
								gitlab {
									namespace = "spacelift-ci"
									id = data.spacelift_gitlab_integration.test.id
								}
							}
						`,
						Check: Resource(resourceName+".test", Attribute("gitlab.0.id", Equals(testConfig.SourceCode.Gitlab.Default.ID))),
					},
				})
			})
		})
	}
}
