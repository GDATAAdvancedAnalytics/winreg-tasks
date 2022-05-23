# SPDX-License-Identifier: MIT

meta:
  id: aligned_bstr_expand_size
  endian: le
  imports:
   - ./u4

seq:
  - id: string_length
    type: aligned_u4
  - id: content
    size: byte_count
    if: string_length.value > 0
    encoding: utf-16le
    type: str
  - id: padding
    size: (8 - (byte_count % 8)) % 8
    if: string_length.value > 0

instances:
  byte_count:
    value: string_length.value * 2 + 2
