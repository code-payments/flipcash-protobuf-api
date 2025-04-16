// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file activity/v1/activity_feed_service.proto (package flipcash.activity.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ActivityFeedType, Notification } from "./model_pb";
import { Auth } from "../../common/v1/common_pb";

/**
 * @generated from message flipcash.activity.v1.GetLatestNotificationsRequest
 */
export class GetLatestNotificationsRequest extends Message<GetLatestNotificationsRequest> {
  /**
   * The activity feed to fetch notifications from
   *
   * @generated from field: flipcash.activity.v1.ActivityFeedType type = 1;
   */
  type = ActivityFeedType.UNKNOWN;

  /**
   * Maximum number of notifications to return. If <= 0, the server default is used
   *
   * @generated from field: int32 max_items = 2;
   */
  maxItems = 0;

  /**
   * @generated from field: flipcash.common.v1.Auth auth = 3;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<GetLatestNotificationsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipcash.activity.v1.GetLatestNotificationsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(ActivityFeedType) },
    { no: 2, name: "max_items", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLatestNotificationsRequest {
    return new GetLatestNotificationsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLatestNotificationsRequest {
    return new GetLatestNotificationsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLatestNotificationsRequest {
    return new GetLatestNotificationsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetLatestNotificationsRequest | PlainMessage<GetLatestNotificationsRequest> | undefined, b: GetLatestNotificationsRequest | PlainMessage<GetLatestNotificationsRequest> | undefined): boolean {
    return proto3.util.equals(GetLatestNotificationsRequest, a, b);
  }
}

/**
 * @generated from message flipcash.activity.v1.GetLatestNotificationsResponse
 */
export class GetLatestNotificationsResponse extends Message<GetLatestNotificationsResponse> {
  /**
   * @generated from field: flipcash.activity.v1.GetLatestNotificationsResponse.Result result = 1;
   */
  result = GetLatestNotificationsResponse_Result.OK;

  /**
   * @generated from field: repeated flipcash.activity.v1.Notification notifications = 2;
   */
  notifications: Notification[] = [];

  constructor(data?: PartialMessage<GetLatestNotificationsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipcash.activity.v1.GetLatestNotificationsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(GetLatestNotificationsResponse_Result) },
    { no: 2, name: "notifications", kind: "message", T: Notification, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLatestNotificationsResponse {
    return new GetLatestNotificationsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLatestNotificationsResponse {
    return new GetLatestNotificationsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLatestNotificationsResponse {
    return new GetLatestNotificationsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetLatestNotificationsResponse | PlainMessage<GetLatestNotificationsResponse> | undefined, b: GetLatestNotificationsResponse | PlainMessage<GetLatestNotificationsResponse> | undefined): boolean {
    return proto3.util.equals(GetLatestNotificationsResponse, a, b);
  }
}

/**
 * @generated from enum flipcash.activity.v1.GetLatestNotificationsResponse.Result
 */
export enum GetLatestNotificationsResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: DENIED = 1;
   */
  DENIED = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(GetLatestNotificationsResponse_Result)
proto3.util.setEnumType(GetLatestNotificationsResponse_Result, "flipcash.activity.v1.GetLatestNotificationsResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "DENIED" },
]);

