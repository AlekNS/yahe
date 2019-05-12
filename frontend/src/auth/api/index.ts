import * as config from '../../config';

import { rxPost } from '../../utils/api';

const apiUrl = `${config.API_URL}/auth`;

/**
 *
 * @param payload
 */
export const loginAction = (payload: {
  email: string,
  password: string,
}) => rxPost(`${apiUrl}/login`, {
  login: payload.email, password: payload.password
});

/**
 * @param payload
 */
export const logoutAction = (payload?: any) => rxPost(`${apiUrl}/logout`, payload);
