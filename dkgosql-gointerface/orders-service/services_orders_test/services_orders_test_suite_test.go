package services_orders_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestServicesOrdersTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ServicesOrdersTest Suite")
}
