provider "google" {
  project = var.project_name
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}




// Shipper
resource "google_cloud_run_service" "shipper" {
  name     = var.shipper_service_name
  location = var.location

  template {
    spec {
      containers {
        image = var.shipper_container_id
        env {
          name = "SRV_ROLE"
          value = "shipper"
        }
        env {
          name = "PATH_TEMPLATE"
          value = "web/public/templates"
        }
        env {
          name = "PATH_STATIC"
          value = "web/public/static/"
        }
        env {
          name = "PROJECT_ID"
          value = var.project_name
        }
        env {
          name = "AUTH_NAME_KEY"
          value = var.auth_name_key
        }
        env {
          name = "AUTH_NAME_VALUE"
          value = var.auth_name_value
        }
        env {
          name = "AUTH_TOKEN_KEY"
          value = var.auth_token_key
        }
        env {
          name = "AUTH_TOKEN_VALUE"
          value = var.auth_token_value
        }
      }
    }
  }

  traffic {
    percent = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.shipper.location
  project = google_cloud_run_service.shipper.project
  service = google_cloud_run_service.shipper.name
  policy_data = data.google_iam_policy.noauth.policy_data
}

output "shipper_url" {
  value = google_cloud_run_service.shipper.status[0].url
}




// Forwarder
resource "google_cloud_run_service" "forwarder" {
  name     = var.shipper_service_name
  location = var.location

  template {
    spec {
      containers {
        image = var.shipper_container_id
        env {
          name = "SRV_ROLE"
          value = "forwarder"
        }
        env {
          name = "PATH_TEMPLATE"
          value = "web/public/templates"
        }
        env {
          name = "PATH_STATIC"
          value = "web/public/static/"
        }
        env {
          name = "PROJECT_ID"
          value = var.project_name
        }
        env {
          name = "AUTH_NAME_KEY"
          value = var.auth_name_key
        }
        env {
          name = "AUTH_NAME_VALUE"
          value = var.auth_name_value
        }
        env {
          name = "AUTH_TOKEN_KEY"
          value = var.auth_token_key
        }
        env {
          name = "AUTH_TOKEN_VALUE"
          value = var.auth_token_value
        }
      }
    }
  }

  traffic {
    percent = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.forwarder.location
  project = google_cloud_run_service.forwarder.project
  service = google_cloud_run_service.forwarder.name
  policy_data = data.google_iam_policy.noauth.policy_data
}

output "forwarder_url" {
  value = google_cloud_run_service.forwarder.status[0].url
}




// GHA
resource "google_cloud_run_service" "gha" {
  name     = var.shipper_service_name
  location = var.location

  template {
    spec {
      containers {
        image = var.shipper_container_id
        env {
          name = "SRV_ROLE"
          value = "gha"
        }
        env {
          name = "PATH_TEMPLATE"
          value = "web/public/templates"
        }
        env {
          name = "PATH_STATIC"
          value = "web/public/static/"
        }
        env {
          name = "PROJECT_ID"
          value = var.project_name
        }
        env {
          name = "AUTH_NAME_KEY"
          value = var.auth_name_key
        }
        env {
          name = "AUTH_NAME_VALUE"
          value = var.auth_name_value
        }
        env {
          name = "AUTH_TOKEN_KEY"
          value = var.auth_token_key
        }
        env {
          name = "AUTH_TOKEN_VALUE"
          value = var.auth_token_value
        }
      }
    }
  }

  traffic {
    percent = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.gha.location
  project = google_cloud_run_service.gha.project
  service = google_cloud_run_service.gha.name
  policy_data = data.google_iam_policy.noauth.policy_data
}

output "gha_url" {
  value = google_cloud_run_service.gha.status[0].url
}



// Carrier
resource "google_cloud_run_service" "carrier" {
  name     = var.shipper_service_name
  location = var.location

  template {
    spec {
      containers {
        image = var.shipper_container_id
        env {
          name = "SRV_ROLE"
          value = "carrier"
        }
        env {
          name = "PATH_TEMPLATE"
          value = "web/public/templates"
        }
        env {
          name = "PATH_STATIC"
          value = "web/public/static/"
        }
        env {
          name = "PROJECT_ID"
          value = var.project_name
        }
        env {
          name = "AUTH_NAME_KEY"
          value = var.auth_name_key
        }
        env {
          name = "AUTH_NAME_VALUE"
          value = var.auth_name_value
        }
        env {
          name = "AUTH_TOKEN_KEY"
          value = var.auth_token_key
        }
        env {
          name = "AUTH_TOKEN_VALUE"
          value = var.auth_token_value
        }
      }
    }
  }

  traffic {
    percent = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.carrier.location
  project = google_cloud_run_service.carrier.project
  service = google_cloud_run_service.carrier.name
  policy_data = data.google_iam_policy.noauth.policy_data
}

output "carrier_url" {
  value = google_cloud_run_service.carrier.status[0].url
}
