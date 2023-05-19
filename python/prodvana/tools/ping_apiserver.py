from prodvana_sdk.client import make_channel
from prodvana_sdk.proto.prodvana.auth import auth_manager_pb2, auth_manager_pb2_grpc


def main() -> None:
    channel = make_channel()
    with channel:
        client = auth_manager_pb2_grpc.AuthSessionManagerStub(channel)
        client.Check(auth_manager_pb2.Empty())
    print("success")


if __name__ == "__main__":
    main()
