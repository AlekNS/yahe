import { Action } from 'redux';

/**
 * Create default action for redux
 *
 * @param type
 */
export const createAction = (type: string) => <T>(payload: T) => ({ type, payload });

/**
 * Create default async action (with using of saga)
 *
 * @param type
 */
export const createAsyncAction = (type: string) => {
  const REQUEST = `${type}_REQUEST`;
  const SUCCESS = `${type}_SUCCESS`;
  const FAILURE = `${type}_FAILURE`;

  return Object.assign(createAction(REQUEST), {
    type,
    success: createAction(SUCCESS),
    failure: createAction(FAILURE),
    REQUEST,
    SUCCESS,
    FAILURE
  });
};

/**
 * Default reducer helper
 *
 * @param handlers
 * @param initialState
 */
export const createReducer = (handlers: {
  [key: string]: Function,
}, initialState: any) =>
  (state = initialState, action?: Action) =>
    (action && handlers[action.type]
      ? handlers[action.type](state, action)
      : state);
