# SPDX-License-Identifier: MIT

meta:
  id: tstime
  endian: le
  imports:
    - filetime

seq:
  - id: is_localized
    type: u1
  - id: pad
    size: 7
  - id: filetime
    type: filetime
