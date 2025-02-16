/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import Component from '@glimmer/component';
import { inject as service } from '@ember/service';

export default class HashiCorpConsul extends Component {
  @service('flashMessages') flashMessages;
  @service('env') env;

  get consulVersion() {
    const suffix = !['', 'oss'].includes(this.env.var('CONSUL_BINARY_TYPE')) ? '+ent' : '';
    return `${this.env.var('CONSUL_VERSION')}${suffix}`;
  }
}
