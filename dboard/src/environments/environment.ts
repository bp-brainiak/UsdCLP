// This file can be replaced during build by using the `fileReplacements` array.
// `ng build` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

import { APP_ID } from "@angular/core";

export const environment = {
  production: false,
  bakendHost: 'http://localhost:9433',
  endPointUF: 'https://api.cmfchile.cl/api-sbifv3/recursos_api/uf?apikey=b6027c6dac746ee3f3f732ba652eaccf36ac9a0d&formato=json',
  endPointIPC: 'https://api.cmfchile.cl/api-sbifv3/recursos_api/ipc?apikey=b6027c6dac746ee3f3f732ba652eaccf36ac9a0d&formato=json',
  endPointUTM: 'https://api.cmfchile.cl/api-sbifv3/recursos_api/utm?apikey=b6027c6dac746ee3f3f732ba652eaccf36ac9a0d&formato=json',
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.

