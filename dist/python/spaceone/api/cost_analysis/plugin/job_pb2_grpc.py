# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from spaceone.api.cost_analysis.plugin import job_pb2 as spaceone_dot_api_dot_cost__analysis_dot_plugin_dot_job__pb2

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
        + f' but the generated code in spaceone/api/cost_analysis/plugin/job_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class JobStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.get_tasks = channel.unary_unary(
                '/spaceone.api.cost_analysis.plugin.Job/get_tasks',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_plugin_dot_job__pb2.GetTasksRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_plugin_dot_job__pb2.TasksInfo.FromString,
                _registered_method=True)


class JobServicer(object):
    """Missing associated documentation comment in .proto file."""

    def get_tasks(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_JobServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'get_tasks': grpc.unary_unary_rpc_method_handler(
                    servicer.get_tasks,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_plugin_dot_job__pb2.GetTasksRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_plugin_dot_job__pb2.TasksInfo.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.cost_analysis.plugin.Job', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.cost_analysis.plugin.Job', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Job(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def get_tasks(request,
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
            '/spaceone.api.cost_analysis.plugin.Job/get_tasks',
            spaceone_dot_api_dot_cost__analysis_dot_plugin_dot_job__pb2.GetTasksRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_plugin_dot_job__pb2.TasksInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
