# SPDX-License-Identifier: MIT

meta:
  id: bstr
  endian: le

seq:
  - id: len
    type: u4
  - id: str
    type: str
    encoding: utf-16le
    size: len
