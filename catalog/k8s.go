package catalog

// k8sDisplayNames maps Kubernetes resource kinds (sas.Technology.Service
// when Provider is "k8s") to their human-friendly names.
var k8sDisplayNames = map[string]string{
	"deployment":  "Kubernetes Deployment",
	"statefulset": "Kubernetes StatefulSet",
	"daemonset":   "Kubernetes DaemonSet",
	"service":     "Kubernetes Service",
	"ingress":     "Kubernetes Ingress",
	"configmap":   "Kubernetes ConfigMap",
	"secret":      "Kubernetes Secret",
	"job":         "Kubernetes Job",
	"cronjob":     "Kubernetes CronJob",
}
