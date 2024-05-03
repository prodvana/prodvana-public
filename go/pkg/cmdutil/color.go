package cmdutil

import (
	"github.com/fatih/color"

	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
)

func ColorForStatus(status ds_model_pb.Status) *color.Color {
	switch status {
	case ds_model_pb.Status_CONVERGED, ds_model_pb.Status_ROLLED_BACK:
		return color.New(color.FgGreen)
	case ds_model_pb.Status_FAILED:
		return color.New(color.FgRed)
	case ds_model_pb.Status_WAITING_MANUAL_APPROVAL:
		return color.New(color.FgYellow)
	default:
		return color.New(color.FgWhite)
	}
}
