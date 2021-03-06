package handler

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	nmstate "github.com/nmstate/kubernetes-nmstate/api/shared"
	nmstatev1alpha1 "github.com/nmstate/kubernetes-nmstate/api/v1alpha1"

	testenv "github.com/nmstate/kubernetes-nmstate/test/env"
)

var _ = Describe("NodeNetworkConfigurationPolicy upgrade", func() {
	Context("when v1alpha1 is populated", func() {
		BeforeEach(func() {
			policy := nmstatev1alpha1.NodeNetworkConfigurationPolicy{
				ObjectMeta: metav1.ObjectMeta{
					Name: TestPolicy,
				},
				Spec: nmstate.NodeNetworkConfigurationPolicySpec{
					DesiredState: linuxBrUp(bridge1),
				},
			}
			Expect(testenv.Client.Create(context.TODO(), &policy)).To(Succeed(), "should success creating a v1alpha1 nncp")
		})
		AfterEach(func() {
			updateDesiredStateAndWait(linuxBrAbsent(bridge1))
			resetDesiredStateForNodes()
		})
		It("should be stored as v1beta1 and end with available state", func() {
			waitForAvailableTestPolicy()
		})
	})

})
