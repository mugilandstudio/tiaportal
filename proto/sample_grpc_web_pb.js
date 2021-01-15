/**
 * @fileoverview gRPC-Web generated client stub for tiaportal
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')
const proto = {};
proto.tiaportal = require('./sample_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tiaportal.SampleServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.tiaportal.SampleServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.tiaportal.EchoRequest,
 *   !proto.tiaportal.EchoResponse>}
 */
const methodDescriptor_SampleService_Echo = new grpc.web.MethodDescriptor(
  '/tiaportal.SampleService/Echo',
  grpc.web.MethodType.UNARY,
  proto.tiaportal.EchoRequest,
  proto.tiaportal.EchoResponse,
  /**
   * @param {!proto.tiaportal.EchoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tiaportal.EchoResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.tiaportal.EchoRequest,
 *   !proto.tiaportal.EchoResponse>}
 */
const methodInfo_SampleService_Echo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.tiaportal.EchoResponse,
  /**
   * @param {!proto.tiaportal.EchoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.tiaportal.EchoResponse.deserializeBinary
);


/**
 * @param {!proto.tiaportal.EchoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.tiaportal.EchoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.tiaportal.EchoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.tiaportal.SampleServiceClient.prototype.echo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/tiaportal.SampleService/Echo',
      request,
      metadata || {},
      methodDescriptor_SampleService_Echo,
      callback);
};


/**
 * @param {!proto.tiaportal.EchoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.tiaportal.EchoResponse>}
 *     Promise that resolves to the response
 */
proto.tiaportal.SampleServicePromiseClient.prototype.echo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/tiaportal.SampleService/Echo',
      request,
      metadata || {},
      methodDescriptor_SampleService_Echo);
};


module.exports = proto.tiaportal;

