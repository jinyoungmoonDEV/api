# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.inventory.v1 import job_task_pb2 as spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2


class JobTaskStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.delete = channel.unary_unary(
                '/spaceone.api.inventory.v1.JobTask/delete',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.inventory.v1.JobTask/get',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.GetJobTaskRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.inventory.v1.JobTask/list',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTasksInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.inventory.v1.JobTask/stat',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class JobTaskServicer(object):
    """Missing associated documentation comment in .proto file."""

    def delete(self, request, context):
        """
        desc: Deletes a specific JobTask. You must specify the `job_task_id` of the JobTask to delete.
        request_example: >-
        {
        "job_task_id": "job-task-123456789012",
        "domain_id": "domain-123456789012"
        }
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """
        desc: Gets a specific JobTask. Prints detailed information about the JobTask, including its state, updated or failure counts, and error log.
        request_example: >-
        {
        "job_task_id": "job-task-123456789012",
        "domain_id": "domain-123456789012"
        }
        response_example: >-
        {
        "job_task_id": "job-task-123456789012",
        "status": "FAILURE",
        "updated_count": 199,
        "failure_count": 1,
        "errors": [
        {
        "error_code": "ERROR_PLUGIN",
        "message": "{\"tags\": [\"Could not interpret the value as a list\"]}",
        "additional": {
        "domain_id": "domain-123456789012",
        "resource_id": "eventarc-us-central1-function",
        "resource_type": "inventory.CloudService",
        "cloud_service_group": "Pub/Sub",
        "cloud_service_type": "Subscription",
        "provider": "google_cloud"
        }
        }
        ],
        "job_id": "job-123456789012",
        "secret_id": "secret-123456789012",
        "provider": "google_cloud",
        "service_account_id": "sa-123456789012",
        "project_id": "project-123456789012",
        "domain_id": "domain-123456789012",
        "created_at": "2022-01-01T11:00:02.588Z",
        "started_at": "2022-01-01T11:00:02.819Z",
        "finished_at": "2022-01-01T11:00:34.398Z"
        }
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """
        desc: Gets a list of all JobTasks in a specific Job. You can use a query to get a filtered list of JobTasks.
        request_example: >-
        {
        "query": {}
        }
        response_example: >-
        {
        "results": [
        {
        "job_task_id": "job_task-69b301d0fbfc",
        "status": "SUCCESS",
        "updated_count": 55,
        "job_id": "job-587a3d3b4db3",
        "secret_id": "secret-957e407bfc15",
        "provider": "aws",
        "service_account_id": "sa-a41ff4765671",
        "project_id": "project-77dffd3f7cd3",
        "domain_id": "domain-58010aa2e451",
        "created_at": "2022-06-17T08:00:00.680Z",
        "started_at": "2022-06-17T08:00:00.914Z",
        "finished_at": "2022-06-17T08:05:48.933Z"
        },
        {
        "job_task_id": "job_task-c5077338cf23",
        "status": "SUCCESS",
        "updated_count": 1921,
        "job_id": "job-587a3d3b4db3",
        "secret_id": "secret-1cd7417c1889",
        "provider": "aws",
        "service_account_id": "sa-5e186fcc7c91",
        "project_id": "project-18655561c535",
        "domain_id": "domain-58010aa2e451",
        "created_at": "2022-06-17T08:00:00.582Z",
        "started_at": "2022-06-17T08:00:00.814Z",
        "finished_at": "2022-06-17T08:07:31.995Z"
        }
        ],
        "total_count": 4839
        }
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_JobTaskServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.GetJobTaskRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTasksInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.inventory.v1.JobTask', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class JobTask(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.JobTask/delete',
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.JobTask/get',
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.GetJobTaskRequest.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def list(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.JobTask/list',
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskQuery.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTasksInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def stat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.JobTask/stat',
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__task__pb2.JobTaskStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
