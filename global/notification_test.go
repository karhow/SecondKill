package global

import "testing"

func TestNotifyUser(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"pushplus 发送消息测试",
			args{
				[]interface{}{"测试信息1", "测试信息2", "测试信息3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NotifyUser(tt.args.v...)
		})
	}
}
