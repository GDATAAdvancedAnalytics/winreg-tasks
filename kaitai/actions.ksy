# SPDX-License-Identifier: MIT

meta:
  id: actions
  endian: le
  imports:
    - ./util/bstr

seq:
  - id: version
    type: u2

  - id: context # must match the principal_id of the JobBucket
    type: bstr
    if: version == 3

  - id: actions
    type: action
    repeat: eos

types:
  action:
    seq:
      - id: magic
        type: u2
      - id: id
        type: bstr
      - id: properties
        type:
          switch-on: magic
          cases:
            0x6666: exe_task_properties
            0x7777: com_handler_properties
            0x8888: email_task_properties
            0x9999: messagebox_task_properties

  com_handler_properties:
    seq:
      - id: clsid
        size: 16
      - id: data
        type: bstr

  exe_task_properties:
    seq:
      - id: command
        type: bstr
      - id: arguments
        type: bstr
      - id: working_directory
        type: bstr
      - id: flags
        type: u2
        if: _root.version == 3

  email_task_properties:
    seq:
      - id: from
        type: bstr
      - id: to
        type: bstr
      - id: cc
        type: bstr
      - id: bcc
        type: bstr
      - id: reply_to
        type: bstr
      - id: server
        type: bstr
      - id: subject
        type: bstr
      - id: body
        type: bstr
      - id: num_attachment_filenames
        type: u4
      - id: attachment_filenames
        type: bstr
        repeat: expr
        repeat-expr: num_attachment_filenames
      - id: num_headers
        type: u4
      - id: headers
        type: key_value_string
        repeat: expr
        repeat-expr: num_headers

  messagebox_task_properties:
    seq:
      - id: caption
        type: bstr
      - id: content
        type: bstr

  key_value_string:
    seq:
      - id: key
        type: bstr
      - id: value
        type: bstr
