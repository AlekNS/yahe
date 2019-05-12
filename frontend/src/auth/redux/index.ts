import { from } from 'seamless-immutable';

import { createReducer, createAction, createAsyncAction } from '../../utils';

export const AUTH_LOGIN = 'auth/login/AUTH_LOGIN';
export const AUTH_RESET = 'auth/login/AUTH_RESET';

export const authLogin = createAsyncAction(AUTH_LOGIN);
export const authReset = createAction(AUTH_RESET);

const initialState = from({
  fetching: false,
  authenticated: false,
  error: null,
  user: {},
});

export default createReducer({
  [AUTH_RESET]: (state: any) => state.merge(initialState),

  [authLogin.REQUEST]: (state: any) => state.merge({ fetching: true, error: null }),

  [authLogin.SUCCESS]: (state: any, payload: any) => state.merge({
    fetching: false,
    authenticated: true,
    error: null,
    user: payload,
  }),

  [authLogin.FAILURE]: (state: any, payload: any) => state.merge({
    fetching: false,
    authenticated: false,
    error: payload,
    user: {},
  }),
}, initialState);
