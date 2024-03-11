// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file pb/chat_role.proto (package stredono, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum stredono.Role
 */
export enum Role {
  /**
   * @generated from enum value: NORMAL = 0;
   */
  NORMAL = 0,

  /**
   * @generated from enum value: VIP = 1;
   */
  VIP = 1,

  /**
   * @generated from enum value: MODERATOR = 2;
   */
  MODERATOR = 2,

  /**
   * @generated from enum value: OWNER = 3;
   */
  OWNER = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Role)
proto3.util.setEnumType(Role, "stredono.Role", [
  { no: 0, name: "NORMAL" },
  { no: 1, name: "VIP" },
  { no: 2, name: "MODERATOR" },
  { no: 3, name: "OWNER" },
]);
