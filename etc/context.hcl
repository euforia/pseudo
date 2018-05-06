//
// This file contains global variables that are available i.e. in scope for a
// script.  The values in this file are of no consequence as this is just a
// structural reference where only keys are used.
//

region {
    id          = "region.id"
    datacenters = ["region.datacenters[0]"]
}

vcs {
    github {
        address = "git+ssh://github.com"
        org     = "org"
    }
}

registry {
    container {
        ecr {
            address = "registry.container.ecr.address"
        }
    }
}

platform {

    env {
        id              = "platform.env.id"
        name            = "platform.env.name"
        internal_domain = "platform.env.internal_domain"
        external_domain = "platform.env.internal_domain"
        dns             = ["platform.env.dns[0]"]
    }

    enclave {
        id = "enclave.id"
    }

}

app {
    name    = "app.name"
    version = "app.version"
    tags    = ["tags[0]"]
}
