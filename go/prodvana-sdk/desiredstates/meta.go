package desiredstates

import (
	"github.com/pkg/errors"
	ds_model_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/desired_state/model"
)

func GetMeta(state *ds_model_pb.State) (*ds_model_pb.Metadata, error) {
	switch inner := state.StateOneof.(type) {
	case *ds_model_pb.State_Service:
		return inner.Service.Meta, nil
	case *ds_model_pb.State_ServiceInstance:
		return inner.ServiceInstance.Meta, nil
	case *ds_model_pb.State_CustomTask:
		return inner.CustomTask.Meta, nil
	case *ds_model_pb.State_ProtectionAttachment:
		return inner.ProtectionAttachment.Meta, nil
	case *ds_model_pb.State_ManualApproval:
		return inner.ManualApproval.Meta, nil
	case *ds_model_pb.State_RuntimeObject:
		return inner.RuntimeObject.Meta, nil
	case *ds_model_pb.State_DeliveryExtension:
		return inner.DeliveryExtension.Meta, nil
	case *ds_model_pb.State_ProtectionLink:
		return inner.ProtectionLink.Meta, nil
	default:
		return nil, errors.Errorf("Unsupported state %T", inner)
	}
}
