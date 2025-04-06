-- Copyright (c) 2020 Khulnasoft GmbH. All rights reserved.
-- Licensed under the GNU Affero General Public License (AGPL). See License.AGPL.txt in the project root for license information.

-- must be idempotent

-- @khulnasoftDB contains name of the DB the script manipulates, and is replaced by the file reader
SET
@khulnasoftDB = IFNULL(@khulnasoftDB, '`__KHULNASOFT_DB_NAME__`');

SET
@statementStr = CONCAT('DROP DATABASE IF EXISTS ', @khulnasoftDB);
PREPARE statement FROM @statementStr;
EXECUTE statement;

SET
@statementStr = CONCAT('CREATE DATABASE ', @khulnasoftDB, ' CHARSET utf8mb4');
PREPARE statement FROM @statementStr;
EXECUTE statement;
