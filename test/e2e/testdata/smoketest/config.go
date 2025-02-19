package smoketest

import (
	"log"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/apecloud/kubeblocks/test/e2e"
	e2eutil "github.com/apecloud/kubeblocks/test/e2e/util"
)

func Config() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	Context("Configure running e2e information", func() {
		It("create a secret to save the access key and ", func() {
			accessKey := e2eutil.ExecCommand("aws configure get aws_access_key_id")
			secretKey := e2eutil.ExecCommand("aws configure get aws_secret_access_key")
			createSecret := "kubectl create secret generic " + ConfigType + "-credential-for-backuprepo \\\n" +
				"  -n kb-system \\\n" +
				"  --from-literal=accessKeyId=" + e2eutil.StringStrip(accessKey) + " \\\n" +
				"  --from-literal=secretAccessKey=" + e2eutil.StringStrip(secretKey)
			b := e2eutil.ExecuteCommand(createSecret)
			Expect(b).Should(BeTrue())
		})

		It(" configure backup-repo", func() {
			var yaml string
			if ConfigType == "oss" {
				yaml = dir + "/testdata/config/backuprepo_oss.yaml"
			} else {
				yaml = dir + "/testdata/config/backuprepo_s3.yaml"
			}
			b := e2eutil.OpsYaml(yaml, "create")
			Expect(b).Should(BeTrue())
		})
	})
}

func DeleteConfig() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	Context("delete e2e config resources", func() {
		It("delete secret and backuprepo", func() {
			deleteSecret := e2eutil.ExecuteCommand("kubectl delete secret " + ConfigType + "-credential-for-backuprepo -n kb-system")
			Expect(deleteSecret).Should(BeTrue())
			deleteBr := e2eutil.ExecuteCommand("kubectl delete backuprepo my-repo")
			Expect(deleteBr).Should(BeTrue())
		})

	})
}
