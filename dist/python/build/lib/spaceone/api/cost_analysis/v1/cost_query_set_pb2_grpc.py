# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.cost_analysis.v1 import cost_query_set_pb2 as spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2


class CostQuerySetStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostQuerySet/create',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CreateCostQuerySetRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.FromString,
                )
        self.update = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostQuerySet/update',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.UpdateCostQuerySetRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.FromString,
                )
        self.delete = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostQuerySet/delete',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostQuerySet/get',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostQuerySet/list',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.CostQuerySet/stat',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class CostQuerySetServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Creates a new CostQuerySet. You can make your own custom query that meets your needs, and input it as an `option` parameter of the resource. Queries such as `group_by` and `granularity` are provided by default.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Updates a specific CostQuerySet. You can make changes in the details of queries.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Deletes a specific CostQuerySet. You must specify the `cost_query_set_id` of the CostQuerySet to delete.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Gets a specific CostQuerySet. Prints detailed information about the CostQuerySet, including the details of queries.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Gets a list of all CostQuerySets. You can use a query to get a filtered list of CostQuerySets.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CostQuerySetServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CreateCostQuerySetRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.UpdateCostQuerySetRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.cost_analysis.v1.CostQuerySet', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CostQuerySet(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostQuerySet/create',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CreateCostQuerySetRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostQuerySet/update',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.UpdateCostQuerySetRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostQuerySet/delete',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetRequest.SerializeToString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostQuerySet/get',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostQuerySet/list',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetQuery.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetsInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.cost_analysis.v1.CostQuerySet/stat',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_cost__query__set__pb2.CostQuerySetStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
