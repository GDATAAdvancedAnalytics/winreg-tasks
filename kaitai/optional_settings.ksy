# SPDX-License-Identifier: MIT

meta:
  id: optional_settings
  endian: le
  imports:
    - ./aligned/u4
    - ./util/tstimeperiod

seq:
  - id: len
    type: aligned_u4
  - id: idle_duration_seconds
    type: u4
    if: len.value > 0
  - id: idle_wait_timeout_seconds
    type: u4
    if: len.value > 0
  - id: execution_time_limit_seconds
    type: u4
    if: len.value > 0
  - id: delete_expired_task_after
    type: u4
    if: len.value > 0
  - id: priority
    type: u4
    if: len.value > 0
  - id: restart_on_failure_delay
    type: u4
    if: len.value > 0
  - id: restart_on_failure_retries
    type: u4
    if: len.value > 0
  - id: network_id
    size: 16
    if: len.value > 0
  - id: padding0
    size: 4
    if: len.value > 0
  - id: privileges
    type: u8
    enum: privilege
    if: len.value == 0x38 or len.value == 0x58
  - id: periodicity
    type: tstimeperiod
    if: len.value == 0x58
  - id: deadline
    type: tstimeperiod
    if: len.value == 0x58
  - id: exclusive
    type: u1
    if: len.value == 0x58
  - id: padding1
    size: 3
    if: len.value == 0x58

enums:
  privilege:
    0x4: se_create_token_privilege
    0x8: se_assign_primary_token_privilege
    0x10: se_lock_memory_privilege
    0x20: se_increase_quota_privilege
    0x40: se_machine_account_privilege
    0x80: se_tcb_privilege
    0x100: se_security_privilege
    0x200: se_take_ownership_privilege
    0x400: se_load_driver_privilege
    0x800: se_system_profile_privilege
    0x1000: se_systemtime_privilege
    0x2000: se_profile_single_process_privilege
    0x4000: se_increase_base_priority_privilege
    0x8000: se_create_pagefile_privilege
    0x10000: se_create_permanent_privilege
    0x20000: se_backup_privilege
    0x40000: se_restore_privilege
    0x80000: se_shutdown_privilege
    0x100000: se_debug_privilege
    0x200000: se_audit_privilege
    0x400000: se_system_environment_privilege
    0x800000: se_change_notify_privilege
    0x1000000: se_remote_shutdown_privilege
    0x2000000: se_undock_privilege
    0x4000000: se_sync_agent_privilege
    0x8000000: se_enable_delegation_privilege
    0x10000000: se_manage_volume_privilege
    0x20000000: se_impersonate_privilege
    0x40000000: se_create_global_privilege
    0x80000000: se_trusted_cred_man_access_privilege
    0x100000000: se_relabel_privilege
    0x200000000: se_increase_working_set_privilege
    0x400000000: se_time_zone_privilege
    0x800000000: se_create_symbolic_link_privilege
    0x1000000000: se_delegate_session_user_impersonate_privilege
