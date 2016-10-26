package models

type T_xiugai_log struct {
	ID          int64
	Xiangmuname string `xorm:"varchar(100)"`
    Time        string `xorm:"varchar(50)"`
	Xiugainame  string `xorm:"varchar(50)"`
	Beizhuinfor string `xorm:"text"`
	_godb       `xorm:"-"`
}

func (t *T_xiugai_log) Init() *T_xiugai_log {
	t._godb._bean = t
	t._godb._beans = make([]T_xiugai_log, 0)
	return t
}
