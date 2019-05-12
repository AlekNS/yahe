import { createStore, applyMiddleware, compose } from 'redux';
import createHistory from 'history/createBrowserHistory';
import { routerMiddleware } from 'react-router-redux';
import { composeWithDevTools } from 'redux-devtools-extension';
import { createLogger } from 'redux-logger';
import { from } from 'seamless-immutable';
import { stateTransformer } from 'redux-seamless-immutable';

import rootReducer from './redux';
import rootEpics from './epics';
import { createEpicMiddleware } from 'redux-observable';
// import { BehaviorSubject } from 'rxjs';

const history = createHistory();
const loggerMiddleware = createLogger({
  stateTransformer,
  collapsed: true
});

const composeEnhancers = composeWithDevTools || compose;

// const epic$ = new BehaviorSubject(rootEpics);

const configureStore = (isProduction: boolean, initialState: any, epicDeps: any) => {
  const epicMiddlewares = createEpicMiddleware(epicDeps);
  const middlewares = [
    createEpicMiddleware,
    routerMiddleware(history),
  ] as any[];
  if (!isProduction) {
    middlewares.unshift(loggerMiddleware)
  }

  const store = createStore(
    rootReducer,
    initialState,
    composeEnhancers(applyMiddleware(...middlewares))
  );

  if (!isProduction) {
    // if ((module as any).hot) {
    //   (module as any).hot.accept('./rootReducer', () => {
    //     const nextReducer = require('./rootReducer').default;
    //     store.replaceReducer(nextReducer);

    //     const nextRootEpic = rootEpics;
    //     epic$.next(nextRootEpic);
    //   });
    // }
  }

  epicMiddlewares.run(rootEpics);

  return store;
};

const initialState = from({});
const epicsDeps = {};
const store = configureStore(process.env.NODE_ENV === 'production', initialState, epicsDeps);

export { store, history };
