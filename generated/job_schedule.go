// Code generated by kaitai-struct-compiler from a .ksy source file. DO NOT EDIT.

package generated

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"


type JobSchedule_TimeMode int
const (
	JobSchedule_TimeMode__OneTime JobSchedule_TimeMode = 0
	JobSchedule_TimeMode__Daily JobSchedule_TimeMode = 1
	JobSchedule_TimeMode__Weekly JobSchedule_TimeMode = 2
	JobSchedule_TimeMode__DaysInMonths JobSchedule_TimeMode = 3
	JobSchedule_TimeMode__DaysInWeeksInMonths JobSchedule_TimeMode = 4
)

type JobSchedule_DayOfWeek int
const (
	JobSchedule_DayOfWeek__Sunday JobSchedule_DayOfWeek = 1
	JobSchedule_DayOfWeek__Monday JobSchedule_DayOfWeek = 2
	JobSchedule_DayOfWeek__Tuesday JobSchedule_DayOfWeek = 4
	JobSchedule_DayOfWeek__Wednesday JobSchedule_DayOfWeek = 8
	JobSchedule_DayOfWeek__Thursday JobSchedule_DayOfWeek = 16
	JobSchedule_DayOfWeek__Friday JobSchedule_DayOfWeek = 32
	JobSchedule_DayOfWeek__Saturday JobSchedule_DayOfWeek = 64
)

type JobSchedule_Months int
const (
	JobSchedule_Months__January JobSchedule_Months = 1
	JobSchedule_Months__February JobSchedule_Months = 2
	JobSchedule_Months__March JobSchedule_Months = 4
	JobSchedule_Months__April JobSchedule_Months = 8
	JobSchedule_Months__May JobSchedule_Months = 16
	JobSchedule_Months__June JobSchedule_Months = 32
	JobSchedule_Months__July JobSchedule_Months = 64
	JobSchedule_Months__August JobSchedule_Months = 128
	JobSchedule_Months__September JobSchedule_Months = 256
	JobSchedule_Months__October JobSchedule_Months = 512
	JobSchedule_Months__November JobSchedule_Months = 1024
	JobSchedule_Months__December JobSchedule_Months = 2048
)
type JobSchedule struct {
	StartBoundary *Tstime
	EndBoundary *Tstime
	Unknown0 *Tstime
	RepetitionIntervalSeconds uint32
	RepetitionDurationSeconds uint32
	ExecutionTimeLimitSeconds uint32
	Mode JobSchedule_TimeMode
	Data1 uint16
	Data2 uint16
	Data3 uint16
	Pad0 uint16
	StopTasksAtDurationEnd uint8
	IsEnabled uint8
	Pad1 uint16
	Unknown1 uint32
	MaxDelaySeconds uint32
	Pad2 uint32
	_io *kaitai.Stream
	_root *JobSchedule
	_parent interface{}
}
func NewJobSchedule() *JobSchedule {
	return &JobSchedule{
	}
}

func (this *JobSchedule) Read(io *kaitai.Stream, parent interface{}, root *JobSchedule) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1 := NewTstime()
	err = tmp1.Read(this._io, this, nil)
	if err != nil {
		return err
	}
	this.StartBoundary = tmp1
	tmp2 := NewTstime()
	err = tmp2.Read(this._io, this, nil)
	if err != nil {
		return err
	}
	this.EndBoundary = tmp2
	tmp3 := NewTstime()
	err = tmp3.Read(this._io, this, nil)
	if err != nil {
		return err
	}
	this.Unknown0 = tmp3
	tmp4, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.RepetitionIntervalSeconds = uint32(tmp4)
	tmp5, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.RepetitionDurationSeconds = uint32(tmp5)
	tmp6, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.ExecutionTimeLimitSeconds = uint32(tmp6)
	tmp7, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Mode = JobSchedule_TimeMode(tmp7)
	tmp8, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.Data1 = uint16(tmp8)
	tmp9, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.Data2 = uint16(tmp9)
	tmp10, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.Data3 = uint16(tmp10)
	tmp11, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.Pad0 = uint16(tmp11)
	tmp12, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.StopTasksAtDurationEnd = tmp12
	tmp13, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.IsEnabled = tmp13
	tmp14, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.Pad1 = uint16(tmp14)
	tmp15, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Unknown1 = uint32(tmp15)
	tmp16, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.MaxDelaySeconds = uint32(tmp16)
	tmp17, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Pad2 = uint32(tmp17)
	return err
}
