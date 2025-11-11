package authentication_operator

import (
	hyperv1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
	oapiv2 "github.com/openshift/hypershift/control-plane-operator/controllers/hostedcontrolplane/v2/oapi"
	component "github.com/openshift/hypershift/support/controlplane-component"
	"github.com/openshift/hypershift/support/util"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	ComponentName = "authentication-operator"
)

var _ component.ComponentOptions = &authenticationOperator{}

type authenticationOperator struct {
}

// IsRequestServing implements controlplanecomponent.ComponentOptions.
func (a *authenticationOperator) IsRequestServing() bool {
	return false
}

// MultiZoneSpread implements controlplanecomponent.ComponentOptions.
func (a *authenticationOperator) MultiZoneSpread() bool {
	return false
}

// NeedsManagementKASAccess implements controlplanecomponent.ComponentOptions.
func (a *authenticationOperator) NeedsManagementKASAccess() bool {
	return true
}

func NewComponent() component.ControlPlaneComponent {
	return component.NewDeploymentComponent(ComponentName, &authenticationOperator{}).
		WithAdaptFunction(adaptDeployment).
		WithPredicate(isOAuthEnabled).
		WithManifestAdapter(
			"role.yaml",
			component.DisableIfAnnotationExist(hyperv1.DisablePKIReconciliationAnnotation),
		).
		WithManifestAdapter(
			"rolebinding.yaml",
			component.DisableIfAnnotationExist(hyperv1.DisablePKIReconciliationAnnotation),
		).
		WithManifestAdapter(
			"serviceaccount.yaml",
		).
		WithDependencies(oapiv2.ComponentName).
		InjectAvailabilityProberContainer(util.AvailabilityProberOpts{
			KubeconfigVolumeName: "guest-kubeconfig",
			// Note: Don't check for Authentication API here as the operator itself installs that CRD
			RequiredAPIs: []schema.GroupVersionKind{},
		}).
		Build()
}

func isOAuthEnabled(cpContext component.WorkloadContext) (bool, error) {
	// Enable authentication-operator when OAuth is enabled
	return util.HCPOAuthEnabled(cpContext.HCP), nil
}
