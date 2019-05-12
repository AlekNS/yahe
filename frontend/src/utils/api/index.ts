import { Observable, Observer } from 'rxjs';
import { ajax, AjaxRequest, AjaxResponse } from 'rxjs/ajax';
import { map } from 'rxjs/operators';

import * as routes from '../../routes';
import { isStoredTokenValid, getToken, removeToken } from '../../auth/utils';

/**
 * Response error
 */
export class ResponseError extends Error {
  errors: string[] | undefined;
  status: number | undefined;

  constructor(error: any) {
    super(error.message || error.error);

    this.errors = error.errors;
    this.status = error.status_code || error.statusCode;
  }
}

/**
 * Authorization error
 */
export class AuthorizationError extends ResponseError { }

/**
 * Parse response body to json
 *
 * @param response
 */
export const parseJSON = (response: AjaxResponse) => {
  let parsedResponse = {};
  try {
    parsedResponse = JSON.parse(response.responseText)
  } catch (err) {
    parsedResponse = {
      error: err,
      statusCode: response.status,
    };
  }

  if (response.status < 400) {
    return parsedResponse;
  }

  if (response.status === 401) {
    removeToken();
    window.location.replace(window.location.hostname + `/${routes.AUTH_LOGIN}`);
  } else if (response.status === 403) {
    throw new AuthorizationError(parsedResponse);
  }

  throw new ResponseError(parsedResponse);
};

/**
 * Get auth header if token is exist in storage
 */
export function authHeader() {
  return isStoredTokenValid()
    ? { Authorization: `Bearer ${getToken()}` }
    : {};
}

/**
 * Fetch wrapper function
 *
 * @param url
 * @param options
 */
export const rxApiCall = (url: string, options: AjaxRequest) => ajax({
  url,
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    ...authHeader()
  },
  ...options
}).pipe(
  map(response => parseJSON)
)

/**
 * Simple stub
 *
 * @param url
 * @param opts
 * @param response
 */
export const stubRxApiCall = (url: string, opts?: any, response: any = {}) => Observable.create((observer: Observer<any>) => {
  console.log('HTTP REQUEST', url, opts);
  setTimeout(
    () => {
      console.log('HTTP RESPONSE', url, opts, response)
      observer.next(response)
      observer.complete()
    },
    Math.round((Math.random() * 500 + 600))
  );
});

/**
 * Fetch wrapper for GET requests
 *
 * @param url
 * @param opts
 */
export const rxGet = (url: string, opts?: AjaxRequest) => (rxApiCall(url, {
  ...opts,
  method: 'GET'
}));

/**
 * Wrapper for POST requests
 *
 * @param url
 * @param body
 * @param opts
 */
export const rxPost = (url: string, body?: any, opts?: AjaxRequest) => (rxApiCall(url, {
  ...opts,
  method: 'POST',
  body: JSON.stringify(body)
}));

/**
 * Wrapper for PATCH requests
 *
 * @param url
 * @param body
 * @param opts
 */
export const rxPatch = (url: string, body?: any, opts?: AjaxRequest) => (rxApiCall(url, {
  ...opts,
  method: 'PATCH',
  body: JSON.stringify(body)
}));

/**
 * Wrapper for PUT requests
 *
 * @param url
 * @param body
 * @param opts
 */
export const rxPut = (url: string, body?: any, opts?: AjaxRequest) => (rxApiCall(url, {
  ...opts,
  method: 'PUT',
  body: JSON.stringify(body)
}));

/**
 * Wrapper for DELETE requests
 *
 * @param url
 * @param opts
 */
export const rxDel = (url: string, opts?: AjaxRequest) => (rxApiCall(url, {
  ...opts,
  method: 'DELETE'
}));
