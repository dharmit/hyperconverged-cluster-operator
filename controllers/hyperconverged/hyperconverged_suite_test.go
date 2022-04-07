package hyperconverged

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/kubevirt/hyperconverged-cluster-operator/controllers/commonTestUtils"
	hcoutil "github.com/kubevirt/hyperconverged-cluster-operator/pkg/util"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	pkgDirectory = "controllers/hyperconverged"
	testFilesLoc = "test-files"
)

func TestHyperconverged(t *testing.T) {
	RegisterFailHandler(Fail)

	var (
		testFilesLocation = getTestFilesLocation() + "/upgradePatches"
		destFile          string
	)

	getClusterInfo := hcoutil.GetClusterInfo

	BeforeSuite(func() {
		hcoutil.GetClusterInfo = func() hcoutil.ClusterInfo {
			return &commonTestUtils.ClusterInfoMock{}
		}

		wd, _ := os.Getwd()
		destFile = path.Join(wd, "upgradePatches.json")
		err := commonTestUtils.CopyFile(destFile, path.Join(testFilesLocation, "upgradePatches.json"))
		Expect(err).ToNot(HaveOccurred())

	})

	AfterSuite(func() {
		hcoutil.GetClusterInfo = getClusterInfo
		err := os.Remove(destFile)
		Expect(err).ToNot(HaveOccurred())
	})

	RunSpecs(t, "Hyperconverged Suite")
}

func getTestFilesLocation() string {
	wd, err := os.Getwd()
	Expect(err).ToNot(HaveOccurred())
	if strings.HasSuffix(wd, pkgDirectory) {
		return testFilesLoc
	}
	return path.Join(pkgDirectory, testFilesLoc)
}
