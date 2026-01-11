/**
 * MemoHome Core Context
 * 
 * Provides a configurable context for core functions to use different storage backends
 */

import type { TokenStorage } from './storage'
import { FileTokenStorage } from './storage/file'

/**
 * Global context for core functions
 */
export interface MemoHomeContext {
  storage: TokenStorage
  currentUserId?: string
}

/**
 * Default context (uses file storage for CLI)
 */
let defaultContext: MemoHomeContext = {
  storage: new FileTokenStorage(),
}

/**
 * Get the current context
 */
export function getContext(): MemoHomeContext {
  return defaultContext
}

/**
 * Set the global context
 * Use this to configure storage backend (e.g., Redis for Telegram bot)
 */
export function setContext(context: Partial<MemoHomeContext>): void {
  defaultContext = { ...defaultContext, ...context }
}

/**
 * Create a new context without modifying the global one
 * Useful for multi-user scenarios
 */
export function createContext(options: {
  storage: TokenStorage
  userId?: string
}): MemoHomeContext {
  return {
    storage: options.storage,
    currentUserId: options.userId,
  }
}

/**
 * Reset context to default (file storage)
 */
export function resetContext(): void {
  defaultContext = {
    storage: new FileTokenStorage(),
  }
}

