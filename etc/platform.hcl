//
// This file contains configurations for your platform.
//
// The environment layout is broken into the platform and the application
// such that application environments have the option to be logical
//
// APPLICATION  test|dev|qa|qa1 | stg|live
// -----------------------------+---------
// PLATFORM     dev             | live
//

platform {
    // Underlying platform environments. These are different from the logical
    // application environments.  These complete independent setups of the
    // underlying infrastructure
    env {
        dev {
            name            = "development"
            internal_domain = "pseudo.dev"
            external_domain = "pseudo.de"
            dns             = ["4.2.2.2", "8.8.8.8"]
        }

        prod {
            name            = "production"
            internal_domain = "pseudo.prod"
            external_domain = "pseudo.io"
            dns             = ["8.8.8.8", "4.2.2.2"]
        }
    }

    // Network enclaves or isolation zones sample
    enclave {
        pci  { cidrs = [] }
        hipa { cidrs = [] }
        mpaa { cidrs = [] }
        sox  { cidrs = [] }
    }

}
