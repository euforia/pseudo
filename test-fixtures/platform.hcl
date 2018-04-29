
platform {

    env {
        dev {
            name            = "development"
            internal_domain = "internal.domain"
            external_domain = "external.domain"
            dns             = ["192.168.96.10", "192.168.98.10"]
        }
    }

    env prod {
        name            = "production"
        internal_domain = "internal.domain"
        external_domain = "external.domain"
        dns             = ["192.168.96.10", "192.168.98.10"]
    }

    enclave {
        test          {}
        sandbox       {}
    }

}
