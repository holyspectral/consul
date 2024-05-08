# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: BUSL-1.1

# This manifest file describes active releases and is consumed by the backport tooling.

schema = 1
active_versions {
  version "1.18" {
    ce_active = true
    lts       = true
  }
  version "1.17" {}
  version "1.16" {}
  version "1.15" {
    lts = true
  }
}
