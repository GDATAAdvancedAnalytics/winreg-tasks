# SPDX-License-Identifier: MIT

meta:
  id: dynamic_info
  endian: le
seq:
  - id: magic
    size: 4
  - id: creation_time
    type: u8
  - id: last_run_time
    type: u8
  - id: task_state  # not used in recent Windows version; always zero
    type: u4
  - id: last_error_code
    type: u4
  - id: last_successful_run_time
    type: u8
