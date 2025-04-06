-- Copyright (c) 2020 Khulnasoft GmbH. All rights reserved.
-- Licensed under the GNU Affero General Public License (AGPL). See License.AGPL.txt in the project root for license information.


-- create test DB user
SET @khulnasoftDbPassword = IFNULL(@khulnasoftDbPassword, 'test');

SET @statementStr = CONCAT(
    'CREATE USER IF NOT EXISTS "khulnasoft"@"%" IDENTIFIED BY "', @khulnasoftDbPassword, '";'
);
SELECT @statementStr ;
PREPARE stmt FROM @statementStr; EXECUTE stmt; DEALLOCATE PREPARE stmt;
