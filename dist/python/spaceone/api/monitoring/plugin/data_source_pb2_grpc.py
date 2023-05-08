# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from spaceone.api.monitoring.plugin import data_source_pb2 as spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2


class DataSourceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.init = channel.unary_unary(
                '/spaceone.api.monitoring.plugin.DataSource/init',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.InitRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.PluginInfo.FromString,
                )
        self.verify = channel.unary_unary(
                '/spaceone.api.monitoring.plugin.DataSource/verify',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.PluginVerifyRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )


class DataSourceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def init(self, request, context):
        """
        desc: Initializes a specific DataSource. During initialization, the DataSource information to be passed to the DataSource user is delivered as `metadata`. DataSource information includes its name and version.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def verify(self, request, context):
        """
        desc: Verifies a specific DataSource. You must specify the parameter `secret_data`, encrypted account data of the DataSource to validate.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DataSourceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'init': grpc.unary_unary_rpc_method_handler(
                    servicer.init,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.InitRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.PluginInfo.SerializeToString,
            ),
            'verify': grpc.unary_unary_rpc_method_handler(
                    servicer.verify,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.PluginVerifyRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.monitoring.plugin.DataSource', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DataSource(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def init(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.monitoring.plugin.DataSource/init',
            spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.InitRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.PluginInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def verify(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.monitoring.plugin.DataSource/verify',
            spaceone_dot_api_dot_monitoring_dot_plugin_dot_data__source__pb2.PluginVerifyRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
