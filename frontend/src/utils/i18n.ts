import i18next from 'i18next';
import XHR from 'i18next-xhr-backend';
import langDetector from 'i18next-browser-languagedetector';

const resources = {
  en: {},
};

i18next
  .use(XHR)
  .use(langDetector)
  .init({
    fallbackLng: 'en',
    debug: process.env.NODE_ENV === 'development',
    react: {
      wait: true,
      nsMode: 'default',
    },
    resources,
    ns: ['app', 'errors'],
    defaultNS: 'app',
  });

export { i18next };
