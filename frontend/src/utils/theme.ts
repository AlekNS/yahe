import { THEME_STORAGE_KEY } from '../config';

/**
 * Get current theme value
 */
export const getThemeFromStorage = (): string => localStorage.getItem(THEME_STORAGE_KEY) || '';

/**
 * Store theme value
 */
export const setThemeToStorage = (theme: string) => localStorage.setItem(THEME_STORAGE_KEY, theme);

/**
 * Available themes
 */
export const THEMES = {
  blue: 'blue',
  dark: 'dark',
};
