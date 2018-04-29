
// This file contains resource configurations within your org.

regions {
    us-west-2 {
        datacenters = ["us-west-2a", "us-west-2b", "us-west-2c"]
    }
    # us-east-1 {
    #     datacenters = ["us-east-1a", "us-east-1b", "us-east-1c", "us-east-1d", "us-east-1e"]
    # }
}

vcs {
    github {
        address = "git+ssh://github.com"
        org     = "org"
    }
}

registry {
    vm {}

    container {
        ecr {
            address = "1234567891011.dkr.ecr.us-west-2.amazonaws.com"
        }
        docker {
            address = "docker.io"
        }
    }

    deployment custom {
        address = "custom.${platform.env.internal_domain}:9999"
    }
}

orchestration {
    nomad {
        address = "http://nomad.${platform.env.internal_domain}:4646"
    }
}

secrets {
    vault {
        address = "https://vault.${platform.env.internal_domain}:8200"
    }
}
