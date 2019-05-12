import { JWT_STORAGE_KEY } from '../../config';

/**
 * Store raw token
 *
 * @param token
 */
export const setToken = (token: string) => {
  localStorage.setItem(JWT_STORAGE_KEY, token);
};

/**
 * Get raw token value
 */
export const getToken = (): string => localStorage.getItem(JWT_STORAGE_KEY) || '';

/**
 * Get info from token
 */
export const getInfo = () => {
  const token = getToken();

  if (token) {
    // @TODO: Use real jwt token
    const decoded = {}; // jwtDecode(token);
    return decoded;
  }

  return null;
};

/**
 * Validate stored token
 */
export const isStoredTokenValid = () => {
  const token = getToken();

  if (token) {
    // @TODO: make expired validation
    // jwtDecode(token);
    return true;
  }

  return false;
};

/**
 * Remove token from storage
 */
export const removeToken = () => {
  localStorage.removeItem(JWT_STORAGE_KEY);
};
