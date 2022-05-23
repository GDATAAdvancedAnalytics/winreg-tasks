# SPDX-License-Identifier: MIT

meta:
  id: user_info
  endian: le
  imports:
    - ./aligned/u1
    - ./aligned/u4
    - ./aligned/buffer
    - ./aligned/bstr

seq:
  - id: skip_user
    type: aligned_u1
  - id: skip_sid
    type: aligned_u1
    if: skip_user.value == 0
  - id: sid_type
    type: aligned_u4
    if: skip_user.value == 0 and skip_sid.value == 0
  - id: sid
    type: aligned_buffer
    if: skip_user.value == 0 and skip_sid.value == 0
  - id: username
    type: aligned_bstr
    if: skip_user.value == 0

enums:
  sid_type: # SID_NAME_USE enum in WinAPI
    1: user
    2: group
    3: domain
    4: alias
    5: well_known_group
    6: deleted_account
    7: invalid
    8: unknown
    9: computer
    10: label
    11: logon_session
