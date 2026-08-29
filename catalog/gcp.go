package catalog

// gcpDisplayNames maps GCP service identifiers (sas.Technology.Service
// when Provider is "gcp") to their human-friendly product names.
var gcpDisplayNames = map[string]string{
	"cloud-run":      "Google Cloud Run",
	"gce":            "Google Compute Engine",
	"gke":            "Google Kubernetes Engine",
	"cloud-sql":      "Google Cloud SQL",
	"gcs":            "Google Cloud Storage",
	"pubsub":         "Google Pub/Sub",
	"bigquery":       "Google BigQuery",
	"cloud-tasks":    "Google Cloud Tasks",
	"secret-manager": "Google Secret Manager",
}
