# SPDX-License-Identifier: MIT

meta:
  id: aligned_buffer
  endian: le
  imports:
    - ./u4

seq:
  - id: size
    type: aligned_u4
  - id: data
    size: size.value
  - id: padding
    size: (8 - (size.value % 8)) % 8
