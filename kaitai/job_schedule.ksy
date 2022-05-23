# SPDX-License-Identifier: MIT

meta:
  id: job_schedule
  endian: le
  imports:
    - ./util/bstr
    - ./util/filetime
    - ./util/tstime

seq:
  - id: start_boundary
    type: tstime
  - id: end_boundary
    type: tstime
  - id: unknown0
    type: tstime
  - id: repetition_interval_seconds
    type: u4
  - id: repetition_duration_seconds
    type: u4
  - id: execution_time_limit_seconds
    type: u4
  - id: mode
    type: u4
    enum: time_mode
  - id: data1
    type: u2
  - id: data2
    type: u2
  - id: data3
    type: u2
  - id: pad0
    type: u2
  - id: stop_tasks_at_duration_end
    type: u1
  - id: is_enabled
    type: u1
  - id: pad1
    type: u2
  - id: unknown1
    type: u4
  - id: max_delay_seconds
    type: u4
  - id: pad2
    type: u4

enums:
  time_mode:
    # run at <start_boundary>
    0: one_time

    # run at <start_boundary> and repeat every <data1> days
    1: daily

    # run on days of week <(data2 as day_of_week bitmap)> every <data1> weeks starting at <start_boundary>
    2: weekly

    # run in months <(data3 as months bitmap> on days <(data2:data1 as day in month bitmap)>
    # starting at <start_boundary>
    3: days_in_months

    # run in months <(data3 as months bitmap> in weeks <(data2 as week bitmap)>
    # on days <(data1 as day_of_week bitmap)> starting at <start_boundary>
    4: days_in_weeks_in_months

  day_of_week:
    0x1: sunday
    0x2: monday
    0x4: tuesday
    0x8: wednesday
    0x10: thursday
    0x20: friday
    0x40: saturday

  months:
    0x1: january
    0x2: february
    0x4: march
    0x8: april
    0x10: may
    0x20: june
    0x40: july
    0x80: august
    0x100: september
    0x200: october
    0x400: november
    0x800: december
