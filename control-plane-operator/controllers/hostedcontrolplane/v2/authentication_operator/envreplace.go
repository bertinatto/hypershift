package authentication_operator

import (
	"github.com/openshift/hypershift/control-plane-operator/controllers/hostedcontrolplane/imageprovider"
	corev1 "k8s.io/api/core/v1"
)

var (
	// map env. variable in authentication-operator Deployment -> name of the image in payload.
	operatorImageRefs = map[string]string{
		"IMAGE_OAUTH_SERVER": "oauth-server",
	}
)

type environmentReplacer struct {
	// map env variable name -> new env. variable value
	values map[string]string
}

func newEnvironmentReplacer(releaseImageProvider, userReleaseImageProvider imageprovider.ReleaseImageProvider) *environmentReplacer {
	er := &environmentReplacer{values: map[string]string{}}

	version := userReleaseImageProvider.Version()
	er.setVersions(version)
	er.setOperatorImageReferences(releaseImageProvider, userReleaseImageProvider)

	return er
}

func (er *environmentReplacer) setOperatorImageReferences(releaseImageProvider, userReleaseImageProvider imageprovider.ReleaseImageProvider) {
	// For authentication-operator, we use images from the management cluster release
	for envVar, payloadName := range operatorImageRefs {
		if imageURL, ok := releaseImageProvider.ImageExist(payloadName); ok {
			er.values[envVar] = imageURL
		}
	}
}

func (er *environmentReplacer) setVersions(ver string) {
	er.values["OPERATOR_IMAGE_VERSION"] = ver
	er.values["OPERAND_IMAGE_VERSION"] = ver
}

func (er *environmentReplacer) replaceEnvVars(envVars []corev1.EnvVar) {
	for i := range envVars {
		if value, ok := er.values[envVars[i].Name]; ok {
			envVars[i].Value = value
		}
	}
}
