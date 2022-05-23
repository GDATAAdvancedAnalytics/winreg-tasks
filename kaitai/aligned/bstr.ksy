# SPDX-License-Identifier: MIT

meta:
  id: aligned_bstr
  endian: le
  imports:
    - ./u4

seq:
  - id: byte_count
    type: aligned_u4
  - id: string
    size: byte_count.value
    encoding: utf-16le
    type: str
  - id: block_padding
    size: (8 - (byte_count.value % 8)) % 8
