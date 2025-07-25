// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file event/v1/event_streaming_service.proto (package flipcash.event.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ForwardEventsRequest, ForwardEventsResponse, StreamEventsRequest, StreamEventsResponse } from "./event_streaming_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service flipcash.event.v1.EventStreaming
 */
export const EventStreaming = {
  typeName: "flipcash.event.v1.EventStreaming",
  methods: {
    /**
     * StreamEvents streams events for the requesting user.
     *
     * @generated from rpc flipcash.event.v1.EventStreaming.StreamEvents
     */
    streamEvents: {
      name: "StreamEvents",
      I: StreamEventsRequest,
      O: StreamEventsResponse,
      kind: MethodKind.BiDiStreaming,
    },
    /**
     * ForwardEvents is an internal RPC for forwarding events to another server.
     *
     * @generated from rpc flipcash.event.v1.EventStreaming.ForwardEvents
     */
    forwardEvents: {
      name: "ForwardEvents",
      I: ForwardEventsRequest,
      O: ForwardEventsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

