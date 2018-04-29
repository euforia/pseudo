job "myjob" {
  region      = "${region.id}"
  datacenters = ["${region.datacenters[0]}"]
  type        = "service"

  constraint {
      attribute = "$${meta.foobar}"
      value = "baz"
  }

  constraint {
      attribute = "$${meta.env_type}"
      value = "${platform.env.id}"
  }

  update {
   stagger      = "20s"
   max_parallel = 1
  }

  group "stack" {
    count = 2

    constraint {
      operator = "distinct_hosts"
    }

    constraint {
      operator  = "distinct_property"
      attribute = "$${attr.platform.aws.placement.availability-zone}"
    }

    vault {
      change_mode = "noop"
      env = false
      policies = ["read-secrets"]
    }

    task "api" {

      driver = "docker"

      config {
        image = "${app.name}:${app.version}"
        port_map {
            default = 9000
        }
        labels {
            service = "$${NOMAD_JOB_NAME}"
        }
        logging {
          type = "syslog"
          config {
            tag = "$${NOMAD_JOB_NAME}-$${NOMAD_TASK_NAME}"
          }
        }
      }

      env {
        APP_VERSION   = "${app.version}"
        AWS_REGION    = "${region.id}"
        AWS_BUCKET    = "bucket-${tld(platform.env.internal_domain)}"
        NOMAD_SERVER  = "nomad.${platform.env.internal_domain}:4646"
      }

      service {
        name = "$${JOB}-$${TASK}"
        port = "default"
        tags = [
            "alert=$${NOMAD_JOB_NAME}",
            "${replace(app.version, ".", "-")}",
            "${app.tags[0]}",
            "latest"
        ]
        check {
          type     = "http"
          path     = "/healthz"
          interval = "10s"
          timeout  = "3s"
        }
      }

      template {
        data = <<EOF
{{ with printf "secret/%s" (env "NOMAD_JOB_NAME") | secret }}{{ range $k, $v := .Data }}{{ $k }}={{ $v }}
{{ end }}
        DB_NAME = "{{.Data.DB_NAME}}"
{{ end }}
EOF
        destination = ".env"
        env = true
      }

      resources {
        cpu    = 100 # MHz
        memory = 128 # MB

        network {
          mbits = 2
          port "default" {}
        }
      }
    }
  }
}
