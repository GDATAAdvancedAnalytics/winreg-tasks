# SPDX-License-Identifier: MIT

meta:
  id: triggers
  endian: le
  imports:
    - ./aligned/bstr
    - ./aligned/bstr_expand_size
    - ./aligned/buffer
    - ./aligned/u1
    - ./aligned/u4
    - ./util/bstr
    - ./util/filetime
    - ./util/tstime
    - ./util/tstimeperiod
    - ./job_schedule
    - ./optional_settings
    - ./user_info

seq:
  - id: header
    type: header

  - id: job_bucket
    type: job_bucket

  - id: triggers
    repeat: eos
    type: trigger

enums:
  job_bucket_flags:
    # 0x1: flags_0x1
    0x2: run_only_if_idle
    0x4: restart_on_idle
    0x8: stop_on_idle_end
    0x10: disallow_start_if_on_batteries
    0x20: stop_if_going_on_batteries
    0x40: start_when_available
    0x80: run_only_if_network_available
    0x100: allow_start_on_demand
    0x200: wake_to_run
    0x400: execute_parallel
    0x800: execute_stop_existing
    0x1000: execute_queue
    0x2000: execute_ignore_new
    0x4000: logon_type_s4u
    # 0x8000: flags_0x8000
    0x10000: logon_type_interactivetoken
    # 0x20000: flags_0x20000
    0x40000: logon_type_password
    0x80000: logon_type_interactivetokenorpassword
    # 0x100000: flags_0x100000
    # 0x200000: flags_0x200000
    0x400000: enabled
    0x800000: hidden
    0x1000000: runlevel_highest_available
    0x2000000: task
    0x4000000: version
    0x8000000: token_sid_type_none
    0x10000000: token_sid_type_unrestricted
    0x20000000: interval
    0x40000000: allow_hard_terminate

  session_state:
    1: console_connect
    2: console_disconnect
    3: remote_connect
    4: remote_disconnect
    5: session_lock
    6: session_unlock

types:
  header:
    seq:
      - id: version
        type: aligned_u1

      - id: start_boundary  # the earliest startBoundary of all triggers
        type: tstime

      - id: end_boundary  # the latest endBoundary of all triggers
        type: tstime

  job_bucket:
    seq:
      - id: flags
        type: aligned_u4
      - id: crc32  # the crc32 hash of the task XML
        type: aligned_u4
      - id: principal_id
        type: aligned_bstr
        if: _root.header.version.value >= 0x16
      - id: display_name
        type: aligned_bstr
        if: _root.header.version.value >= 0x17
      - id: user_info
        type: user_info
      - id: optional_settings
        type: optional_settings

  trigger:
    seq:
      - id: magic
        type: aligned_u4

      - id: properties
        type:
          switch-on: magic.value
          cases:
            0x6666: wnf_state_change_trigger
            0x7777: session_change_trigger
            0x8888: registration_trigger
            0xaaaa: logon_trigger
            0xcccc: event_trigger
            0xdddd: time_trigger
            0xeeee: idle_trigger
            0xffff: boot_trigger

  generic_trigger_data:
    seq:
      - id: start_boundary
        type: tstime
      - id: end_boundary
        type: tstime
      - id: delay_seconds
        type: u4
      - id: timeout_seconds
        type: u4
      - id: repetition_interval_seconds
        type: u4
      - id: repetition_duration_seconds
        type: u4
      - id: repetition_duration_seconds_2
        type: u4
      - id: stop_at_duration_end
        type: u1
      - id: padding
        size: 3
      - id: enabled
        type: aligned_u1
      - id: unknown
        size: 8
      - id: trigger_id
        type: bstr
        if: _root.header.version.value >= 0x16
      - id: pad_to_block
        size: (8 - (trigger_id.len + 4)) % 8
        if: _root.header.version.value >= 0x16

  wnf_state_change_trigger:
    seq:
      - id: generic_data
        type: generic_trigger_data
      - id: state_name
        size: 8
      - id: len_data
        type: aligned_u4
      - id: data
        size: len_data.value

  session_change_trigger:
    seq:
      - id: generic_data
        type: generic_trigger_data
      - id: state_change
        type: u4
        enum: session_state
      - id: padding
        size: 4
      - id: user
        type: user_info

  registration_trigger:
    seq:
      - id: generic_data
        type: generic_trigger_data

  logon_trigger:
    seq:
      - id: generic_data
        type: generic_trigger_data
      - id: user
        type: user_info

  event_trigger:
    seq:
      - id: generic_data
        type: generic_trigger_data
      - id: subscription
        type: aligned_bstr_expand_size
      - id: unknown0
        type: u4
      - id: unknown1
        type: u4
      - id: unknown2
        type: aligned_bstr_expand_size
      - id: len_value_queries
        type: aligned_u4
      - id: value_queries
        type: value_query
        repeat: expr
        repeat-expr: len_value_queries.value

  time_trigger:
    seq:
      - id: job_schedule
        type: job_schedule
      - id: trigger_id
        type: bstr
        if: _root.header.version.value >= 0x16
      - id: padding
        size: (8 - (trigger_id.len + 4)) % 8
        if: _root.header.version.value >= 0x16

  idle_trigger:
    seq:
      - id: generic_data
        type: generic_trigger_data

  boot_trigger:
    seq:
      - id: generic_data
        type: generic_trigger_data

  value_query:
    seq:
      - id: name
        type: aligned_bstr_expand_size
      - id: value
        type: aligned_bstr_expand_size
