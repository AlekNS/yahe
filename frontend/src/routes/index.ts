//
// AUTH
//
export const BASE_AUTH = '/auth';
export const AUTH_LOGIN = `${BASE_AUTH}/login`;
export const AUTH_LOGOUT = `${BASE_AUTH}/logout`;

//
// EXPLORER
//
export const BASE_APP = '/explorer';

export const APP_DASHBOARD = `${BASE_APP}/dashboard`;

export const APP_NETWORK_LIST = `${BASE_APP}/networks`;
export const APP_NETWORK = `${APP_NETWORK_LIST}/:networkId`;
export const APP_NETWORK_CHANNEL_LIST = `${APP_NETWORK}/channels`;
export const APP_NETWORK_CHANNEL = `${APP_NETWORK_CHANNEL_LIST}/:channelId`;

export const APP_CHANNEL_LIST = `${APP_NETWORK}/channels`;
export const APP_CHAINCODE_LIST = `${APP_NETWORK_CHANNEL}/chaincodes`
export const APP_BLOCK_LIST = `${APP_NETWORK_CHANNEL}/blocks`;
export const APP_TRANSACTION_LIST = `${APP_NETWORK_CHANNEL}/transactions`;

export const APP_USER = `${BASE_APP}/user`;
export const APP_USER_SETTINGS = `${APP_USER}/settings`;
