// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file pb/event.proto (package stredono, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum stredono.SubTier
 */
export enum SubTier {
  /**
   * @generated from enum value: PRIME = 0;
   */
  PRIME = 0,

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
  { no: 0, name: "PRIME" },
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
}
// Retrieve enum metadata with: proto3.getEnumType(EventType)
proto3.util.setEnumType(EventType, "stredono.EventType", [
  { no: 0, name: "TIP" },
  { no: 1, name: "CHEER" },
  { no: 2, name: "FOLLOW" },
  { no: 3, name: "SUB" },
  { no: 4, name: "SUB_GIFT" },
  { no: 5, name: "RAID" },
]);

/**
 * @generated from message stredono.Event
 */
export class Event extends Message<Event> {
  /**
   * @generated from field: string Id = 1;
   */
  Id = "";

  /**
   * @generated from field: string Channel = 2;
   */
  Channel = "";

  /**
   * @generated from field: string Username = 3;
   */
  Username = "";

  /**
   * @generated from field: string TTSUrl = 4;
   */
  TTSUrl = "";

  /**
   * @generated from field: stredono.EventType Type = 5;
   */
  Type = EventType.TIP;

  /**
   * @generated from oneof stredono.Event.Payload
   */
  Payload: {
    /**
     * @generated from field: stredono.Event.TipPayload Tip = 6;
     */
    value: Event_TipPayload;
    case: "Tip";
  } | {
    /**
     * @generated from field: stredono.Event.CheerPayload Cheer = 7;
     */
    value: Event_CheerPayload;
    case: "Cheer";
  } | {
    /**
     * @generated from field: stredono.Event.SubPayload Sub = 8;
     */
    value: Event_SubPayload;
    case: "Sub";
  } | {
    /**
     * @generated from field: stredono.Event.SubGiftPayload SubGift = 9;
     */
    value: Event_SubGiftPayload;
    case: "SubGift";
  } | {
    /**
     * @generated from field: stredono.Event.RaidPayload Raid = 10;
     */
    value: Event_RaidPayload;
    case: "Raid";
  } | {
    /**
     * @generated from field: stredono.Event.ChatMessagePayload ChatMessage = 11;
     */
    value: Event_ChatMessagePayload;
    case: "ChatMessage";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Event>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Event";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "Channel", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "Username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "TTSUrl", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "Type", kind: "enum", T: proto3.getEnumType(EventType) },
    { no: 6, name: "Tip", kind: "message", T: Event_TipPayload, oneof: "Payload" },
    { no: 7, name: "Cheer", kind: "message", T: Event_CheerPayload, oneof: "Payload" },
    { no: 8, name: "Sub", kind: "message", T: Event_SubPayload, oneof: "Payload" },
    { no: 9, name: "SubGift", kind: "message", T: Event_SubGiftPayload, oneof: "Payload" },
    { no: 10, name: "Raid", kind: "message", T: Event_RaidPayload, oneof: "Payload" },
    { no: 11, name: "ChatMessage", kind: "message", T: Event_ChatMessagePayload, oneof: "Payload" },
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

/**
 * @generated from message stredono.Event.SubGiftPayload
 */
export class Event_SubGiftPayload extends Message<Event_SubGiftPayload> {
  /**
   * @generated from field: string Recipient = 1;
   */
  Recipient = "";

  /**
   * @generated from field: stredono.SubTier Tier = 2;
   */
  Tier = SubTier.PRIME;

  constructor(data?: PartialMessage<Event_SubGiftPayload>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Event.SubGiftPayload";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Recipient", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "Tier", kind: "enum", T: proto3.getEnumType(SubTier) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event_SubGiftPayload {
    return new Event_SubGiftPayload().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event_SubGiftPayload {
    return new Event_SubGiftPayload().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event_SubGiftPayload {
    return new Event_SubGiftPayload().fromJsonString(jsonString, options);
  }

  static equals(a: Event_SubGiftPayload | PlainMessage<Event_SubGiftPayload> | undefined, b: Event_SubGiftPayload | PlainMessage<Event_SubGiftPayload> | undefined): boolean {
    return proto3.util.equals(Event_SubGiftPayload, a, b);
  }
}

/**
 * @generated from message stredono.Event.SubPayload
 */
export class Event_SubPayload extends Message<Event_SubPayload> {
  /**
   * @generated from field: string Message = 1;
   */
  Message = "";

  /**
   * @generated from field: stredono.SubTier Tier = 2;
   */
  Tier = SubTier.PRIME;

  constructor(data?: PartialMessage<Event_SubPayload>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Event.SubPayload";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "Tier", kind: "enum", T: proto3.getEnumType(SubTier) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event_SubPayload {
    return new Event_SubPayload().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event_SubPayload {
    return new Event_SubPayload().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event_SubPayload {
    return new Event_SubPayload().fromJsonString(jsonString, options);
  }

  static equals(a: Event_SubPayload | PlainMessage<Event_SubPayload> | undefined, b: Event_SubPayload | PlainMessage<Event_SubPayload> | undefined): boolean {
    return proto3.util.equals(Event_SubPayload, a, b);
  }
}

/**
 * @generated from message stredono.Event.RaidPayload
 */
export class Event_RaidPayload extends Message<Event_RaidPayload> {
  /**
   * @generated from field: int32 Viewers = 1;
   */
  Viewers = 0;

  constructor(data?: PartialMessage<Event_RaidPayload>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Event.RaidPayload";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Viewers", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event_RaidPayload {
    return new Event_RaidPayload().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event_RaidPayload {
    return new Event_RaidPayload().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event_RaidPayload {
    return new Event_RaidPayload().fromJsonString(jsonString, options);
  }

  static equals(a: Event_RaidPayload | PlainMessage<Event_RaidPayload> | undefined, b: Event_RaidPayload | PlainMessage<Event_RaidPayload> | undefined): boolean {
    return proto3.util.equals(Event_RaidPayload, a, b);
  }
}

/**
 * @generated from message stredono.Event.TipPayload
 */
export class Event_TipPayload extends Message<Event_TipPayload> {
  /**
   * @generated from field: string Message = 1;
   */
  Message = "";

  /**
   * @generated from field: double Amount = 2;
   */
  Amount = 0;

  /**
   * @generated from field: string Currency = 3;
   */
  Currency = "";

  constructor(data?: PartialMessage<Event_TipPayload>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Event.TipPayload";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "Amount", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "Currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event_TipPayload {
    return new Event_TipPayload().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event_TipPayload {
    return new Event_TipPayload().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event_TipPayload {
    return new Event_TipPayload().fromJsonString(jsonString, options);
  }

  static equals(a: Event_TipPayload | PlainMessage<Event_TipPayload> | undefined, b: Event_TipPayload | PlainMessage<Event_TipPayload> | undefined): boolean {
    return proto3.util.equals(Event_TipPayload, a, b);
  }
}

/**
 * @generated from message stredono.Event.CheerPayload
 */
export class Event_CheerPayload extends Message<Event_CheerPayload> {
  /**
   * @generated from field: string Message = 1;
   */
  Message = "";

  /**
   * @generated from field: int32 Amount = 2;
   */
  Amount = 0;

  constructor(data?: PartialMessage<Event_CheerPayload>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Event.CheerPayload";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "Amount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event_CheerPayload {
    return new Event_CheerPayload().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event_CheerPayload {
    return new Event_CheerPayload().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event_CheerPayload {
    return new Event_CheerPayload().fromJsonString(jsonString, options);
  }

  static equals(a: Event_CheerPayload | PlainMessage<Event_CheerPayload> | undefined, b: Event_CheerPayload | PlainMessage<Event_CheerPayload> | undefined): boolean {
    return proto3.util.equals(Event_CheerPayload, a, b);
  }
}

/**
 * @generated from message stredono.Event.ChatMessagePayload
 */
export class Event_ChatMessagePayload extends Message<Event_ChatMessagePayload> {
  /**
   * @generated from field: string Message = 1;
   */
  Message = "";

  constructor(data?: PartialMessage<Event_ChatMessagePayload>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "stredono.Event.ChatMessagePayload";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event_ChatMessagePayload {
    return new Event_ChatMessagePayload().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event_ChatMessagePayload {
    return new Event_ChatMessagePayload().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event_ChatMessagePayload {
    return new Event_ChatMessagePayload().fromJsonString(jsonString, options);
  }

  static equals(a: Event_ChatMessagePayload | PlainMessage<Event_ChatMessagePayload> | undefined, b: Event_ChatMessagePayload | PlainMessage<Event_ChatMessagePayload> | undefined): boolean {
    return proto3.util.equals(Event_ChatMessagePayload, a, b);
  }
}

