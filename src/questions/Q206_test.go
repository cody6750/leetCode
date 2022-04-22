package questions

import (
	"log"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *LinkedList
	}
	tests := []struct {
		name string
		args args
		want *LinkedList
	}{
		{
			name: "works",
			args: args{
				head: &LinkedList{Val: 1, Next: &LinkedList{Val: 2, Next: &LinkedList{Val: 3, Next: nil}}},
			},
			want: &LinkedList{Val: 3, Next: &LinkedList{Val: 2, Next: &LinkedList{Val: 1, Next: nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			if !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
			log.Printf("reverseList() = %v, want %v", got, tt.want)
			log.Printf("reverseList() = %v, want %v", got.Next, tt.want.Next)

		})
	}
}
