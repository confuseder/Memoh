/**
 * Token Storage Interface
 * 
 * Abstraction for storing authentication tokens in different backends
 */

export interface Config {
  apiUrl: string
  token?: string
}

export interface TokenStorage {
  /**
   * Get the API URL
   */
  getApiUrl(): Promise<string> | string

  /**
   * Set the API URL
   */
  setApiUrl(url: string): Promise<void> | void

  /**
   * Get the authentication token for a user
   * @param userId - User identifier (optional for single-user storage)
   */
  getToken(userId?: string): Promise<string | null> | string | null

  /**
   * Set the authentication token for a user
   * @param token - The authentication token
   * @param userId - User identifier (optional for single-user storage)
   */
  setToken(token: string, userId?: string): Promise<void> | void

  /**
   * Clear the authentication token for a user
   * @param userId - User identifier (optional for single-user storage)
   */
  clearToken(userId?: string): Promise<void> | void

  /**
   * Load full configuration (if applicable)
   */
  loadConfig?(): Promise<Config> | Config

  /**
   * Save full configuration (if applicable)
   */
  saveConfig?(config: Config): Promise<void> | void
}

