// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file pb/event.proto (package stredono, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from enum stredono.SubTier
 */
export enum SubTier {
  /**
   * @generated from enum value: TIER_UNKNOWN = 0;
   */
  TIER_UNKNOWN = 0,

  /**
   * @generated from enum value: TIER_1 = 1;
   */
  TIER_1 = 1,

  /**
   * @generated from enum value: TIER_2 = 2;
   */
  TIER_2 = 2,

  /**
   * @generated from enum value: TIER_3 = 3;
   */
  TIER_3 = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(SubTier)
proto3.util.setEnumType(SubTier, "stredono.SubTier", [
  { no: 0, name: "TIER_UNKNOWN" },
  { no: 1, name: "TIER_1" },
  { no: 2, name: "TIER_2" },
  { no: 3, name: "TIER_3" },
]);

/**
 * @generated from enum stredono.EventType
 */
export enum EventType {
  /**
   * @generated from enum value: TIP = 0;
   */
  TIP = 0,

  /**
   * @generated from enum value: CHEER = 1;
   */
  CHEER = 1,

  /**
   * @generated from enum value: FOLLOW = 2;
   */
  FOLLOW = 2,

  /**
   * @generated from enum value: SUB = 3;
   */
  SUB = 3,

  /**
   * @generated from enum value: SUB_GIFT = 4;
   */
  SUB_GIFT = 4,

  /**
   * @generated from enum value: RAID = 5;
   */
  RAID = 5,

  /**
   * @generated from enum value: CHAT_TTS = 6;
   */
  CHAT_TTS = 6,
}
// Retrieve enum metadata with: proto3.getEnumType(EventType)
proto3.util.setEnumType(EventType, "stredono.EventType", [
  { no: 0, name: "TIP" },
  { no: 1, name: "CHEER" },
  { no: 2, name: "FOLLOW" },
  { no: 3, name: "SUB" },
  { no: 4, name: "SUB_GIFT" },
  { no: 5, name: "RAID" },
  { no: 6, name: "CHAT_TTS" },
]);

/**
 * @generated from message stredono.Event
 */
export class Event extends Message<Event> {
  /**
   * @generated from field: string ID = 1;
   */
  ID = "";

  /**
   * @generated from field: stredono.EventType Type = 2;
   */
  Type = EventType.TIP;

  /**
   * @generated from field: string Uid = 3;
   */
  Uid = "";

  /**
   * @generated from field: string ProviderID = 4;
   */
  ProviderID = "";

  /**
   * @generated from field: string SenderID = 5;
   */
  SenderID = "";

  /**
   * @generated from field: string SenderName = 6;
   */
  SenderName = "";

  /**
   * @generated from field: string Provider = 7;
   */
  Provider = "";

  /**
   * @generated from field: string TTSUrl = 8;
   */
  TTSUrl = "";

  /**
   * @generated from field: bool IsApproved = 9;
   */
  IsApproved = false;

  /**
   * @generated from field: int64 Timestamp = 10;
   */
  Timestamp = protoInt64.zero;

  /**
   * @generated from field: map<string, string> Data = 11;
   */
  Data: { [key: string]: string } = {};

  /**
   * @generated from field: bool WasShown = 12;
   */
  WasShown = false;

  constructor(data?: PartialMessage<Event>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Event";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "Type", kind: "enum", T: proto3.getEnumType(EventType) },
    { no: 3, name: "Uid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "ProviderID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "SenderID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "SenderName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "Provider", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "TTSUrl", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "IsApproved", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 10, name: "Timestamp", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 11, name: "Data", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 12, name: "WasShown", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event {
    return new Event().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event {
    return new Event().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event {
    return new Event().fromJsonString(jsonString, options);
  }

  static equals(a: Event | PlainMessage<Event> | undefined, b: Event | PlainMessage<Event> | undefined): boolean {
    return proto3.util.equals(Event, a, b);
  }
}

