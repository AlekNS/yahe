import { Observable } from 'rxjs';
import { map, switchMap, catchError } from 'rxjs/operators';
import { ActionsObservable, combineEpics } from 'redux-observable';

import * as api from '../api';
import { authLogin } from '../redux';

/**
 * @param action$
 */
const login = (action$: any) => {
  return (action$.ofType(authLogin.REQUEST) as Observable<any>)
    .pipe(
      map((action: any) => action.payload),
      switchMap((credentials: any) =>
        api.loginAction(credentials).pipe(
          map((response: any) => authLogin.success(response)),
          catchError((error: any) => authLogin.failed(error))
        )
      )
    )
};

export default combineEpics(
  login,
);
