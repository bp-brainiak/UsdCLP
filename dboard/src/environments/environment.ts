// This file can be replaced during build by using the `fileReplacements` array.
// `ng build` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

import { APP_ID } from "@angular/core";

export const environment = {
  production: false,
  urlEndpoint: 'https://yfapi.net/v6/finance/quote?region=US&lang=es&symbols=USDCLP%3DX',
  urlApiKey: '',
  chartEndpoint: 'https://yfapi.net/v8/finance/chart/USDCLP%3DX?',
  bakendHost: 'http://localhost:9433'
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.

