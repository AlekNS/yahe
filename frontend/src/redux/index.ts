import { combineReducers, routerReducer } from 'redux-seamless-immutable';

import auth from '../auth/redux';

export default combineReducers({
  routing: routerReducer,

  auth,
});
