# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: archivematica/ccp/admin/v1beta1/deprecated.proto
# Protobuf Python Version: 5.27.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    1,
    '',
    'archivematica/ccp/admin/v1beta1/deprecated.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from client.gen.archivematica.ccp.admin.v1beta1 import admin_pb2 as archivematica_dot_ccp_dot_admin_dot_v1beta1_dot_admin__pb2
from client.gen.buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n0archivematica/ccp/admin/v1beta1/deprecated.proto\x12\x1f\x61rchivematica.ccp.admin.v1beta1\x1a+archivematica/ccp/admin/v1beta1/admin.proto\x1a\x1b\x62uf/validate/validate.proto\"U\n\x11\x41pproveJobRequest\x12\x1f\n\x06job_id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01R\x05jobId\x12\x1f\n\x06\x63hoice\x18\x02 \x01(\tB\x07\xbaH\x04r\x02\x10\x01R\x06\x63hoice\"\x14\n\x12\x41pproveJobResponse\"\x89\x01\n\x1c\x41pproveTransferByPathRequest\x12\x1c\n\tdirectory\x18\x01 \x01(\tR\tdirectory\x12K\n\x04type\x18\x02 \x01(\x0e\x32-.archivematica.ccp.admin.v1beta1.TransferTypeB\x08\xbaH\x05\x82\x01\x02\x10\x01R\x04type\"9\n\x1d\x41pproveTransferByPathResponse\x12\x18\n\x02id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01R\x02id\"9\n\x1d\x41pprovePartialReingestRequest\x12\x18\n\x02id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01R\x02id\" \n\x1e\x41pprovePartialReingestResponseB\xc2\x02\n#com.archivematica.ccp.admin.v1beta1B\x0f\x44\x65precatedProtoP\x01Zkgithub.com/artefactual/archivematica/hack/ccp/internal/api/gen/archivematica/ccp/admin/v1beta1;adminv1beta1\xa2\x02\x03\x41\x43\x41\xaa\x02\x1f\x41rchivematica.Ccp.Admin.V1beta1\xca\x02\x1f\x41rchivematica\\Ccp\\Admin\\V1beta1\xe2\x02+Archivematica\\Ccp\\Admin\\V1beta1\\GPBMetadata\xea\x02\"Archivematica::Ccp::Admin::V1beta1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'archivematica.ccp.admin.v1beta1.deprecated_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n#com.archivematica.ccp.admin.v1beta1B\017DeprecatedProtoP\001Zkgithub.com/artefactual/archivematica/hack/ccp/internal/api/gen/archivematica/ccp/admin/v1beta1;adminv1beta1\242\002\003ACA\252\002\037Archivematica.Ccp.Admin.V1beta1\312\002\037Archivematica\\Ccp\\Admin\\V1beta1\342\002+Archivematica\\Ccp\\Admin\\V1beta1\\GPBMetadata\352\002\"Archivematica::Ccp::Admin::V1beta1'
  _globals['_APPROVEJOBREQUEST'].fields_by_name['job_id']._loaded_options = None
  _globals['_APPROVEJOBREQUEST'].fields_by_name['job_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_APPROVEJOBREQUEST'].fields_by_name['choice']._loaded_options = None
  _globals['_APPROVEJOBREQUEST'].fields_by_name['choice']._serialized_options = b'\272H\004r\002\020\001'
  _globals['_APPROVETRANSFERBYPATHREQUEST'].fields_by_name['type']._loaded_options = None
  _globals['_APPROVETRANSFERBYPATHREQUEST'].fields_by_name['type']._serialized_options = b'\272H\005\202\001\002\020\001'
  _globals['_APPROVETRANSFERBYPATHRESPONSE'].fields_by_name['id']._loaded_options = None
  _globals['_APPROVETRANSFERBYPATHRESPONSE'].fields_by_name['id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_APPROVEPARTIALREINGESTREQUEST'].fields_by_name['id']._loaded_options = None
  _globals['_APPROVEPARTIALREINGESTREQUEST'].fields_by_name['id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_APPROVEJOBREQUEST']._serialized_start=159
  _globals['_APPROVEJOBREQUEST']._serialized_end=244
  _globals['_APPROVEJOBRESPONSE']._serialized_start=246
  _globals['_APPROVEJOBRESPONSE']._serialized_end=266
  _globals['_APPROVETRANSFERBYPATHREQUEST']._serialized_start=269
  _globals['_APPROVETRANSFERBYPATHREQUEST']._serialized_end=406
  _globals['_APPROVETRANSFERBYPATHRESPONSE']._serialized_start=408
  _globals['_APPROVETRANSFERBYPATHRESPONSE']._serialized_end=465
  _globals['_APPROVEPARTIALREINGESTREQUEST']._serialized_start=467
  _globals['_APPROVEPARTIALREINGESTREQUEST']._serialized_end=524
  _globals['_APPROVEPARTIALREINGESTRESPONSE']._serialized_start=526
  _globals['_APPROVEPARTIALREINGESTRESPONSE']._serialized_end=558
# @@protoc_insertion_point(module_scope)
