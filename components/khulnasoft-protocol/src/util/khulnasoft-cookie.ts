/**
 * Copyright (c) 2021 Khulnasoft GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */
import * as cookie from "cookie";

/**
 * This cookie indicates whether the connected client is a Khulnasoft user (= "has logged in within the last year") or not.
 * This is used by "khulnasoft.com" and "www.khulnasoft.com" to display different content/buttons.
 */
export const NAME = "khulnasoft-user";
export const VALUE = "true";

/**
 * @param domain The domain the Khulnasoft installation is installed onto
 * @returns
 */
export function options(domain: string): cookie.CookieSerializeOptions {
    // Reference: https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies
    return {
        path: "/", // make sure we send the cookie to all sub-pages
        httpOnly: false,
        secure: false,
        maxAge: 60 * 60 * 24 * 365, // 1 year
        sameSite: "lax", // default: true. "Lax" needed to ensure we see cookies from users that navigate to khulnasoft.com from external sites
        domain: `.${domain}`, // explicitly include subdomains to not only cover "khulnasoft.com", but also "www.khulnasoft.com" or workspaces
    };
}

export function generateCookie(domain: string): string {
    return cookie.serialize(NAME, VALUE, options(domain));
}

export function isPresent(cookies: string): boolean {
    // needs to match the old (khulnasoft-user=loggedIn) and new (khulnasoft-user=true) values to ensure a smooth transition during rollout.
    return !!cookies.match(`${NAME}=`);
}
