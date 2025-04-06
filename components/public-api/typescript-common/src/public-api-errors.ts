/**
 * Copyright (c) 2023 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { ApplicationError, ErrorCodes } from "@khulnasoft/khulnasoft-protocol/lib/messaging/error";
import {
    InvalidKhulnasoftYMLError as InvalidKhulnasoftYMLErrorData,
    RepositoryNotFoundError as RepositoryNotFoundErrorData,
    RepositoryUnauthorizedError as RepositoryUnauthorizedErrorData,
} from "@khulnasoft/public-api/lib/khulnasoft/v1/error_pb";

export class RepositoryNotFoundError extends ApplicationError {
    constructor(readonly info: PlainMessage<RepositoryNotFoundErrorData>) {
        // on gRPC we remap to PRECONDITION_FAILED, all error code for backwards compatibility with the dashboard
        super(ErrorCodes.NOT_FOUND, "Repository not found.", info);
    }
}
export class UnauthorizedRepositoryAccessError extends ApplicationError {
    constructor(readonly info: PartialMessage<RepositoryUnauthorizedErrorData>) {
        // on gRPC we remap to PRECONDITION_FAILED, all error code for backwards compatibility with the dashboard
        super(ErrorCodes.NOT_AUTHENTICATED, "Repository unauthorized.", info);
    }
}
export class InvalidKhulnasoftYMLError extends ApplicationError {
    constructor(readonly info: PlainMessage<InvalidKhulnasoftYMLErrorData>) {
        // on gRPC we remap to PRECONDITION_FAILED, all error code for backwards compatibility with the dashboard
        super(ErrorCodes.INVALID_KHULNASOFT_YML, "Invalid khulnasoft.yml: " + info.violations.join(","), info);
    }
}
