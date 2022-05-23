# SPDX-License-Identifier: MIT

meta:
  id: tstimeperiod
  endian: le

seq:
  - id: year
    type: u2
  - id: month
    type: u2
  - id: week
    type: u2
  - id: day  # if used in conjunction with week this is "day of week"
    type: u2
  - id: hour
    type: u2
  - id: minute
    type: u2
  - id: second
    type: u2
