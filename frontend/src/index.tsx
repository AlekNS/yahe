import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { ConnectedRouter } from 'react-router-redux';
import { I18nextProvider } from 'react-i18next';

import './index.css';
import App from './containers/App/App';
import * as serviceWorker from './serviceWorker';
import { i18next } from './utils';
import { history, store } from './configureStore';

ReactDOM.render(
  <Provider store={store}>
    <I18nextProvider i18n={i18next}>
      <ConnectedRouter history={history}>
        <App />
      </ConnectedRouter>
    </I18nextProvider>
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
