# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from spaceone.api.monitoring.plugin import event_pb2 as spaceone_dot_api_dot_monitoring_dot_plugin_dot_event__pb2

GRPC_GENERATED_VERSION = '1.64.0'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in spaceone/api/monitoring/plugin/event_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class EventStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.parse = channel.unary_unary(
                '/spaceone.api.monitoring.plugin.Event/parse',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_event__pb2.ParseRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_event__pb2.EventsInfo.FromString,
                _registered_method=True)


class EventServicer(object):
    """Missing associated documentation comment in .proto file."""

    def parse(self, request, context):
        """Parses a specific Event resource and transforms the resource from external data collected to data available in Cloudforet internally. The `parse` method standardizes all different data from external monitoring systems.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EventServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'parse': grpc.unary_unary_rpc_method_handler(
                    servicer.parse,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_event__pb2.ParseRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_plugin_dot_event__pb2.EventsInfo.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.monitoring.plugin.Event', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.monitoring.plugin.Event', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Event(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def parse(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.monitoring.plugin.Event/parse',
            spaceone_dot_api_dot_monitoring_dot_plugin_dot_event__pb2.ParseRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_plugin_dot_event__pb2.EventsInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
