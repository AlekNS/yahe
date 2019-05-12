import authEpics from '../auth/epics';
import { combineEpics } from 'redux-observable';

const rootEpics = combineEpics(
  authEpics
);

export default rootEpics;
