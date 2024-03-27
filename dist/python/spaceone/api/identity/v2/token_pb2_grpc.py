# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from spaceone.api.identity.v2 import token_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2


class TokenStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.issue = channel.unary_unary(
                '/spaceone.api.identity.v2.Token/issue',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.IssueTokenRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.TokenInfo.FromString,
                )
        self.grant = channel.unary_unary(
                '/spaceone.api.identity.v2.Token/grant',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.GrantTokenRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.GrantTokenInfo.FromString,
                )


class TokenServicer(object):
    """Missing associated documentation comment in .proto file."""

    def issue(self, request, context):
        """+noauth
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def grant(self, request, context):
        """+noauth
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TokenServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'issue': grpc.unary_unary_rpc_method_handler(
                    servicer.issue,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.IssueTokenRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.TokenInfo.SerializeToString,
            ),
            'grant': grpc.unary_unary_rpc_method_handler(
                    servicer.grant,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.GrantTokenRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.GrantTokenInfo.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v2.Token', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Token(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def issue(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.Token/issue',
            spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.IssueTokenRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.TokenInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def grant(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.Token/grant',
            spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.GrantTokenRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_token__pb2.GrantTokenInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
