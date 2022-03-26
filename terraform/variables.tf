variable project_name {
  type = string
  description = "GCP Project ID"
  nullable = false
}

variable location {
  type = string
  description = "Valid GCP location"
  nullable = false
}

variable auth_name_key {
  type = string
}

variable auth_name_value {
  type = string
}

variable auth_token_key {
  type = string
}

variable auth_token_value {
  type = string
}

variable shipper_container_id {
  type = string
  nullable = false
}

variable forwarder_container_id {
  type = string
  nullable = false
}

variable gha_container_id {
  type = string
  nullable = false
}

variable carrier_container_id {
  type = string
  nullable = false
}

// Optional config
variable shipper_service_name {
  type = string
  default = "shipper"
  nullable = false
}

variable forwarder_service_name {
  type = string
  default = "forwarder"
  nullable = false
}

variable gha_service_name {
  type = string
  default = "gha"
  nullable = false
}

variable carrier_service_name {
  type = string
  default = "carrier"
  nullable = false
}