package authentication_operator

import (
	component "github.com/openshift/hypershift/support/controlplane-component"
	"github.com/openshift/hypershift/support/util"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func adaptDeployment(cpContext component.WorkloadContext, deployment *appsv1.Deployment) error {
	util.UpdateContainer(ComponentName, deployment.Spec.Template.Spec.Containers, func(c *corev1.Container) {
		// Replace environment variables with actual image references from release
		envReplacer := newEnvironmentReplacer(cpContext.ReleaseImageProvider, cpContext.UserReleaseImageProvider)
		envReplacer.replaceEnvVars(c.Env)
	})

	return nil
}
