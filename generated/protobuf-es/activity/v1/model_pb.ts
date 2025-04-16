// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file activity/v1/model.proto (package flipcash.activity.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { PaymentAmount } from "../../common/v1/common_pb";

/**
 * ActivityFeedType enables multiple activity feeds, where notifications may be
 * split across different parts of the app
 *
 * @generated from enum flipcash.activity.v1.ActivityFeedType
 */
export enum ActivityFeedType {
  /**
   * @generated from enum value: UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * Activity feed displayed under the Balance tab
   *
   * @generated from enum value: TRANSACTION_HISTORY = 1;
   */
  TRANSACTION_HISTORY = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(ActivityFeedType)
proto3.util.setEnumType(ActivityFeedType, "flipcash.activity.v1.ActivityFeedType", [
  { no: 0, name: "UNKNOWN" },
  { no: 1, name: "TRANSACTION_HISTORY" },
]);

/**
 * The ID of the notification, which is guaranteed to be consistent for grouped
 * events. Updates to a notification with the same ID should result in re-ordering
 * within the activity feed using the latest content.
 *
 * @generated from message flipcash.activity.v1.NotificationId
 */
export class NotificationId extends Message<NotificationId> {
  /**
   * @generated from field: bytes value = 1;
   */
  value = new Uint8Array(0);

  constructor(data?: PartialMessage<NotificationId>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipcash.activity.v1.NotificationId";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NotificationId {
    return new NotificationId().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NotificationId {
    return new NotificationId().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NotificationId {
    return new NotificationId().fromJsonString(jsonString, options);
  }

  static equals(a: NotificationId | PlainMessage<NotificationId> | undefined, b: NotificationId | PlainMessage<NotificationId> | undefined): boolean {
    return proto3.util.equals(NotificationId, a, b);
  }
}

/**
 * Notification is a message that is displayed in an activity feed
 *
 * @generated from message flipcash.activity.v1.Notification
 */
export class Notification extends Message<Notification> {
  /**
   * The ID of this notification
   *
   * @generated from field: flipcash.activity.v1.NotificationId id = 1;
   */
  id?: NotificationId;

  /**
   * The localized title text for the notification
   *
   * @generated from field: string localized_text = 2;
   */
  localizedText = "";

  /**
   * If a payment applies, the amount that was paid
   *
   * @generated from field: flipcash.common.v1.PaymentAmount payment_amount = 3;
   */
  paymentAmount?: PaymentAmount;

  /**
   * The timestamp of this notification
   *
   * @generated from field: google.protobuf.Timestamp ts = 4;
   */
  ts?: Timestamp;

  /**
   * Additional metadata for this notification specific to the notification
   *
   * @generated from oneof flipcash.activity.v1.Notification.additional_metadata
   */
  additionalMetadata: {
    /**
     * @generated from field: flipcash.activity.v1.WelcomeBonusNotificationMetadata welcome_bonus = 5;
     */
    value: WelcomeBonusNotificationMetadata;
    case: "welcomeBonus";
  } | {
    /**
     * @generated from field: flipcash.activity.v1.GaveUsdcNotificationMetadata gave_usdc = 6;
     */
    value: GaveUsdcNotificationMetadata;
    case: "gaveUsdc";
  } | {
    /**
     * @generated from field: flipcash.activity.v1.ReceivedUsdcNotificationMetadata received_usdc = 7;
     */
    value: ReceivedUsdcNotificationMetadata;
    case: "receivedUsdc";
  } | {
    /**
     * @generated from field: flipcash.activity.v1.WithdrewUsdcNotificationMetadata withdrew_usdc = 8;
     */
    value: WithdrewUsdcNotificationMetadata;
    case: "withdrewUsdc";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Notification>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipcash.activity.v1.Notification";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "message", T: NotificationId },
    { no: 2, name: "localized_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "payment_amount", kind: "message", T: PaymentAmount },
    { no: 4, name: "ts", kind: "message", T: Timestamp },
    { no: 5, name: "welcome_bonus", kind: "message", T: WelcomeBonusNotificationMetadata, oneof: "additional_metadata" },
    { no: 6, name: "gave_usdc", kind: "message", T: GaveUsdcNotificationMetadata, oneof: "additional_metadata" },
    { no: 7, name: "received_usdc", kind: "message", T: ReceivedUsdcNotificationMetadata, oneof: "additional_metadata" },
    { no: 8, name: "withdrew_usdc", kind: "message", T: WithdrewUsdcNotificationMetadata, oneof: "additional_metadata" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Notification {
    return new Notification().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Notification {
    return new Notification().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Notification {
    return new Notification().fromJsonString(jsonString, options);
  }

  static equals(a: Notification | PlainMessage<Notification> | undefined, b: Notification | PlainMessage<Notification> | undefined): boolean {
    return proto3.util.equals(Notification, a, b);
  }
}

/**
 * @generated from message flipcash.activity.v1.WelcomeBonusNotificationMetadata
 */
export class WelcomeBonusNotificationMetadata extends Message<WelcomeBonusNotificationMetadata> {
  constructor(data?: PartialMessage<WelcomeBonusNotificationMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipcash.activity.v1.WelcomeBonusNotificationMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WelcomeBonusNotificationMetadata {
    return new WelcomeBonusNotificationMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WelcomeBonusNotificationMetadata {
    return new WelcomeBonusNotificationMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WelcomeBonusNotificationMetadata {
    return new WelcomeBonusNotificationMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: WelcomeBonusNotificationMetadata | PlainMessage<WelcomeBonusNotificationMetadata> | undefined, b: WelcomeBonusNotificationMetadata | PlainMessage<WelcomeBonusNotificationMetadata> | undefined): boolean {
    return proto3.util.equals(WelcomeBonusNotificationMetadata, a, b);
  }
}

/**
 * @generated from message flipcash.activity.v1.GaveUsdcNotificationMetadata
 */
export class GaveUsdcNotificationMetadata extends Message<GaveUsdcNotificationMetadata> {
  constructor(data?: PartialMessage<GaveUsdcNotificationMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipcash.activity.v1.GaveUsdcNotificationMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GaveUsdcNotificationMetadata {
    return new GaveUsdcNotificationMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GaveUsdcNotificationMetadata {
    return new GaveUsdcNotificationMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GaveUsdcNotificationMetadata {
    return new GaveUsdcNotificationMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: GaveUsdcNotificationMetadata | PlainMessage<GaveUsdcNotificationMetadata> | undefined, b: GaveUsdcNotificationMetadata | PlainMessage<GaveUsdcNotificationMetadata> | undefined): boolean {
    return proto3.util.equals(GaveUsdcNotificationMetadata, a, b);
  }
}

/**
 * @generated from message flipcash.activity.v1.ReceivedUsdcNotificationMetadata
 */
export class ReceivedUsdcNotificationMetadata extends Message<ReceivedUsdcNotificationMetadata> {
  constructor(data?: PartialMessage<ReceivedUsdcNotificationMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipcash.activity.v1.ReceivedUsdcNotificationMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReceivedUsdcNotificationMetadata {
    return new ReceivedUsdcNotificationMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReceivedUsdcNotificationMetadata {
    return new ReceivedUsdcNotificationMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReceivedUsdcNotificationMetadata {
    return new ReceivedUsdcNotificationMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: ReceivedUsdcNotificationMetadata | PlainMessage<ReceivedUsdcNotificationMetadata> | undefined, b: ReceivedUsdcNotificationMetadata | PlainMessage<ReceivedUsdcNotificationMetadata> | undefined): boolean {
    return proto3.util.equals(ReceivedUsdcNotificationMetadata, a, b);
  }
}

/**
 * @generated from message flipcash.activity.v1.WithdrewUsdcNotificationMetadata
 */
export class WithdrewUsdcNotificationMetadata extends Message<WithdrewUsdcNotificationMetadata> {
  constructor(data?: PartialMessage<WithdrewUsdcNotificationMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipcash.activity.v1.WithdrewUsdcNotificationMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WithdrewUsdcNotificationMetadata {
    return new WithdrewUsdcNotificationMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WithdrewUsdcNotificationMetadata {
    return new WithdrewUsdcNotificationMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WithdrewUsdcNotificationMetadata {
    return new WithdrewUsdcNotificationMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: WithdrewUsdcNotificationMetadata | PlainMessage<WithdrewUsdcNotificationMetadata> | undefined, b: WithdrewUsdcNotificationMetadata | PlainMessage<WithdrewUsdcNotificationMetadata> | undefined): boolean {
    return proto3.util.equals(WithdrewUsdcNotificationMetadata, a, b);
  }
}

