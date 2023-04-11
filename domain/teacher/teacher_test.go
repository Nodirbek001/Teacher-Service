package teacher

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var (
	testSubjectId = uuid.New()
	testTecherId  = uuid.New()
)

func TestNew(t *testing.T) {
	type args struct {
		firstName   string
		lastname    string
		email       string
		phoneNumber string
		subjectID   uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Teacher
		wantErr bool
	}{
		struct {
			name    string
			args    args
			want    Teacher
			wantErr bool
		}{
			name: "should pass",
			args: args{
				firstName:   "Ali",
				lastname:    "vali",
				email:       "@damd",
				phoneNumber: "62652",
				subjectID:   testSubjectId,
			},
			want: Teacher{
				id:          testTecherId,
				firstName:   "Ali",
				lastName:    "vali",
				email:       "@damd",
				phoneNumber: "62652",
				subjectID:   testSubjectId,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.firstName, tt.args.lastname, tt.args.email, tt.args.phoneNumber, tt.args.subjectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.id=testTecherId
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
